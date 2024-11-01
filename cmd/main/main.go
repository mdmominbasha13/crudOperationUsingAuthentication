package main

import (
	"github.com/CrudOperationUsingAuthentication/pkg/logger"
	"github.com/CrudOperationUsingAuthentication/pkg/router"

	"github.com/CrudOperationUsingAuthentication/pkg/config"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")

	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
