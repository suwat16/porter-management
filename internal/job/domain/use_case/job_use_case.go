package usecase

import (
	"porter-management/config"
	"porter-management/internal/job/domain/entity"

	"github.com/streadway/amqp"

	JobRepo "porter-management/internal/job/infra/repository"
)

type JobUseCase struct {
	Uow  *config.Uow
	Amqp *amqp.Channel
}

func NewJobUseCase(uow *config.Uow, amqp *amqp.Channel) *JobUseCase {
	return &JobUseCase{
		Uow:  uow,
		Amqp: amqp,
	}
}

func (jobUseCase *JobUseCase) ExecuteNewJob(jobName string, requester entity.Requester, destination entity.Destination, equipment entity.Equipment) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination, equipment)
	if err != nil {
		return nil, err
	}

	err = jobUseCase.Uow.DoInTransaction(func() error {
		jobRepo := JobRepo.NewJobRepository(jobUseCase.Uow.Tx)
		_, err := jobRepo.Save(&job)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	jobUseCase.Amqp.Publish("job", "create-job", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(`{"name": "Job 1", "description": "This is job 1"}`),
	})

	return &job, nil
}
