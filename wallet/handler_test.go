package wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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

        // Ensure that the response body is valid JSON
        var response map[string]interface{}
        err := json.Unmarshal(rec.Body.Bytes(), &response)
        assert.NoError(t, err)

        // Check the response fields
        assert.Equal(t, float64(200), response["code"])
        assert.Equal(t, "Delete Success", response["message"])
    }

    mockService.AssertExpectations(t)
}

func TestCreateWalletHandler(t *testing.T) {
	
	t.Run("given walletRequest to create wallet should return 201 and wallet struct", func(t *testing.T) {
	
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

	})


	t.Run("given dup walletRequest to create wallet should return 500 and message duplicated", func(t *testing.T) {
        // Set up your handler, mock service, and any necessary middleware
        mockService := new(MockService)
        handler := NewHandler(mockService)
        handlerWithMiddleware := apperrs.CustomErrorMiddleware(handler.CreateWalletHandler)

        // Create a sample wallet request
        reqBody := WalletRequest{
            UserID:     1,
            UserName:   "User1",
            WalletName: "Wallet1",
            WalletType: "Type1",
            Balance:    100.0,
        }

        // Configure the mock service to return nil and an error indicating duplication
        errorMessage := "Duplicated wallets"
        mockService.On("CreateWallet", &reqBody).Return(&Wallet{}, apperrs.NewInternalServerError(errorMessage))

        // Prepare the HTTP request
		e := echo.New()
        reqBodyBytes, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewReader(reqBodyBytes))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)


        // Invoke the handler function and assert the response
        assert.NoError(t, handlerWithMiddleware(c)) // Ensure that handler returns an error
        assert.Equal(t, http.StatusInternalServerError, rec.Code)

        // Optionally, you can assert the response body to ensure it contains the error message
        assert.Contains(t, rec.Body.String(), errorMessage)

        // Assert that the CreateWallet method was called with the correct parameters
        mockService.AssertExpectations(t)
    })


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
