package main

import (
	"context"
	output "github.com/rudSarkar/PortMonitor/helper/api"
	config "github.com/rudSarkar/PortMonitor/helper/database"
	"os"
)

func main() {
	URI := os.Getenv("MONGODB_URI")
	defer config.IsConnected(URI).Disconnect(context.TODO())

	output.GetOutput()
}
