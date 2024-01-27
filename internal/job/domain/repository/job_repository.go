package repository

import "porter-management/internal/job/domain/entity"

type JobRepository interface {
	Save(job *entity.Job) error
	FindOneById(id string) (*entity.Job, error)
}
