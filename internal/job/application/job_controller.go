package application

import (
	"database/sql"
	"porter-management/internal/job/domain/entity"
	usecase "porter-management/internal/job/domain/use_case"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type JobController struct {
	Gin *gin.Engine
	Db  *sql.DB
	Ch  *amqp.Channel
}

func NewJobController(gin *gin.Engine, db *sql.DB, ch *amqp.Channel) *JobController {
	return &JobController{Gin: gin, Db: db, Ch: ch}
}

func (jc *JobController) RegisterRoutes(routePath string) {
	//Exchange Declare
	err := jc.Ch.ExchangeDeclare("job", "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	jc.Gin.POST(routePath, jc.CreateJob)
}

func (jc *JobController) CreateJob(c *gin.Context) {
	requester := entity.Requester{
		Name:     "John Doe",
		Position: "Manager",
	}

	destination := entity.Destination{
		Building: "A",
		Floor:    "1",
		Room:     "101",
	}

	equipment := entity.Equipment{
		Name:     "Laptop",
		Quantity: 1,
	}

	usecase := usecase.NewJobUseCase(jc.Db, jc.Ch)
	_, err := usecase.ExecuteNewJob("Job 1", requester, destination, equipment)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	// convert job to json and return
	c.JSON(200, gin.H{
		"message": "success",
		"body":    "{}",
	})
}
