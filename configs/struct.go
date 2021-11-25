package configs

// 定义框架结构体
type config struct {
	AppVersion  int   `yaml:"AppVersion" json:"AppVersion"`
	MysqlConfig mysql `yaml:"MYSQLConfig" json:"MYSQLConfig"`
}

// mysql配置环境
type mysql struct {
	UserConfig  configMysql `json:"user" yaml:"UserConfig"`
	UserDBTable []string    `yaml:"UserDBTable" json:"UserDBTable"`
}

// 数据库账号配置文件内容
type configMysql struct {
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Architecture string `yaml:"Architecture"`
}
