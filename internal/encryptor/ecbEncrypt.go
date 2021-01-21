package encryptor

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"log"
)

func EcbEncrypt(key []byte, message string) (string) {

	plaintext := utils.PKCS5Padding([]byte(message), aes.BlockSize)

	if len(plaintext)%aes.BlockSize != 0 {
		log.Fatal("plaintext is not a multiple of the block size")
	}

	block, cipherErr := aes.NewCipher(key)
	if cipherErr != nil {
		log.Fatal("Error creating new cipher")
	}

	cipherText := make([]byte, len(plaintext))
	mode := NewECBEncrypter(block)

	mode.CryptBlocks(cipherText, plaintext)

	encrypted := base64.StdEncoding.EncodeToString(cipherText)

	return encrypted
}
