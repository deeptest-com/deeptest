<template>
  <div class="processor-header">
    <div class="left" :class="{'hide-arrow' : !showArrowScenarioType.includes(record.processorType)}">
      <!-- ::::通用的场景图标 和 场景名称 -->
      <IconSvg :type="DESIGN_TYPE_ICON_MAP[record.processorType]" class="processor-icon-svg"/>
      <a-typography-text strong v-if="!record.processorType.includes('processor_logic_')">
        {{ scenarioTypeMapToText[record.processorType] }}
        {{ record.name ? ` - ${record.name}`: '' }}
      </a-typography-text>
      <a-typography-text
          strong
          v-else
          style=" display: inline-block; text-align: left;margin-right: 4px;"
          :type="record.processorType === 'processor_logic_if' ? 'success' : 'danger'">
        {{
          record.processorType === 'processor_logic_if' ? 'if' : 'else'
        }}
        {{ record.name ? ` - ${record.name}`: '' }}
      </a-typography-text>
    </div>

    <div class="summary">
      <!-- ::::数据迭代 -->
      <template v-if="record.processorType === 'processor_data_default'">
        <p class="text">从 <a  :href="downloadUrl">{{`${detail?.url}`}}</a> 中取值赋给变量<code>{{ `${detail?.variableName}` }}</code>， 且重复 <code>{{ `${detail?.repeatTimes}` }}</code> 次
        </p>
      </template>
      <!-- ::::迭代次数：processor_loop_time -->
      <template v-if="record.processorType === 'processor_loop_time'">
        <p class="text">迭代<code>{{ `${detail?.times}` }}</code>次</p>
      </template>
      <!-- ::::循环列表 -->
      <template v-if="record.processorType === 'processor_loop_in'">
        <p class="text">从 <code>{{ `${detail?.list}` }}</code>中 <code>{{ `${!detail?.isRand ? '顺序' : '随机'}` }}</code>取值赋给变量 <code>{{ `${detail?.variableName}` }}</code></p>
      </template>
      <!-- ::::循环直到 -->
      <template v-if="record.processorType === 'processor_loop_until'">
        <p class="text">迭代直到 <code>{{ `${detail?.untilExpression}` }}</code> 为  <code>{{ `true` }}</code>时</p>
      </template>
      <!-- ::::循环区间 -->
      <template v-if="record.processorType === 'processor_loop_range'">
        <p class="text">从区间 <code>{{ `[${detail?.range}]` }}</code> 中 <code>{{ `${!detail?.isRand ? '顺序' : '随机'}` }}</code> 取值赋给变量 <code>{{ `${detail?.variableName}` }}</code></p>
      </template>
      <!-- ::::跳出循环 -->
      <template v-if="record.processorType === 'processor_loop_break'">
        <p class="text">满足条件 <code>{{ `${detail?.expression}` }}</code> 时， 跳出迭代</p>
      </template>
      <!-- ::::条件分支-如果 -->
      <template v-if="record.processorType === 'processor_logic_if'">
        <p class="text">如果 <code>{{ `${detail?.expression}` }}</code> 为 <code>true</code></p>
      </template>
      <!-- ::::否则 -->
      <template v-if="record.processorType === 'processor_logic_else'">
        <p class="text">否则</p>
      </template>
      <!-- ::::等待时间 -->
      <template v-if="record.processorType === 'processor_time_default'">
        <p class="text">等待 <code>{{ `${detail?.sleepTime}` }}</code> 秒</p>
      </template>
      <!-- ::::设置Cookie -->
      <template v-if="record.processorType === 'processor_cookie_set'">
        <p class="text">设置 <code>{{ `${detail?.cookieName}` }}</code> 为 <code>{{ `${detail?.variableValue}` }}</code></p>
      </template>
      <!-- ::::清空Cookie -->
      <template v-if="record.processorType === 'processor_cookie_clear'">
        <p class="text">清除 <code>{{ `${detail?.cookieName}` }}</code></p>
      </template>
      <!-- ::::设置变量 -->
      <template v-if="record.processorType === 'processor_variable_set'">
        <p class="text">设置变量 <code>{{ `${detail?.variableName}` }}</code>为<code>{{ `${detail?.variableValue}` }}</code></p>
      </template>
      <!-- ::::删除变量 -->
      <template v-if="record.processorType === 'processor_variable_clear'">
        <p class="text">删除变量<code>{{ `${detail?.variableName}` }}</code></p>
      </template>
      <!-- ::::输出 -->
      <template v-if="record.processorType === 'processor_print_default'">
        <p class="text">输出：{{ `${detail?.result}` }} </p>
      </template>
      <!-- ::::断言 -->
      <template v-if="record.processorType === 'processor_assertion_default'">
        <p class="text">断言表达式：<code>{{ `${detail?.expression}` }}</code></p>
      </template>
      <!-- ::::自定义代码 -->
      <template v-if="record.processorType === 'processor_custom_code'">
        <p class="text"><code>{{ `${detail?.content}` }}</code></p>
      </template>
    </div>

    <div class="status" v-if="showScenarioExecStatus.hasOwnProperty(record.processorType)">
      <span v-if="detail?.result" class="success">{{ showScenarioExecStatus[record.processorType]?.success }}</span>
      <span v-else class="fail">{{ showScenarioExecStatus[record.processorType]?.fail }}</span>
    </div>



    <div class="right" @click.stop="clickMore">
      详情
      <RightOutlined/>
    </div>
    <LogContentDrawer
        :data="data"
        :visible="visible"
        @onClose="visible = false"/>
  </div>
