package wallet

import (
	"errors"
	"fmt"
	"strings"
)

const (
	minUserNameLength   = 3
	maxUserNameLength   = 255
	minWalletNameLength = 3
	maxWalletNameLength = 255
	minBalance          = 500
	maxBalance          = 10000
)

var validWalletTypes = []string{"Savings", "Credit Card", "Crypto Wallet"}


func ValidateWalletRequest(wallet *WalletRequest) error {
	var errMsgs []string

	if wallet.UserID <= 0 {
		errMsgs = append(errMsgs, "UserID must be greater than 0")
	}

	if len(wallet.UserName) < minUserNameLength || len(wallet.UserName) > maxUserNameLength {
		errMsgs = append(errMsgs, fmt.Sprintf("UserName must be between %d and %d characters", minUserNameLength, maxUserNameLength))
	}

	if len(wallet.WalletName) < minWalletNameLength || len(wallet.WalletName) > maxWalletNameLength {
		errMsgs = append(errMsgs, fmt.Sprintf("WalletName must be between %d and %d characters", minWalletNameLength, maxWalletNameLength))
	}

	if !contains(validWalletTypes, wallet.WalletType) {
		errMsgs = append(errMsgs, fmt.Sprintf("WalletType must be one of: %s", strings.Join(validWalletTypes, ", ")))
	}

	if wallet.Balance <= 0 || wallet.Balance < minBalance || wallet.Balance > maxBalance {
		errMsgs = append(errMsgs, fmt.Sprintf("Balance must be greater than 0 and between %d and %d", minBalance, maxBalance))
	}

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, "; "))
	}


	return nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}