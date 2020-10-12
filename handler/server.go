package handler

import (
	"fmt"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

const DefaultPort = 8080

func (s *Service) StartServer() *echo.Echo {
	e := echo.New()
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	group := e.Group("api")
	s.CustomerHandler.Mount(group)

	listenerPort := fmt.Sprintf(":%v", DefaultPort)
	if err := e.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}); err != nil {
		e.Logger.Fatal(err.Error())
	}

	return e
}
