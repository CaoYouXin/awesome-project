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
      <view class="awe-popup__content">
        <scroll-view class="awe-popup__scroll-wrapper" scroll-y>
          <view v-for="(item, index) in detailsList" :key="index" class="awe-popup__item">
            <view class="awe-popup__name fit-content hanz">{{ item.text }}</view>
            <view class="awe-popup__desc">{{ item.consult }}</view>
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
@import '@/components/awe-popup/awe-popup-list.scss';

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
</style>
