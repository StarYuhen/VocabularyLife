package MYSQL

// 用于存储数据表的结构体文件

// adminuserAccount 这个数据表所有的字段
type adminuserAccount struct {
	ID                  uint64 `json:"id"`                    // ID主键
	Uid                 uint64 `json:"uid"`                   // uid 一个用户拥有的一个uid
	User                string `json:"user"`                  // 账号
	UserName            string `json:"user_name"`             // 用户名
	PassWord            string `json:"pass_word"`             // 密码
	UserLV              int    `json:"user_lv"`               // 用户等级
	ImgUrl              string `json:"img_url"`               // 用户头像
	AccountTime         string `json:"account_time"`          // 用户注册时间
	ContinuousCheckIn   uint64 `json:"continuous_check_in"`   // 连续签到
	AccumulativeCheckIn uint64 `json:"accumulative_check_in"` // 累计签到
	MaximumContinuous   uint64 `json:"maximum_continuous"`    // 最大连续
	WordLimit           uint64 `json:"word_limit"`            // 拥有的单词上限
	PlannedWords        uint64 `json:"planned_words"`         // 计划的单词
	UsedWords           uint64 `json:"used_words"`            // 已经用过了的单词上限
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


