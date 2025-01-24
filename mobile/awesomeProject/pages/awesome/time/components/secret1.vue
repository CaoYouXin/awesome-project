<template>
  <view class="secret1-content">
    <scroll-view class="secret1-scroll-wrapper" scroll-y>
      <view v-for="(item, index) in list" :key="index" class="secret1-item">
        <view class="secret1-name">{{ item.name }}</view>
        <view class="secret1-desc">
          <view v-for="(desc, idx) in item.desc" :key="idx" class="secret1-desc-item">
            <view class="secret1-desc-key" style="width: 2em">{{ desc.key }}</view>
            <view class="secret1-desc-key">&nbsp;:&nbsp;</view>
            <view class="secret1-desc-value">{{ desc.value }}</view>
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
          name: 1,
          desc: [
            { key: '肺', value: '脏--阴' },
            { key: '天赋', value: '金刚手--疗愈者' },
            { key: '能力', value: '自信、独立，阳性能量很强，领导力' },
            { key: '图腾', value: '玉皇大帝、牛魔王' },
            { key: '负面', value: '执着，顽固' },
            { key: '缺漏', value: '长得慢，也容易是单亲家庭，缺母 80% 概率' },
            { key: '关系', value: '负责脑波大于 9.7Hz（白天醒着）时，统管七魄' },
          ],
        },
        {
          name: 2,
          desc: [
            { key: '大肠', value: '腑--阳' },
            { key: '天赋', value: '同理心、感受物质震动、把握细节强' },
            { key: '能力', value: '服务意识强，具有同理心，细节能力强、组织文字能力好' },
            { key: '图腾', value: '服务官' },
            { key: '负面', value: '服务过量，啰嗦；传递负面能量' },
            { key: '缺漏', value: '不会收拾东西、杂乱；没有时间观念，不能吃过量物质性食物' },
          ],
        },
        {
          name: 3,
          desc: [
            { key: '心', value: '脏--阴' },
            { key: '天赋', value: '大心神、心想事成、慈悲；跨纬度；远程疗育' },
            { key: '能力', value: '慈悲、宽恕、善解人意' },
            { key: '图腾', value: '大心神' },
            { key: '负面', value: '斤斤计较、算计、小心眼、财迷' },
            { key: '缺漏', value: '无法理解别人想法、对声音反应迟钝' },
            {
              key: '关系',
              value: '家庭内与 4、7、8 冲突；7 岁前魂容易丢，导致呼吸障碍；7 岁后，魂通过心进入松果体，开启特异功能',
            },
          ],
        },
        {
          name: 4,
          desc: [
            { key: '肝', value: '脏--阴' },
            { key: '天赋', value: '高效率执行、雷厉风行' },
            { key: '能力', value: '独断、效率、缺乏耐心、容易怒；容易成为商品行业龙头管理人' },
            { key: '图腾', value: '独裁者' },
            { key: '负面', value: '独裁、不讲道理、愤怒、指责' },
            { key: '缺漏', value: '无法化解血毒，大肠亚硝酸毒会影响肝，不能饮酒、不能吃红肉' },
            {
              key: '关系',
              value:
                '家庭内与 29112 不冲突，其它冲突；7 的格可以管理 4 的格。负责脑波小于 7.23Hz（晚上睡觉）时，统管三魂（爽灵、幽精、胎光）',
            },
          ],
        },
        {
          name: 5,
          desc: [
            { key: '胆', value: '腑--阳' },
            { key: '天赋', value: '上通下达、执行力' },
            { key: '能力', value: '胆量、中正、决断' },
            { key: '图腾', value: '中正官' },
            {
              key: '负面',
              value: '焦虑、优柔寡断，太爱讲道理，要结果忽略过程；忽略情感；胆量过时：胆大包天，不顾他人感受',
            },
            { key: '缺漏', value: '无法消化食物，不能吃含油脂的食物、干果；决断力缺乏' },
            {
              key: '关系',
              value: '家庭内与 4、7、8 冲突；诚服与 8 号人格',
            },
          ],
        },
        {
          name: 6,
          desc: [
            { key: '肾', value: '脏--阴' },
            { key: '天赋', value: '千里耳、松果体成像；耳朵可通过松果体，开启特异功能' },
            { key: '能力', value: '执行、显化、精力旺盛、耐久力强' },
            { key: '图腾', value: '法官' },
            {
              key: '负面',
              value: '恐惧、害怕；色、无法控制情绪。色弱人比例较高',
            },
            { key: '缺漏', value: '脑髓不足或空虚，易发生健忘、精神不振或未老先衰、精力不振。听力不好、尿毒症风险' },
            {
              key: '关系',
              value: '家庭内与 4 难相处，适合做老师，精通技巧',
            },
          ],
        },
        {
          name: 7,
          desc: [
            { key: '胃', value: '腑--阳，仓廪，地表的财神财库' },
            { key: '天赋', value: '度量、承载' },
            { key: '能力', value: '目标性强' },
            { key: '图腾', value: '牛魔王' },
            {
              key: '负面',
              value: '达到自己的标准才执行',
            },
            { key: '缺漏', value: '容易胃底的幽门螺旋杆菌感染财富运，亲密无间容易出问题' },
            {
              key: '关系',
              value: '家庭内，达到自己的标准才执行（容易各种点炮）',
            },
          ],
        },
        {
          name: 8,
          desc: [
            { key: '脾', value: '脏--阴' },
            { key: '天赋', value: '统帅、管理、显化' },
            { key: '能力', value: '管理能力、统筹规划、喜欢指挥、音乐天赋' },
            { key: '图腾', value: '统帅' },
            {
              key: '负面',
              value: '抱怨、委屈',
            },
            { key: '缺漏', value: '忌口甜食及含糖食物、没注意、不能管理团队' },
            {
              key: '关系',
              value: '统领、统帅，家庭内与 4、7 冲突',
            },
          ],
        },
        {
          name: 9,
          desc: [
            { key: '膀胱', value: '腑--阳' },
            { key: '天赋', value: '天才、通松果体' },
            { key: '能力', value: '单纯、服务他人，自愈能力强、下载信息、化繁为简' },
            { key: '图腾', value: '天使' },
            {
              key: '负面',
              value: '炫耀、太务虚、看他人眼色、做事不落地',
            },
            { key: '缺漏', value: '生长发育缓慢，学习、接受知识、信息能力弱' },
            {
              key: '关系',
              value: '家庭内与 4、7、8 冲突',
            },
          ],
        },
        {
          name: 0,
          desc: [
            { key: '小肠', value: '腑--阳' },
            { key: '天赋', value: '道、转化、先天认知能力、直觉、方向感' },
            { key: '能力', value: '先天具有自我学习能力，不用老师；对事务的理解精准；归纳总结能力强；判断力强；隐忍' },
            { key: '图腾', value: '0、道、灵' },
            {
              key: '负面',
              value: '哀悯过度，精于算计、太刨根问底、强迫症',
            },
            { key: '缺漏', value: '没有转化能力，理解能力差、接受知识、信息能力弱' },
            {
              key: '关系',
              value: '家庭内与 4、7、8 冲突',
            },
          ],
        },
      ],
    };
  },
};
</script>

