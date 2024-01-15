package domain

import "time"

// Job object value
type Job struct {
	Id   int64
	Name string
}

// Porter object value
type Porter struct {
	Id   int64
	Name string
}

// JobRequestStatus enum
type JobRequestStatus int

const (
	WAITING_FOR_PORTER JobRequestStatus = iota
	PORTER_ACCEPTED
)

func (jobRequestStatus JobRequestStatus) String() string {
	return [...]string{"WAITING_FOR_PORTER", "PORTER_ACCEPTED"}[jobRequestStatus]
}

// JobRequest aggregate root
type JobRequest struct {
	Id        int64
	Version   int64
	Job       Job
	Porter    *Porter
	Status    JobRequestStatus
	Aggregate Aggregate
}

type Aggregate struct {
	Event []map[string]interface{}
}

func CreateNewJobRequest(job Job) (JobRequest, error) {
	id := time.Now().Unix()

	jobRequest := &JobRequest{
		Id:      id,
		Version: 1,
		Job:     job,
		Status:  WAITING_FOR_PORTER,
	}

	jobRequest.jobRequestPushEvent("NOTIFY_PORTER")
	return *jobRequest, nil
}

// Push notification to porters
func (jobRequest *JobRequest) jobRequestPushEvent(eventName string) {
	event := map[string]interface{}{
		"event": eventName,
		"data":  jobRequest,
	}

	jobRequest.Aggregate.Event = append(jobRequest.Aggregate.Event, event)
}

func PorterAcceptJobRequest(jobRequest *JobRequest, porter *Porter) (JobRequest, error) {
	jobRequest.Porter = porter
	jobRequest.Version = jobRequest.Version + 1
	jobRequest.Status = PORTER_ACCEPTED

	// Push event update job status is delivery
	jobRequest.jobRequestPushEvent("PORTER_ACCEPTED")

	return *jobRequest, nil
}
