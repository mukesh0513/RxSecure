package encryptor

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"
)

func EcbEncrypt(key []byte, message string) (string, error) {

	plaintext := utils.PKCS5Padding([]byte(message), aes.BlockSize)

	if len(plaintext)%aes.BlockSize != 0 {
		logrus.Info(map[string]interface{}{
			"component":       "EcbEncrypt",
			"message": 		   "plaintext is not a multiple of the block size",
		})
		return "", errors.New("plaintext is not a multiple of the block size")
	}

	block, cipherErr := aes.NewCipher(key)
	if cipherErr != nil {
		logrus.Info(map[string]interface{}{
			"component":       "EcbEncrypt",
			"message": 		   "Error creating new cipher",
		})
		return "", errors.New("Error creating new cipher")
	}

	cipherText := make([]byte, len(plaintext))
	mode := NewECBEncrypter(block)

	mode.CryptBlocks(cipherText, plaintext)

	encrypted := base64.StdEncoding.EncodeToString(cipherText)

	return encrypted, nil
}
