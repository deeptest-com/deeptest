<template>
  <HomeLayout>
    <div class="sys-setting-main">
      <RouteTab />
      <router-view></router-view>
    </div>
  </HomeLayout>
</template>

<script setup lang="ts">
import RouteTab from './components/common/Tab.vue';
import HomeLayout from "@/layouts/HomeLayout.vue";

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
