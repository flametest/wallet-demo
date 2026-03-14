package dto

import "github.com/shopspring/decimal"

type WalletTransferRequest struct {
	FromDisplayId string          `json:"from_display_id" validate:"required"`
	ToDisplayId   string          `json:"to_display_id" validate:"required"`
	Amount        decimal.Decimal `json:"amount" validate:"required,gte=0"`
}
