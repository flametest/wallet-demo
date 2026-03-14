package dto

type WalletTransferReq struct {
	FromDisplayId string `json:"from_display_id" validate:"required"`
	ToDisplayId   string `json:"to_display_id" validate:"required"`
	Amount        string `json:"amount" validate:"required"`
}
