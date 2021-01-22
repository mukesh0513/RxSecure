package redis

import (
	"github.com/gomodule/redigo/redis"
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
