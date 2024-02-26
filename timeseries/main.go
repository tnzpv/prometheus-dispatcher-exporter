package timeseries

import (
	"github.com/prometheus/client_golang/prometheus"
	"prometheus-dispatcher-exporter/model"
	"strconv"
)

var JobMetrics = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "dispatcher_jobs",
		Help: "Jobs metrics in the dispatcher",
	},
	[]string{
		"Id",
		"Project",
		"Department",
		"Name",
		"User",
		"Start",
		"End",
		"Frames",
		"Status",
	},
)

var TaskMetrics = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "dispatcher_tasks",
		Help: "Tasks metrics in the dispatcher",
	},
	[]string{
		"Id",
		"JobID",
		"Project",
		"Department",
		"JobName",
		"User",
		"Start",
		"End",
		"Frames",
		"Status",
		"Worker",
	},
)

var WorkerMetrics = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "dispatcher_workers",
		Help: "Workers metrics in the dispatcher",
	},
	[]string{
		"Name",
		"Status",
		"Enable",
	},
)

func IngestMetrics(jobs []model.Job, workers []model.Worker) {

	JobMetrics.Reset()
	TaskMetrics.Reset()
	WorkerMetrics.Reset()

	for _, job := range jobs {
		job_start := float64(job.Start.Unix())
		job_end := float64(job.End.Unix())
		JobMetrics.WithLabelValues(
			job.Id,
			job.Project,
			job.Department,
			job.Name,
			job.User,
			strconv.Itoa(int(job_start)),
			strconv.Itoa(int(job_end)),
			strconv.Itoa(job.Frames),
			job.Status,
		).Set(job.TimeSpent.Seconds())

		for _, task := range job.Tasks {
			task_start := float64(task.Start.Unix())
			task_end := float64(task.End.Unix())
			TaskMetrics.WithLabelValues(
				strconv.Itoa(task.Id),
				job.Id,
				job.Project,
				job.Department,
				job.Name,
				job.User,
				strconv.Itoa(int(task_start)),
				strconv.Itoa(int(task_end)),
				strconv.Itoa(task.Frames),
				task.Status,
				task.Worker,
			).Set(task.TimeSpent.Seconds())
		}
	}

	for _, worker := range workers {
		var enable int8
		if worker.Enable {
			enable = 1
		} else {
			enable = 0
		}

		WorkerMetrics.WithLabelValues(
			worker.Name,
			worker.Status,
			strconv.Itoa(int(enable)),
		).Set(float64(enable))
	}
}
