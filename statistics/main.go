package statistics

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"prometheus-dispatcher-exporter/model"
)

func PushToPsql(jobs []model.Job, host string, user string, password string, port int, dbname string) {
	dbName := dbname
	tableName := "jobs"
	dbHost := host
	dbUser := user
	dbPassword := password
	dbPort := port

	// CONNECT TO DATABASE
	conninfo := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable dbname=%s port=%d", dbUser, dbPassword, dbHost, dbName, dbPort)
	psql, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}

	// CREATE TABLE
	rows, err := psql.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id SERIAL,jid text,project text,department text,name text,\"user\" text,submitted timestamp,start timestamp,\"end\" timestamp,frames integer,status text,time_spent real)", tableName))
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	for _, job := range jobs {
		var count int
		req := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE jid=$1", tableName)
		row := psql.QueryRow(req, job.Id)
		if row.Err() != nil {
			log.Fatal(row.Err())
		}
		row.Scan(&count)
		if count >= 1 {
			// found
			req = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE project=$2 AND department=$3 AND name=$4 AND \"user\"=$5 AND submitted=$11 AND start=$6 AND \"end\"=$7 AND frames=$8 AND status=$9 AND time_spent=$10 AND jid=$1", tableName)
			row := psql.QueryRow(
				req,
				job.Id,
				job.Project,
				job.Department,
				job.Name,
				job.User,
				job.Start,
				job.End,
				job.Frames,
				job.Status,
				job.TimeSpent.Seconds(),
				job.Submitted,
			)
			if row.Err() != nil {
				log.Fatal(row.Err())
			}
			row.Scan(&count)
			if count == 0 {
				// update
				log.Printf("POSTGRESQL Update job %s...", job.Id)
				req = fmt.Sprintf("UPDATE %s SET project=$2,department=$3,name=$4,\"user\"=$5,submitted=$11,start=$6,\"end\"=$7,frames=$8,status=$9,time_spent=$10 WHERE jid=$1", tableName)
				rows, err := psql.Query(
					req,
					job.Id,
					job.Project,
					job.Department,
					job.Name,
					job.User,
					job.Start,
					job.End,
					job.Frames,
					job.Status,
					job.TimeSpent.Seconds(),
					job.Submitted,
				)
				if err != nil {
					log.Fatal(err)
				}
				rows.Close()
				log.Printf("POSTGRESQL Update job %s... Done.", job.Id)
			}
		} else {
			// insert
			log.Printf("POSTGRESQL Insert job %s...", job.Id)
			req = fmt.Sprintf("INSERT INTO %s (jid,project,department,name,\"user\",submitted,start,\"end\",frames,status,time_spent) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id", tableName)
			rows, err := psql.Query(req, job.Id, job.Project, job.Department, job.Name, job.User, job.Submitted, job.Start, job.End, job.Frames, job.Status, job.TimeSpent.Seconds())
			if err != nil {
				log.Fatal(err)
			}
			rows.Close()
			log.Printf("POSTGRESQL Insert job %s... Done.", job.Id)
		}
	}
	psql.Close()
}
