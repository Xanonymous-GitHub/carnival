package api

import (
	"context"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/proto/entpb"
	"github.com/Xanonymous-GitHub/carnival/server/internal/grpcsvc"
	"github.com/Xanonymous-GitHub/carnival/server/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
)

type ApplicationHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type applicationHandler struct{}

func NewApplicationHandler() ApplicationHandler {
	return &applicationHandler{}
}

func (h *applicationHandler) Create(c *gin.Context) {
	var applicationData *entpb.Application

	err := c.BindJSON(&applicationData)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if err := service.CreateApplication(ctx, grpcsvc.ApplicationSvcClient, applicationData); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (h *applicationHandler) Get(c *gin.Context) {
	id := strings.TrimSpace(c.Query("id"))
	if len(id) != 16 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	application, err := service.GetApplication(ctx, grpcsvc.ApplicationSvcClient, id)
	if err != nil {
		log.Println(err)
		se, _ := status.FromError(err)
		if se.Code() == codes.NotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, application)
}
