<template>
  <div class="container">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="服务管理">
        <ServiceSetting/>
      </a-tab-pane>

      <a-tab-pane key="2" tab="环境管理">
        <EnvSetting/>
      </a-tab-pane>

      <a-tab-pane key="3" tab="数据池" >
        <div style="height: 90vh;">
        </div>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import ServiceSetting from './components/ServiceSetting/index.vue';
import EnvSetting from './components/EnvSetting/index.vue';

const activeKey = ref('1');

</script>

<style scoped lang="less">
.container{
  margin: 16px;
  background: #ffffff;
  min-height: calc(100vh - 92px);
}
::v-deep{
  .ant-tabs-bar{
    margin: 0;
  }
}
</style>
