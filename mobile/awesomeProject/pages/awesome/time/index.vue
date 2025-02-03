<template>
  <view class="awesome-time-page">
    <view class="form-container">
      <up-form labelPosition="left" :model="form" :rules="rules" ref="uForm">
        <up-form-item label="名称" prop="name" borderBottom>
          <up-input v-model="form.name" placeholder="请输入名称"></up-input>
        </up-form-item>
        <up-form-item label="时间" prop="birthday" borderBottom>
          <up-input v-model="form.birthday" placeholder="请输入时间"></up-input>
        </up-form-item>
        <up-row>
          <up-col span="6">
            <up-form-item label="农历" prop="isLunar" borderBottom>
              <up-switch v-model="form.isLunar"></up-switch>
            </up-form-item>
          </up-col>
          <up-col span="6">
            <up-form-item label="闰月" prop="isFlag" borderBottom>
              <up-switch v-model="form.isFlag" :disabled="isFlagDisabled"></up-switch>
            </up-form-item>
          </up-col>
        </up-row>
        <up-button @click="submitForm()" text="提交"></up-button>
      </up-form>
    </view>

    <!-- <uni-table border stripe emptyText="暂无更多数据">
      <uni-tr>
        <uni-th align="center">姓名</uni-th>
        <uni-th align="center">时辰</uni-th>
        <uni-th align="center">阳历生日</uni-th>
        <uni-th align="left">阳格</uni-th>
        <uni-th align="center">阴历生日</uni-th>
        <uni-th align="left">阴格</uni-th>
        <uni-th align="left">操作</uni-th>
      </uni-tr>
      <uni-tr v-for="(item, index) in tableData" :key="index">
        <uni-td>{{ item.name }}</uni-td>
        <uni-td>{{ item.hour.hour }}（{{ item.hour.element }}）</uni-td>
        <uni-td>{{ item.solarDate }}</uni-td>
        <uni-td>{{ item.geYang }}</uni-td>
        <uni-td>{{ item.lunarDate }}</uni-td>
        <uni-td>{{ item.geYin }}</uni-td>
        <uni-td>
          <up-button
            @click="handleDelete(index)"
            text="删除"
            type="error"
          ></up-button>
        </uni-td>
      </uni-tr>
    </uni-table> -->

    <up-card
      v-for="(item, index) in tableData"
      :key="index"
      :class="[hideFoot[index] && 'hide-foot']"
      padding="10"
      @click="toggleFoot(index)"
    >
      <template #head>
        <view class="line between">
          <view class="line start flex1">
            <view class="flex0" style="margin-right: 10rpx">
              <up-text type="primary" :text="item.name"></up-text>
            </view>
            <view class="flex1" style="margin-right: 10rpx">
              <up-text type="success" :text="`${item.hour.hour}(${item.hour.element})`"></up-text>
            </view>
          </view>
          <view @click.stop="handleDelete(index)">
            <up-icon name="trash"></up-icon>
          </view>
        </view>
      </template>
      <template #body>
        <uni-table border stripe emptyText="暂无更多数据">
          <uni-tr>
            <uni-th align="center" style="width: 50%">阳历：{{ item.solarDate }}</uni-th>
            <uni-th align="center" style="width: 50%">阴历：{{ item.lunarDate }}</uni-th>
          </uni-tr>
          <uni-tr>
            <uni-td align="center">{{ item.geYang }}</uni-td>
            <uni-td align="center">{{ item.geYin }}</uni-td>
          </uni-tr>
        </uni-table>
      </template>
      <template #foot>
        <uni-table border stripe emptyText="暂无更多数据">
          <uni-tr>
            <uni-th align="center" style="width: 50%">男：此时主{{ getNow(index).man }}格</uni-th>
            <uni-th align="center" style="width: 50%">女：此时主{{ getNow(index).woman }}格</uni-th>
          </uni-tr>
          <uni-tr v-for="(step, idx) in getNow(index).tableData" :key="idx">
            <uni-td align="center" v-html="step.man"></uni-td>
            <uni-td align="center" v-html="step.woman"></uni-td>
          </uni-tr>
        </uni-table>
      </template>
    </up-card>

    <awe-popup ref="popupLevel" title="数字的密码">
      <secret1 />
    </awe-popup>

    <awe-popup ref="popupStar" title="标星小密码">
      <secret2 />
    </awe-popup>

    <up-float-button :isMenu="true" bottom="50rpx" :list="list" @item-click="itemClick"></up-float-button>
  </view>
