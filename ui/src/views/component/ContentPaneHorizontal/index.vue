<!--
  适用于 上下 两个面板的布局
-->
<template>
  <div class="container" :style="containerStyle">
    <div class="content">
      <multipane class="vertical-panes" layout="horizontal" @paneResizeStop="handlePaneResizeTop">
        <div class="pane top" :style="{ minHeight: '46px', height: '60%'}">
          <slot name="top"/>
        </div>
        <multipane-resizer/>
        <div class="pane bottom" :style="{ minHeight: '36px', flexGrow:1}">
          <slot name="bottom"/>
        </div>
      </multipane>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps,ref } from 'vue'
import {Multipane, MultipaneResizer} from '@/components/Resize/index';
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
//TODO 备注：使用 topPercent 和 bottomPercent 有问题，拖动后会再次重置高度，后面再解决，另外需要加 FlexGrow,自动填充高度，否则也会有问题
const props = defineProps({
  containerStyle: {
    type: Object,
    required: false,
    default: {} as Object,
  },
  topPercent: {
    type: String,
    required: false,
    default: '60%',
  },
  bottomPercent: {
    type: String,
    required: false,
    default: '40%',
  },
})

const handlePaneResizeTop = () => {
  bus.emit(settings.paneResizeTop);
}
</script>

<style lang="less" scoped>
.container {
  margin: 16px;
  background: #ffffff;
  //height: 100%;
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
    .top {
      height: 0;
      overflow-y: scroll;
      position: relative;
    }

    .bottom {
      height: 0;
      overflow-y: scroll;
    }
  }
}

.vertical-panes {
  height: 100%;
  width: 100%;
}

</style>

