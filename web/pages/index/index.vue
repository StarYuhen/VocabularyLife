
<!--复习界面 -->
<template>
	<view>
		<text> {{rsa}}</text>
	</view>
</template>

<script>
	import cryptorsa from "@/assets/js/crypto.js"
	export default {
		data() {
			return {
				rsa:""
			}
		},
		onLoad() {
			this.rsa=cryptorsa.RSATest()
		},
		methods: {

		}
	}
</script>

