package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func Encrypt(stringToEncrypt string) (encryptedString string, err error) {
	key, _ := hex.DecodeString(bitoffset)
	plaintext := []byte(stringToEncrypt)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(encryptedString string) (decryptedString string, err error) {

	key, _ := hex.DecodeString(bitoffset)
	enc, _ := hex.DecodeString(encryptedString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func DecryptString(encVal string, defaultVal string) (val string) {
	val, err := Decrypt(encVal)
	if err != nil {
		return defaultVal
	}
	return val
}

func EncryptFloat64(val float64) (encryptedString string, err error) {
	sVal := fmt.Sprintf("%f8", val)
	return Encrypt(sVal)
}

func DecryptFloat64(encVal string, defaultVal float64) (val float64) {
	v, err := Decrypt(encVal)
	if err != nil {
		return defaultVal
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		return s
	}
	return defaultVal
}

func EncryptInt64(val int64) (encryptedString string, err error) {
	sVal := fmt.Sprintf("%d", val)
	return Encrypt(sVal)
}

func DecryptInt64(encVal string, defaultVal int64) (val int64) {
	v, err := Decrypt(encVal)
	if err != nil {
		return defaultVal
	}

	if s, err := strconv.ParseInt(v, 10, 64); err == nil {
		return s
	}
	return defaultVal
}
