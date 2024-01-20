package domain

type JobRequestRepo interface {
	Save(jobRequest *JobRequest) error
	GetById(id string) (*JobRequest, error)
}
