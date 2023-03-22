package middleware

import (
	"bytes"
	"encoding/json"
	"ginws/config"
	"ginws/repository"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogMiddleware(d *config.Dependencies, c *gin.Context, customType interface{}) error {

	logInfo := &repository.LogInfo{
		Username:   c.GetString("username"),
		IPAddress:  c.ClientIP(),
		Handler:    c.FullPath(),
		BodyParams: customType,
		ErrorInfo:  nil,
	}

	err := repository.SaveLog(d, logInfo)

	if err != nil {
		return err
	}
	return nil
}

type RequestParams struct {
	QueryParams map[string]string      `json:"query_params"`
	Body        map[string]interface{} `json:"body"`
}

func LogMiddleware2(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the query parameters
		queryParams := c.Request.URL.Query()
		queryParamsMap := make(map[string]string)
		for key, values := range queryParams {
			if len(values) > 0 {
				queryParamsMap[key] = values[0]
			}
		}

		// Read and store the request body
		var requestBody map[string]interface{}
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.Error(err)
				return
			}
			// Restore the request body
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			err = json.Unmarshal(body, &requestBody)
			if err != nil {
				c.Error(err)
				return
			}
		}

		c.Next()

		requestParams := &RequestParams{
			QueryParams: queryParamsMap,
			Body:        requestBody,
		}

		logInfo := &repository.LogInfo{
			Username:   c.GetString("username"),
			IPAddress:  c.ClientIP(),
			Handler:    c.FullPath(),
			BodyParams: requestParams,
			ErrorInfo:  nil,
		}

		err := repository.SaveLog(d, logInfo)

		if err != nil {
			c.Error(err)
		}
	}
}
