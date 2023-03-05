package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AssertEnvForError(env string, err error) string {

	result := "Invalid request payload"

	if env == "TEST" {
		result = err.Error()
	}

	return result
}

func GenerateAccessToken() string {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		panic(err)
	}
	token := hex.EncodeToString(tokenBytes)

	return token
}

func IsValidAccessToken(c *gin.Context) bool {

	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return false
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return false
	}

	token := authHeaderParts[1]

	if len(token) != 64 {
		return false
	}

	match, err := regexp.MatchString(`^[0-9a-fA-F]+$`, token)
	if err != nil || !match {
		return false
	}

	return true
}

func IsValidCustomerID(cust_id string) bool {

	if len(cust_id) != 9 {
		return false
	}

	match, err := regexp.MatchString(`^\d{9}$`, cust_id)
	if err != nil || !match {
		return false
	}

	return true
}

func IsValidUsername(username string) bool {

	if len(username) > 20 {
		return false
	}

	match, err := regexp.MatchString(`^[a-zA-Z0-9]{1,20}$`, username)
	if err != nil || !match {
		return false
	}

	return true
}

func IsValidEGN(egn string) bool {
	if len(egn) != 10 {
		return false
	}

	// Use regular expression to check if egn contains only digits
	match, err := regexp.MatchString(`^\d{10}$`, egn)
	if err != nil || !match {
		return false
	}

	var weights = []int{2, 4, 8, 5, 10, 9, 7, 3, 6}
	var sum int

	for i := 0; i < len(weights); i++ {
		digit, err := strconv.Atoi(string(egn[i]))
		if err != nil {
			return false
		}
		sum += digit * weights[i]
	}

	var controlDigit = (sum % 11) % 10
	lastDigit, err := strconv.Atoi(string(egn[9]))
	if err != nil {
		return false
	}

	return controlDigit == lastDigit
}
