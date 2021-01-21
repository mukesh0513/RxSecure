package decryptor

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/spf13/cast"
	"log"
)

func EcbDecrypt(key []byte, message string) (string) {

	cipherText, decodeErr := base64.StdEncoding.DecodeString(message)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	if len(cipherText) < aes.BlockSize {
		log.Fatal("cipherText block is too short")
	}

	cipherBlock, cipherErr := aes.NewCipher(key)
	if cipherErr != nil {
		log.Fatal("Error occurred creating cipher")
	}

	if len(cipherText)%aes.BlockSize != 0 {
		log.Fatal("ciphertext is not a multiple of the block size")
	}

	mode := NewDecrypterWithModeECB(cipherBlock)

	mode.CryptBlocks(cipherText, cipherText)

	unpaddedCipherText := utils.PKCS5UnPadding(cipherText)

	plaintext := cast.ToString(unpaddedCipherText)

	return plaintext
}