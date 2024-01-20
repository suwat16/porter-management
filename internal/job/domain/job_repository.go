package domain

type JobRepository interface {
	Save(job *Job) error
	FindOneById(id string) (*Job, error)
}
