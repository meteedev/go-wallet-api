package main

import (
	
	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	appvalidate "github.com/KKGo-Software-engineering/fun-exercise-api/validate"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Wallet API
//	@version		1.0
//	@description	Sophisticated Wallet API
//	@host			localhost:1323
func main() {

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	walletService := wallet.NewService(p)
	handler := wallet.NewHandler(walletService)
	
	e := echo.New()

	e.Use(apperrs.CustomErrorMiddleware)
	
	e.Validator = &appvalidate.CustomValidator{Validator: validator.New()}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	
	e.GET("/api/v1/wallets", handler.WalletHandler)
	e.POST("/api/v1/wallets", handler.CreateWalletHandler)
	e.PUT("/api/v1/wallets/:id",handler.UpdateWalletHandler)

	e.GET("/api/v1/users/:id/wallets", handler.WalletByUserIdHandler)
	e.POST("/api/v1/users/:id/wallets", handler.DeleteWalletHandler)
	
	e.Logger.Fatal(e.Start(":1323"))
}
