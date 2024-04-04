package wallet

import (
	"strings"
	"testing"
)

func TestValidateWalletRequest(t *testing.T) {
    testCases := []struct {
        name      string
        wallet    *WalletRequest
        wantError bool
    }{
        {
            name: "Valid wallet request",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: false,
        },
        {
            name: "Invalid UserID (negative)",
            wallet: &WalletRequest{
                UserID:     -1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid UserName (too short)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JD",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid UserName (too long)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   strings.Repeat("a", maxUserNameLength+1),
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid WalletName (too short)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "W1",
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid WalletName (too long)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: strings.Repeat("a", maxUserNameLength+1),
                WalletType: "Savings",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid WalletType",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "InvalidType",
                Balance:    1000,
            },
            wantError: true,
        },
        {
            name: "Invalid Balance (negative)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    -100,
            },
            wantError: true,
        },
        {
            name: "Invalid Balance (below minimum)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    200,
            },
            wantError: true,
        },
        {
            name: "Invalid Balance (above maximum)",
            wallet: &WalletRequest{
                UserID:     1,
                UserName:   "JohnDoe",
                WalletName: "Savings",
                WalletType: "Savings",
                Balance:    11000,
            },
            wantError: true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            validateWalletRequestTest(t, tc.wallet, tc.wantError)
        })
    }
}

func validateWalletRequestTest(t *testing.T, wallet *WalletRequest, wantError bool) {
    err := ValidateWalletRequestCreate(wallet)
    if (err != nil) != wantError {
        t.Errorf("ValidateWalletRequest(%v) returned error: %v, wantError: %t", wallet, err, wantError)
    }
}
