package decryptor

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/spf13/cast"
	"log"
	"reflect"
)

func AESDecrypt(key string, message string) (interface{}) {

	inputParams := map[string]interface{}{
		"key":       key,
		"encryptedText": message,
	}

	assertRule := map[string]reflect.Kind{
		"key":           reflect.String,
		"encryptedText": reflect.String,
	}

	utils.Validate(inputParams, assertRule)

	decodedBytes, ok := hex.DecodeString(key) //hexDecode
	if ok != nil {
		log.Fatal(ok)
	}

	base64DecodedMessageBytes, ok := base64.StdEncoding.DecodeString(message)
	if ok != nil {
		log.Fatal(ok)
	}

	base64DecodedString := cast.ToString(base64DecodedMessageBytes)

	decryptedString := EcbDecrypt(decodedBytes, base64DecodedString)

	return decryptedString
}