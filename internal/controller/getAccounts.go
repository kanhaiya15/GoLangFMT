package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/GoLangFMT/internal/db/sqldb"
)

// GetAccounts GetAccounts
func GetAccounts(c *gin.Context) {
	var user User
	dbConn := sqldb.GetConnection()
	if dbConn != nil {
		dbConn.First(&user)
	}
	c.JSON(http.StatusOK, user)
}
