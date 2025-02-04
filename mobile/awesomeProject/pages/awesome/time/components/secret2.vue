<template>
  <view class="awe-popup__content">
    <scroll-view class="awe-popup__scroll-wrapper" scroll-y>
      <view v-for="(item, index) in list" :key="index" class="awe-popup__item">
        <view class="awe-popup__name hanz">{{ item.name }}</view>
        <view v-if="typeof item.desc === 'string'" class="awe-popup__desc">{{ item.desc }}</view>
        <view v-else-if="typeof item.desc === 'object' && item.desc.length" class="awe-popup__desc">
          <view v-for="(desc, idx) in item.desc" :key="idx" class="awe-popup__desc-item">
            <view class="awe-popup__desc-key" style="width: 2em">{{ desc.key }}</view>
            <view class="awe-popup__desc-key">&nbsp;:&nbsp;</view>
            <view class="awe-popup__desc-value">{{ desc.value }}</view>
          </view>
        </view>
      </view>
    </scroll-view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      list: [
        {
          name: '※',
          desc: '自出生开始计算阴阳：阳格周期内，阳格为显性（阴格为个体无意识）；阴格周期内，阴格为显性（阳格为个体无意识）；能开发出：松果体、心轮能力的人，可依靠呼吸调整阴阳格；如果没有调整意识系统能力者，按照程序显示的计算。',
        },
        {
          name: '※',
          desc: '如果计算周期内：生孩子，按照生孩子日期所在格，重新按照第0年，重新累计本周期时间。',
        },
        {
          name: '※',
          desc: '家族遗传 RNA 是灵魂部分；这部分遵循双方家族遗传惯例，部分缺的孩子会缺，部分多的孩子也会继承；父亲遗传 61.8% 情感部分，母亲遗传 61.8% 智慧部分；灵魂簇遵循男女男女继承规则；家族灵魂簇具有继承性；',
        },
        {
          name: '※',
          desc: [
            { key: '总述', value: '家族灵魂簇具有继承性，孩子与父母九宫格在自然分娩时遵守' },
            { key: '天格', value: '0 生 9 或 0 ，9 生 6 ，6 生 3 ； 101 童子也可以生天地格（ 7 或者 9 ）' },
            { key: '地格', value: '1 生 4 ，4 生 7 ，7 生 1 ；特殊情况： 7347 生 347 ，也容易得胃病抑郁症' },
            { key: '人格', value: '8 生 5 ，5 生 2 ，2 生 8' },
            { key: '特例', value: '带 3 副魂，念想比较强，会生出自己一样格（例如： 325 生 325 ）' },
          ],
        },
        {
          name: '※',
          desc: '通常一个灵魂需要 80-120 次学会人类载具课程；如果：没有修完情绪体，则需要重复修，直到突破能量限制；修结束的密码——以“爱”的方式回馈对方的情绪变化；',
        },
        {
          name: '※',
          desc: '基本规则：4 克 369；4 为独裁者，与 2 合，其余克。家庭财富：至少有 7。同一灵魂簇为 6 和，不同灵魂簇按时区关系，分为冲、平、和。特殊规则优先，五行生克关系劣后：特殊关系为无序向有序转化、包容为优先关系；正面情绪体为优先关系。',
        },
        {
          name: '6 和',
          desc: [
            {
              key: '性',
              value:
                '是否缺 6,6 是主魂（ 336 与 426/246/156 和）辅魂（ 167/268/369 ）平（魄里面带 6 ）缺（魂魄都缺 6 ）',
            },
            { key: '格', value: '303 、 336 、 369 、后天童子和 325 、 235 、 358 和' },
            {
              key: '自我',
              value:
                '4 、 6 公平效率冲突 246 、 426 ； 28 冲突，服务官与统帅 268 、 28101 ，这种格自己冲突，最多 5 和，除非修行到禅二以上，自己调整管理情绪体',
            },
            {
              key: '时间',
              value: '男性比女性大 2.3 岁或者 9.7 岁和；时间反着（女大 3 或者大 9 岁），根据生理、换魂周期反而冲、克',
            },
          ],
        },
        {
          name: '※',
          desc: '地球是镜子法则，二元性部分刚好反过来， 147 生克关系刚好相反； 5D 载具没有： 45/78 （肝胆脾胃），九宫格缺：4578 反而是好事，只需要忌口 + 辟谷',
        },
        {
          name: '生克',
          desc: [
            { key: '相生', value: '金生水，水生木，木生火，火生土，土生金' },
            { key: '相克', value: '金克木，木克土，土克水，水克火，火克金' },
          ],
        },
        {
          name: '※',
          desc: '情义，事业属于义的部分；情绪体上，做事需要合伙人，每个人性格担任的角色不同；事业合伙人关系——需要 3 和。',
        },
        {
          name: '※',
          desc: '什么样的团队，都需要资源：至少有一个 7 是辅助神；',
        },
        {
          name: '※',
          desc: [
            { key: '总述', value: '不同行业重要性标准不同，一般类公司，至少需要以下配置' },
            { key: '8号', value: '统帅，一位，人力资源管理，董事会秘书，总经理助理' },
            { key: '6号', value: '执行官，1-2 位，负责落地执行' },
            { key: '9号', value: '如果是研发类机构：文化、研发等创新工作，至少需要一个 9' },
            { key: '偏重', value: '147 偏重物质；258 偏重人；789 偏重资源' },
          ],
        },
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
@import '@/components/awe-popup/awe-popup-list.scss';
</style>
