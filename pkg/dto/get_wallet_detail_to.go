package dto

type GetWalletDetailReq struct {
	DisplayID string `param:"display_id" validate:"required" json:"display_id"`
}
