package api

import (
	"net/http"

	"github.com/flametest/vita/vserver"
	"github.com/flametest/wallet-demo/internal/api/handler"
	"github.com/flametest/wallet-demo/internal/container"
	"github.com/flametest/wallet-demo/internal/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	container.Container
}

func NewApp(c container.Container) *App {
	return &App{c}
}

func (a *App) Router(server vserver.Server) vserver.Server {
	srv := server.(*vserver.EchoServer)
	e := srv.GetEchoServer()
	walletService := service.NewWalletService(a.Container)
	walletHandler := handler.NewWalletHandler(walletService)
	e.Add("GET", "/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "up")
	})
	e.Add("POST", "/wallets", walletHandler.CreateWallet)
	e.Add("GET", "/wallets/:displayId", walletHandler.GetWalletDetail)
	e.Add("POST", "/wallets/transfer", walletHandler.TransferFund)
	return srv
}
