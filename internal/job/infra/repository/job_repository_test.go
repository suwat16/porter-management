package repository

import (
	"porter-management/internal/job/domain/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	repo := NewJobRepository(tx)

	jobName := "job name"

	requester := entity.Requester{
		Name:     "requester name",
		Position: "requester position",
	}

	destination := entity.Destination{
		Building: "destination building",
		Floor:    "destination floor",
		Room:     "destination room",
	}

	equipment := entity.Equipment{
		Name:     "equipment name",
		Quantity: 1,
	}

	job, _ := entity.CreateNewJob(jobName, requester, destination, equipment)

	mock.ExpectExec("INSERT INTO jobs").
		WithArgs(job.Id, job.Name, job.Requester.Name, job.Requester.Position, job.Destination.Building, job.Destination.Floor, job.Destination.Room, job.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Save(&job)
	if err != nil {
		t.Errorf("error was not expected while saving job: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
