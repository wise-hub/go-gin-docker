package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateMiddleware(c *gin.Context, customType interface{}) error {

	// validate keys (fields)
	if err := c.ShouldBind(customType); err != nil {
		return err
	}

	// validate values
	validate := validator.New()
	if err := validate.Struct(customType); err != nil {
		return err
	}

	return nil

}
