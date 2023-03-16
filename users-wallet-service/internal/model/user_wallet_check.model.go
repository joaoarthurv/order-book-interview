package model

type UserWalletCheck struct {
	UserId          string  `json:"userId,omitempty"`
	ProductType     string  `json:"productType,omitempty"`
	Price           float64 `json:"price,omitempty"`
	ProductQuantity int     `json:"productQuantity,omitempty"`
}

type UserWalletCheckResponse struct {
	CheckValue bool `json:"checkValue"`
}
