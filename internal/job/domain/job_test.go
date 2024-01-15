package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewJob(t *testing.T) {
	assert := assert.New(t)

	jobName := "job name"

	requester := Requester{
		Name:     "requester name",
		Position: "requester position",
	}

	destination := Destination{
		Building: "destination building",
		Floor:    "destination floor",
		Room:     "destination room",
	}

	job, err := CreateNewJob(jobName, requester, destination)

	// assert error is nil
	assert.Nil(err, "error should be nil")

	// assert job status is NEW
	assert.Equal(job.Status, JobStaus(WAITING), "job status should be NEW")

	// assert job aggregate event is not empty
	assert.NotEmpty(job.Aggregate.Event, "job aggregate event should not be empty")
}
