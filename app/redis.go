package app

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"fmt"
)

var (
	RedisConn 	redis.Conn
	err			error

)

func init() {
	url := fmt.Sprintf("%s:%s",Config.GetString("redis_host"),Config.GetString("redis_port"))
	RedisConn,err = redis.Dial("tcp",url)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	_,err = RedisConn.Do("auth",Config.GetString("redis_pass"))
	if err != nil {
		log.Printf("redis AUTH 失败 %s",err.Error())
		panic(err)
	}
}