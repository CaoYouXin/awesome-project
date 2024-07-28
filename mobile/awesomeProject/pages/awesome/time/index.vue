<template>
  <view class="awesome-time-page">
    <view class="form-container">
      <u-form labelPosition="left" :model="form" :rules="rules" ref="uForm">
        <u-form-item label="姓名" prop="name" borderBottom>
          <u-input v-model="form.name" placeholder="请输入姓名"></u-input>
        </u-form-item>
        <u-form-item label="生日" prop="birthday" borderBottom>
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
    <uni-table border stripe emptyText="暂无更多数据">
      <!-- 表头行 -->
      <uni-tr>
        <uni-th align="center">姓名</uni-th>
        <uni-th align="center">时辰</uni-th>
        <uni-th align="center">阳历生日</uni-th>
        <uni-th align="left">阳格</uni-th>
        <uni-th align="center">阴历生日</uni-th>
        <uni-th align="left">阴格</uni-th>
        <uni-th align="left">操作</uni-th>
      </uni-tr>
      <!-- 表格数据行 -->
      <uni-tr v-for="(item, index) in tableData" :key="index">
        <uni-td>{{ item.name }}</uni-td>
        <uni-td>{{ item.hour.hour }}（{{ item.hour.element }}）</uni-td>
        <uni-td>{{ item.solarDate }}</uni-td>
        <uni-td>{{ item.geYang }}</uni-td>
        <uni-td>{{ item.lunarDate }}</uni-td>
        <uni-td>{{ item.geYin }}</uni-td>
        <uni-td>
          <u-button
            @click="tableData.splice(index, 1)"
            text="删除"
            type="error"
          ></u-button>
        </uni-td>
      </uni-tr>
    </uni-table>
  </view>
</template>

<script>
import { calcAwesome } from "@/utils/awesome";

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
    };
  },
  onShow() {
    this.resetForm();
    this.tableData = JSON.parse(localStorage.getItem("timeHistory")) || [];
  },
  onHide() {
    console.log("saving");
    localStorage.setItem("timeHistory", JSON.stringify(this.tableData));
  },
  methods: {
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
          this.tableData.push({
            ...this.form,
            ...res,
          });

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
.awesome-time-page {
  .form-container {
    padding: $to-border;

    button {
      margin: $to-border 0;
    }
  }
}
</style>