<style lang="scss" scoped>
.secret1-content {
  padding: $to-border;

  .secret1-scroll-wrapper {
    height: 50vh;

    .secret1-item {
      border: solid 1rpx white;
      border-radius: 50rpx;
      min-height: 100rpx;
      display: flex;
      justify-content: flex-start;
      align-items: flex-start;

      &:nth-of-type(odd) {
        background-color: #f39999;
        flex-direction: row;

        .secret1-name {
          border-color: #f39999;
          box-shadow: inset 0 0 20rpx #f39999;
        }

        .secret1-desc-item {
          flex-direction: row;

          .secret1-desc-value {
            text-align: left;
          }
        }
      }

      &:nth-of-type(even) {
        background-color: #8ab5f5;
        flex-direction: row-reverse;

        .secret1-name {
          border-color: #8ab5f5;
          box-shadow: inset 0 0 20rpx #8ab5f5;
        }

        .secret1-desc-item {
          flex-direction: row-reverse;

          .secret1-desc-value {
            text-align: right;
          }
        }
      }

      & + .secret1-item {
        margin-top: $near-border;
      }

      .secret1-name {
        flex-grow: 0;
        flex-shrink: 0;
        width: 100rpx;
        height: 100rpx;
        border-radius: 50%;
        background-color: white;
        border: $near-border solid transparent;
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .secret1-desc {
        flex-grow: 1;
        width: 1rpx;
        color: white;
        padding: $near-border $to-border;

        .secret1-desc-item {
          display: flex;
          justify-content: flex-start;
          align-items: flex-start;

          & + .secret1-desc-item {
            margin-top: $near-border;
          }

          .secret1-desc-key {
            flex-grow: 0;
            flex-shrink: 0;
            color: gold;
            text-align: center;
          }
        }
      }
    }
  }
}
</style>
