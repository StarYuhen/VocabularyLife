<template>
	<!-- 设置页面 -->
	<!-- 必须只能有一个view当主盒子 类似于body ，template只是tag标签，不是body-->
	<view>
		<view class="AccountSetUp">
			<!-- 
			用户头像 昵称 设置 使用按钮 
			使用vue的监听命令，使这个view视图可点击 
			-->
			<navigator url='/pages/SetUp/PersonalData'>
			<view class="Account" hover-start-time="50" hover-stay-time="200" >
				<u-image class="ImgSetup" width="150rpx" height="150rpx" shape="square" border-radius="30rpx"
					:src="Account.SrcLogin" :lazy-load="true" :fade="true" duration="450">
					<u-loading slot="loading"></u-loading>
					<view slot="error" style="font-size: 24rpx;">加载失败</view>
					</u-image>
				<view class="TextSetup">
					<text id="SetUpName">{{Account.TextName}}</text>
					<view class="TextDataSetUp">
						<text id="SetUpUID">UID:{{Account.TextUID}}</text>
						<text id="SetUpLVL">LVL{{Account.TextLVL}}</text>
					</view>
				</view>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpicon"></u-icon>
			</view>
			</navigator>

			<!-- 签到统计界面 -->
			<view class="AccountSignIn">
				<view id="SingInSetUp_ONE">
					<text id="SingIn">{{Account.AccumulativeCheckIn}}</text>
					<view>连续签到</view>
				</view>
				<view id="SingInSetUp_TWO">
					<text id="SingIn">{{Account.ContinuousCheckIn}}</text>
					<view>累计签到</view>
				</view>
				<view id="SingInSetUp_THREE">
					<text id="SingIn">{{Account.MaxContinuous}}</text>
					<view>最大连续</view>
				</view>
			</view>
		</view>

		<!-- 学习设置 -->
		<view class="LearningSettings">
			<text id="LearningText">学习统计</text>
		</view>

		<!-- 下面的选项按钮 -->
		<view class="LearningSettingsInit">
			<view class="DailyLearningTextClass">
				<view class="LearningViewText">
					<text id="DailyLearning_TextInit">每日学习量</text>
					<view id="LearningView">当日复习的单词量+当日新学单词量</view>
				</view>
				<u-input id="Ufield" v-model="value" :type="type" :border="true" placeholder=" " border-color="#38b29b"
					height="50" />
			</view>
			<view id="DivLive"></view>


			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">单词上限总量</text>
				<view id="LearningClickInit">
					<text id="TextIDSetUp">{{LearningSettingsSetUp.WordNumber}}</text>
				</view>
			</view>
			<view id="DivLive"></view>

			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">每日复习提醒时间</text>
				<view id="LearningClickInit">
					<text id="TextIDSetUp">{{LearningSettingsSetUp.WordNumber}}</text>
				</view>
			</view>
			<view id="DivLive"></view>


			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">已规划记忆的单词量</text>
				<view id="LearningClickInit">
					<text id="TextIDSetUp">{{LearningSettingsSetUp.WordNumber}}</text>
				</view>
			</view>
		</view>

		<view class="LearningSettings">
			<text id="LearningText">学习情况</text>
		</view>
		
		
		





		<!-- 单词上限 -->
		<view class="LearningSettings">
			<text id="LearningText">其他设置</text>
		</view>

		<!--其他设置下布局界面-->
		<view class="LearningSettingsInit">
			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">更多设置</text>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>

			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">词汇量测试</text>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>

			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">邀请有礼</text>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>

			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">退出登录</text>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>
			
			<view class="DailyLearningTextClass">
				<text id="DailyLearning_Text">关于</text>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>

		</view>



	</view>
</template>


<script>
	// template script style 都是一级节点，无论放在哪里，优先级都最高
	import axios from "axios"
	// 设置 axios 全局
	//	axios.defaults.baseURL = 'http://localhost';
	//	axios.defaults.withCredentials = true;
	import setupJs from "@/assets/js/SetUp.js"

	export default {
		data() {
			return {
				Account: {
					//  头像地址 名称 UID 等级 连续签到 累计签到 最大连续
					SrcLogin: "https://www.yuhenm.com/imgs/yuhen.png",
					TextName: "玉衡",
					TextUID: "1008611",
					TextLVL: "10",
					ContinuousCheckIn: "20",
					AccumulativeCheckIn: "20",
					MaxContinuous: "20",
				},

				LearningSettingsSetUp: {
					WordNumber: "40"
				},
				value: '10',
				type: 'number',
				border: true
			}
		},
		onLoad(opention) {
			// 这里时同步函数，页面渲染完后会立即渲染的函数
			setupJs.AccountLoad(this.Account);
		},
		methods: {
			// AccountClick: setupJs.AccountClick,
			AccountSetUp: setupJs.AccountSetUp,
			AccountInit: function() {}
		},
	}
</script>
