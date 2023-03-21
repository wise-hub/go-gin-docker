package model_in

type InLogin struct {
	Username string `json:"username" binding:"required" validate:"required,min=3,max=20,alphanum"`
	Password string `json:"password" binding:"required" validate:"required,min=3,max=50"`
}
