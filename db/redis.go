package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

const (
	addr  = "r-bp1nk23h10ng2relnqpd.redis.rds.aliyuncs.com:6379"
	psw   = "Lht86052516"
	uname = "default"
)

var cli *redis.Client

func init() {
	cli = redis.NewClient(&redis.Options{
		Username: uname,
		Addr:     addr,
		Password: psw, // no password set
		DB:       0,   // use default DB
	})
	ping := cli.Ping(context.Background())
	if _, err := ping.Result(); err != nil {
		log.Fatal(fmt.Sprintf("redis init err:%v", err))
	}
}

const (
	openId = "openId"
)

func GetRedis() *redis.Client {
	return cli
}
