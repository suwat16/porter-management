package repository

import (
	"database/sql"
	"porter-management/internal/job-request/domain"
)

type JobRepository struct {
	db *sql.Tx
}

func NewJobRepository(db *sql.Tx) *JobRepository {
	return &JobRepository{db: db}
}

func (repo *JobRepository) Save(job *domain.Job) (domain.Job, error) {
	row := repo.db.QueryRow("select id from job where id = $1", job.Id)

	var id int64
	if row.Scan(&id) == nil {
		// _, err := repo.db.Exec("update job set status = $1, updated_at = $2 where id = $3",
		// job.Status, job.Id)
		// return err
	}

	// _, err := repo.db.Exec("insert into job (id, status, created_at, updated_at) values ($1, $2, $3, $4)", job.Id, job.Status, job.CreatedAt, job.UpdatedAt)

	// if err != nil {
	// return err
	// }

	return *job, nil
}

func (repo *JobRepository) FindOneById(id int64) (domain.Job, error) {
	return domain.Job{}, nil
}