</template>
<script setup lang="ts">
import {defineProps, h, defineEmits, computed, toRefs, ref} from 'vue';
import {RightOutlined, LoadingOutlined, ExclamationCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {responseCodes} from '@/config/constant';
import IconSvg from "@/components/IconSvg";
import {formatWithSeconds} from '@/utils/datetime';
import LogContentDrawer from './LogContentDrawer/index.vue';
import {
  scenarioTypeMapToText,
  showArrowScenarioType,
  DESIGN_TYPE_ICON_MAP,
  showScenarioExecStatus,
} from "@/views/scenario/components/Design/config";

const props = defineProps(['record'])

const emits = defineEmits(['more']);
const visible = ref(false);
const data:any = computed(() => {
  return props.record;
})

const downloadUrl = computed(() => {
  return `${window.location.origin}/${detail?.value?.url}`
})

const detail = computed(() => {
  return JSON.parse(props.record?.detail || '{}');
})

function clickMore() {
  visible.value = true;
  // emits('more', props.record);
}

</script>
<style scoped lang="less">
.processor-header {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .processor-icon-svg {
    display: inline-block;
    margin-right: 4px;
  }

  .left {
    margin-right: 8px;
    width: 100px;
    white-space: nowrap; /* 禁止换行 */
    overflow: hidden; /* 超出部分隐藏 */
    text-overflow: ellipsis; /* 显示省略号 */
    &.hide-arrow {
      margin-left: 28px;
    }

  }

  .summary {
    flex: 1;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    margin-right: 16px;
    white-space: nowrap; /* 禁止换行 */
    overflow: hidden; /* 超出部分隐藏 */
    text-overflow: ellipsis; /* 显示省略号 */
  }

  .right {
    font-weight: normal;
  }

  .text {
    display: inline-block;
    margin: 0 2px;
    white-space: nowrap; /* 禁止换行 */
    overflow: hidden; /* 超出部分隐藏 */
    text-overflow: ellipsis; /* 显示省略号 */
  }

  .status {
    width: 60px;
    text-align: left;
    //min-width: 40px;
    height: 22px;
    font-size: 14px;
    border-radius: 2px;
    line-height: 22px;
    margin-right: 16px;

    .success {
      background: #E6FFF4;
      color: #04C495;

    }

    .fail {
      background: #FFF2F0;;
      color: #F63838;
    }

  }

  code {
    margin: 0 1px;
    padding: 0.2em 0.4em;
    font-size: .9em;
    background: #f2f4f5;
    border: 1px solid #f0f0f0;
    border-radius: 3px;
  }
}
</style>
