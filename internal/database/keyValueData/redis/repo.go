package redis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisConnection struct{}

var (
	redisConn         redis.Conn
	connPool *redis.Pool
	//RedisConnProvider keyValueData.IKeyValueDatabase
)

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 50,
		// max number of connections
		MaxActive: 10000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func init()  {
	connPool = newPool()
	//conn, err := redis.Dial("tcp", "localhost:6379")
	//if err != nil {
	//	fmt.Println("db err: ", err)
	//}
	//redisConn = conn
}
//func Initialize(conn redis.Conn) {
//	if redisConn == nil {
//		redisConn = conn
//		//RedisConnProvider = RedisConnection{}
//	}
//}

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
