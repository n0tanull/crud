package app

import (
	"main/internal/controller"
	"main/internal/repo"
	"main/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func Run() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	log := logrus.New()
	log.Level = logrus.DebugLevel
	controller := controller.NewController(service.NewServ(repo.NewRepository(ctx)), log)
	r := gin.Default()
	r.Group("/api")
	{
		r.GET("/users", controller.GetUser)
		r.POST("/users", controller.AddUser)
		r.DELETE("/users/:id", controller.DeleteUser)
		r.PUT("/users/:id", controller.ChangeUser)
	}
	r.Run(":8080")
}
