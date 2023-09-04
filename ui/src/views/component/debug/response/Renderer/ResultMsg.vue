<template>
    <div v-for="(item, index) in responseData"
           :key="index"
           :class="['item', customClass, getResultClass(item)]">

        <span v-if="item.resultStatus===ResultStatus.Pass">
          <CheckCircleOutlined />
        </span>

        <span v-if="item.resultStatus===ResultStatus.Fail">
          <CloseCircleOutlined />
        </span>&nbsp;

        <span>{{item.resultMsg}}</span>
      </div>
</template>

<script lang = "ts" setup>
import {defineProps} from 'vue';
import { CheckCircleOutlined, CloseCircleOutlined} from '@ant-design/icons-vue';

import {ResultStatus} from "@/utils/enum";
const props = defineProps(['responseData', 'customClass'])

const getResultClass = (item) => {
  return item.resultStatus===ResultStatus.Pass? 'pass':
      item.resultStatus===ResultStatus.Fail ? 'fail' : ''
}
</script>

<style lang="less" scoped>

.item {
    margin: 3px;
    padding: 5px;

    &.trans {
      background-color: transparent !important;
    }
    &.pass {
      color: #14945a;
      background-color: #F1FAF4;
    }
    &.fail {
      color: #D8021A;
      background-color: #FFECEE;
    }
  }
</style>