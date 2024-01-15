package domain

import "time"

type JobStaus int

const (
	AVAILABLE JobStaus = iota
	WORKING
)

func (jobStatus JobStaus) String() string {
	return [...]string{"AVAILABLE", "WORKING"}[jobStatus]
}

// Porter is aggregate root
type Porter struct {
	Id       int64
	Name     string
	JobStaus JobStaus
}

func CreateNewPorter(name string) (Porter, error) {
	id := time.Now().Unix()

	porter := &Porter{
		Id:       id,
		Name:     name,
		JobStaus: JobStaus(AVAILABLE),
	}

	return *porter, nil
}

// Update job status
func (porter *Porter) UpdateJobStatus(jobStatus JobStaus) {
	porter.JobStaus = jobStatus
}
