package handler

import (
	"net/http"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/wallet-demo/pkg/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *WalletHandler) CreateWallet(c echo.Context) error {
	req := &dto.CreateWalletReq{}
	binder := echo.DefaultBinder{}
	err := binder.BindBody(c, req)
	if err != nil {
		return verrors.BadRequestError(err.Error())
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return verrors.BadRequestError(err.Error())
	}

	wallet, err := h.walletService.CreateWallet(c.Request().Context(), req)
	if err != nil {
		return err
	}
	walletDto := &dto.WalletDto{
		Name:      wallet.Name,
		Balance:   wallet.Balance,
		DisplayId: wallet.DisplayId,
	}
	return c.JSON(http.StatusOK, walletDto)
}
