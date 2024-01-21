package domain

type JobRequestRepository interface {
	Save(jobRequest *JobRequest) error
	GetById(id string) (*JobRequest, error)
	Remove(id string) error
}
