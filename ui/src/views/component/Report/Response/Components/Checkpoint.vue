<template>
  <div class="response-checkpoint-main">
    <div class="head">
      <a-row type="flex" class="item">
        <a-col flex="50px">编号</a-col>
        <a-col flex="60px">类型</a-col>
        <a-col flex="150px">变量 / 键值 / 表达式</a-col>
        <a-col flex="60px">运算符</a-col>
        <a-col flex="100px">数值</a-col>
        <a-col flex="80px">实际结果</a-col>
        <a-col flex="60px">状态</a-col>
      </a-row>
    </div>

    <div class="body">
      <a-row v-for="(item, idx) in checkpointsData" :key="idx" type="flex" class="item">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="60px">{{ t(item.type) }}</a-col>
        <a-col flex="150px">{{ item.type === CheckpointType.extractor ? item.extractorVariable : item.expression }} </a-col>
        <a-col flex="60px">{{ t(item.operator) }}</a-col>
        <a-col flex="100px">{{ item.value }}</a-col>
        <a-col flex="80px" style="width: 0; word-break: break-word;">
          {{ item.actualResult }}
        </a-col>

        <a-col flex="60px" :class="getResultCls(item.resultStatus)">
          {{ item.resultStatus ? t(item.resultStatus) : '' }}
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineProps} from "vue";
import {useI18n} from "vue-i18n";
import { PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {CheckpointType} from "@/utils/enum";
import {getResultCls} from "@/utils/dom"

const {t} = useI18n();

const props = defineProps({
  checkpointsData: {
    type: []
  }
})

</script>

<style lang="less">
.response-checkpoint-main {
}
</style>

<style lang="less" scoped>
.response-checkpoint-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    padding: 6px;
    height: calc(100% - 30px);
    overflow-y: auto;

    .item {
      .ant-col {
        padding: 0 3px;
        word-break: break-all;
      }
    }
  }
}
</style>