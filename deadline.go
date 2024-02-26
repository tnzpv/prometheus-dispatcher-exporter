package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"prometheus-dispatcher-exporter/model"
	"prometheus-dispatcher-exporter/model/deadline"
	"strconv"
	"strings"
	"time"
)

func deadline_jobs(deadline_api string) deadline.Jobs {
	url := deadline_api + "/Jobs"

	req, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error when get http request %s : %s\n", url, err)
		os.Exit(1)
	}
	var result deadline.Jobs
	json.NewDecoder(req.Body).Decode(&result)
	return result
}

func deadline_dispatcher_jobs(deadline_api string) []model.Job {

	var jobs []model.Job

	for _, dispatcher_job := range deadline_jobs(deadline_api) {
		// ************
		// Job duration
		var frame_start int
		var frame_end int
		var duration int

		result := strings.Split(
			dispatcher_job.Props.Frames,
			"-",
		)
		if len(result) > 1 {
			frame_start, _ = strconv.Atoi(result[0])
			frame_end, _ = strconv.Atoi(result[1])
		}
		duration = frame_end - frame_start + 1
		// ************

		job := new(model.Job)
		job.Id = dispatcher_job.ID
		job.Project = dispatcher_job.Props.Ex0
		job.Department = dispatcher_job.Props.Dept
		job.Name = dispatcher_job.Props.Name
		job.User = dispatcher_job.Props.User
		job.Submitted = dispatcher_job.Date
		job.Start = dispatcher_job.DateStart
		job.End = dispatcher_job.DateComp
		job.Frames = duration
		job.Status = deadline.Job_Status[dispatcher_job.Stat]

		dtasks := deadline_tasks(dispatcher_job.ID, deadline_api)

		for _, t := range dtasks.Tasks {

			// ************
			// Task duration
			var frame_start int
			var frame_end int
			var duration int

			result := strings.Split(
				t.Frames,
				"-",
			)
			if len(result) > 1 {
				frame_start, _ = strconv.Atoi(result[0])
				frame_end, _ = strconv.Atoi(result[1])
			}
			duration = frame_end - frame_start + 1
			// ************

			task := new(model.Task)
			task.Id = t.TaskID
			task.Status = deadline.Task_Status[t.Stat]
			task.Worker = t.Slave
			task.Start = t.Start
			task.End = t.Comp
			task.Frames = duration
			task.TimeSpent = task.End.Sub(task.Start)
			if task.TimeSpent < time.Duration(0) {
				task.TimeSpent = time.Duration(0)
			}

			job.Tasks = append(job.Tasks, *task)
		}

		// Sum of all tasks "time spent"
		for _, t := range job.Tasks {
			job.TimeSpent += t.TimeSpent
		}
		jobs = append(jobs, *job)
	}

	return jobs
}

func deadline_tasks(JobID string, deadline_api string) deadline.Tasks {
	url := deadline_api + "/tasks?JobID=" + JobID

	req, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error when get http request %s : %s\n", url, err)
		os.Exit(1)
	}
	var result deadline.Tasks
	json.NewDecoder(req.Body).Decode(&result)
	return result
}

func deadline_workers(deadline_api string) deadline.Workers {
	url := deadline_api + "/Slaves"

	req, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error when get http request %s : %s\n", url, err)
		os.Exit(1)
	}
	var result deadline.Workers
	json.NewDecoder(req.Body).Decode(&result)
	return result
}

func deadline_dispatcher_workers(deadline_api string) []model.Worker {
	var workers []model.Worker

	for _, dispatcher_worker := range deadline_workers(deadline_api) {
		worker := new(model.Worker)
		worker.Name = dispatcher_worker.Info.Name
		worker.Status = deadline.Worker_Status[dispatcher_worker.Info.Stat]
		worker.Enable = dispatcher_worker.Settings.Enable

		workers = append(workers, *worker)
	}

	return workers
}
