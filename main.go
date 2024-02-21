package main

import (
	"banking/app"
	"banking/logger"
)

func main() {

	logger.Info("Starting our application...")
	app.Start()

}
