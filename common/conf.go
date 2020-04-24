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
	Password string `toml:"password"`
}

type ScarletConf struct {
	Addr string `toml:"addr"`
}

type Conf struct {
	Mysql   MysqlConf   `toml:"mysql"`
	Redis   RedisConf   `toml:"redis"`
	Scarlet ScarletConf `toml:"scarlet"`
}

var DEVELOP = true

var AbbrFuncMap = map[string]string{"waf": "waf_status", "get": "get_args_check", "post": "post_args_check", "cookie": "cookie_check", "ua": "ua_check", "blackip": "ip_blacklist", "whiteip": "ip_whitelist", "cc": "cc_defense", "sql": "libsqli_token_check"}
