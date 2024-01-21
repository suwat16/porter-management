package usecase

import "porter-management/internal/job-request/domain"

type AccepJobRequestUseCase struct {
	JobRequestRepository domain.JobRequestRepository
}

func (c *AccepJobRequestUseCase) Execute(jobRequest *domain.JobRequest) (domain.JobRequest, error) {
	err := c.JobRequestRepository.Save(jobRequest)
	if err != nil {
		return domain.JobRequest{}, err
	}

	return *jobRequest, nil
}
