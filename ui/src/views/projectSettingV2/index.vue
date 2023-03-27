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
        <ServiceSetting/>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script setup lang="ts">
import { computed, defineComponent, reactive, Ref, ref, UnwrapRef } from 'vue';
import { CheckOutlined, EditOutlined,DownOutlined } from '@ant-design/icons-vue';
import ServiceSetting from './components/ServiceSetting/index.vue';
import EnvSetting from './components/EnvSetting/index.vue';
import SchemaEditor from '@/components/SchemaEditor/index.vue'
const expandedKeys = ref<string[]>(['0-0-0']);
const selectedKeys = ref<string[]>([]);

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
//.content{
//  margin: 20px;
//  .header{
//    display: flex;
//    align-items: center;
//    justify-content: space-between;
//    margin-bottom: 8px;
//  }
//}
</style>
