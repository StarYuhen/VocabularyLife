package flow

import "time"

// 数据库配置信息
type adminaccount struct {
	ID                  uint64    `json:"id" gorm:"primaryKey;"`                     // ID主键
	Uid                 string    `json:"uid"`                                       // uid 一个用户拥有的一个uid
	User                string    `json:"user"`                                      // 账号
	UserName            string    `json:"UserName" gorm:"type:text;column:userName"` // 用户名
	Password            string    // 密码字段不允许导出
	UserLV              int       `json:"UserLV" gorm:"column:userLV"`                           // 用户等级
	ImgUrl              string    `json:"ImgUrl" gorm:"column:imgUrl"`                           // 用户头像
	AccountTime         time.Time `json:"account_time" gorm:"column:accountTime"`                // 用户注册时间
	ContinuousCheckIn   uint64    `json:"ContinuousCheckIn" gorm:"column:ContinuousCheckIn"`     // 连续签到
	AccumulativeCheckIn uint64    `json:"AccumulativeCheckIn" gorm:"column:AccumulativeCheckIn"` // 累计签到
	MaximumContinuous   uint64    `json:"MaximumContinuous" gorm:"column:MaximumContinuous"`     // 最大连续
	WordLimit           uint64    `json:"WordLimit" gorm:"column:WordLimit"`                     // 拥有的单词上限
	PlannedWords        uint64    `json:"PlannedWords" gorm:"column:PlannedWords"`               // 计划的单词
	UsedWords           uint64    `json:"UsedWords" gorm:"column:UsedWords"`                     // 已经用过了的单词上限
}

// 登录信息表
type userinfo struct {
	Uid string `json:"uid"`
	Ip  string `json:"ip"`
	Jwt string `json:"jwt"`
}

// WriteIO 用于集合数据库流程--只限于查询和更新
type WriteIO struct {
	Admin       adminaccount
	UserInfo    userinfo
	UpdateInter UpdateInterface
}

// UpdateInterface  每一次执行更新和查询主次都需要检查是否存在
type UpdateInterface interface {
	selectAccount(user string, password string) bool
	ReadUserInfo() bool
	SelectAccountUser(user string) bool
	UpdateUserInfoIp(ip, jwt string) error
	UpdateAdminAccount(imgurl string) error
	UpdateAdminAccountPassword(password string) error
}
