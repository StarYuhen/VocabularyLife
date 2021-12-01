package configs

// 定义框架结构体
type config struct {
	AppVersion    appversion `yaml:"AppVersion" json:"AppVersion"`
	HTTPConfig    http       `yaml:"HTTPConfig" json:"HTTPConfig"`
	CryptoConfig  crypto     `yaml:"CryptoConfig" json:"CryptoConfig"`
	CaptchaConfig captcha    `yaml:"CaptchaConfig" json:"CaptchaConfig"`
	MysqlConfig   mysql      `yaml:"MYSQLConfig" json:"MYSQLConfig"`
	RedisConfig   redis      `yaml:"RedisConfig" json:"RedisConfig"`
}

// 项目前后端版本内容
type appversion struct {
	Backend int `yaml:"Backend"`
	Web     int `yaml:"Web"`
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
type http struct {
	Ip   string `yaml:"Ip"`
	Port string `yaml:"Port"`
}

// 接口加密参数配置
type crypto struct {
	Time       int      `yaml:"Time"`
	Data       int      `yaml:"Data"`
	Jwt        string   `yaml:"Jwt"`
	JwtName    string   `yaml:"JwtName"`
	Cookie     string   `yaml:"Cookie"`
	CookieName string   `yaml:"CookieName"`
	CookieList []string `yaml:"CookieList"`
}

// 图形验证码配置
type captcha struct {
	NumberKeyLong int `yaml:"NumberKeyLong" json:"NumberKeyLong"`
	ImgWidth      int `yaml:"ImgWidth" json:"ImgWidth"`
	ImgHeight     int `yaml:"ImgHeight" json:"ImgHeight"`
}

// redis 配置
type redis struct {
	Addr         string `yaml:"Addr" json:"Addr"`
	Password     string `yaml:"Password" json:"Password"`
	DB           []int  `yaml:"DB" json:"DB"`
	PoolSize     int    `yaml:"PoolSize" json:"PoolSize"`
	MinIdleConns int    `yaml:"MinIdleConns" json:"MinIdleConns"`
}
