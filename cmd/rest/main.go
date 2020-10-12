package main

import (
	"SkeletonAPI/handler"
	"context"
	"fmt"
	configEnv "github.com/joho/godotenv"
	"os"
	"os/signal"
	"time"
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
	server := service.StartServer()

	// Shutdown with gracefull handler
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
