package main

import (
	"SkeletonAPI/handler"
	"fmt"
	configEnv "github.com/joho/godotenv"
	"os"
)

func main()  {
	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	// Init dependencies
	service := handler.MakeHandler()

	// start echo server
	service.StartServer()

	// Shutdown with gracefull handler
	service.ShutdownServer()
}
