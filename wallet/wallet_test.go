package wallet

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
)

type StubService struct {
	Wallets        []Wallet
	Wallet        	Wallet
	Err            error
	DeletedRow     int64
	DuplicateCount int
}

// GetAllWallets mocks the GetAllWallets method.
func (s StubService) GetAllWallets() ([]Wallet, error) {
	return s.Wallets, s.Err
}

// GetWalletsByWalletType mocks the GetWalletsByWalletType method.
func (s StubService) GetWalletsByWalletType(walletType string) ([]Wallet, error) {
	return s.Wallets, s.Err
}

// GetWalletsByUserId mocks the GetWalletsByUserId method.
func (s StubService) GetWalletsByUserId(userId int) ([]Wallet, error) {
	return s.Wallets, s.Err
}

// CreateWallet mocks the CreateWallet method.
func (s StubService) CreateWallet(wallet *WalletRequest) (*Wallet, error) {
	return &s.Wallet, s.Err
}

// DeleteWalletByUserId mocks the DeleteWalletByUserId method.
func (s StubService) DeleteWalletByUserId(userId string) (int64, error) {
	return s.DeletedRow, s.Err
}

// GetCountWalletsByCriteria mocks the GetCountWalletsByCriteria method.
func (s StubService) GetCountWalletsByCriteria(criteria *Wallet) (int, error) {
	return s.DuplicateCount, s.Err
}

// UpdateWalletByWalletId mocks the UpdateWalletByWalletId method.
func (s StubService) UpdateWalletByWalletId(walletId int, request *WalletRequest) (*Wallet, error) {
	return &s.Wallet, s.Err
}

func TestWallet(t *testing.T) {
    t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
        // Setup
        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/", nil)
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
        rec := httptest.NewRecorder()

        c := e.NewContext(req, rec)
        c.SetPath("/api/v1/wallets")

        // Register the CustomErrorMiddleware with the Echo instance
        e.Use(apperrs.CustomErrorMiddleware)

        // Create the handler with a stub that always returns an internal server error
        stubService := StubService{Err: apperrs.NewInternalServerError("Internal Server Error")}
        h := NewHandler(&stubService)

        // Wrap the handler with the middleware
        handlerWithMiddleware := apperrs.CustomErrorMiddleware(h.WalletHandler)

        // Act: Execute the handler
        handlerWithMiddleware(c)

        // Assert: Check the response status code and body
        assert.Equal(t, http.StatusInternalServerError, rec.Code, "status code should be 500")
        assert.Contains(t, rec.Body.String(), "Internal Server Error", "Message should contain Internal Server Error")
    })

	t.Run("given user able to get wallets should return list of wallets", func(t *testing.T) {
        // Setup
        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/", nil)
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
        rec := httptest.NewRecorder()

        c := e.NewContext(req, rec)
        c.SetPath("/api/v1/wallets")

        // Define the expected wallets
        expectedWallets := []Wallet{
            {ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Credit Card", Balance: 100.00},
            {ID: 2, UserID: 2, UserName: "Jane Doe", WalletName: "Jane's Wallet", WalletType: "Debit Card", Balance: 150.00},
        }

        // Create a stub service with the expected wallets
        stubService := StubService{Wallets: expectedWallets}

        // Create the handler with the stub service
        h := NewHandler(&stubService)

        // Act: Execute the handler
        err := h.WalletHandler(c)

        // Assert: Check the response status code and body
        assert.NoError(t, err, "no error should occur")
        assert.Equal(t, http.StatusOK, rec.Code, "status code should be 200")

        // Unmarshal the response body to check the returned wallets
        var got []Wallet
        if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
            t.Errorf("unable to unmarshal response JSON: %v", err)
        }

        // Compare each field of expected and received wallets
        for i, expected := range expectedWallets {
            assert.Equal(t, expected.ID, got[i].ID, "ID should match")
            assert.Equal(t, expected.UserID, got[i].UserID, "UserID should match")
            assert.Equal(t, expected.UserName, got[i].UserName, "UserName should match")
            assert.Equal(t, expected.WalletName, got[i].WalletName, "WalletName should match")
            assert.Equal(t, expected.WalletType, got[i].WalletType, "WalletType should match")
            assert.InDelta(t, expected.Balance, got[i].Balance, 0.01, "Balance should match")
        }
    })


}


