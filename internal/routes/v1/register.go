package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/GoLangFMT/internal/controller"
	"github.com/kanhaiya15/GoLangFMT/internal/middleware"
	"github.com/kanhaiya15/GoLangFMT/pkg/lumber"
)

var logger lumber.Logger

// RegisterRoutes registers routes and logger in the package
func RegisterRoutes(router *gin.RouterGroup, logger lumber.Logger) {
	router.GET("/hello", middleware.Authentication(), versionedHello)
	router.GET("/getAccounts", controller.GetAccounts)
	logger.Infof("v1.0 routes injected")
}

func versionedHello(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, []byte("{\"status\": \"VALID\"}"))
}
