package wallet

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandler_CreateWalletHandler(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CreateWalletHandler(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateWalletHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
