<template>
  <view class="secret1-content">
    <scroll-view class="secret1-scroll-wrapper" scroll-y>
      <view v-for="(item, index) in list" :key="index" class="secret1-item">
        <view class="secret1-name">{{ item.name }}</view>
        <view v-if="typeof item.desc === 'string'" class="secret1-desc">{{ item.desc }}</view>
        <view v-else-if="typeof item.desc === 'object' && item.desc.length" class="secret1-desc">
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
          desc: '自出生开始计算阴阳：阳格周期内，阳格为显性（阴格为个体无意识）；阴格周期内，阴格为显性（阳格为个体无意识）；能开发出：松果体、心轮能力的人，可依靠呼吸调整阴阳格；如果没有调整意识系统能力者，按照程序显示的计算。',
        },
        {
          name: 2,
          desc: '如果计算周期内：生孩子，按照生孩子日期所在格，重新按照第0年，重新累计本周期时间。',
        },
        {
          name: 3,
          desc: '家族遗传 RNA 是灵魂部分；这部分遵循双方家族遗传惯例，部分缺的孩子会缺，部分多的孩子也会继承；父亲遗传 61.8% 情感部分，母亲遗传 61.8% 智慧部分；灵魂簇遵循男女男女继承规则；家族灵魂簇具有继承性；',
        },
        {
          name: 4,
          desc: '通常一个灵魂需要 80-120 次学会人类载具课程；如果：没有修完情绪体，则需要重复修，直到突破能量限制；修结束的密码——以“爱”的方式回馈对方的情绪变化；',
        },
        {
          name: 5,
          desc: '基本规则：4 克 369；4 为独裁者，与 2 合，其余克。家庭财富：至少有 7。同一灵魂簇为 6 和，不同灵魂簇按时区关系，分为冲、平、和。',
        },
        {
          name: 6,
          desc: '情义，事业属于义的部分；情绪体上，做事需要合伙人，每个人性格担任的角色不同；事业合伙人关系——需要 3 和。',
        },
        {
          name: 7,
          desc: '什么样的团队，都需要资源：至少有一个 7 是辅助神；',
        },
        {
          name: 8,
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

        .secret1-desc {
          text-align: justify;

          .secret1-desc-item {
            flex-direction: row;

            .secret1-desc-value {
              text-align: justify;
            }
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

        .secret1-desc {
          text-align: justify;

          .secret1-desc-item {
            flex-direction: row-reverse;

            .secret1-desc-value {
              text-align: justify;
            }
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
