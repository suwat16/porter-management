package usecase

import "porter-management/internal/job/domain"

type CreateJobUseCase struct {
	JobRepository domain.JobRepository
}

func (c *CreateJobUseCase) Execute(jobName string, requester domain.Requester, destination domain.Destination) (*domain.Job, error) {
	job, err := domain.CreateNewJob(jobName, requester, destination)
	if err != nil {
		return nil, err
	}

	err = c.JobRepository.Save(&job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
