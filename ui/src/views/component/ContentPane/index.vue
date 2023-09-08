<!--
  适用于 左侧目录树 + 右边区域表格筛选，且左侧目录树可伸缩
-->
<template>
  <div class="container" :style="containerStyle || {}">
    <div :class="['content', showExpand && 'expand-content']">
      <multipane class="vertical-panes" layout="vertical" @paneResize="handlePaneResize">
        <div ref="paneLeft" :class="['pane', 'left', !isFold && 'unfold']">
          <slot name="left"></slot>
        </div>
        <multipane-resizer />
        <div :class="['pane', 'right', !isFold && 'unfold']" :style="{ flexGrow: 1  }">
          <slot name="right"></slot>
          <div v-if="showExpand" class="expand-icon" @click="toggle">
            <img :src="PutAway" />
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
import PutAway from '@/assets/images/put-away.png';

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

    &.expand-content {
      .right {
        overflow: unset;
      }
    }

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
      z-index: 2;

      &.unfold {
        overflow: scroll;

        .expand-icon {
          transform: rotate(180deg);
          left: -14px;
        }
      }

      &:has(.expand-icon:hover) {
        overflow: unset;
        z-index: 2;
      }

      .expand-icon {
        position: absolute;
        top: 6px;
        left: -16px;
        width: 30px;
        height: 30px;

        img {
          width: 100%;
          height: 100%;
          image-rendering: pixelated; 
        }
      }
    }
  }
}

.vertical-panes {
  height: 100%;
  width: 100%;
}

</style>
