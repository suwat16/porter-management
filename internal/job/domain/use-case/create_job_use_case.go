package usecase

import "porter-management/internal/job/domain"

type CreateJobUseCase struct {
	JobRepository domain.JobRepository
}

func (c *CreateJobUseCase) Execute(job *domain.Job) (*domain.Job, error) {
	err := c.JobRepository.Save(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
