package controller

import (
	"porter-management/config"
	entity "porter-management/internal/job/domain/entity"
	usecase "porter-management/internal/job/domain/use_case"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	JobUseCase usecase.JobUseCase
	Uow        config.Uow
}

func NewJobController(jobUseCase usecase.JobUseCase, uow config.Uow) *JobController {
	return &JobController{
		JobUseCase: jobUseCase,
		Uow:        uow,
	}
}

func (jobC *JobController) InitController(router *gin.Engine) {
	JobController := router.Group("/job")
	{
		JobController.POST("/new")
	}
}

func (jobC *JobController) NewJob(c *gin.Context) {

	jobName := c.PostForm("jobName")

	requester := entity.Requester{
		Name: c.PostForm("requesterName"),
	}

	destination := entity.Destination{
		Building: c.PostForm("building"),
	}

	equipment := entity.Equipment{
		Name: c.PostForm("equipmentName"),
	}

	jobC.JobUseCase.ExecuteNewJob(jobC.Uow, jobName, requester, destination, equipment)

}
