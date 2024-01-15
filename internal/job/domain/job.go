package domain

import "time"

type Job struct {
	Id        int64
	Version   int64
	Status    string
	Aggregate Aggregate
}

type Aggregate struct {
	Event []map[string]interface{}
}

func CreateNewJob() (Job, error) {

	job := &Job{
		Id:      time.Now().Unix(),
		Version: 1,
		Status:  "NEW",
	}

	job.jobPushEvent()
	return *job, nil
}

func (job *Job) jobPushEvent() {
	event := map[string]interface{}{
		"event": "NEW",
		"date":  time.Now().Unix(),
	}
	job.Aggregate.Event = append(job.Aggregate.Event, event)
}
