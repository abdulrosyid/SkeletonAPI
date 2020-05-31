package main

import "SkeletonAPI/handler"

func main()  {
	// Init dependencies
	service := handler.MakeHandler()

	// start echo server
	service.StartServer()

	// Shutdown with gracefull handler
	service.ShutdownServer()
}
