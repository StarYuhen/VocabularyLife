package configs

// 定义框架结构体
type config struct {
	AppVersion    int         `yaml:"AppVersion" json:"AppVersion"`
	HTTPConfig    httplocal   `yaml:"HTTPConfig" json:"HTTPConfig"`
	CryptoConfig  cryptolocal `yaml:"CryptoConfig" json:"CryptoConfig"`
	CaptchaConfig captcha     `yaml:"CaptchaConfig" json:"CaptchaConfig"`
	MysqlConfig   mysql       `yaml:"MYSQLConfig" json:"MYSQLConfig"`
}

// mysql配置环境
type mysql struct {
	UserConfig  configMysql `yaml:"UserConfig" json:"user" `
	UserDBTable []string    `yaml:"UserDBTable" json:"UserDBTable"`
}

// 数据库账号配置文件内容
type configMysql struct {
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Architecture string `yaml:"Architecture"`
}

// 服务器网络配置内容
type httplocal struct {
	Ip   string `yaml:"Ip"`
	Port string `yaml:"Port"`
}

// 接口加密参数配置
type cryptolocal struct {
	Time int `yaml:"Time"`
	Data int `yaml:"Data"`
}

// 图形验证码配置
type captcha struct {
	NumberKeyLong int `yaml:"NumberKeyLong" json:"NumberKeyLong"`
	ImgWidth      int `yaml:"ImgWidth" json:"ImgWidth"`
	ImgHeight     int `yaml:"ImgHeight" json:"ImgHeight"`
}
