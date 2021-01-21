package encryptor

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"log"
	"reflect"
)

func AESEncrypt(inputParams map[string]interface{}) (interface{}) {

	assertRule := map[string]reflect.Kind{
		"key":       reflect.String,
		"plainText": reflect.String,
	}

	utils.Validate(inputParams, assertRule)

	key := inputParams["key"].(string)
	message := inputParams["plainText"].(string)

	decodedBytes, ok := hex.DecodeString(key) //hexDecode
	if ok != nil {
		log.Fatal(ok)
	}
	encryptedString := EcbEncrypt(decodedBytes, message)

	encoded := base64.StdEncoding.EncodeToString([]byte(encryptedString))

	return encoded
}