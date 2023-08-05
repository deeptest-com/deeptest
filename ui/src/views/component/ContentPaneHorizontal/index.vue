<!--
  适用于 上下 两个面板的布局
-->
<template>
  <div class="container" :style="containerStyle || {}">
    <div class="content">
      <multipane class="vertical-panes" layout="horizontal">
        <div class="pane top" :style="{ minHeight: '150px', height: '300px', maxHeight: '700px' }">
          <slot name="top"/>
        </div>
        <multipane-resizer/>
        <div class="pane bottom" :style="{ flexGrow: 1  }">
          <slot name="bottom"/>
        </div>
      </multipane>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import {Multipane, MultipaneResizer} from '@/components/Resize/index';
const props = defineProps(['containerStyle'])
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
    .top {
      overflow-y: scroll;
      position: relative;
    }

    .bottom {
      flex: 1;
      overflow-y: scroll;
    }
  }
}

.vertical-panes {
  height: 100%;
  width: 100%;
}

</style>

