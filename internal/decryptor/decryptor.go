package decryptor

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"reflect"
)

func AESDecrypt(key string, message string) (interface{}, error) {

	inputParams := map[string]interface{}{
		"key":       key,
		"encryptedText": message,
	}

	assertRule := map[string]reflect.Kind{
		"key":           reflect.String,
		"encryptedText": reflect.String,
	}

	validateError := utils.Validate(inputParams, assertRule)
	if validateError != nil {
		return nil, validateError
	}

	decodedBytes, ok := hex.DecodeString(key) //hexDecode
	if ok != nil {
		logrus.Info(map[string]interface{}{
			"component":       "AESDecrypt",
			"message": 		   ok,
		})
		return "", errors.New(ok.Error())
	}

	base64DecodedMessageBytes, ok := base64.StdEncoding.DecodeString(message)
	if ok != nil {
		logrus.Info(map[string]interface{}{
			"component":       "AESDecrypt",
			"message": 		   ok,
		})
		return "", errors.New(ok.Error())
	}

	base64DecodedString := cast.ToString(base64DecodedMessageBytes)

	decryptedString, err := EcbDecrypt(decodedBytes, base64DecodedString)
	if err != nil {
		return "", err
	}

	return decryptedString, nil
}