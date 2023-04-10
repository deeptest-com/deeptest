<template>
  <div class="container">
    <a-tabs v-model:activeKey="activeKey" @change="handleTabClick">
      <a-tab-pane key="enviroment" tab="环境管理">
      </a-tab-pane>

      <a-tab-pane key="data-pool" tab="数据池" >
      </a-tab-pane>
      <a-tab-pane key="service-setting" tab="服务管理">
        <!-- <ServiceSetting :params='router?.currentRoute.value?.query'/> -->
      </a-tab-pane>
    </a-tabs>
    <router-view></router-view>
  </div>
</template>
<script setup lang="ts">

import { ref,watch } from 'vue';
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { StateType as ProjectSettingStateType } from './store';
import ServiceSetting from './components/ServiceSetting/index.vue';
import EnvSetting from './components/EnvSetting/index.vue';
const expandedKeys = ref<string[]>(['0-0-0']);
const selectedKeys = ref<string[]>([]);

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

function handleTabClick(e: string) {
  console.log(e);
  if (e === 'enviroment') {
    router.push(`/project-setting/${e}/var`);
  } else {
    router.push(`/project-setting/${e}`);
  }
}

// 监听路由中是否携带参数，用于security模块
watch(() => {
  return router.currentRoute.value.query;
}, async (newVal) => {
 if(router.currentRoute.value.query?.firtab){
    const  firtab:any=router.currentRoute.value.query.firtab
    activeKey.value=firtab
  }
  if(router.currentRoute.value.path === '/project-setting/enviroment') {
    router.push('/project-setting/enviroment/var')
  }   
}, {
  immediate: true
})
</script>

<style scoped lang="less">
.container{
  margin: 16px;
  background: #ffffff;
  min-height: calc(100vh - 92px);
  min-width: 1200px;
}

:deep(.ant-tabs-bar) {
  margin: 0;
}
</style>
