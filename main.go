package main

import (
	"os"

	r "github.com/liam-lai/service-show-filter/app/router"
	"github.com/liam-lai/service-show-filter/app/util"
)

func main() {
	var port, logLevel string

	if logLevel = os.Getenv("LOG_LEVEL"); logLevel == "" {
		logLevel = "info"
	}

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	util.LogInit(logLevel)
	r := r.Router()
	r.Run(":" + port)
}
