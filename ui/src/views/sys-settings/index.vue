<template>
  <div class="sys-setting-main">
    <Tab />
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import Tab from './components/common/Tab.vue';

</script>

<style scoped lang="less">
.sys-setting-main {
  background: #ffffff;
  height: 100%;
  overflow: hidden;
}

:deep(.ant-tabs-bar) {
  margin: 0;
}
</style>
