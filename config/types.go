package config

type email struct {
	Port     int
	Host     string
	UserName string `mapstructure:"user_name"`
	Secret   string
}

type secret struct {
	JwtSecret string `mapstructure:"jwt_secret"`
	Issuer    string
}

type oss struct {
	DefaultAvatarPath  string `mapstructure:"default_avatar_path"`
	AvatarPrefix       string `mapstructure:"avatar_prefix"`
	OssAccessKeyID     string `mapstructure:"oss_access_key_id"`
	OssAccessKeySecret string `mapstructure:"oss_access_key_secret"`
	Region             string
	Bucket             string
}

type mysql struct {
	UserName string
	Password string
	Address  string
	Database string
	Charset  string
}

type redis struct {
	Host     string
	Port     string
	Database int
	Network  string
	Password string
}

type config struct {
	Email  email
	Secret secret
	OSS    oss
	Mysql  mysql
	Redis  redis
}
