package helpers

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRemoteAddr(c *gin.Context) string {

	// Try the X-Real-IP header first
	if xrip := c.Request.Header.Get("X-Real-IP"); xrip != "" {
		if ValidateIP(xrip) {
			return xrip
		}
	}

	// Try the X-Forwarded-For header next
	if xff := c.Request.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		// Use the first IP address in the list
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if ValidateIP(ip) {
				return ip
			}
		}
	}

	// Finally - use gin Request RemoteAddr
	finIp := strings.Split(c.Request.RemoteAddr, ":")[0]
	if ValidateIP(finIp) {
		return finIp
	}

	return "not_valid_ip"

}

func ValidateIP(ip string) bool {
	// Use a regular expression to validate the IP address
	match, err := regexp.MatchString(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, ip)
	if err != nil {
		return false
	}
	if !match {
		return false
	}
	return true
}

func AssertEnvForError(env string, err error) string {

	result := "Invalid request payload"

	if env == "TEST" || env == "PROD" {
		result = err.Error()
	}

	return result
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
