<template>
  <div class="processor-header">
    <div class="left" :class="{'hide-arrow' : !showArrowScenarioType.includes(record.processorType)}">
      <!-- ::::通用的场景图标 和 场景名称 -->
      <IconSvg :type="DESIGN_TYPE_ICON_MAP[record.processorType]" class="processor-icon-svg"/>
      <a-typography-text strong v-if="!record.processorType.includes('processor_logic_')">
        {{ scenarioTypeMapToText[record.processorType] }}
      </a-typography-text>
      <a-typography-text
          strong
          v-else
          style=" display: inline-block; text-align: left;margin-right: 4px;"
          :type="record.processorType === 'processor_logic_if' ? 'success' : 'danger'">{{
          record.processorType === 'processor_logic_if' ? 'if' : 'else'
        }}
      </a-typography-text>
    </div>

    <div class="summary">
      <!-- ::::数据迭代 -->
      <template v-if="record.processorType === 'processor_data_default'">
        <span class="text">
          从<a class="text" :href="downloadUrl">文件名</a>中{{ `随机` }}读取变量<code>{{ `a` }}</code>， 且重复 <code>{{ 3 }}</code> 次
        </span>
      </template>
      <!-- ::::迭代次数：processor_loop_time -->
      <template v-if="record.processorType === 'processor_loop_time'">
        <span class="text">迭代<code>{{ 3 }}</code>次</span>
      </template>
      <!-- ::::循环列表 -->
      <template v-if="record.processorType === 'processor_loop_in'">
        <span class="text">在 <code>{{ `a,b,c,d` }}</code> 中 {{ `随机` }} 读取变量 <code>{{ `a` }}</code></span>
      </template>

      <!-- ::::循环直到 -->
      <template v-if="record.processorType === 'processor_loop_until'">
        <span class="text">循环直到 <code>{{ `a-b` }}</code> 为  <code>{{ `true` }}</code></span>
      </template>

      <!-- ::::循环区间 -->
      <template v-if="record.processorType === 'processor_loop_range'">
        <span class="text">在区间 <code>{{ `[a-b]` }}</code> 中 {{ `随机` }} 读取变量 <code>{{ `a` }}</code></span>
      </template>

      <!-- ::::跳出循环 -->
      <template v-if="record.processorType === 'processor_loop_break'">
        <span class="text">满足条件 <code>{{ `a-b>0` }}</code> 时， 跳出迭代</span>
      </template>

      <!-- ::::条件分支-如果 -->
      <template v-if="record.processorType === 'processor_logic_if'">
        <span class="text">如果 <code>{{ `a-b` }}</code> 为 <code>true</code></span>
      </template>

      <!-- ::::否则 -->
      <template v-if="record.processorType === 'processor_logic_else'">
        <span class="text">否则</span>
      </template>

      <!-- ::::等待时间 -->
      <template v-if="record.processorType === 'processor_time_default'">
        <span class="text">等待 <code>{{ `2` }}</code> 秒</span>
      </template>


      <!-- ::::设置Cookie -->
      <template v-if="record.processorType === 'processor_cookie_set'">
        <span class="text">设置 <code>{{ `Cookie1` }}</code> 为 <code>{{ `122` }}</code></span>
      </template>

      <!-- ::::清空Cookie -->
      <template v-if="record.processorType === 'processor_cookie_clear'">
        <span class="text">清除 <code>{{ `Cookie1` }}</code></span>
      </template>

      <!-- ::::设置变量 -->
      <template v-if="record.processorType === 'processor_variable_set'">
        <span class="text">设置变量 <code>{{ `x` }}</code>为<code>{{ `1` }}</code></span>
      </template>

      <!-- ::::删除变量 -->
      <template v-if="record.processorType === 'processor_variable_clear'">
        <span class="text">清除变量<code>{{ `x` }}</code></span>
      </template>

      <!-- ::::输出 -->
      <template v-if="record.processorType === 'processor_print_default'">
        <span class="text">输出：{{ record?.Detail?.['结果'] || '无输出数据' }} x</span>
      </template>

      <!-- ::::断言 -->
      <template v-if="record.processorType === 'processor_assertion_default'">
        <span class="text">断言表达式：<code>{{ `a - x` }}</code> x</span>
      </template>

      <!-- ::::自定义代码 -->
      <template v-if="record.processorType === 'processor_custom_code'">
        <span class="text"><code>{{ `hello worold console.log` }}</code></span>
      </template>
    </div>


    <div class="status" v-if="showScenarioExecStatus.hasOwnProperty(record.processorType)">
      <span v-if="true" class="success">{{ showScenarioExecStatus[record.processorType]?.success }}</span>
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
const data = computed(() => {
  return props.record;
})

const downloadUrl = computed(() => {
  return `${window.location.origin}/${props?.record?.datail?.path}`
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

    &.hide-arrow {
      margin-left: 28px;
    }

  }

  .summary {
    flex: 1;
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }

  .right {
    font-weight: normal;
  }

  .text {
    display: inline-block;
    margin: 0 2px;
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
