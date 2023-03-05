package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	secretKey = []byte("XmBkb3mjBMXUEw8iRZjAtWApn8RU5wkh")
)

func EncryptData(username string, role string, expirationDate time.Time) (string, error) {
	// Pad the plaintext data to a multiple of the block size

	//plaintext := []byte(fmt.Sprintf("%s,%s,%s", username, role, expirationDate))

	plaintext := []byte(fmt.Sprintf("%s,%s,%s", username, role, expirationDate.Format(time.RFC3339)))

	blockSize := aes.BlockSize
	padding := blockSize - (len(plaintext) % blockSize)
	padtext := make([]byte, len(plaintext)+padding)
	copy(padtext, plaintext)
	for i := len(plaintext); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	// Generate a new AES cipher using the provided key
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	// Generate a new GCM cipher with a random nonce
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// Encrypt the padded plaintext using the GCM cipher and nonce
	ciphertext := gcm.Seal(nonce, nonce, padtext, nil)

	// Encode the ciphertext and key as a base64 string separated by a colon
	ciphertextEncoded := base64.URLEncoding.EncodeToString(ciphertext)

	return ciphertextEncoded, nil
}

func DecryptData(token string) (*TokenData, error) {
	// Decode the base64-encoded string into ciphertext and nonce
	ciphertext, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	// Generate a new AES cipher using the provided key
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	// Generate a new GCM cipher with the nonce extracted from the ciphertext
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("invalid ciphertext")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext using the GCM cipher and nonce
	padtext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	// Unpad the plaintext data
	padding := int(padtext[len(padtext)-1])
	plaintext := padtext[:len(padtext)-padding]

	// Parse the plaintext data into its individual components
	parts := strings.Split(string(plaintext), ",")
	if len(parts) != 3 {
		return nil, errors.New("invalid plaintext")
	}

	var tokenData TokenData

	expDate, err := time.Parse(time.RFC3339, parts[2])
	if err != nil {
		return nil, errors.New("invalid expiration date")
	}
	tokenData.ExpDate = expDate
	tokenData.User = parts[0]
	tokenData.Role = parts[1]
	tokenData.Token = token

	return &tokenData, nil
}
