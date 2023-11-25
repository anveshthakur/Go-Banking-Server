package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	logger.Info("Starting the application on 8000")
	app.Start()
}
