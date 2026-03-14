package handler

import (
	"net/http"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/wallet-demo/pkg/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *WalletHandler) GetWalletDetail(c echo.Context) error {
	req := &dto.GetWalletDetailReq{}
	binder := echo.DefaultBinder{}
	err := binder.BindPathParams(c, req)
	if err != nil {
		return verrors.BadRequestError(err.Error())
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return verrors.BadRequestError(err.Error())
	}
	wallet, err := h.walletService.GetByDisplayId(c.Request().Context(), req.DisplayID)
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
