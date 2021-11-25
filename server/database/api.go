package database

// POSTAccountLogin  获取当前用户的账号信息，也是登录接口
/*
 请求参数列表：

sign:加密参数 账号+密码生成的MD5数据，加密前格式： 账号|密码|time内容 RSA
data:当前设备信息 加密前格式: 设备信息|1314|time内容 RSA
time:时间戳加上（520），加密前格式：时间戳|520 RSA



--- 请在axios配置拦截器，这个加密属于中间件解析的内容

TODO 加密和整体架构待全部完善
*/
func POSTAccountLogin() {

}
