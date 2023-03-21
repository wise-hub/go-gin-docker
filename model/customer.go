package model

type Customer struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	EGN        string `json:"egn"`
}
