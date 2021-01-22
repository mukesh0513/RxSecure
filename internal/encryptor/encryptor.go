package encryptor

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"
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

	validateError := utils.Validate(inputParams, assertRule)
	if validateError != nil {
		return nil
	}

	decodedBytes, ok := hex.DecodeString(key) //hexDecode
	if ok != nil {
		logrus.Info(map[string]interface{}{
			"component":       "AESEncrypt",
			"message": 		   ok,
		})
		return ""
	}
	encryptedString, err := EcbEncrypt(decodedBytes, plainText)
	if err != nil {
		return ""
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(encryptedString))

	return encoded
}
