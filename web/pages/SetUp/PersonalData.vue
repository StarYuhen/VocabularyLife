<template>
  <view>
    <!-- 灰白色分割-->
    <!-- <view class="Databackhr">
    </view> -->
    <!--
    账号绑定资料设置
    -->
    <view class="PersonalData">
      <view class="PersonalDataSetUp">
        <u-cell-group>
          <u-cell-item title="头像">
            <u-image class="ImgSetup" width="110rpx" height="110rpx" border-radius="20rpx"
                     :src="Account.SrcLogin" :lazy-load="true" margin-right="5px">
              <u-loading slot="loading"></u-loading>
              <view slot="error" style="font-size: 24px;">
                <text display="inline-block" margin-top="25rpx ">加载失败</text>
              </view>
            </u-image>
            <!-- 修改头像文字（暂未添加） -->
          </u-cell-item>
          <u-cell-item title="UID" hover-class="cell-hover-class" :arrow="false">{{ Account.TextUID }}</u-cell-item>
          <u-cell-item title="等级" hover-class="cell-hover-class">{{ Account.TextLVL }}</u-cell-item>
          <u-cell-item title="个人昵称" hover-class="cell-hover-class">{{ Account.TextName }}</u-cell-item>
          <u-cell-item title="签名" hover-class="cell-hover-class">{{ Account.TextSign }}</u-cell-item>
        </u-cell-group>
        <!-- 灰白色分割线 -->
        <view class="Databackhr"></view>
        <!-- 第二列表 -->
        <u-cell-group class="SecondList">
          <u-cell-item title="邮箱" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="手机" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="微信" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="性别" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="地区" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="生日" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="学校" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
          <u-cell-item title="星座" hover-class="cell-hover-class"> {{ Account.TextMail }}</u-cell-item>
        </u-cell-group>
        <view class="Databackhr"></view>

        <u-cell-group class="SecondList">
          <u-cell-item title="修改密码" hover-class="cell-hover-class"></u-cell-item>
          <u-cell-item title="退出登录" hover-class="cell-hover-class"></u-cell-item>
        </u-cell-group>
      </view>
      <view id="DivLive"></view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      Account: {
        SrcLogin: "https://www.yuhenm.com/imgs/yuhen.png",
        TextUID: "1008611",
        TextLVL: "10",
        TextName: "玉衡",
        TextSign: "这个人还没有什么签名呢",
        TextMail: "未设置"
      },
    }
  },
  methods: {}
}
</script>

<style>

</style>
