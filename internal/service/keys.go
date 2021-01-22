package service

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mukesh0513/RxSecure/internal/decryptor"
	"github.com/mukesh0513/RxSecure/internal/encryptor"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/utils"
)

const (
	MASTER_KEY        = "1D8AB5CC1426BE1A2F519877F1D92595EEA5A908A25975F60080D22E7F2583E3"
	MaxEncryptionKeys = 1000
)

func CreateToken(c *gin.Context, payload *model.Payload) (string, error) {

	//Generate Random 32-byte token
	var token = utils.GenerateToken()
	fmt.Println("token -> " + token)
	//Generate Hash for the token to fetch index of encryption key
	var hash = utils.GenerateHash(token)
	//var hash = 1
	var encryptedKey = IKeyFactory("mysql").GetEncryptedKey(int64(hash))
	decryptKeyInputParams := make(map[string]interface{})
	decryptKeyInputParams["key"] = MASTER_KEY
	decryptKeyInputParams["encryptedText"] = encryptedKey
	var decryptedKey = decryptor.AESDecrypt(decryptKeyInputParams)
	var encryptedPayload = encryptor.AESEncrypt(decryptedKey.(string), payload.Payload)
	//data := new(model.Data)
	//data.Token = token
	//data.Payload = encryptedPayload.(string)
	//
	IPayloadFactory("mysql").SetEncryptedData(token, encryptedPayload.(string))
	//database.DB.Create(data)
	return token, nil
}

func GetToken(c *gin.Context, token string) (string, error) {

	//Generate Random 32-byte token
	//var token = args.Token
	//fmt.Println("token -> " + token)
	//Generate Hash for the token to fetch index of encryption key
	var hash = utils.GenerateHash(token)
	//var hash = 1
	var encryptedKey = IKeyFactory("mysql").GetEncryptedKey(int64(hash))
	decryptKeyInputParams := make(map[string]interface{})
	decryptKeyInputParams["key"] = MASTER_KEY
	decryptKeyInputParams["encryptedText"] = encryptedKey
	var decryptedKey = decryptor.AESDecrypt(decryptKeyInputParams)

	encryptedPayloadData:= IPayloadFactory("mysql").GetEncryptedData(token)

	decryptPayloadInputParams := make(map[string]interface{})
	decryptPayloadInputParams["key"] = decryptedKey
	decryptPayloadInputParams["encryptedText"] = encryptedPayloadData
	var decryptedPayload = decryptor.AESDecrypt(decryptPayloadInputParams)

	return decryptedPayload.(string), nil
}

func GenerateEncryptionKeys(c *gin.Context) error {
	for i := 0; i < MaxEncryptionKeys; i++ {
		key := make([]byte, 32)
		_, err := rand.Read(key)
		if err != nil {
			return err
		}

		encryptionRow := model.EncKeys{}
		encryptionRow.EncKey = encryptor.AESEncrypt(MASTER_KEY, fmt.Sprintf("%x", key)).(string)
		IKeyFactory("mysql").SetEncryptedKey(int64(i), encryptionRow.EncKey)
		//database.DB.Create(&encryptionRow)
	}
	return nil
}