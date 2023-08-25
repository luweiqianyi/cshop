package store

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
	"time"
)

func TestRedisUtil(t *testing.T) {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		PingTimeout: time.Second * 30,
	}
	MustUseRedisStore(conf)
	cli := getStoreInstance()
	success := cli.Ping()
	fmt.Println(success)

	if cli == nil {
		fmt.Println("end!")
	}
	err := Set("key", "TestRedisUtil", 3600)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestSet(t *testing.T) {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		PingTimeout: time.Second * 30,
	}
	MustUseRedisStore(conf)
	err := getStoreInstance().Set("zhangSan", "111")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func TestGet(t *testing.T) {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		PingTimeout: time.Second * 30,
	}
	MustUseRedisStore(conf)
	val, err := Get("zhangSan")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("val: %v\n", val)
}

func TestDelete(t *testing.T) {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		PingTimeout: time.Second * 30,
	}
	MustUseRedisStore(conf)

	err := Delete("zhangSan")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
