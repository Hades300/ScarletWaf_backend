package common

type MysqlConf struct {
	Addr     string `toml:"addr"`
	Database string `toml:"database"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type RedisConf struct {
	Addr     string `toml:"addr"`
	Database int    `toml:"database"`
}

type Conf struct {
	Mysql MysqlConf `toml:"mysql"`
	Redis RedisConf `toml:"redis"`
}

var DEVELOP = true
