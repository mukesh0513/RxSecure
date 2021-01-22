package encryptor

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"log"
	"reflect"
)

func AESEncrypt(key string, plainText string) interface{} {

	inputParams := map[string]interface{}{
		"key":       key,
		"plainText": plainText,
	}

	assertRule := map[string]reflect.Kind{
		"key":       reflect.String,
		"plainText": reflect.String,
	}

	utils.Validate(inputParams, assertRule)

	decodedBytes, ok := hex.DecodeString(key) //hexDecode
	if ok != nil {
		log.Fatal(ok)
	}
	encryptedString := EcbEncrypt(decodedBytes, plainText)

	encoded := base64.StdEncoding.EncodeToString([]byte(encryptedString))

	return encoded
}
