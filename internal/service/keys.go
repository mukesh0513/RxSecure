package service

import (
	"errors"
	"fmt"
	"github.com/mukesh0513/RxSecure/internal/cache"
	"github.com/mukesh0513/RxSecure/internal/config"
	"github.com/mukesh0513/RxSecure/internal/decryptor"
	"github.com/mukesh0513/RxSecure/internal/encryptor"
	"github.com/mukesh0513/RxSecure/internal/model"
	"github.com/mukesh0513/RxSecure/internal/utils"
	"github.com/sirupsen/logrus"

	"crypto/rand"
	"github.com/gin-gonic/gin"
)

const (
	MASTER_KEY        = "1D8AB5CC1426BE1A2F519877F1D92595EEA5A908A25975F60080D22E7F2583E3"
	MaxEncryptionKeys = 1000
)

func CreateToken(c *gin.Context, payload *model.Payload) (string, error) {

	//Generate Random 32-byte token
	token := utils.GenerateToken()
	//Generate Hash for the token to fetch index of encryption key
	hash := utils.GenerateHash(token)
	var encryptedKey string

	key := cache.Get(string(hash)); if key != nil{
		encryptedKey = key.(string)
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from cache",
		})
	} else {
		encryptedKey = IKeyFactory(config.GetConfig().DatabaseSelection.Keys).GetEncryptedKey(int64(hash))
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from DB",
		})

		cache.Set(string(hash),encryptedKey,cache.NoExpiration)
	}

	decryptedKey, err := decryptor.AESDecrypt(MASTER_KEY, encryptedKey)
	if err != nil {
		return "", errors.New("Unable to fetch key")
	}
	encryptedPayload := encryptor.AESEncrypt(decryptedKey.(string), payload.Payload)

	cache.Set(token, encryptedPayload.(string), cache.NoExpiration)

	//data := new(model.Data)
	//data.Token = token
	//data.Payload = encryptedPayload.(string)
	//database.DB.Create(data)
	IPayloadFactory(config.GetConfig().DatabaseSelection.Payload).SetEncryptedData(token, encryptedPayload.(string))

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
		var encryptedKey = IKeyFactory(config.GetConfig().DatabaseSelection.Keys).GetEncryptedKey(int64(hash))
		logrus.Info(map[string]interface{}{
			"component":       "GetToken",
			"message": 		   "Fetching from DB",
		})
		cache.Set(string(hash),encryptedKey,cache.NoExpiration)
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
		encryptedPayloadData := IPayloadFactory(config.GetConfig().DatabaseSelection.Payload).GetEncryptedData(token)
		cache.Set(token, encryptedPayloadData, cache.NoExpiration)
	}


	decryptedPayload, err := decryptor.AESDecrypt(decryptedKey.(string), encPayload)
	if err != nil {
		return "", errors.New("Error while decrypting the payload")
	}

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
		IKeyFactory(config.GetConfig().DatabaseSelection.Keys).SetEncryptedKey(int64(i), encryptionRow.EncKey)
		//encryptionRow.EncKey = encKey.(string)
		//database.DB.Create(&encryptionRow)
	}
	return nil
}