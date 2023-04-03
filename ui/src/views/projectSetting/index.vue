<template>
  <div class="container">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="环境管理">
        <EnvSetting/>
      </a-tab-pane>

      <a-tab-pane key="2" tab="数据池" >
        <div style="height: 90vh;">
        </div>
      </a-tab-pane>
      <a-tab-pane key="3" tab="服务管理">
        <ServiceSetting :params='router?.currentRoute.value?.query'/>
      </a-tab-pane>
    </a-tabs>
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


const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const activeKey = ref('1');
const router = useRouter();


getUserList();
async function getUserList() {
  await store.dispatch('ProjectSetting/getUserOptionsList')
}

// 监听路由中是否携带参数，用于security模块
watch(() => {
  return router.currentRoute.value.query;
}, async (newVal) => {
 if(router.currentRoute.value.query?.firtab){
    const  firtab:any=router.currentRoute.value.query.firtab
    activeKey.value=firtab
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
