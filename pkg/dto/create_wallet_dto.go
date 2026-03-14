package dto

type CreateWalletReq struct {
	Name string `json:"name" validate:"required,max=64"`
}
