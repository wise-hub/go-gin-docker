package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"ginws/config"
	"ginws/model"
	"ginws/repository"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	secretKey1      = "XmBkb3mjBMXUEw8iRZjAtWApn8RU5wkh"
	gcmNonceLength  = 12  // GCM recommends a 12-byte nonce
	gcmTagLength    = 128 // GCM default tag length
	plaintextLength = 48  // plaintext with padding
)

func ValidateToken(c *gin.Context, d *config.Dependencies) (*model.TokenParams, error) {

	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		fmt.Println("1")
		return nil, errors.New("Unauthorized (0)")
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		fmt.Println("2")
		return nil, errors.New("Unauthorized (1)")

	}

	token := authHeaderParts[1]

	if len(token) > 168 { // for len 32 user and len 32 role + standard datetime
		return nil, errors.New("Unauthorized (2)")
	}

	match, err := regexp.MatchString(`^[a-zA-Z0-9\.\-\=\_\+]+$`, token)
	if err != nil || !match {
		return nil, errors.New("Unauthorized (3)")
	}

	tokenParams, err := decryptToken(token)
	if err != nil {
		return nil, errors.New("Unauthorized (4)")
	}

	if !repository.ValidateTokenOnline(d.Db, token) {

		return nil, errors.New("Unauthorized (5)")
	}

	return tokenParams, nil
}

func EncryptToken(username, role string) (string, error) {
	remainingLength := plaintextLength - len(username) - len(role) - 2 // two commas
	padding := ""
	if remainingLength > 0 {
		padding = strings.Repeat("0", remainingLength)
	}
	plaintext := fmt.Sprintf("%s,%s,%s", username, role, padding)
	fmt.Println(plaintext) // del after
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

func decryptToken(token string) (*model.TokenParams, error) {
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

	tokenParams := &model.TokenParams{
		User: parts[0],
		Role: parts[1],
	}
	fmt.Println(parts[0])
	fmt.Println(parts[1])

	return tokenParams, nil
}
