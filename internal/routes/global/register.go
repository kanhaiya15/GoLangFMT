package global

import (
	"net/http"

	"github.com/kanhaiya15/GoLangFMT/pkg/lumber"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes and logger in the package
func RegisterRoutes(router *gin.RouterGroup, logger lumber.Logger) {

	router.GET("/health", health)
	router.GET("/status", status)
	logger.Infof("global routes injected")
}

// health API
func health(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEPlain, []byte("OK"))
}

// status api
func status(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, []byte("{\"status\": \"OK\"}"))
}
