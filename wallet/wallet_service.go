package wallet

import (
	"log"

	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
)

type WalletService struct {
	WalletStore postgres.Storer
}


func NewService(db postgres.Storer) WalletService {
	return WalletService{WalletStore: db}
}

func (s WalletService) GetAllWallets() ([]Wallet, error){

	wallets, err := s.WalletStore.FindAll()

	if err != nil {
		return nil, apperrs.NewInternalServerError(err.Error())
	}

	if len(wallets)==0{
		return nil, apperrs.NewNotFoundError("wallet not found")
	}

	var walletResponses []Wallet
	for _, w := range wallets {
		walletResponses = append(walletResponses, Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}

	return walletResponses,nil

}


func (s WalletService) GetWalletsByWalletType(walletType string) ([]Wallet, error){

	wallets, err := s.WalletStore.FindByWalletType(walletType)

	if err != nil {
		return nil, apperrs.NewInternalServerError(err.Error())
	}

	if len(wallets)==0{
		return nil, apperrs.NewNotFoundError("wallet not found")
	}

	var walletResponses []Wallet
	for _, w := range wallets {
		walletResponses = append(walletResponses, Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}

	return walletResponses,nil

}


func (s WalletService) GetWalletsByUserId(userId int) ([]Wallet, error){

	wallets, err := s.WalletStore.FindByUserId(userId)

	if err != nil {
		return nil, apperrs.NewInternalServerError(err.Error())
	}

	if len(wallets)==0{
		return nil, apperrs.NewNotFoundError("wallet not found")
	}

	var walletResponses []Wallet
	for _, w := range wallets {
		walletResponses = append(walletResponses, Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}

	return walletResponses,nil

}


func (ws WalletService) CreateWallet(request *WalletRequest) (*Wallet,error){
	
	
	err := ValidateWalletRequest(request)

	if err != nil{
		log.Println(err)
		return nil,apperrs.NewBadRequestError(err.Error())
	}

	wallet := postgres.Wallet{
		UserID:     request.UserID,
		UserName:   request.UserName,
		WalletName: request.WalletName,
		WalletType: request.WalletType,
		Balance:    request.Balance,
	}

	isDuplicated , err := ws.checkDuplicated(wallet)

	if err != nil{
		log.Println(err)
		return nil,apperrs.NewInternalServerError(err.Error())
	}

	if isDuplicated {
		log.Printf("Duplicated wallet userid=%d userName=%s walletname=%s walletType=%s",wallet.UserID,wallet.UserName,wallet.WalletName,wallet.WalletType)
		return nil,apperrs.NewInternalServerError("Duplicated wallets")
	}

	
	w , err := ws.WalletStore.Create(&wallet)
	
	if err != nil{
		log.Println(err)
		return nil,apperrs.NewInternalServerError("Create wallet failed")
	}

	walletResponses := Wallet{
		ID:         w.ID,
		UserID:     w.UserID,
		UserName:   w.UserName,
		WalletName: w.WalletName,
		WalletType: w.WalletType,
		Balance:    w.Balance,
		CreatedAt:  w.CreatedAt,
	}


	return &walletResponses,nil
}

func (ws WalletService) checkDuplicated(wallet postgres.Wallet) (bool,error){
	rowCount , err := ws.WalletStore.CountByCriteria(wallet)
	
	if err!= nil{
		return false,err
	}

	isDup := rowCount > 0

	return isDup , nil
}



func (ws WalletService) DeleteWalletByUserId(userId string)(int64,error){
	
	deleteRow , err := ws.WalletStore.DeleteByUserId(userId)
	

	if err != nil{
		log.Println(err)
		return 0,apperrs.NewInternalServerError("Delete wallet failed")
	}
	
	if deleteRow == 0{
		log.Println("delete affected ",deleteRow)
		return 0,apperrs.NewUnprocessableEntity("Delete wallet failed")
	}

	return deleteRow,nil
}


func (ws WalletService) UpdateWalletByWalletId(walletId int,request *WalletRequest) (*Wallet,error){
	
	err := ValidateWalletRequest(request)

	if err != nil{
		log.Println(err)
		return nil,apperrs.NewBadRequestError(err.Error())
	}
	
	
	wallet := postgres.Wallet{
		UserID:     request.UserID,
		UserName:   request.UserName,
		WalletName: request.WalletName,
		WalletType: request.WalletType,
		Balance:    request.Balance,
	}

	updateRow , err := ws.WalletStore.UpdateByWalletId(walletId,wallet)
	

	if err != nil{
		log.Println(err)
		return nil,apperrs.NewInternalServerError("Update wallet failed")
	}
	
	if updateRow == 0{
		log.Println("update affected ",updateRow)
		return nil,apperrs.NewUnprocessableEntity("Delete wallet failed")
	}

	w , err := ws.WalletStore.FindByWalletId(walletId)

	if err != nil{
		log.Println(err)
		return nil,apperrs.NewInternalServerError("Update wallet failed")
	}

	walletResponses := Wallet{
		ID:         w.ID,
		UserID:     w.UserID,
		UserName:   w.UserName,
		WalletName: w.WalletName,
		WalletType: w.WalletType,
		Balance:    w.Balance,
		CreatedAt:  w.CreatedAt,
	}

	return &walletResponses,nil
}


