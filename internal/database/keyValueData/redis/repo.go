package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type RedisConnection struct{}

var (
	redisConn         redis.Conn
	//RedisConnProvider keyValueData.IKeyValueDatabase
)

func Initialize(conn redis.Conn) {
	if redisConn == nil {
		redisConn = conn
		//RedisConnProvider = RedisConnection{}
	}
}

type KeyService struct{}

func (service KeyService) GetEncryptedKey(hash int64) string {
	result, err := redis.String(redisConn.Do("HGET", "enc_keys:"+string(hash), "enc_key"))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (service KeyService) SetEncryptedKey(hash int64, enc_key string) string {
	result, err := redisConn.Do("HSET", "enc_keys:"+string(hash), "enc_key", "123")
	if err != nil {
		log.Fatal(err)
	}
	return result.(string)
}

func NewKeyService() KeyService {
	return KeyService{}
}
//
//func (db RedisConnection) FindByKey(table string, key string) interface{} {
//	result, err := redis.String(redisConn.Do("HGET", table, key))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return result
//}
//
//func (db RedisConnection) Create(table string, args interface{}) interface{} {
//	result, err := redisConn.Do("HSET", redis.Args{}.Add(table).AddFlat(args))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return result
//}
//
//func (db RedisConnection) Delete() {
//
//}
//
//func GetRedisConn() redis.Conn {
//	return redisConn
//}
