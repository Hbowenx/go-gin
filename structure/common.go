package structure

const ConfigPath = "/root/go/cell_phone_store/config.toml"
//const ConfigPath = "config.toml"
const PageSize = 10

type WxAppConfig struct {
	AppId string `toml:"appId"`
	AppSecret string `toml:"appSecret"`
}

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Oss OssConfig `toml:"oss"`
	WxApp WxAppConfig `toml:"wxApp"`
}

type Pagination struct {
	Page     int `json:"page"`
}

type DatabaseConfig struct{
	Type     string `toml:"type"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type OssConfig  struct {
	Url 	     string	`toml:"url"`
	SecretId     string `toml:"secretId"`
	SecretKey    string `toml:"secretKey"`
}
