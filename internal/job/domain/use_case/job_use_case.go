package usecase

import (
	"porter-management/config"
	"porter-management/internal/job/domain/entity"
	"porter-management/internal/job/infra/repository"
)

type JobUseCase struct {
	JobRepository repository.JobRepository
}

func NewJobUseCase(jobRepository repository.JobRepository) *JobUseCase {
	return &JobUseCase{
		JobRepository: jobRepository,
	}
}

func (jobUseCase *JobUseCase) ExecuteNewJob(uow config.Uow, jobName string, requester entity.Requester, destination entity.Destination, equipment entity.Equipment) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination, equipment)
	if err != nil {
		return nil, err
	}

	err = uow.DoInTransaction(func() error {
		_, err := jobUseCase.JobRepository.Save(&job)
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
