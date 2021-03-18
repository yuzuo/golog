package traceback

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var MongoConn *mgo.Session

type Message struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string
	Level     string
	LogTime   string
	Content   map[string]string `json:"content"`
}

func (m *Message) GetChannel() string {
	/*
		获取消息队列名称
	*/
	return m.Name
}

func SetUpMongo(dbhost string, authdb string, authuser string, authpass string, poollimit int) *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{dbhost}, //数据库地址 dbhost: mongodb://user@123456:127.0.0.1:27017
		Timeout:   60 * time.Second, // 连接超时时间 timeout: 60 * time.Second
		Source:    authdb,           // 设置权限的数据库 authdb: admin
		Username:  authuser,         // 设置的用户名 authuser: user
		Password:  authpass,         // 设置的密码 authpass: 123456
		PoolLimit: poollimit,        // 连接池的数量 poollimit: 100
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	MongoConn = s
	return s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := MongoConn.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func Insert(db, collection string, doc interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(doc)
}

func (m *Message) Resolve() error {
	/*
	 消费队列
	*/
	err := Insert("logs", "vanwei", &m)
	if err != nil {
		return err
	}
	fmt.Printf("consumed %+v\n", m.Content)
	return nil
}

func (m *Message) Marshal() ([]byte, error) {
	/*
		序列化
	*/
	return jsoniter.Marshal(m)
}

func (m *Message) Unmarshal(reply []byte) (IMessage, error) {
	/*
		反序列化
	*/
	var msg Message
	err := jsoniter.Unmarshal(reply, &msg)
	return &msg, err
}
