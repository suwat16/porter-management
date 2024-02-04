package repository

import (
	"database/sql"
	"porter-management/internal/job/domain/entity"
)

type JobRepository struct {
	tx *sql.Tx
}

func NewJobRepository(tx *sql.Tx) *JobRepository {
	return &JobRepository{tx: tx}
}

func (repo *JobRepository) Save(job *entity.Job) (entity.Job, error) {
	_, err := repo.tx.Exec("INSERT INTO jobs (id, version, name, status, requester_name, requester_position, destination_building, destination_floor, destination_room) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", job.Id, job.Version, job.Name, job.Status, job.Requester.Name, job.Requester.Position, job.Destination.Building, job.Destination.Floor, job.Destination.Room)
	if err != nil {
		return entity.Job{}, err
	}

	return *job, nil
}
