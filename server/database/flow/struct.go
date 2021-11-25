package flow

import "time"

// 数据库配置信息
type adminaccount struct {
	ID                  uint64    `json:"id" gorm:"primaryKey"`                       // ID主键
	Uid                 uint64    `json:"uid"`                                        // uid 一个用户拥有的一个uid
	User                string    `json:"user"`                                       // 账号
	UserName            string    `json:"user_name" gorm:"type:text;column:userName"` // 用户名
	PassWord            string    `gorm:"column:passWord"`                            // 密码
	UserLV              int       `json:"user_lv" gorm:"column:userLV"`               // 用户等级
	ImgUrl              string    `json:"img_url" gorm:"column:imgUrl"`               // 用户头像
	AccountTime         time.Time `json:"account_time" gorm:"column:accountTime"`     // 用户注册时间
	ContinuousCheckIn   uint64    `json:"continuous_check_in"`                        // 连续签到
	AccumulativeCheckIn uint64    `json:"accumulative_check_in"`                      // 累计签到
	MaximumContinuous   uint64    `json:"maximum_continuous"`                         // 最大连续
	WordLimit           uint64    `json:"word_limit"`                                 // 拥有的单词上限
	PlannedWords        uint64    `json:"planned_words"`                              // 计划的单词
	UsedWords           uint64    `json:"used_words"`                                 // 已经用过了的单词上限
}

// 设置界面 -- 账号
type account struct {
	Uid                 uint64 `json:"uid"`                   // uid 一个用户拥有的一个uid
	UserName            string `json:"user_name"`             // 用户名
	UserLV              int    `json:"user_lv"`               // 用户等级
	ImgUrl              string `json:"img_url"`               // 用户头像
	ContinuousCheckIn   uint64 `json:"continuous_check_in"`   // 连续签到
	AccumulativeCheckIn uint64 `json:"accumulative_check_in"` // 累计签到
	MaximumContinuous   uint64 `json:"maximum_continuous"`    // 最大连续
}
