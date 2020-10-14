package presenter

import (
	"SkeletonAPI/modules/customer/model"
	"SkeletonAPI/modules/customer/usecase"
	useCaseMock "SkeletonAPI/modules/customer/usecase/mocks"
	"database/sql"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPMessageHandler_Mount(t *testing.T) {
	e := echo.New()
	n := NewHTTPHandler(new(useCaseMock.CustomerUseCase))
	n.Mount(e.Group("test"))
}

func TestHTTPMessageHandler_CheckSaldo(t *testing.T) {
	t.Run("Test Success Handler CheckSaldo", func(t *testing.T) {
		result := usecase.ResultUseCase{}
		result.Result = model.CustomerAccount{}
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("CheckSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:account_number")
		c.SetParamNames("account_number")
		c.SetParamValues("555001")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.CheckSaldo(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Test Failed Handler CheckSaldo param", func(t *testing.T) {
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:account_number")
		c.SetParamNames("account_number")
		c.SetParamValues("555001aaa")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.CheckSaldo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Test Failed Handler CheckSaldo ErrNoRows", func(t *testing.T) {
		result := usecase.ResultUseCase{}
		result.Error = sql.ErrNoRows
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("CheckSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:account_number")
		c.SetParamNames("account_number")
		c.SetParamValues("555001")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.CheckSaldo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestHTTPMessageHandler_TransferSaldo(t *testing.T) {
	t.Run("Test Success Handler TransferSaldo", func(t *testing.T) {
		body := `{"to_account_number":555002,"amount":100}`
		result := usecase.ResultUseCase{}
		result.Result = model.CustomerAccount{}
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("TransferSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:from_account_number/transfer")
		c.SetParamNames("from_account_number")
		c.SetParamValues("555001")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.TransferSaldo(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Test Failed Handler TransferSaldo body param", func(t *testing.T) {
		body := `{"to_account_number":"555002","amount":100}`
		result := usecase.ResultUseCase{}
		result.Result = model.CustomerAccount{}
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("TransferSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:from_account_number/transfer")
		c.SetParamNames("from_account_number")
		c.SetParamValues("555001")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.TransferSaldo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Test Failed Handler TransferSaldo param", func(t *testing.T) {
		body := `{"to_account_number":555002,"amount":100}`
		result := usecase.ResultUseCase{}
		result.Result = model.CustomerAccount{}
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("TransferSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:from_account_number/transfer")
		c.SetParamNames("from_account_number")
		c.SetParamValues("555001aa")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.TransferSaldo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Test Failed Handler TransferSaldo ErrNoRows", func(t *testing.T) {
		body := `{"to_account_number":555002,"amount":100}`
		result := usecase.ResultUseCase{}
		result.Error = sql.ErrNoRows
		mockCustomerUseCase := new(useCaseMock.CustomerUseCase)
		mockCustomerUseCase.On("TransferSaldo", mock.Anything).Return(result)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/account/:from_account_number/transfer")
		c.SetParamNames("from_account_number")
		c.SetParamValues("555001")

		handler := NewHTTPHandler(mockCustomerUseCase)
		handler.TransferSaldo(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
