package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type KeyService struct{}

func (service KeyService) GetEncryptedKey(hash int64) string {
	//conn, err := redis.Dial("tcp", "localhost:6379")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer conn.Close()
	//
	//result, err := redis.String(conn.Do("HGET", "enc_keys:"+string(hash), "enc_key"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//return result

	result, err := redis.String(redisConn.Do("HGET", "enc_keys:"+string(hash), "enc_key"))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (service KeyService) SetEncryptedKey(hash int64, enc_key string) string {
	//Initialising every time as concurrency issue coming up.

	//conn, err := redis.Dial("tcp", "localhost:6379")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer conn.Close()
	//
	//_, err = conn.Do("HSET", "enc_keys:"+string(hash), "enc_key", enc_key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//return  ""

	result, err := redisConn.Do("HSET", "enc_keys:"+string(hash), "enc_key", enc_key)
	if err != nil {
		log.Fatal(err)
	}
	return result.(string)
}

func NewKeyService() KeyService {
	return KeyService{}
}