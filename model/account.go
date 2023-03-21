package model

type Account struct {
	IBAN    string  `json:"iban"`
	Balance float32 `json:"balance"`
}
