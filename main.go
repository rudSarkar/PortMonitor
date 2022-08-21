package main

import (
	"context"
	output "github.com/rudSarkar/PortMonitor/helper/api"
	config "github.com/rudSarkar/PortMonitor/helper/database"
	"github.com/thep0y/go-logger/log"
	"os"
)

func main() {
	if checkDotEnv() {
		log.Fatal("environment file not found")
	} else {
		URI := os.Getenv("MONGODB_URI")
		defer config.IsConnected(URI).Disconnect(context.TODO())
		output.GetOutput()
	}

}
