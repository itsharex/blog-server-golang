package config

import "time"

type Config struct {
	Server Server	`yaml:"server"`
	Db Db			`yaml:"db"`
	Mylog Mylog		`yaml:"mylog"`
	Redis Redis		`yaml:"redis"`
	JWT	  JWT		`yaml:"jwt"`
}

type Server struct {
	Address string	`yaml:"address"`
	Model string	`yaml:"model"`
}

type Db struct {
	Dialects string	`yaml:"dialects"`
  	Host string		`yaml:"host"`
  	Port int		`yaml:"port"`
  	Db string		`yaml:"db"`
  	Username string	`yaml:"username"`
  	Password string	`yaml:"password"`
  	Charset string	`yaml:"charset"`

	MaxIdle int		`yaml:"max-idle-conns"`
	MaxOpen int		`yaml:"max-open-conns"`
}

type Mylog struct{
  	Path string	`yaml:"path"`
  	Name string	`yaml:"name"`
	Model string	`yaml:"model"`
	Format string	`yaml:"format"`
	Level string	`yaml:"level"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int  `yaml:"port"`
	Db int `yaml:"db"`
	Password string `yaml:"password"`
}

type JWT struct {
	SigningKey 	string	`mapstructure:"signing-key" yaml:"signing-key"`
	ExpireTime 	time.Duration	`mapstructure:"expires-time" yaml:"expires-time"`
	Issuer 		string 	`mapstructure:"issuer" yaml:"issuer"`
	JwtBlacklistGracePeriod int64 	`mapstructure:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"`
}
