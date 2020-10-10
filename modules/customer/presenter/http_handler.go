package presenter

import (
	"SkeletonAPI/helper"
	"SkeletonAPI/modules/customer/model"
	"SkeletonAPI/modules/customer/usecase"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type HTTPMessageHandler struct {
	CustomerUseCase usecase.CustomerUseCase
}

func NewHTTPHandler(customer usecase.CustomerUseCase) *HTTPMessageHandler {
	return &HTTPMessageHandler{CustomerUseCase: customer}
}

func (h *HTTPMessageHandler) Mount(group *echo.Group) {
	group.POST("/account/:from_account_number/transfer", h.TransferSaldo)
	group.GET("/account/:account_number", h.CheckSaldo)
}

func (h *HTTPMessageHandler) TransferSaldo(c echo.Context) error {
	sourceAccount, err := strconv.Atoi(c.Param("from_account_number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	dataRequest := model.TransferRequest{}
	err = c.Bind(&dataRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	dataRequest.FromAccountNumber = sourceAccount
	result := h.CustomerUseCase.TransferSaldo(dataRequest)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(false, http.StatusBadRequest, result.Error.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseDetailOutput(true, http.StatusCreated, "", nil))
}

func (h *HTTPMessageHandler) CheckSaldo(c echo.Context) error {
	accountNumber, err := strconv.Atoi(c.Param("account_number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(false, http.StatusBadRequest, err.Error(), nil))
	}

	result := h.CustomerUseCase.CheckSaldo(accountNumber)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(false, http.StatusBadRequest, result.Error.Error(), nil))
	}

	customerAccount := result.Result.(model.CustomerAccount)

	return c.JSON(http.StatusOK, helper.ResponseDetailOutput(true, http.StatusOK, "", customerAccount))
}
