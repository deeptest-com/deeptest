<!--
  适用于 左侧目录树 + 右边区域表格筛选，且左侧目录树可伸缩
-->
<template>
  <div class="container" :style="containerStyle || {}">
    <div class="content">
      <multipane class="vertical-panes" layout="vertical" @paneResize="handlePaneResize">
        <div ref="paneLeft" :class="['pane', 'left', !isFold && 'unfold']" :style="{ minWidth: '150px', width: '300px', maxWidth: '600px' }">
          <slot name="left"></slot>
        </div>
        <multipane-resizer v-if="isFold"/>
        <div class="pane right" :style="{ flexGrow: 1  }">
          <slot name="right"></slot>
        </div>
        <div v-if="showExpand" class="expand-icon" @click="toggle" :style="{ marginLeft: foldIconLeft }">
          <menu-fold-outlined v-if="isFold" />
          <menu-unfold-outlined v-else />
        </div>
      </multipane>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import { defineProps, nextTick, ref } from 'vue'
import { MenuFoldOutlined, MenuUnfoldOutlined } from "@ant-design/icons-vue";

import {Multipane, MultipaneResizer} from '@/components/Resize/index';

const {t} = useI18n();
const props = defineProps(['containerStyle', 'showExpand'])
const isFold = ref(true);
const paneLeft = ref();
const foldIconLeft = ref('288px');

const toggle = async () => {
  isFold.value = !isFold.value;
  await nextTick();
  const width = paneLeft.value ? paneLeft.value.getBoundingClientRect().width : 0;
  foldIconLeft.value = !isFold.value ? '-12px' : `${width - 12}px`;
};

const handlePaneResize = (...args) => {
  const width = Number(args[2].split('px')[0]);
  foldIconLeft.value = `${(width < 150 ? 150 : width > 300 ? 300 : width) - 12}px`;
};

</script>

<style lang="less" scoped>
.container {
  margin: 16px;
  background: #ffffff;
  height: calc(100vh - 96px);
  overflow: hidden;
  :deep(.ant-pagination) {
    margin-right: 24px;
  }
  .content {
    display: flex;
    width: 100%;
    position: relative;
    height: 100%;

    .expand-icon {
      position: fixed;
      font-size: 20px;
    }
    .left {
      overflow-y: scroll;
      position: relative;

      &.unfold {
        width: 0 !important;
        min-width: 0 !important;
      }

    }

    .right {
      flex: 1;
      overflow: scroll;
    }
  }
}

.vertical-panes {
  height: 100%;
  width: 100%;
}

</style>
