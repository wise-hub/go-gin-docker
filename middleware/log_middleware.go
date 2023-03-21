package middleware

import (
	"ginws/config"
	"ginws/repository"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	QueryParams map[string]string      `json:"query_params"`
	Body        map[string]interface{} `json:"body"`
}

func LogMiddleware(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		queryParams := c.Request.URL.Query()

		queryParamsMap := make(map[string]string)
		for key, values := range queryParams {
			if len(values) > 0 {
				queryParamsMap[key] = values[0]
			}
		}

		// Get the request body as a map[string]interface{}
		var requestBody map[string]interface{}
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			// Handle error if needed
		}

		// Combine the GET parameters and request body
		requestParams := &RequestParams{
			QueryParams: queryParamsMap,
			Body:        requestBody,
		}

		// You can customize this to include additional information as needed
		logInfo := &repository.LogInfo{
			Username:   c.GetString("username"), // Replace this with the actual username
			IPAddress:  c.ClientIP(),
			Handler:    c.HandlerName(),
			BodyParams: requestParams, // You might need to adjust this based on the request format
			ErrorInfo:  nil,           // Replace this with an actual error if one occurs
		}

		err = repository.SaveLog(d, logInfo)

		if err != nil {
			c.Error(err)
		}
	}
}
