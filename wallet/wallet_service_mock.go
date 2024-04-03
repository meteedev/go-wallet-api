package wallet

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of the Service interface
type MockService struct {
	mock.Mock
}

func (m *MockService) GetAllWallets() ([]Wallet, error) {
	args := m.Called()
	return args.Get(0).([]Wallet), args.Error(1)
}

func (m *MockService) GetWalletsByWalletType(walletType string) ([]Wallet, error) {
	args := m.Called(walletType)
	return args.Get(0).([]Wallet), args.Error(1)
}

func (m *MockService) GetWalletsByUserId(userId int) ([]Wallet, error) {
	args := m.Called(userId)
	return args.Get(0).([]Wallet), args.Error(1)
}

func (m *MockService) CreateWallet(wallet *WalletRequest) (*Wallet, error) {
	args := m.Called(wallet)
	return args.Get(0).(*Wallet), args.Error(1)
}

func (m *MockService) DeleteWalletByUserId(userId string) (int64, error) {
	args := m.Called(userId)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockService) UpdateWalletByWalletId(walletId int, request *WalletRequest) (*Wallet, error) {
	args := m.Called(walletId, request)
	return args.Get(0).(*Wallet), args.Error(1)
}



// Helper function to convert WalletRequest to Wallet
func toWallet(request *WalletRequest) *Wallet {
	return &Wallet{
		ID:         0,
		UserID:     request.UserID,
		UserName:   request.UserName,
		WalletName: request.WalletName,
		WalletType: request.WalletType,
		Balance:    0,
		CreatedAt:  time.Time{},
	}
}
