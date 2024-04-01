package wallet

import (
	"net/http"
	"strconv"

	"github.com/KKGo-Software-engineering/fun-exercise-api/apperrs"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {

	walletType := c.QueryParam("wallet_type")

	var wallets []Wallet
	var err error

	if walletType != "" {
		wallets, err = h.service.GetWalletsByWalletType(walletType)
	} else {
		wallets, err = h.service.GetAllWallets()
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, wallets)

}



// UserHandler
//	@Summary		Get user wallets
//	@Description	Get user wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/{id}/wallets [get]
//	@Param			id	path	string	true	"user id"
//	@Failure		500	{object}	apperrs.CustomError
//	@Failure		400	{object}	apperrs.CustomError
func (h *Handler) WalletByUserIdHandler(c echo.Context) error {
	userId , err := strconv.Atoi(c.Param("id"))

	if err != nil {
        return apperrs.NewBadRequestError("invalid user ID")
    }

	wallets, err := h.service.GetWalletsByUserId(userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, wallets)

}



// CreateWallet
// @Summary Create user wallets
// @Description Create user wallets
// @Tags wallet
// @Accept json
// @Produce json
// @Param WalletCreateRequest body WalletRequest true "WalletRequest"
// @Success 201 {object} Wallet
// @Router /api/v1/wallets/ [post]
// @Failure 500 {object} apperrs.CustomError
// @Failure 400 {object} apperrs.CustomError
func (h *Handler) CreateWalletHandler(c echo.Context) error {

	req := new(WalletRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	wallet, err := h.service.CreateWallet(req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, wallet)
}


// DeleteWallet
// @Summary Delete user wallets
// @Description Delete user wallets by user id
// @Tags wallet
// @Accept json
// @Produce json
// @Router	/api/v1/users/{id}/wallets [delete]
// @Param	id	path	string	true	"user id"
// @Success 200 {object} Wallet
// @Failure 500 {object} apperrs.CustomError
// @Failure 400 {object} apperrs.CustomError
func (h *Handler) DeleteWalletHandler(c echo.Context) error {

	userId := c.Param("id")

	if len(userId) == 0 {
		return apperrs.NewBadRequestError("user id required")
	}

	rowAffected, err := h.service.DeleteWalletByUserId(userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, rowAffected)

}


// UpdateWallet
// @Summary Update user wallets
// @Description Update user wallets by wallet id
// @Tags wallet
// @Accept json
// @Produce json
// @Router /api/v1/wallets/{id} [put]
// @Param	id	path	string	true	"wallet id"
// @Param WalletCreateRequest body WalletRequest true "WalletRequest"
// @Success 200 {object} Wallet
// @Failure 500 {object} apperrs.CustomError
// @Failure 400 {object} apperrs.CustomError
func (h *Handler) UpdateWalletHandler(c echo.Context) error {

	walletId , err := strconv.Atoi(c.Param("id"))

	if err != nil {
        return apperrs.NewBadRequestError("invalid wallet ID")
    }


	req := new(WalletRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	walletResponse, err := h.service.UpdateWalletByWalletId(walletId, req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, walletResponse)

}
