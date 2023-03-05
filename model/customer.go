package model

type Customer struct {
	CustomerID string `json:"customerId"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	EGN        string `json:"egn"`
}
