package wallet

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
}
