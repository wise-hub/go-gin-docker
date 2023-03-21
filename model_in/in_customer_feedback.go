package model_in

type InCustomerFeedback struct {
	CustomerID string `json:"customer_id" binding:"required" validate:"required,min=9,max=9,numeric"`
	Feedback   string `json:"feedback" binding:"required"`
}
