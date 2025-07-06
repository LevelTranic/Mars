package v2

type Configure struct {
	Version        string            `yaml:"config_version"`
	AuthToken      string            `yaml:"auth_token"`
	ActiveGC       bool              `yaml:"active_gc"`
	AllRecycled    bool              `yaml:"all_recycled"`
	ActiveSniffing bool              `yaml:"active_sniffing"`
	Json           ConfigureJson     `yaml:"json"`
	Server         ConfigureServer   `yaml:"server"`
	Redis          ConfigureRedis    `yaml:"redis"`
	Database       ConfigureDatabase `yaml:"database"`
}

type ConfigureJson struct {
	Runtime string `yaml:"runtime"`
	Sonic   struct {
		FastMode    bool `yaml:"fast_mode"`
		JITPretouch bool `yaml:"jit_pretouch"`
	} `yaml:"sonic"`
}

type ConfigureServer struct {
	TLSInCert          string `yaml:"tls_in_cert"`
	TLSInKey           string `yaml:"tls_in_key"`
	ListenMode         string `yaml:"listen_mode"`
	ListenAddress      string `yaml:"listen_address"`
	ListenPort         int    `yaml:"listen_port"`
	MaxRequestBodySize int    `yaml:"max_request_body_size"`
	RequestBodySize    int    `yaml:"req_body_size"`
	ResponseBodySize   int    `yaml:"resp_body_size"`
}

type ConfigureRedis struct {
	Enabled  bool     `yaml:"enabled"`
	Driver   string   `yaml:"driver"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Servers  []string `yaml:"servers"`
}

type ConfigureDatabase struct {
	Driver string                  `yaml:"driver"`
	Global ConfigureDatabaseGlobal `yaml:"global"`
	Sqlite ConfigureDatabaseSqlite `yaml:"sqlite"`
	Mysql  ConfigureDatabaseMysql  `yaml:"mysql"`
}

type ConfigureDatabaseGlobal struct {
	MaxIdleConns    int `yaml:"max_idle_conns"`
	ConnMaxLifetime int `yaml:"conn_max_lifetime"`
	MaxOpenConns    int `yaml:"max_open_conns"`
	ConnMaxIdleTime int `yaml:"conn_max_idle_time"`
}

type ConfigureDatabaseSqlite struct {
	Name string `yaml:"name"`
}

type ConfigureDatabaseMysql struct {
	User                      string `yaml:"user"`
	Pass                      string `yaml:"pass"`
	Addr                      string `yaml:"addr"`
	DBName                    string `yaml:"dbname"`
	Charset                   string `yaml:"charset"`
	DisableDatetimePrecision  bool   `yaml:"disable_datetime_precision"`
	DontSupportRenameIndex    bool   `yaml:"dont_support_rename_index"`
	DontSupportRenameColumn   bool   `yaml:"dont_support_rename_column"`
	SkipInitializeWithVersion bool   `yaml:"skip_initialize_with_version"`
}
