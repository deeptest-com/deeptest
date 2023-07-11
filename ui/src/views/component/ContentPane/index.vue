<!--
  适用于 左侧目录树 + 右边区域表格筛选，且左侧目录树可伸缩
-->
<template>
  <div class="container">
    <div class="content">
      <multipane class="vertical-panes" layout="vertical">
        <div class="pane left" :style="{ minWidth: '150px', width: '300px', maxWidth: '600px' }">
          <Tree/>
          <slot name="left"></slot>
        </div>
        <multipane-resizer/>
        <div class="pane right" :style="{ flexGrow: 1 }">
          <slot name="right"></slot>
        </div>
      </multipane>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {Multipane, MultipaneResizer} from '@/components/Resize/index';
const {t} = useI18n();
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

    .left {
      height: calc(100vh - 80px);
      overflow-y: scroll;
      position: relative;
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
