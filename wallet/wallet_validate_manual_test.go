package wallet_test

import (
	"errors"
    "testing"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/stretchr/testify/assert"
)

func TestValidateWalletRequest(t *testing.T) {
    // Test cases
    testCases := []struct {
        name     string
        wallet   *wallet.WalletRequest
        expected error
    }{
        // Valid wallet request
        {
            name: "Valid Wallet Request",
            wallet: &wallet.WalletRequest{
                UserID:     123,
                UserName:   "username",
                WalletName: "Wallet",
                WalletType: "Savings",
                Balance:    1000,
            },
            expected: nil,
        },
        // Invalid UserID
        {
            name: "Invalid UserID",
            wallet: &wallet.WalletRequest{
                UserID:     0,
                UserName:   "username",
                WalletName: "Wallet",
                WalletType: "Savings",
                Balance:    1000,
            },
            expected: errors.New("UserID must be greater than 0"),
        },
        // Invalid WalletName
        {
            name: "Invalid WalletName",
            wallet: &wallet.WalletRequest{
                UserID:     123,
                UserName:   "username",
                WalletName: "W",
                WalletType: "Savings",
                Balance:    1000,
            },
            expected: errors.New("WalletName must be between 3 and 255 characters"),
        },
        // Add more test cases as needed
    }

    // Iterate over test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Call the function under test
            err := wallet.ValidateWalletRequest(tc.wallet)

            // Assert the result
            assert.Equal(t, tc.expected, err)
        })
    }
}
