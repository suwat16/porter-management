package domain

import "time"

// Job enum status
type JobStaus int

const (
	WAITING JobStaus = iota
	DELIVERY
	DONE
)

func (jobStatus JobStaus) String() string {
	return [...]string{"WAITING", "DELIVERY", "DONE"}[jobStatus]
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
		Status:      JobStaus(WAITING),
	}

	job.jobPushEvent("CREATE_JOB")
	return *job, nil
}

// Push event to job request
func (job *Job) jobPushEvent(eventName string) {
	event := map[string]interface{}{
		"event": eventName,
		"data":  job,
	}

	job.Aggregate.Event = append(job.Aggregate.Event, event)
}
