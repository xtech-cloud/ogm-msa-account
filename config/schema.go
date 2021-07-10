package config

type Logger_ struct {
	Level string `yaml:level`
	Dir   string `yaml:dir`
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
	Driver string  `yaml:driver`
	MySQL  MySQL_  `yaml:mysql`
	SQLite SQLite_ `yaml:sqlite`
}

type Encrypt_ struct {
	Secret string `yaml:secret`
}

type JWT_ struct {
	Expiry int64 `yaml:expiry`
}

type Token_ struct {
	JWT JWT_ `yaml:jwt`
}

type Service_ struct {
	Name     string `yaml:name`
	TTL      int64  `yaml:ttl`
	Interval int64  `yaml:interval`
	Address  string `yaml:address`
}

type ConfigSchema_ struct {
	Service  Service_  `yaml:service`
	Logger   Logger_   `yaml:logger`
	Database Database_ `yaml:database`
	Encrypt  Encrypt_  `yaml:encrypt`
	Token    Token_    `yaml:token`
}
