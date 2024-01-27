package usecase

import (
	infra "porter-management/internal/job-request/infra/repository"
	"porter-management/internal/job/domain/entity"
)

type CreateJobUseCase struct {
	JobRepository infra.JobRequestRepo
}

func (c *CreateJobUseCase) Execute(jobName string, requester entity.Requester, destination entity.Destination) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
