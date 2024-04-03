package main

import (
	"context"
	"os"
	"net/http"
	"os/signal"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Wallet API
//	@version		1.0
//	@description	Sophisticated Wallet API
//	@host			localhost:1323
func main() {

	//init database connection
	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	//add database to service
	walletService := wallet.NewService(p)

	//add service to handler
	handler := wallet.NewHandler(walletService)
	
	e := echo.New()

	//set up error
	e.Use(apperrs.CustomErrorMiddleware)
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	
	e.GET("/api/v1/wallets", handler.WalletHandler)
	e.POST("/api/v1/wallets", handler.CreateWalletHandler)
	e.PUT("/api/v1/wallets/:id",handler.UpdateWalletHandler)

	e.GET("/api/v1/users/:id/wallets", handler.WalletByUserIdHandler)
	e.POST("/api/v1/users/:id/wallets", handler.DeleteWalletHandler)
	
	//e.Logger.Fatal(e.Start(":1323"))

	//graceful shutdown
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
