import axios from "axios"
// 设置 axios 全局
axios.defaults.baseURL = 'http://localhost';
axios.defaults.withCredentials = true;

module.exports = {
	// 返回请求获取的配置数据
	AccountSetUp: function(userName) {
		console.log("sdfsd")
	},
	// 动态更改网页内容
	AccountLoad: async function(Account){
		let data;
		await axios.post('/api/accountSetup', {
			"userName": "admin",
		}).then(function(response) {
			// 请求更改数据
			data = response.data;
		});
		// 将数据赋值
		Account.TextName = data.user_name;
		Account.TextUID = data.uid;
		Account.SrcLogin = data.img_url;
		Account.TextLVL = data.user_lv;
		Account.ContinuousCheckIn = data.continuous_check_in;
		Account.AccumulativeCheckIn = data.accumulative_check_in;
		Account.MaxContinuous = data.maximum_continuous;
		console.log(data)
	},
	// 点击后跳转的个人设置界面，路由跳转
	AccountClick:function(){
		console.log("sdfds")
	},
}
