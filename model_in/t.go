package model_in

type T struct {
	CustomerID string `json:"customer_id" binding:"required" validate:"required,min=9,max=9,numeric"`
}

type F struct {
	FeedbackID string `json:"feedback_id" binding:"required" validate:"required,min=1,max=999999999,numeric"`
}
