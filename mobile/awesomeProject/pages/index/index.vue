<template>
  <view class="container">
    <view class="intro">
      <view class="title">我的小工具</view>
      <view class="desc">学习九宫格，使用九宫格</view>
    </view>
    <text class="intro">百度：</text>
    <uni-link :href="href" :text="href"></uni-link>
    <view style="height: 200rpx"></view>
    <u-button
      @click="goTo('/pages/awesome/time/index')"
      text="时间密码"
    ></u-button>
    <u-button
      @click="goTo('/pages/awesome/phone/index')"
      text="手机号密码"
    ></u-button>
  </view>
</template>

<script>
export default {
  data() {
    return {
      href: "https://www.baidu.com/s?wd=道医九宫格",
    };
  },
  methods: {
    goTo(url) {
      uni.navigateTo({
        url,
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.container {
  padding: $to-border;
  font-size: $font-size;
  line-height: $line-height;

  .title {
    margin-bottom: $font-size-huge;
    font-size: $font-size-huge;
    line-height: $line-height-huge;
    text-align: center;
  }

  button + button {
    margin-top: $to-border;
  }
}
</style>
