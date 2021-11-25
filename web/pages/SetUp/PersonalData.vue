<template>
	<view>
		<!-- 灰白色分割-->
		<view class="Databackhr">
		</view>
		<!--
		账号绑定资料设置
		-->
		<view class="PersonalData">
			<view class="PersonalDataSetUp">
				<text id="DailyLearning_Text">头像</text>
				<u-image class="ImgSetup" width="100rpx" height="100rpx"  border-radius="30rpx"
					:src="Account.SrcLogin" :lazy-load="true">
					<u-loading slot="loading"></u-loading>
					<view slot="error" style="font-size: 24rpx;">加载失败</view>
					</u-image>
				<u-icon name="arrow-right" size="40rpx" color="#c8c9cc" id="SetUpiconList"></u-icon>
			</view>
			<view id="DivLive"></view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				Account:{
					SrcLogin:"https://www.yuhenm.com/imgs/yuhen.png"
				},
			}
		},
		methods: {

		}
	}
</script>

<style>

</style>
