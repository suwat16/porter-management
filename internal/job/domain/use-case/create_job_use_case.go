package usecase

import (
	"porter-management/internal/job/domain/entity"
	"porter-management/internal/job/infra/repository"
)

type CreateJobUseCase struct {
	JobRepository repository.JobRepository
}

func (c *CreateJobUseCase) Execute(jobName string, requester entity.Requester, destination entity.Destination) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination)
	if err != nil {
		return nil, err
	}

	_, err = c.JobRepository.Save(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
