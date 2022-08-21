package main

import (
	"context"
	output "github.com/rudSarkar/PortMonitor/helper/api"
	config "github.com/rudSarkar/PortMonitor/helper/database"
	"github.com/thep0y/go-logger/log"
	"os"
)

func main() {
	URI := os.Getenv("MONGODB_URI")
	if checkDotEnv() {
		log.Fatal("environment file not found")
	} else {
		defer config.IsConnected(URI).Disconnect(context.TODO())
		output.GetOutput()
	}

}
