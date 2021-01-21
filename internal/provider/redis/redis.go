package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type RedisConnection struct{}

var redisConn redis.Conn

func Initialize(conn redis.Conn) {
	if redisConn == nil {
		redisConn = conn
	}
}

func (db *RedisConnection) Find(table string, model interface{}, args interface{}) interface{} {
	result, err := redis.String(redisConn.Do("HGET", table, args))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (db *RedisConnection) Create(table string, model interface{}, args interface{}) interface{} {
	result, err := redisConn.Do("HSET", redis.Args{}.Add(table).AddFlat(args))
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *RedisConnection) Delete() {

}

func GetRedisConn() redis.Conn {
	return redisConn
}
