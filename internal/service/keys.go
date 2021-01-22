package service

import (
	"errors"
	"fmt"
	"github.com/mukesh0513/RxSecure/internal/cache"
	"github.com/mukesh0513/RxSecure/internal/database"
	"github.com/mukesh0513/RxSecure/internal/decryptor"
	"github.com/mukesh0513/RxSecure/internal/encryptor"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"
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
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from cache",
		})
	} else {
		encryptedKey, ok := GetEncryptedKey(c, hash)
		if ok != nil{
			return "", errors.New(ok.Error())
		}
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from DB",
		})
		encKey = encryptedKey.EncKey
		cache.Set(string(hash),encKey,cache.NoExpiration)
	}

	decryptedKey, err := decryptor.AESDecrypt(MASTER_KEY, encKey)
	if err != nil {
		return "", errors.New("Unable to fetch key")
	}
	encryptedPayload, err := encryptor.AESEncrypt(decryptedKey.(string), payload.Payload)
	if err != nil {
		return "", errors.New("Error while encrypting the payload")
	}

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
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from cache",
		})
	} else {
		encryptedKey, ok := GetEncryptedKey(c, hash)
		if ok != nil{
			return "", errors.New(ok.Error())
		}
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from DB",
		})
		encKey = encryptedKey.EncKey
		cache.Set(string(hash),encKey,cache.NoExpiration)
	}

	decryptedKey, err := decryptor.AESDecrypt(MASTER_KEY, encKey)
	if err != nil {
		return "", errors.New("Unable to fetch key")
	}

	var encPayload string
	encryptedPayload := cache.Get(token); if encryptedPayload != nil {
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from cache",
		})
		encPayload = encryptedPayload.(string)
	}else{
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from DB",
		})
		encryptedPayloadData, ok := GetDataValue(c, token)
		if ok != nil{
			return "", errors.New(ok.Error())
		}

		encPayload = encryptedPayloadData.Payload
		cache.Set(token, encPayload, cache.NoExpiration)
	}


	decryptedPayload, err := decryptor.AESDecrypt(decryptedKey.(string), encPayload)
	if err != nil {
		return "", errors.New("Error while decrypting the payload")
	}

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
		encKey, ok := encryptor.AESEncrypt(MASTER_KEY, fmt.Sprintf("%x", key))
		if ok != nil{
			return errors.New(ok.Error())
		}

		encryptionRow.EncKey = encKey.(string)
		database.DB.Create(&encryptionRow)
	}
	return nil
}