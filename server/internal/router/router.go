package router

import (
	"github.com/Xanonymous-GitHub/carnival/server/internal/api"
	. "github.com/Xanonymous-GitHub/carnival/server/pkg/vp"
	"github.com/gin-gonic/gin"
)

func routerMode() string {
	isDebugMode := Vp.GetBool("debug")

	if isDebugMode {
		return gin.DebugMode
	}

	return gin.ReleaseMode
}

func NewRouter() *gin.Engine {
	mode := routerMode()
	gin.SetMode(mode)

	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	const (
		apiRootPath        = "/api"
		applicationApiPath = apiRootPath + "/application"
	)

	// Application APIs
	applicationApiHandlers := api.NewApplicationHandler()
	applicationApiGroup := r.Group(applicationApiPath)
	{
		applicationApiGroup.POST("", applicationApiHandlers.Create)
	}

	return r
}
