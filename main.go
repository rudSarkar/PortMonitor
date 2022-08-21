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
	if URI == "" {
		log.Fatal("MONGODB_URI not found in the .env")
	} else {
		defer config.IsConnected(URI).Disconnect(context.TODO())
		output.GetOutput()
	}

}
