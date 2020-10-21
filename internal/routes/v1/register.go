package v1

import (
	"net/http"

	"github.com/LambdaTest/mould/internal/middleware"
	"github.com/LambdaTest/mould/pkg/lumber"
	"github.com/gin-gonic/gin"
)

var logger lumber.Logger

// RegisterRoutes registers routes and logger in the package
func RegisterRoutes(router *gin.RouterGroup, logger lumber.Logger) {
	router.POST("/hello", middleware.Authentication(), versionedHello)
	logger.Infof("v1.0 routes injected")
}

func versionedHello(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, []byte("{\"status\": \"VALID\"}"))
}
