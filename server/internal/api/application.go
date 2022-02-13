package api

import (
	"context"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	"github.com/Xanonymous-GitHub/carnival/server/internal/grpcsvc"
	"github.com/Xanonymous-GitHub/carnival/server/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ApplicationHandler interface {
	Create(c *gin.Context)
}

type applicationHandler struct{}

func NewApplicationHandler() ApplicationHandler {
	return &applicationHandler{}
}

func (h *applicationHandler) Create(c *gin.Context) {
	var applicationData *entpb.Application

	err := c.ShouldBindJSON(applicationData)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	ctx := context.Background()
	service.CreateApplication(ctx, grpcsvc.ApplicationSvcClient, applicationData)
}
