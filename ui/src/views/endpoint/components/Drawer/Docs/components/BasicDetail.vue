<!--
 项目基本信息
-->
<template>
  <a-descriptions class="basic-info">
    <template #title>
      <div class="title" @click="switchExpand">
        <RightOutlined class="icon" v-if="!expand"/>
        <DownOutlined class="icon" v-if="expand"/>
        <span class="text">基本信息</span>
      </div>
    </template>
    <template v-if="expand">
      <a-descriptions-item v-for="item in items" :key="item.label" :label="item.label">{{ item.value }}</a-descriptions-item>
    </template>
  </a-descriptions>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';

import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';


const props = defineProps({
  items: {
    required: true,
    type: Object,
  },
})


const emit = defineEmits(['ok', 'close', 'refreshList']);

const expand = ref(true);

function switchExpand() {
  expand.value = !expand.value;
}


</script>
<style lang="less" scoped>
.basic-info {
  position: relative;
  padding: 0 24px;

  :deep(.ant-descriptions-header) {
    padding-bottom: 16px;
  }

}


.title {
  position: absolute;
  cursor: pointer;

  .icon {
    position: absolute;
    left: -18px;
    top: 4px;
    font-size: 14px;

  }

  .text {

  }
}


</style>
