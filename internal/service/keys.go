package service

import (
	"fmt"
	"github.com/mukesh0513/RxSecure/internal/database"
	"github.com/mukesh0513/RxSecure/internal/decryptor"
	"github.com/mukesh0513/RxSecure/internal/encryptor"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"log"

	"crypto/rand"
	"github.com/gin-gonic/gin"
)

const (
	MASTER_KEY = "1D8AB5CC1426BE1A2F519877F1D92595EEA5A908A25975F60080D22E7F2583E3"
	MaxEncryptionKeys = 1000
	)

func GetTokenizeValue(c *gin.Context, args model.GetApiParams) (model.EncKeys, error) {
	var post model.EncKeys

	if err := database.DB.Where("token = ?", args.Token).Find(&post).Error; err != nil {
		log.Println(err)
		return post, err
	}
	return post, nil
}

func GetDataValue(c *gin.Context, token string) (model.Data, error) {
	var data model.Data

	if err := database.DB.Where("token = ?", token).Find(&data).Error; err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

func GetEncryptedKey(c *gin.Context, index int) (model.EncKeys, error) {
	var post model.EncKeys

	if err := database.DB.Where("id = ?", index).Find(&post).Error; err != nil {
		log.Println(err)
		return post, err
	}
	return post, nil
}

func CreateToken(c *gin.Context, payload *model.Payload) (string, error) {

	//Generate Random 32-byte token
	var token = utils.GenerateToken()
	fmt.Println("token -> " + token)
	//Generate Hash for the token to fetch index of encryption key
	var hash = utils.GenerateHash(token)
	//var hash = 1
	var encryptedKey, _ = GetEncryptedKey(c, hash)
	decryptKeyInputParams := make(map[string]interface{})
	decryptKeyInputParams["key"] = MASTER_KEY
	decryptKeyInputParams["encryptedText"] = encryptedKey.EncKey
	var decryptedKey = decryptor.AESDecrypt(decryptKeyInputParams)
	var encryptedPayload = encryptor.AESEncrypt(decryptedKey.(string), payload.Payload)
	data := new(model.Data)
	data.Token = token
	data.Payload = encryptedPayload.(string)
	database.DB.Create(data)
	return token, nil
}

func GetToken(c *gin.Context, data *model.Data) (string, error) {

	//Generate Random 32-byte token
	var token = data.Token
	fmt.Println("token -> " + token)
	//Generate Hash for the token to fetch index of encryption key
	var hash = utils.GenerateHash(token)
	//var hash = 1
	var encryptedKey, _ = GetEncryptedKey(c, hash)
	decryptKeyInputParams := make(map[string]interface{})
	decryptKeyInputParams["key"] = MASTER_KEY
	decryptKeyInputParams["encryptedText"] = encryptedKey.EncKey
	var decryptedKey = decryptor.AESDecrypt(decryptKeyInputParams)

	encryptedPayloadData, _ := GetDataValue(c, token)

	decryptPayloadInputParams := make(map[string]interface{})
	decryptPayloadInputParams["key"] = decryptedKey
	decryptPayloadInputParams["encryptedText"] = encryptedPayloadData.Payload
	var decryptedPayload = decryptor.AESDecrypt(decryptPayloadInputParams)

	return decryptedPayload.(string), nil
}

func GenerateEncryptionKeys(c *gin.Context) error {
	for i:=0; i< MaxEncryptionKeys; i++{
		key := make([]byte, 32)
		_, err := rand.Read(key)
		if err != nil {
			return err
		}

		encryptionRow := model.EncKeys{}
		encryptionRow.EncKey = encryptor.AESEncrypt(MASTER_KEY, fmt.Sprintf("%x", key)).(string)
		database.DB.Create(&encryptionRow)
	}
	return nil
}