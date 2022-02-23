package customAES

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES(t *testing.T) {

	t.Run("Encrypts and decrypts ", func(t *testing.T) {

		plainTexts := []string{"1234567890", "123456789012345678901234567890123456789012345678901234567890", "1", ""}

		aes := NewCustomAES("Hello")

		for _, plainText := range plainTexts {

			encrypted := aes.Encrypt(plainText)

			decrypted := aes.Decrypt(encrypted)

			assert.Equal(t, plainText, decrypted)
		}
	})
}
