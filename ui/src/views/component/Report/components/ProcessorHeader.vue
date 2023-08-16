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
      <!-- ::::迭代次数：processor_loop_time -->
      <template v-if="record.processorType === 'processor_loop_time'">
        <span class="text">{{ 1 }}次</span>
<!--        <span class="text">{{ record.name }}</span>-->
      </template>

      <!-- ::::循环列表 -->
      <template v-if="record.processorType === 'processor_loop_in'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::循环直到 -->
      <template v-if="record.processorType === 'processor_loop_until'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::循环区间 -->
      <template v-if="record.processorType === 'processor_loop_range'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::跳出循环 -->
      <template v-if="record.processorType === 'processor_loop_break'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::跳出循环 -->
      <template v-if="record.processorType === 'processor_loop_break'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::条件分支-如果 -->
      <template v-if="record.processorType === 'processor_logic_if'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::否则 -->
      <template v-if="record.processorType === 'processor_logic_else'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::等待时间 -->
      <template v-if="record.processorType === 'processor_time_default'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::数据迭代 -->
      <template v-if="record.processorType === 'processor_data_default'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>


      <!-- ::::设置Cookie -->
      <template v-if="record.processorType === 'processor_cookie_set'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::清空Cookie -->
      <template v-if="record.processorType === 'processor_cookie_clear'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::设置变量 -->
      <template v-if="record.processorType === 'processor_variable_set'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::删除变量 -->
      <template v-if="record.processorType === 'processor_variable_clear'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::输出 -->
      <template v-if="record.processorType === 'processor_print_default'">
        <span class="text">输出：{{ record?.Detail?.['结果'] || '无输出数据' }} x</span>
      </template>

      <!-- ::::断言 -->
      <template v-if="record.processorType === 'processor_assertion_default'">
        <span class="text">{{ record.x }} x</span>
        <span class="text">{{ record.name }}</span>
        <a class="text" href="">文件地址</a>
        <span class="text">重复{{ record.name }}次</span>
      </template>

      <!-- ::::自定义代码 -->
      <template v-if="record.processorType === 'processor_custom_code'">
        <span class="text">{{ record?.Detail?.['结果'] || '无输出数据' }} x</span>
      </template>

    </div>


    <div class="status" v-if="showScenarioExecStatus.hasOwnProperty(record.processorType)">
      <span v-if="true" class="success">{{ showScenarioExecStatus[record.processorType]?.success }}</span>
      <span v-else class="fail">{{ showScenarioExecStatus[record.processorType]?.fail }}</span>
    </div>


    <div class="right" @click.stop="clickMore">
      详情<RightOutlined/>
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
}
</style>
