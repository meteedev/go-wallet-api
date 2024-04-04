package wallet

import (
	"errors"
	"fmt"
	"strings"
)

// Constants for validation
const (
	minUserNameLength   = 3
	maxUserNameLength   = 255
	minWalletNameLength = 3
	maxWalletNameLength = 255
	minBalance          = 500
	maxBalance          = 10000
)

// Valid wallet types
var validWalletTypes = []string{"Savings", "Credit Card", "Crypto Wallet"}


// ValidateWalletRequest validates a wallet request
func ValidateWalletRequestCreate(wallet *WalletRequest) error {
	var errMsgs []string

	validateUserID(wallet.UserID, &errMsgs)
	validateUserName(wallet.UserName, &errMsgs)
	validateWalletName(wallet.WalletName, &errMsgs)
	validateWalletType(wallet.WalletType, &errMsgs)
	validateBalanceGreaterThanZero(wallet.Balance, &errMsgs)
	validateBalanceRangeMinMax(wallet.Balance, &errMsgs)
	

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, "; "))
	}

	return nil
}


func ValidateWalletRequestUpdate(wallet *WalletRequest) error {
	var errMsgs []string

	validateUserID(wallet.UserID, &errMsgs)
	validateUserName(wallet.UserName, &errMsgs)
	validateWalletName(wallet.WalletName, &errMsgs)
	validateWalletType(wallet.WalletType, &errMsgs)
	validateBalanceGreaterEqualZero(wallet.Balance, &errMsgs)

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, "; "))
	}

	return nil
}



// Helper functions for individual validations

func validateUserID(userID int, errMsgs *[]string) {
	if userID <= 0 {
		*errMsgs = append(*errMsgs, "UserID must be greater than 0")
	}
}

func validateUserName(userName string, errMsgs *[]string) {
	if len(userName) < minUserNameLength || len(userName) > maxUserNameLength {
		*errMsgs = append(*errMsgs, fmt.Sprintf("UserName must be between %d and %d characters", minUserNameLength, maxUserNameLength))
	}
}

func validateWalletName(walletName string, errMsgs *[]string) {
	if len(walletName) < minWalletNameLength || len(walletName) > maxWalletNameLength {
		*errMsgs = append(*errMsgs, fmt.Sprintf("WalletName must be between %d and %d characters", minWalletNameLength, maxWalletNameLength))
	}
}

func validateWalletType(walletType string, errMsgs *[]string) {
	if !contains(validWalletTypes, walletType) {
		*errMsgs = append(*errMsgs, fmt.Sprintf("WalletType must be one of: %s", strings.Join(validWalletTypes, ", ")))
	}
}

func validateBalanceRangeMinMax(balance float64, errMsgs *[]string) {
	if  balance < minBalance || balance > maxBalance {
		*errMsgs = append(*errMsgs, fmt.Sprintf("Balance between %d and %d", minBalance, maxBalance))
	}
}


func validateBalanceGreaterThanZero(balance float64, errMsgs *[]string) {
	if balance <= 0  {
		*errMsgs = append(*errMsgs, "Balance must be greater than 0 ")
	}
}


func validateBalanceGreaterEqualZero(balance float64, errMsgs *[]string) {
	if balance < 0  {
		*errMsgs = append(*errMsgs, fmt.Sprintf("Balance must be equal or greater than 0 "))
	}
}



func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