</template>

<script>
import Secret1 from './components/secret1.vue';
import Secret2 from './components/secret2.vue';
import { calcDateArgs, calcAwesome2, calcNow } from '@/utils/awesome';
import dayjs from 'dayjs';

export default {
  components: {
    Secret1,
    Secret2,
  },
  data() {
    return {
      form: {
        name: '',
        birthday: '',
        isLunar: false,
        isFlag: false,
      },
      rules: {
        name: {
          type: 'string',
          required: true,
          message: '请填写名称',
          trigger: ['blur', 'change'],
        },
      },
      tableData: [],
      hideFoot: [],
      list: [
        { key: 'level', name: 'level', color: '#fff', backgroundColor: 'red' },
        { key: 'star', name: 'star', color: '#fff', backgroundColor: 'green' },
      ],
    };
  },
  computed: {
    nowTableData() {
      return this.tableData.map(item => calcNow(item.geYang, item.geYin, item.solarDate));
    },
    isFlagDisabled() {
      return !this.form.isLunar;
    },
  },
  onShow() {
    this.resetForm();
    this.tableData = JSON.parse(localStorage.getItem('timeHistory')) || [];
    this.hideFoot = this.tableData.map(() => true);
  },
  onHide() {
    this.save();
  },
  methods: {
    getNow(idx) {
      return this.nowTableData[idx];
    },
    toggleFoot(idx) {
      if (!this.hideFoot[idx]) {
        this.hideFoot[idx] = true;
      } else {
        this.resetFoot(idx);
      }
    },
    resetFoot(except = -1) {
      this.hideFoot = this.tableData.map(() => true);
      if (except < 0 || except >= this.tableData.length) {
        return;
      }
      this.hideFoot[except] = false;
    },
    save() {
      console.log('saving');
      localStorage.setItem('timeHistory', JSON.stringify(this.tableData));
    },
    handleDelete(index) {
      this.tableData.splice(index, 1);
      this.save();
    },
    resetForm() {
      this.form = {
        name: '',
        birthday: new dayjs().format('YYYYMMDDHH'),
      };
      this.$nextTick(() => {
        this.$refs.uForm.clearValidate();
      });
    },
    validateBirthday() {
      const errMsg = '日期格式不正确: ';

      if (this.form.birthday.length !== '2025020317'.length) {
        throw new Error(errMsg + this.form.birthday.length);
      }

      const hour = parseInt(this.form.birthday.slice(8, 10));
      if (hour < 0 || hour > 23) {
        throw new Error(errMsg + hour);
      }

      return [hour, ...calcDateArgs(this.form.birthday, this.form.isLunar, this.form.isFlag)];
    },
    submitForm() {
      this.$refs.uForm
        .validate()
        .then(_ => {
          const args = this.validateBirthday();
          uni.showToast({ title: '校验通过' });

          console.log(args);
          const res = calcAwesome2(...args);
          this.tableData.unshift({
            ...this.form,
            ...res,
          });
          this.save();
          this.resetForm();
          this.resetFoot();
        })
        .catch(errors => {
          uni.showToast({ title: '校验失败', icon: 'error' });
          console.error(errors);
        });
    },
    itemClick(e) {
      const key = e.key;
      const capitalize = key.charAt(0).toUpperCase() + key.slice(1);
      const ref = this.$refs['popup' + capitalize];
      if (ref && ref.open) {
        ref.open();
      }
      console.log('itemClick', 'popup' + capitalize);
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

    button {
      margin: $to-border 0;
    }
  }

  .u-card.hide-foot {
    ::v-deep .u-card__foot {
      display: none;
    }
  }
}
</style>
