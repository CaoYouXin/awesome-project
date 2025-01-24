<style lang="scss">
// 工具样式表
@import '@/static/style.scss';
// 自定义图标样式表
@import '@/static/customicons.css';
/*每个页面公共css */
@import '@/uni_modules/uni-scss/index.scss';
// uview css entry
@import 'uview-plus/index.scss';
// Bug 修复样式表
@import '@/static/style-fixed.scss';

// 设置整个项目的背景色
page {
  background-color: #f5f5f5;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans',
    'Helvetica Neue', sans-serif;
}
</style>

<script>
export default {
  onLaunch: function () {
    console.log('App Launch');
  },
  onShow: function () {
    console.log('App Show');
  },
  onHide: function () {
    console.log('App Hide');
  },
};
</script>
