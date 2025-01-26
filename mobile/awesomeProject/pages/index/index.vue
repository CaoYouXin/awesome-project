<template>
  <view class="container">
    <view class="intro">
      <view class="title">我的小工具</view>
      <view class="desc">
        <text style="font-weight: bold">我与你之间的距离，就是心与心之间的距离！</text>
        <text>以下是我学习</text>
        <uni-link href="https://www.baidu.com/s?wd=道医九宫格" text="道医九宫格"></uni-link>
        <text>、</text>
        <uni-link href="https://www.baidu.com/s?wd=小六壬" text="小六壬"></uni-link>
        <text>后，开发的两个小工具，仅供自娱自乐！</text>
        <uni-link href="https://github.com/CaoYouXin/awesome-project" text="@v1.7"></uni-link>
      </view>
    </view>

    <view class="bottom">
      <up-row justify="space-between" gutter="10">
        <up-col span="6">
          <up-button @click="goTo('/pages/awesome/time/index')" text="道医九宫格"></up-button>
        </up-col>
        <up-col span="6">
          <up-button @click="goTo('/pages/awesome/ren/index')" text="小六壬"></up-button>
        </up-col>
      </up-row>
      <view style="height: 20rpx"></view>
      <up-row justify="space-between" gutter="10">
        <up-col span="12">
          <up-button @click="goTo('/pages/awesome/phone/index')" text="手机号小密码"></up-button>
        </up-col>
      </up-row>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {};
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

  .bottom {
    padding: $to-border;
    position: fixed;
    bottom: 0;
    right: 0;
    left: 0;
  }
}
</style>
