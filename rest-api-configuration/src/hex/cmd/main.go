package main

import (
	"log"
	"os"
	"restapi/v2/internal/adapters/app/api"
	"restapi/v2/internal/adapters/core/arithmetic"
	"restapi/v2/internal/adapters/framework/right/db"
	"restapi/v2/internal/ports"

	gRPC "restapi/v2/internal/adapters/framework/left/grpc"
)

func main() {

	var err error

	// ports

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)

	if err != nil {
		log.Fatalf("failed to initialize dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
