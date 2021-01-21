package service

import (
	"fmt"
	"github.com/mukesh0513/RxSecure/internal/cache"
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
	token := utils.GenerateToken()
	//Generate Hash for the token to fetch index of encryption key
	hash := utils.GenerateHash(token)
	var encKey string

	key := cache.Get(string(hash)); if key != nil{
		encKey = key.(string)
	} else {
		encryptedKey, _ := GetEncryptedKey(c, hash)
		encKey = encryptedKey.EncKey
		cache.Set(string(hash),encKey,cache.NoExpiration)
	}

	decryptedKey := decryptor.AESDecrypt(MASTER_KEY, encKey)
	encryptedPayload := encryptor.AESEncrypt(decryptedKey.(string), payload.Payload)

	cache.Set(token, encryptedPayload.(string), cache.NoExpiration)

	data := new(model.Data)
	data.Token = token
	data.Payload = encryptedPayload.(string)
	database.DB.Create(data)
	return token, nil
}

func GetToken(c *gin.Context, token string) (string, error) {

	//Generate Hash for the token to fetch index of encryption key
	var hash = utils.GenerateHash(token)

	var encKey string
	key := cache.Get(string(hash)); if key != nil{
		encKey = key.(string)
	} else {
		encryptedKey, _ := GetEncryptedKey(c, hash)
		encKey = encryptedKey.EncKey
		cache.Set(string(hash),encKey,cache.NoExpiration)
	}

	decryptedKey := decryptor.AESDecrypt(MASTER_KEY, encKey)

	var encPayload string
	encryptedPayload := cache.Get(token); if encryptedPayload != nil {
		encPayload = encryptedPayload.(string)
	}else{
		encryptedPayloadData, _ := GetDataValue(c, token)
		encPayload = encryptedPayloadData.Payload
		cache.Set(token, encPayload, cache.NoExpiration)
	}


	decryptedPayload := decryptor.AESDecrypt(decryptedKey.(string), encPayload)

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