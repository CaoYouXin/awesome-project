<template>
  <up-popup :show="show" @close="close" mode="bottom" round="24rpx">
    <view>
      <view class="awe-popup__header">
        <view class="awe-popup__title">{{ title }}</view>
        <view class="awe-popup__close" @click="close">
          <uni-icons type="closeempty" size="24" color="#f39999"></uni-icons>
        </view>
      </view>
      <view class="awe-popup__slot">
        <slot></slot>
      </view>
    </view>
  </up-popup>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      default: '标题',
    },
  },
  data() {
    return {
      show: false,
    };
  },
  methods: {
    open() {
      this.show = true;
    },
    close() {
      this.show = false;
    },
  },
};
</script>

<style lang="scss" scoped>
.awe-popup__header {
  padding: 20rpx;
  border-bottom: solid 1rpx $color-primary;
  position: relative;

  .awe-popup__title {
    text-align: center;
    line-height: 42rpx;
    font-size: 32rpx;
    color: #333;
  }

  .awe-popup__close {
    position: absolute;
    right: $to-border;
    top: 50%;
    transform: translateY(-50%);
  }
}
</style>
