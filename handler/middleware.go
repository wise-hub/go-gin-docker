package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"ginws/config"
	"ginws/repository"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	secretKey1      = "0Fdkw7dFuW8LkIi359vGkwp0x7Yej3vZ"
	gcmNonceLength  = 12  // GCM recommends a 12-byte nonce
	gcmTagLength    = 128 // GCM default tag length
	plaintextLength = 48  // max length of plaintext
)

type TokenParams struct {
	User string
	Role string
}

type RequestParams struct {
	QueryParams map[string]string      `json:"query_params"`
	Body        map[string]interface{} `json:"body"`
}

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

func RoleMiddleware(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetString("role") == "RANDOMxxx" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Insufficient Rights (0)"})
			return
		}

		// add more logic
		c.Next()

	}
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

func EncryptToken(username, role string) (string, error) {
	remainingLength := plaintextLength - len(username) - len(role) - 2 // two commas
	padding := ""
	if remainingLength > 0 {
		padding = strings.Repeat("0", remainingLength)
	}
	plaintext := fmt.Sprintf("%s,%s,%s", username, role, padding)
	//fmt.Println(plaintext) // del after
	plainBytes := []byte(plaintext)

	block, err := aes.NewCipher([]byte(secretKey1))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, gcmNonceLength)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcmNonceLength)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nil, nonce, plainBytes, nil)
	cipherTextWithNonce := make([]byte, gcmNonceLength+len(cipherText))
	copy(cipherTextWithNonce, nonce)
	copy(cipherTextWithNonce[gcmNonceLength:], cipherText)

	return base64.RawURLEncoding.EncodeToString(cipherTextWithNonce), nil
}

func decryptToken(token string) (*TokenParams, error) {
	cipherTextWithNonce, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	if len(cipherTextWithNonce) < gcmNonceLength {
		return nil, errors.New("Invalid ciphertext")
	}

	nonce := cipherTextWithNonce[:gcmNonceLength]
	cipherText := cipherTextWithNonce[gcmNonceLength:]

	block, err := aes.NewCipher([]byte(secretKey1))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, gcmNonceLength)
	if err != nil {
		return nil, err
	}

	plainBytes, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	plaintext := string(plainBytes)
	parts := strings.Split(plaintext, ",")
	if len(parts) != 3 {
		return nil, errors.New("Invalid plaintext")
	}

	tokenParams := &TokenParams{
		User: parts[0],
		Role: parts[1],
	}
	//fmt.Println(parts[0])
	//fmt.Println(parts[1])

	return tokenParams, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware(d *config.Dependencies) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (0)"})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (1)"})
			return
		}

		token := authHeaderParts[1]

		if len(token) > 128 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (2)"})
			return
		}

		match, err := regexp.MatchString(`^[a-zA-Z0-9\.\-\=\_\+]+$`, token)
		if err != nil || !match {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (3)"})
			return
		}

		tokenParams, err := decryptToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (4)"})
			return
		}

		if d.Cfg.TokenDbCheck != "N" {
			if !repository.ValidateTokenOnline(d.Db, token) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"result": "Unauthorized (5)"})
				return
			}
		}

		c.Set("username", tokenParams.User)
		c.Set("role", tokenParams.Role)
		c.Next()
	}
}
