package dto

import "github.com/shopspring/decimal"

type WalletDto struct {
	Name      string          `json:"name"`
	DisplayId string          `json:"display_id"`
	Balance   decimal.Decimal `json:"balance"`
}
