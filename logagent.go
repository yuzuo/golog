package main

import (
	"flag"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/hpcloud/tail"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	traceback "logtrace/traceback"
	"time"
)

func GetUuid4() string {
	return uuid.Must(uuid.NewV4()).String()
}

var logFile = flag.String("log", "./a.log", "eg:-log=./a.log")

var pool *redis.Pool

func SendLog(content string) error {
	fmt.Println(content)
	queue := &traceback.Queue{Pool: pool}
	logtime := traceback.GetLogTime(content)
	loglevel := traceback.GetLogLevelStr(content)
	msg := &traceback.Message{Name: "logAgent", CreatedAt: time.Now(), Content: map[string]string{"content": content}, LogTime: logtime, Level: loglevel}
	for {
		err := queue.Delivery(msg)
		return err
	}
}

func main() {
	flag.Parse()

	Config := new(traceback.Config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}

	pool = &redis.Pool{
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

	t, _ := tail.TailFile(*logFile, tail.Config{Follow: true})
	for line := range t.Lines {
		go SendLog(line.Text)
	}
}
