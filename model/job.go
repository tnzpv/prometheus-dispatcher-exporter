package model

import (
	"time"
)

type Job struct {
	Id         string
	Project    string
	Department string
	Name       string
	User       string
	Submitted  time.Time
	Start      time.Time
	End        time.Time
	Frames     int
	Status     string
	TimeSpent  time.Duration
	Tasks      []Task
}

type Task struct {
	Id        int
	Start     time.Time
	End       time.Time
	Frames    int
	Status    string
	Worker    string
	TimeSpent time.Duration
}
