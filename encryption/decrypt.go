package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

const (
	//CertificateFilePath the absolute filepath that stores the pfx file for encrypting and decrypting
	CertificateFilePath = "/secrets/internal-encryption-certificate/cert.pfx"
	//CertificatePasswordFilePath the absolute filepath that stores the password for the pfx key
	CertificatePasswordFilePath = "/secrets/internal-encryption-certificate/cert-pfx-password"
)

//Decrypt Decrypts the encrypted value using the CertificateFilePath and CertificatePasswordFilePath CONST values
func Decrypt(encData string) (string, error) {

	key, err := ioutil.ReadFile(CertificateFilePath)
	if err != nil {
		return "", fmt.Errorf("Error opening private key file: %w", err)
	}
	pass, err := ioutil.ReadFile(CertificatePasswordFilePath)
	if err != nil {
		return "", fmt.Errorf("Error opening private key password file: %w", err)
	}
	password := strings.TrimSpace(string(pass))
	priv, _, err := pkcs12.Decode(key, password)
	if err != nil {
		return "", fmt.Errorf("Error decoding private key file: %w", err)
	}
	if err := priv.(*rsa.PrivateKey).Validate(); err != nil {
		return "", fmt.Errorf("Error validating private key file: %w", err)
	}
	ed, err := base64.StdEncoding.DecodeString(encData)
	if err != nil {
		return "", fmt.Errorf("Error base64 decoding value: %w", err)
	}
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), ed)
	if err != nil {
		return "", fmt.Errorf("Error decrypting value: %w", err)
	}
	value := string(decryptedData)
	return value, nil
}
