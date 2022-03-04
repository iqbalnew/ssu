package customAES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type CustomAES struct {
	passphrase string
}

func NewCustomAES(passphrase string) *CustomAES {
	return &CustomAES{
		passphrase: passphrase,
	}
}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 100000, 32, sha256.New), salt
}

func (t *CustomAES) Encrypt(plaintext string) (text string, err error) {
	key, salt := deriveKey(t.passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return "", err
	}
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data), nil
}

func (t *CustomAES) Decrypt(ciphertext string) (text string, err error) {
	if ciphertext == "" {
		return ciphertext, nil
	}
	arr := strings.Split(ciphertext, "-")
	salt, err := hex.DecodeString(arr[0])
	if err != nil {
		return "", err
	}
	iv, err := hex.DecodeString(arr[1])
	if err != nil {
		return "", err
	}
	data, err := hex.DecodeString(arr[2])
	if err != nil {
		return "", err
	}
	key, _ := deriveKey(t.passphrase, salt)

	b, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return "", err
	}
	data, err = aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
