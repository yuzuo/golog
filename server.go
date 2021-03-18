package main

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"logtrace/traceback"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	Config := new(traceback.Config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	pool := &redis.Pool{
		MaxIdle: 200,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial(Config.RedisAgent.Net, Config.RedisAgent.Addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	traceback.SetUpMongo(Config.Storage.Addr, Config.Storage.Db, Config.Storage.User, Config.Storage.Password, 10)

	queue := &traceback.Queue{Pool: pool}

	msg := &traceback.Message{
		Name: "logAgent",
	}
	// make a main context
	ctx := context.Background()

	cancelFunc := queue.InitReceiver(ctx, msg, 10)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)

	for {
		switch <-quit {
		case syscall.SIGINT:
			cancelFunc()
			return
		}
	}
}
