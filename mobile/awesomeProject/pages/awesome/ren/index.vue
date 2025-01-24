<template>
  <view class="awesome-time-page">
    <view class="form-container" @click="showCanlendar = true">
      <uni-datetime-picker type="date" :clear-icon="false" v-model="day" @change="change" />
    </view>

    <up-list>
      <up-list-item v-for="(item, index) in indexList" :key="index">
        <up-cell isLink :value="item.ren.text" @click="showDetail(item)">
          <template #title>
            <up-tag>{{ item.hour.hour }} : {{ item.hour.clock }}</up-tag>
          </template>
        </up-cell>
      </up-list-item>
    </up-list>

    <view class="bottom">农历：{{ lunarDate }}</view>

    <awe-popup ref="popupDetails" :title="detailsTitle">
      <view class="details-content">
        <scroll-view class="details-scroll-wrapper" scroll-y>
          <view v-for="(item, index) in detailsList" :key="index" class="details-item">
            <view class="details-name">{{ item.text }}</view>
            <view class="details-desc">{{ item.consult }}</view>
          </view>
        </scroll-view>
      </view>
    </awe-popup>
  </view>
</template>

<script>
import dayjs from 'dayjs';
import { calcRen } from '@/utils/awesome';

export default {
  data() {
    return {
      day: '',
      details: null,
    };
  },
  onShow() {
    this.resetForm();
  },
  computed: {
    lunarDate() {
      if (!this.indexList || !this.indexList.length) return '';

      return this.indexList[0].lunarDate;
    },
    detailsTitle() {
      if (!this.details) return '';

      return `${this.details.hour.hour}:${this.details.hour.clock}:${this.details.ren.text}`;
    },
    detailsList() {
      if (!this.details) return [];

      return this.details.ren.consultList;
    },
    indexList() {
      if (!this.day) return [];

      return new Array(12).fill(0).map((_, index) => {
        const datetime = this.day + ' ' + `${index * 2}`.padStart(2, '0');

        return calcRen(dayjs(datetime, 'YYYY-MM-DD HH'));
      });
    },
  },
  methods: {
    resetForm() {
      this.day = dayjs().format('YYYY-MM-DD');
    },
    change(e) {
      this.day = e;
    },
    showDetail(item) {
      this.details = JSON.parse(JSON.stringify(item));
      this.$refs.popupDetails.open();
    },
  },
};
</script>

<style lang="scss" scoped>
.awesome-time-page {
  padding-bottom: 200rpx;

  .form-container {
    padding: $to-border;
    border-bottom: solid 1rpx $color-primary;
  }

  .bottom {
    box-shadow: 0 0 20rpx rgba(0, 0, 0, 0.1);
    background-color: white;
    text-align: center;
    padding: $to-border;
    position: fixed;
    bottom: 0;
    right: 0;
    left: 0;
  }
}

.details-content {
  padding: $to-border;

  .details-scroll-wrapper {
    height: 50vh;

    .details-item {
      border: solid 1rpx white;
      border-radius: 50rpx;
      min-height: 100rpx;
      display: flex;
      justify-content: flex-start;
      align-items: flex-start;

      &:nth-of-type(odd) {
        background-color: #f39999;
        flex-direction: row;

        .details-name {
          border-color: #f39999;
          box-shadow: inset 0 0 20rpx #f39999;
        }

        .details-desc {
          text-align: left;
        }
      }

      &:nth-of-type(even) {
        background-color: #8ab5f5;
        flex-direction: row-reverse;

        .details-name {
          border-color: #8ab5f5;
          box-shadow: inset 0 0 20rpx #8ab5f5;
        }

        .details-desc {
          text-align: right;
        }
      }

      & + .details-item {
        margin-top: $near-border;
      }

      .details-name {
        flex-grow: 0;
        flex-shrink: 0;
        font-size: 0.75em;
        padding: 0 $near-border;
        width: fit-content;
        height: 100rpx;
        border-radius: 50rpx;
        background-color: white;
        border: $near-border solid transparent;
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .details-desc {
        flex-grow: 1;
        width: 1rpx;
        color: white;
        padding: $near-border $to-border;
      }
    }
  }
}
</style>
