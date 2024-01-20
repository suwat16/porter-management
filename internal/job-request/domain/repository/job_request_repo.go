package domain

import "porter-management/internal/job-request/domain"

type JobRequestRepo interface {
	Save(jobRequest *domain.JobRequest) error
	GetById(id string) (*domain.JobRequest, error)
	Remove(id string) error
}
