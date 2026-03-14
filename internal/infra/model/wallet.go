package model

import "github.com/shopspring/decimal"

type Wallet struct {
	Base
	Name      string          `gorm:"uniqueIndex;column:name" json:"name"`
	DisplayId string          `gorm:"uniqueIndex;column:display_id" json:"display_id"`
	Balance   decimal.Decimal `gorm:"column:balance" json:"balance"`
}
