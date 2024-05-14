package controller

import (
	"context"
	"main/internal/entity"
	"main/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	ChangeUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type Con struct {
	service.Servicer
	log logrus.Logger
}

func NewController(s service.Servicer, l logrus.Logger) Controller {
	return &Con{s, l}
}
func (c *Con) AddUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	cx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = c.Servicer.AddUser(user, cx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	ctx.JSON(200, user)
}
func (c *Con) GetUser(ctx *gin.Context) {
	queryId := ctx.Query("id")
	id, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	cx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	user, err := c.Servicer.GetUser(int(id), cx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	ctx.JSON(200, user)
}
func (c *Con) ChangeUser(ctx *gin.Context) {
	var user entity.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	cx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = c.Servicer.ChangeUser(user, cx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	ctx.JSON(200, user)
}
func (c *Con) DeleteUser(ctx *gin.Context) {
	queryId := ctx.Query("id")
	id, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	cx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = c.Servicer.DeleteUser(int(id), cx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		c.log.Error(err)
	}
	ctx.JSON(200, gin.H{"id": id})
}
