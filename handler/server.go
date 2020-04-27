package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http
func (s *Service) HTTPServerMain() *echo.Echo {
	e := echo.New()
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	group := e.Group("wp/api")
	s.MessageHandler.Mount(group)

	return e
}

func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", DefaultPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
