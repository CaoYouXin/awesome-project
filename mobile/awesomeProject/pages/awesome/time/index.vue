<template>
  <view class="awesome-time-page">
    <view class="form-container">
      <u-form labelPosition="left" :model="form" :rules="rules" ref="uForm">
        <u-form-item label="名称" prop="name" borderBottom>
          <u-input v-model="form.name" placeholder="请输入名称"></u-input>
        </u-form-item>
        <u-form-item label="时间" prop="birthday" borderBottom>
          <u-datetime-picker
            v-model="form.birthday"
            mode="datetime"
            hasInput
            :minDate="new Date(1930, 0, 1).getTime()"
            :maxDate="new Date().getTime()"
          ></u-datetime-picker>
        </u-form-item>
        <u-button @click="submitForm()" text="提交"></u-button>
      </u-form>
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
          <u-button
            @click="handleDelete(index)"
            text="删除"
            type="error"
          ></u-button>
        </uni-td>
      </uni-tr>
    </uni-table> -->

    <u-card
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
              <u-text type="primary" :text="item.name"></u-text>
            </view>
            <view class="flex1" style="margin-right: 10rpx">
              <u-text
                type="success"
                :text="`${item.hour.hour}(${item.hour.element})`"
              ></u-text>
            </view>
          </view>
          <u-icon name="trash" @click.stop="handleDelete(index)"></u-icon>
        </view>
      </template>
      <template #body>
        <uni-table border stripe emptyText="暂无更多数据">
          <uni-tr>
            <uni-th align="center" style="width: 50%">
              阳历：{{ item.solarDate }}
            </uni-th>
            <uni-th align="center" style="width: 50%">
              阴历：{{ item.lunarDate }}
            </uni-th>
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
            <uni-th align="center" style="width: 50%">
              男：此时主{{ getNow(index).man }}格
            </uni-th>
            <uni-th align="center" style="width: 50%">
              女：此时主{{ getNow(index).woman }}格
            </uni-th>
          </uni-tr>
          <uni-tr v-for="(step, idx) in getNow(index).tableData" :key="idx">
            <uni-td align="center" v-html="step.man"></uni-td>
            <uni-td align="center" v-html="step.woman"></uni-td>
          </uni-tr>
        </uni-table>
      </template>
    </u-card>
  </view>
</template>

<script>
import { calcAwesome, calcNow } from "@/utils/awesome";

export default {
  data() {
    return {
      form: {
        name: "",
        birthday: "",
      },
      rules: {
        name: {
          type: "string",
          required: true,
          message: "请填写姓名",
          trigger: ["blur", "change"],
        },
      },
      tableData: [],
      hideFoot: [],
    };
  },
  computed: {
    nowTableData() {
      return this.tableData.map((item) =>
        calcNow(item.geYang, item.geYin, item.solarDate)
      );
    },
  },
  onShow() {
    this.resetForm();
    this.tableData = JSON.parse(localStorage.getItem("timeHistory")) || [];
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
        this.hideFoot = this.hideFoot.map(() => true);
        this.hideFoot[idx] = false;
      }
    },
    save() {
      console.log("saving");
      localStorage.setItem("timeHistory", JSON.stringify(this.tableData));
    },
    handleDelete(index) {
      this.tableData.splice(index, 1);
      this.save();
    },
    resetForm() {
      this.form = {
        name: "",
        birthday: new Date().getTime(),
      };
    },
    submitForm() {
      this.$refs.uForm
        .validate()
        .then((_) => {
          uni.showToast({ title: "校验通过" });

          console.log(this.form.birthday);
          const res = calcAwesome(this.form.birthday);
          this.tableData.unshift({
            ...this.form,
            ...res,
          });
          this.save();
          this.resetForm();
        })
        .catch((errors) => {
          uni.showToast({ title: "校验失败", icon: "error" });
          console.error(errors);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
@import "@/static/style.scss";

.awesome-time-page {
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
