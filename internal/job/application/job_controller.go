package application

import (
	"database/sql"

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

func (jobController *JobController) RegisterRoutes() {
	//Exchange Declare
	err := jobController.Ch.ExchangeDeclare("job", "direct", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	jobController.Gin.GET("/job", jobController.GetJobs)
}

func (jobController *JobController) GetJobs(c *gin.Context) {
	// uow := config.NewUnitOfWork(jobController.Db)

	// jobU := jobUseCase.NewJobUseCase(uow, jobController.Ch)

	// requester := Job.Requester{
	// 	Name:     "John Doe",
	// 	Position: "Manager",
	// }

	// destination := Job.Destination{
	// 	Building: "A",
	// 	Floor:    "1",
	// 	Room:     "101",
	// }

	// equipment := Job.Equipment{
	// 	Name:     "Laptop",
	// 	Quantity: 1,
	// }

	// newJob, err := jobU.ExecuteNewJob("Job 1", requester, destination, equipment)

	// if err != nil {
	// 	c.JSON(500, gin.H{
	// 		"message": err.Error(),
	// 	})
	// }

	// c.JSON(200, gin.H{
	// 	"message": "Hello World",
	// 	"data":    newJob,
	// })
}
