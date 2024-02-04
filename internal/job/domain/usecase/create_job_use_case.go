package usecase

import (
	"porter-management/config"
	"porter-management/internal/job/domain/entity"
	"porter-management/internal/job/infra/repository"
)

type CreateJobUseCase struct {
	JobRepository repository.JobRepository
}

func (c *CreateJobUseCase) Execute(uow config.Uow, jobName string, requester entity.Requester, destination entity.Destination, equipment entity.Equipment) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination, equipment)
	if err != nil {
		return nil, err
	}

	err = uow.DoInTransaction(func() error {
		_, err := c.JobRepository.Save(&job)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &job, nil
}
