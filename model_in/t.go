package model_in

type T struct {
	CustomerID string `json:"customer_id" binding:"required" validate:"required,min=9,max=9,numeric"`
}
