package handler

import (
	"net/http"

	"github.com/flametest/vita/verrors"
	"github.com/flametest/wallet-demo/pkg/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *WalletHandler) TransferFund(c echo.Context) error {
	req := &dto.WalletTransferReq{}
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
	err = h.walletService.TransferFund(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Transfer Success")
}
