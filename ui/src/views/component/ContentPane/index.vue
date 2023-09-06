<!--
  适用于 左侧目录树 + 右边区域表格筛选，且左侧目录树可伸缩
-->
<template>
  <div class="container" :style="containerStyle || {}">
    <div class="content">
      <multipane class="vertical-panes" layout="vertical" @paneResize="handlePaneResize">
        <div ref="paneLeft" :class="['pane', 'left', !isFold && 'unfold']">
          <slot name="left"></slot>
        </div>
        <multipane-resizer />
        <div class="pane right" :style="{ flexGrow: 1  }">
          <slot name="right"></slot>
          <div v-if="showExpand" class="expand-icon" @click="toggle">
            <menu-fold-outlined v-if="isFold" />
            <menu-unfold-outlined v-else />
          </div>
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
};

const handlePaneResize = (...args) => {
  isFold.value = true;
};

</script>

<style lang="less" scoped>
.container {
  margin: 16px;
  background: #ffffff;
  height: calc(100vh - 96px);

  :deep(.ant-pagination) {
    margin-right: 24px;
  }
  .content {
    display: flex;
    width: 100%;
    position: relative;
    height: 100%;

    .left {
      overflow-y: scroll;
      position: relative;
      min-width: 150px;
      width: 300px;
      max-width: 600px;

      &.unfold {
        width: 0 !important;
        min-width: 0 !important;
      }

    }

    .right {
      flex: 1;
      overflow: scroll;
      position: relative;

      &:has(.expand-icon:hover) {
        overflow: unset;
        z-index: 2;
      }

      .expand-icon {
        position: absolute;
        top: 6px;
        left: -8px;
        font-size: 20px;
        color: #1890ff;
      }
    }
  }
}

.vertical-panes {
  height: 100%;
  width: 100%;
}

</style>
