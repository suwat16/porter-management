package domain

import "time"

// Job enum status
type JobStaus int

const (
	NEW JobStaus = iota
	PROGRESS
	DONE
)

func (jobStatus JobStaus) String() string {
	return [...]string{"NEW", "PROGRESS", "DONE"}[jobStatus]
}

// Requester is Object value
type Requester struct {
	Name     string
	Position string
}

// Destination is Object value
type Destination struct {
	Building string
	Floor    string
	Room     string
}

// Job status is enum
type Job struct {
	Id          int64
	Version     int64
	Name        string
	Status      JobStaus
	Requester   Requester
	Destination Destination
	Aggregate   Aggregate
}

type Aggregate struct {
	Event []map[string]interface{}
}

func CreateNewJob(name string, requester Requester, destination Destination) (Job, error) {
	id := time.Now().Unix()

	job := &Job{
		Id:          id,
		Version:     1,
		Name:        name,
		Requester:   requester,
		Destination: destination,
		Status:      JobStaus(NEW),
	}

	job.jobPushEvent()
	return *job, nil
}

// Push event to job request
func (job *Job) jobPushEvent() {
	event := map[string]interface{}{
		"event": "NEW",
		"date":  time.Now().Unix(),
	}
	job.Aggregate.Event = append(job.Aggregate.Event, event)
}
