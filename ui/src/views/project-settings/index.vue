<template>
  <div class="container">
    <RouteTab />
    <router-view></router-view>
  </div>
</template>
<script setup lang="ts">

import { ref, watch } from 'vue';
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { StateType as ProjectSettingStateType } from './store';
import RouteTab from './components/common/Tab.vue';

const router = useRouter();

const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const activeKey = ref(setActiveKey());


getUserList();
async function getUserList() {
  await store.dispatch('ProjectSetting/getUserOptionsList')
}

function setActiveKey() {
  const routePath = router.currentRoute.value.path;
  const res = routePath.split('/');
  return res[2];
}

// 监听路由中是否携带参数，用于security模块
watch(() => {
  return router.currentRoute.value.query;
}, async () => {
  if (router.currentRoute.value.path === '/project-setting/environment') {
    router.push('/project-setting/environment/var')
    store.dispatch('ProjectSetting/setEnvDetail', null);
    activeKey.value = 'environment';
  }
}, {
  immediate: true
})
</script>

<style scoped lang="less">
.container {
  background: #ffffff;
  height: 100%;
  overflow: hidden;
}

:deep(.ant-tabs-bar) {
  margin: 0;
}
</style>
