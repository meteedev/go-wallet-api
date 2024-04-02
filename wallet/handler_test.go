package wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func TestWalletHandler(t *testing.T) {
	mockService := new(MockService)
	handler := NewHandler(mockService)

	mockWallets := []Wallet{
		{ID: 1, UserID: 1, UserName: "User1", WalletName: "Wallet1", WalletType: "Type1", Balance: 100.0},
		{ID: 2, UserID: 2, UserName: "User2", WalletName: "Wallet2", WalletType: "Type2", Balance: 200.0},
	}

	mockService.On("GetAllWallets").Return(mockWallets, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.WalletHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseWallets []Wallet
		err := json.Unmarshal(rec.Body.Bytes(), &responseWallets)
		assert.NoError(t, err)

		assert.Equal(t, mockWallets, responseWallets)
	}

	mockService.AssertExpectations(t)
}

func TestWalletByUserIdHandler(t *testing.T) {
	mockService := new(MockService)
	handler := NewHandler(mockService)

	userID := "123"
	mockUserID, _ := strconv.Atoi(userID)

	mockWallets := []Wallet{
		{ID: 1, UserID: mockUserID, UserName: "User1", WalletName: "Wallet1", WalletType: "Type1", Balance: 100.0},
		{ID: 2, UserID: mockUserID, UserName: "User2", WalletName: "Wallet2", WalletType: "Type2", Balance: 200.0},
	}

	mockService.On("GetWalletsByUserId", mockUserID).Return(mockWallets, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+userID+"/wallets", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/:id/wallets")
	c.SetParamNames("id")
	c.SetParamValues(userID)

	if assert.NoError(t, handler.WalletByUserIdHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseWallets []Wallet
		err := json.Unmarshal(rec.Body.Bytes(), &responseWallets)
		assert.NoError(t, err)

		assert.Equal(t, mockWallets, responseWallets)
	}

	mockService.AssertExpectations(t)
}

func TestCreateWalletHandler(t *testing.T) {
	mockService := new(MockService)
	handler := NewHandler(mockService)

	reqBody := WalletRequest{
		UserID:     1,
		UserName:   "User1",
		WalletName: "Wallet1",
		WalletType: "Type1",
		Balance:    100.0,
	}

	mockWallet := Wallet{
		ID:         1,
		UserID:     reqBody.UserID,
		UserName:   reqBody.UserName,
		WalletName: reqBody.WalletName,
		WalletType: reqBody.WalletType,
		Balance:    reqBody.Balance,
	}

	mockService.On("CreateWallet", &reqBody).Return(&mockWallet, nil)

	e := echo.New()
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewReader(reqBodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CreateWalletHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var responseWallet Wallet
		err := json.Unmarshal(rec.Body.Bytes(), &responseWallet)
		assert.NoError(t, err)

		assert.Equal(t, mockWallet, responseWallet)
	}

	mockService.AssertExpectations(t)
}

func TestDeleteWalletHandler(t *testing.T) {
	mockService := new(MockService)
	handler := NewHandler(mockService)

	userID := "123"

	mockService.On("DeleteWalletByUserId", userID).Return(int64(1), nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+userID+"/wallets", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/users/:id/wallets")
	c.SetParamNames("id")
	c.SetParamValues(userID)

	if assert.NoError(t, handler.DeleteWalletHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response int64
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, int64(1), response)
	}

	mockService.AssertExpectations(t)
}

func TestUpdateWalletHandler(t *testing.T) {
	mockService := new(MockService)
	handler := NewHandler(mockService)

	walletID := "123"
	reqBody := WalletRequest{
		UserID:     1,
		UserName:   "User1",
		WalletName: "Wallet1",
		WalletType: "Type1",
		Balance:    100.0,
	}

	mockWallet := Wallet{
		ID:         123,
		UserID:     reqBody.UserID,
		UserName:   reqBody.UserName,
		WalletName: reqBody.WalletName,
		WalletType: reqBody.WalletType,
		Balance:    reqBody.Balance,
	}

	mockService.On("UpdateWalletByWalletId", 123, &reqBody).Return(&mockWallet, nil)

	e := echo.New()
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/wallets/"+walletID, bytes.NewReader(reqBodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/wallets/:id")
	c.SetParamNames("id")
	c.SetParamValues(walletID)

	if assert.NoError(t, handler.UpdateWalletHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseWallet Wallet
		err := json.Unmarshal(rec.Body.Bytes(), &responseWallet)
		assert.NoError(t, err)

		assert.Equal(t, mockWallet, responseWallet)
	}

	mockService.AssertExpectations(t)
}
