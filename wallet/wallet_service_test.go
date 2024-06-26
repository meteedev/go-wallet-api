package wallet_test

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
    "github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

// MockWalletStore is a mock implementation of postgres.Storer for testing
type MockWalletStore struct {
    mock.Mock
}

func (m *MockWalletStore) FindAll() ([]postgres.Wallet, error) {
    args := m.Called()
    return args.Get(0).([]postgres.Wallet), args.Error(1)
}

func (m *MockWalletStore) FindByWalletType(walletType string) ([]postgres.Wallet, error) {
    args := m.Called(walletType)
    return args.Get(0).([]postgres.Wallet), args.Error(1)
}

func (m *MockWalletStore) FindByWalletId(walletId int) (*postgres.Wallet, error) {
    args := m.Called(walletId)
    return args.Get(0).(*postgres.Wallet), args.Error(1)
}

func (m *MockWalletStore) FindByUserId(userId int) ([]postgres.Wallet, error) {
    args := m.Called(userId)
    return args.Get(0).([]postgres.Wallet), args.Error(1)
}

func (m *MockWalletStore) Create(wallet *postgres.Wallet) (*postgres.Wallet, error) {
    args := m.Called(wallet)
    return args.Get(0).(*postgres.Wallet), args.Error(1)
}

func (m *MockWalletStore) CountByCriteria(wallet postgres.Wallet) (int, error) {
    args := m.Called(wallet)
    return args.Int(0), args.Error(1)
}

func (m *MockWalletStore) DeleteByUserId(userId string) (int64, error) {
    args := m.Called(userId)
    return args.Get(0).(int64), args.Error(1)
}

func (m *MockWalletStore) UpdateByWalletId(walletId int, wallet postgres.Wallet) (int64, error) {
    args := m.Called(walletId, wallet)
    return args.Get(0).(int64), args.Error(1)
}

func TestGetAllWallets(t *testing.T) {
    // Define test data
    storeWallet := []postgres.Wallet{
        {ID: 1, UserID: 123, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 100.00},
        {ID: 2, UserID: 456, UserName: "user2", WalletName: "wallet2", WalletType: "type2", Balance: 200.00},
    }

    testWallets := []wallet.Wallet{
        {ID: 1, UserID: 123, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 100.00},
        {ID: 2, UserID: 456, UserName: "user2", WalletName: "wallet2", WalletType: "type2", Balance: 200.00},
    }

    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("FindAll").Return(storeWallet, nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    wallets, err := walletService.GetAllWallets()

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, testWallets, wallets)
    mockStore.AssertExpectations(t)
}

func TestGetWalletsByWalletType(t *testing.T) {
    // Define test data
    walletType := "type1"
    
    storeWallets := []postgres.Wallet{
        {ID: 1, UserID: 123, UserName: "user1", WalletName: "wallet1", WalletType: walletType, Balance: 100.00},
    }

    testWallets := []wallet.Wallet{
        {ID: 1, UserID: 123, UserName: "user1", WalletName: "wallet1", WalletType: walletType, Balance: 100.00},
    }



    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("FindByWalletType", walletType).Return(storeWallets, nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    wallets, err := walletService.GetWalletsByWalletType(walletType)

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, testWallets, wallets)
    mockStore.AssertExpectations(t)
}

func TestGetWalletsByUserId(t *testing.T) {
    // Define test data
    userID := 123

    storeWallets := []postgres.Wallet{
        {ID: 1, UserID: userID, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 100.00},
    }

    testWallets := []wallet.Wallet{
        {ID: 1, UserID: userID, UserName: "user1", WalletName: "wallet1", WalletType: "type1", Balance: 100.00},
    }

    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("FindByUserId", userID).Return(storeWallets, nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    wallets, err := walletService.GetWalletsByUserId(userID)

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, testWallets, wallets)
    mockStore.AssertExpectations(t)
}

func TestCreateWallet(t *testing.T) {
    // Define test data
    request := &wallet.WalletRequest{
        UserID:     123,
        UserName:   "user1",
        WalletName: "wallet1",
        WalletType: "Savings",
        Balance:    600.00,
    }
    
    createWallet := &postgres.Wallet{
        ID:         1,
        UserID:     request.UserID,
        UserName:   request.UserName,
        WalletName: request.WalletName,
        WalletType: request.WalletType,
        Balance:    request.Balance,
    }

    testWallet := &wallet.Wallet{
        ID:         1,
        UserID:     request.UserID,
        UserName:   request.UserName,
        WalletName: request.WalletName,
        WalletType: request.WalletType,
        Balance:    request.Balance,
    }

    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("Create", mock.AnythingOfType("*postgres.Wallet")).Return(createWallet, nil)
    mockStore.On("CountByCriteria", mock.AnythingOfType("postgres.Wallet")).Return(0, nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    walletResponse, err := walletService.CreateWallet(request)

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, testWallet, walletResponse)
    mockStore.AssertExpectations(t)
}


func TestDeleteWalletByUserId(t *testing.T) {
    // Define test data
    userID := "123"

    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("DeleteByUserId", userID).Return(int64(1), nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    deletedRows, err := walletService.DeleteWalletByUserId(userID)

    // Assert the result
    assert.NoError(t, err)
    assert.Equal(t, int64(1), deletedRows)
    mockStore.AssertExpectations(t)
}

func TestUpdateWalletByWalletId(t *testing.T) {
    // Define test data
    walletID := 1
    request := &wallet.WalletRequest{
        UserID:     123,
        UserName:   "user1",
        WalletName: "updated_wallet1",
        WalletType: "Savings",
        Balance:    650.00,
    }



    testWallet := &postgres.Wallet{
        ID:         walletID,
        UserID:     123,
        UserName:   "user1",
        WalletName: "updated_wallet1",
        WalletType: "Savings",
        Balance:    650.00,
    }

    // Create a mock instance
    mockStore := new(MockWalletStore)
    mockStore.On("UpdateByWalletId", walletID, mock.AnythingOfType("postgres.Wallet")).Return(int64(1), nil)
    mockStore.On("FindByWalletId", walletID).Return(testWallet, nil)

    // Create WalletService with mock store
    walletService := wallet.WalletService{WalletStore: mockStore}

    // Call the function under test
    updatedWallet, err := walletService.UpdateWalletByWalletId(walletID, request)

    // Assert the result
    assert.NoError(t, err)
    assert.NotNil(t, updatedWallet)
    assert.Equal(t, request.UserID, updatedWallet.UserID)
    assert.Equal(t, request.UserName, updatedWallet.UserName)
    assert.Equal(t, request.WalletName, updatedWallet.WalletName)
    assert.Equal(t, request.WalletType, updatedWallet.WalletType)
    assert.Equal(t, request.Balance, updatedWallet.Balance)
    mockStore.AssertExpectations(t)
}
