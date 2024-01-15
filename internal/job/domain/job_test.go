package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewJob(t *testing.T) {
	assert := assert.New(t)

	job, err := CreateNewJob()

	// assert error is nil
	assert.Nil(err, "error should be nil")

	// assert job status is NEW
	assert.Equal(job.Status, "NEW", "status should be NEW")

	// assert job aggregate have one length
	assert.Equal(len(job.Aggregate.Event), 1, "aggregate should have length 1")
}
