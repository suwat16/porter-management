package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewJobRequest(t *testing.T) {
	assert := assert.New(t)

	JobRequest, err := CreateNewJobRequest(Job{})

	// assert error is nil
	assert.Nil(err, "error should be nil")

	// assert job status is WAITING_FOR_PORTER
	assert.Equal(JobRequest.Status, WAITING_FOR_PORTER, "job status should be WAITING_FOR_PORTER")

	// assert job aggregate event is not empty
	assert.NotEmpty(JobRequest.Aggregate.Event, "job aggregate event should not be empty")
}

func TestJobRequestAccept(t *testing.T) {

}
