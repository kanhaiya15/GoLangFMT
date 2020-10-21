package sqldb

import (
	"context"
	"fmt"
	"sync"

	config "github.com/kanhaiya15/GoLangFMT/cfg"
	"github.com/kanhaiya15/GoLangFMT/pkg/lumber"

	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
var (
	DbConn *gorm.DB
	logger lumber.Logger
)

// Setup initializes all crons on service startup
func Setup(ctx context.Context, config *config.Config, wg *sync.WaitGroup, initializedLogger lumber.Logger) {
	defer wg.Done()
	var err error
	logger = initializedLogger
	logger.Infof("Setting up mySQL")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBConf.User, config.DBConf.Password, config.DBConf.Host, config.DBConf.Name)
	logger = logger.WithFields(lumber.Fields{"dsn": dsn})
	logger.Infof("mysql - creating connection to database")
	DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql - error while creating connection to database", err.Error())
		os.Exit(1)
	}
	sqlDB, err := DbConn.DB()
	if err != nil {
		logger.Errorf("mysql - error while creating connection to database", err.Error())
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(time.Hour)

	select {
	case <-ctx.Done():
		sqlDB.Close()
		logger.Infof("Caller has requested graceful shutdown. Returning.....")
		return
	}
}

// GetConnection GetConnection
func GetConnection() *gorm.DB {
	return DbConn
}
