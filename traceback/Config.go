package traceback

type ConfigMongoStorage struct {
	Network    string `yaml:"net"`
	Addr       string `yaml:"addr"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Db         string `yaml:"db"`
	Collection string `yaml:"collection"`
}

type ConfigRedis struct {
	Net      string `yaml:"net"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

type Config struct {
	RedisAgent ConfigRedis        `yaml:"redis-agent"`
	Storage    ConfigMongoStorage `yaml:"mongo-storage"`
}
