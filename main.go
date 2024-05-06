package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"prometheus-dispatcher-exporter/statistics"
	"prometheus-dispatcher-exporter/timeseries"
	"strconv"
	"time"
)

var Version string
var dispatcher_jobs = deadline_dispatcher_jobs
var dispatcher_workers = deadline_dispatcher_workers


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}


func main() {
	log.Printf("Dispatcher exporter starting.")
	log.Printf("Registering metrics objects...")
	prometheus.MustRegister(timeseries.JobMetrics)
	prometheus.MustRegister(timeseries.TaskMetrics)
	prometheus.MustRegister(timeseries.WorkerMetrics)
	log.Printf("Registering metrics objects... Done.")

	// environment variables support, default values setting
	prometheus_enabled, _ := strconv.ParseBool(getEnv("PROMETHEUS_ENABLED", "true"))
	prometheus_port, _ := strconv.Atoi(getEnv("PROMETHEUS_PORT", "9101"))
	postgresql_enabled, _ := strconv.ParseBool(getEnv("POSTGRESQL_ENABLED", "false"))
	postgresql_host := getEnv("POSTGRESQL_HOST", "")
	postgresql_user := getEnv("POSTGRESQL_USER", "")
	postgresql_password := getEnv("POSTGRESQL_PASSWORD", "")
	postgresql_port, _ := strconv.Atoi(getEnv("POSTGRESQL_PORT", "5432"))
	postgresql_db := getEnv("POSTGRESQL_DB", "dispatcher")
	heartbeat_metrics, _ := strconv.Atoi(getEnv("HEARTBEAT_METRICS", "300"))
	heartbeat_statistics, _ := strconv.Atoi(getEnv("HEARTBEAT_STATISTICS", "300"))
	dispatcher_api := getEnv("DISPATCHER_API", "")

	// command line arguments support (use default values from environment variables)
	var exporterFlag = flag.NewFlagSet("prometheus-dispatcher-exporter", flag.ExitOnError)
	var exportVersion = exporterFlag.Bool("version", false, "Farm exporter version")
	var prometheusExporterEnable = exporterFlag.Bool("prometheus", prometheus_enabled, "Launch prometheus exporter")
	var prometheusExporterPort = exporterFlag.Int("port", prometheus_port, "Prometheus exporter listening port")

	var postgresqlExporterEnable = exporterFlag.Bool("postgresql", postgresql_enabled, "Launch postgresql exporter")
	var postgresqlExporterHost = exporterFlag.String("postgresql-host", postgresql_host, "Postgresql host")
	var postgresqlExporterUser = exporterFlag.String("postgresql-user", postgresql_user, "Postgresql user")
	var postgresqlExporterPassword = exporterFlag.String("postgresql-password", postgresql_password, "Postgresql password")
	var postgresqlExporterPort = exporterFlag.Int("postgresql-port", postgresql_port, "Postgresql port")
	var postgresqlExporterDbName = exporterFlag.String("postgresql-dbname", postgresql_db, "Postgresql database name")

	var heartbeatMetricsFlag = exporterFlag.Int("heartbeat-metrics", heartbeat_metrics, "Metrics heartbeat (seconds)")
	var heartbeatStatisticsFlag = exporterFlag.Int("heartbeat-statistics", heartbeat_statistics, "Statistics heartbeat (seconds)")
	
	var dispatcherApi = exporterFlag.String("dispatcher-api", dispatcher_api, "Dispatcher API host")

	exporterFlag.Parse(os.Args[1:])

	if *exportVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	if !*prometheusExporterEnable && !*postgresqlExporterEnable {
		fmt.Println("No exporter activated.")
		os.Exit(1)
	}

	go func() {
		for {
			if *prometheusExporterEnable {
				log.Printf("Get dispatcher jobs...")
				jobs := dispatcher_jobs(*dispatcherApi)
				log.Printf("Get dispatcher jobs... Done.")
				log.Printf("Get dispatcher workers...")
				workers := dispatcher_workers(*dispatcherApi)
				log.Printf("Get dispatcher workers... Done.")

				log.Printf("Expose jobs and workers for prometheus...")
				timeseries.IngestMetrics(jobs, workers)
				log.Printf("Expose jobs and workers for prometheus... Done.")
				
				log.Printf("Sleep...")
				time.Sleep(time.Duration(*heartbeatMetricsFlag) * time.Second)
				log.Printf("Sleep... Done.")
			}
		}
	}()

	go func() {
		for {
			if *postgresqlExporterEnable {
				log.Printf("Get dispatcher jobs...")
				jobs := dispatcher_jobs(*dispatcherApi)
				log.Printf("Get dispatcher jobs... Done.")

				log.Printf("Push jobs to postgresql...")
				statistics.PushToPsql(jobs, *postgresqlExporterHost, *postgresqlExporterUser, *postgresqlExporterPassword, *postgresqlExporterPort, *postgresqlExporterDbName)
				log.Printf("Push jobs and workers to postgresql... Done.")
				
				log.Printf("Sleep...")
				time.Sleep(time.Duration(*heartbeatStatisticsFlag) * time.Second)
				log.Printf("Sleep... Done.")
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *prometheusExporterPort), nil))
}
