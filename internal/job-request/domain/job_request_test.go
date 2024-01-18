package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewJobRequest(t *testing.T) {
	assert := assert.New(t)

	job := Job{
		Id:   1,
		Name: "Job 1",
	}

	JobRequest, err := CreateNewJobRequest(job)

	// assert error is nil
	assert.Nil(err, "error should be nil")

	// assert job status is WAITING_FOR_PORTER
	assert.Equal(JobRequest.Status, WAITING_FOR_PORTER, "job status should be WAITING_FOR_PORTER")

	// assert job aggregate event is not empty
	assert.NotEmpty(JobRequest.Aggregate.Event, "job aggregate event should not be empty")
}

func TestJobRequestAccept(t *testing.T) {
	assert := assert.New(t)

	job := Job{
		Id:   1,
		Name: "Job 1",
	}

	porter := Porter{
		Id:   1,
		Name: "Porter 1",
	}

	JobRequest, _ := CreateNewJobRequest(job)

	JobRequest.PorterAcceptJobRequest(&porter)

	// assert job status is PORTER_ACCEPTED
	assert.Equal(JobRequest.Status, PORTER_ACCEPTED, "job status should be PORTER_ACCEPTED")

	// assert job aggregate event is not empty
	assert.NotEmpty(JobRequest.Aggregate.Event, "job aggregate event should not be empty")
}
