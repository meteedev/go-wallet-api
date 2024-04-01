package wallet

import (
	"time"
)

type Wallet struct {
	ID         int       `json:"id" example:"1"`
	UserID     int       `json:"user_id" example:"1"`
	UserName   string    `json:"user_name" example:"John Doe"`
	WalletName string    `json:"wallet_name" example:"John's Wallet"`
	WalletType string    `json:"wallet_type" example:"Create Card"`
	Balance    float64   `json:"balance" example:"100.00"`
	CreatedAt  time.Time `json:"created_at" example:"2024-03-25T14:19:00.729237Z"`
}

type WalletRequest struct {
	UserID     int     `json:"user_id" example:"1" validate:"required"`
	UserName   string  `json:"user_name" example:"John Doe" validate:"required,min=3,max=255"`
	WalletName string  `json:"wallet_name" example:"John's Wallet" validate:"required,min=3,max=255"`
	WalletType string  `json:"wallet_type" example:"Credit Card" validate:"required,oneof='Savings' 'Credit Card' 'Crypto Wallet'"`
	Balance    float64 `json:"balance" example:"100.00" validate:"required,gt=0,min=500,max=10000"`
}


type Service interface {
	GetAllWallets() ([]Wallet, error)
	
	GetWalletsByWalletType(walletType string) ([]Wallet, error)
	
	GetWalletsByUserId(userId int) ([]Wallet, error)
	
	CreateWallet(wallet *WalletRequest)(*Wallet,error)
	
	DeleteWalletByUserId(userId string)(int64,error)
	
	UpdateWalletByWalletId(walletId int,request *WalletRequest) (*Wallet,error)
}

