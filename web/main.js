import App from './App'

// #ifndef VUE3
import Vue from 'vue'

// 引用uView框架
import uview from 'uview-ui';

//使用vue中间件导入
Vue.use(uview)
Vue.config.productionTip = false
App.mpType = 'app'
const app = new Vue({
	...App
})
app.$mount()
// #endif

// #ifdef VUE3
import {
	createSSRApp
} from 'vue'




export function createApp() {
	const app = createSSRApp(App)
	return {
		app
	}
}
// #endif
