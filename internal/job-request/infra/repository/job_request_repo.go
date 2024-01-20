package infra

import (
	"database/sql"
	"porter-management/internal/job-request/domain"
	"time"
)

type JobRequestRepo struct {
	db *sql.DB
}

func NewJobRequestRepo(db *sql.DB) *JobRequestRepo {
	return &JobRequestRepo{db: db}
}

func (repo *JobRequestRepo) Save(jobRequest *domain.JobRequest) error {

	// select id from job_request where id = $1
	row := repo.db.QueryRow("select id from job_request where id = $1", jobRequest.Id)

	// if id exists, update
	var id int64
	if row.Scan(&id) == nil {
		// update job_request set status = $1, updated_at = $2 where id = $3
		_, err := repo.db.Exec("update job_request set status = $1, updated_at = $2 where id = $3",
			jobRequest.Status, jobRequest.Id)
		return err
	}

	// else insert
	_, err := repo.db.Exec("insert into job_request (id, status, created_at, updated_at) values ($1, $2, $3, $4)", jobRequest.Id, jobRequest.Status, time.Now(), time.Now())

	// return error if any
	if err != nil {
		return err
	}

	return nil
}

func (repo *JobRequestRepo) FindOneById(id int64) (domain.JobRequest, error) {
	// select id, status, created_at, updated_at from job_request where id = $1
	row := repo.db.QueryRow("select id, status, created_at, updated_at from job_request where id = $1", id)

	// if id exists, return JobRequest
	var jobRequest domain.JobRequest
	if row.Scan(&jobRequest.Id, &jobRequest.Status) == nil {
		return jobRequest, nil
	}

	// else return error
	return jobRequest, nil
}
