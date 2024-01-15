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
	assert.Equal(job.Status, JobStaus(NEW), "job status should be NEW")

	// assert job aggregate have one length
	assert.Equal(len(job.Aggregate.Event), 1, "aggregate should have length 1")
}
