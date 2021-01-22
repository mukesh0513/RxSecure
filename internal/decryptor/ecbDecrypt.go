package decryptor

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

func EcbDecrypt(key []byte, message string) (string, error) {

	cipherText, decodeErr := base64.StdEncoding.DecodeString(message)
	if decodeErr != nil {
		logrus.Info(map[string]interface{}{
			"component":       "EcbDecrypt",
			"message": 		   decodeErr,
		})
		return "", errors.New(decodeErr.Error())
	}

	if len(cipherText) < aes.BlockSize {
		logrus.Info(map[string]interface{}{
			"component":       "EcbDecrypt",
			"message": 		   "cipherText block is too short",
		})
		return "", errors.New("cipherText block is too short")
	}

	cipherBlock, cipherErr := aes.NewCipher(key)
	if cipherErr != nil {
		logrus.Info(map[string]interface{}{
			"component":       "EcbDecrypt",
			"message": 		   "Error occurred creating cipher",
		})
		return "", errors.New("Error occurred creating cipher")
	}

	if len(cipherText)%aes.BlockSize != 0 {
		logrus.Info(map[string]interface{}{
			"component":       "EcbDecrypt",
			"message": 		   "ciphertext is not a multiple of the block size",
		})
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	mode := NewDecrypterWithModeECB(cipherBlock)

	mode.CryptBlocks(cipherText, cipherText)

	unpaddedCipherText := utils.PKCS5UnPadding(cipherText)

	plaintext := cast.ToString(unpaddedCipherText)

	return plaintext, nil
}