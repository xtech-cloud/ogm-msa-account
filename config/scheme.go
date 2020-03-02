package config

type Logger_ struct {
	Level int `yaml:level`
}

type SQLite_ struct {
	Path string `yaml:path`
}

type MySQL_ struct {
	Address  string `yaml:address`
	User     string `yaml:user`
	Password string `yaml:password`
	DB       string `yaml:db`
}

type Database_ struct {
	Lite   bool    `yaml:lite`
	MySQL  MySQL_  `yaml:mysql`
	SQLite SQLite_ `yaml:sqlite`
}

type Encrypt_ struct {
	Secret string `yaml:secret`
}

type Service_ struct {
	TTL      int64 `yaml:ttl`
	Interval int64 `yaml:interval`
}

type ConfigScheme_ struct {
	Service  Service_  `yaml:service`
	Logger   Logger_   `yaml:logger`
	Database Database_ `yaml:database`
	Encrypt  Encrypt_  `yaml:encrypt`
}
