package model

type CustomerFeedback struct {
	CustomerID string `json:"customer_id"`
	Feedback   string `json:"feedback"`
	InsertDate string `json:"insert_date"`
	UserName   string `json:"user_name"`
}
