package usecase

import (
	"database/sql"
	"porter-management/internal/job/domain/entity"
	"porter-management/internal/job/infra/repository"
	"porter-management/pkg/uow"

	"github.com/streadway/amqp"
)

type JobUseCase struct {
	DB   *sql.DB
	Amqp *amqp.Channel
}

func NewJobUseCase(db *sql.DB, amqp *amqp.Channel) *JobUseCase {
	return &JobUseCase{
		DB:   db,
		Amqp: amqp,
	}
}

func (u *JobUseCase) ExecuteNewJob(jobName string, requester entity.Requester, destination entity.Destination, equipment entity.Equipment) (*entity.Job, error) {
	job, err := entity.CreateNewJob(jobName, requester, destination, equipment)
	if err != nil {
		return nil, err
	}

	uow := uow.NewUnitOfWork(u.DB)
	err = uow.DoInTransaction(func() error {
		jobRepo := repository.NewJobRepository(uow.Tx)
		_, err := jobRepo.Save(&job)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	u.Amqp.Publish("job", "create-job", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(`{"name": "Job 1", "description": "This is job 1"}`),
	})

	return &job, nil
}
