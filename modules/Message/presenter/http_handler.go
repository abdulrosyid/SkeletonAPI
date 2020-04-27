package presenter

import (
	"WarungPintarAPI/modules/Message/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type HTTPMessageHandler struct {
	MessageUseCase usecase.MessageUseCase
}

func NewHTTPHandler(message usecase.MessageUseCase) *HTTPMessageHandler {
	return &HTTPMessageHandler{MessageUseCase: message}
}

func (h *HTTPMessageHandler) Mount(group *echo.Group) {
	group.POST("/message/add", h.AddMessage)
	group.GET("/message/get", h.GetMessage)
}

func (h *HTTPMessageHandler) AddMessage(c echo.Context) error {
	message := c.QueryParam("message")
	result := h.MessageUseCase.AddMessage(message)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}

	return c.JSON(http.StatusOK, "add message success")
}

func (h *HTTPMessageHandler) GetMessage(c echo.Context) error {
	result := h.MessageUseCase.GetMessage()
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(http.StatusOK, result.Result)
}
