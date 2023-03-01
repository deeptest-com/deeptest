<template>
  <div class="container">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="测试环境">测试环境</a-tab-pane>
      <a-tab-pane key="2" tab="数据池" force-render>数据池</a-tab-pane>
      <a-tab-pane key="3" tab="服务管理">
        <ServiceSetting/>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script setup lang="ts">
import { computed, defineComponent, reactive, Ref, ref, UnwrapRef } from 'vue';
import { CheckOutlined, EditOutlined,DownOutlined } from '@ant-design/icons-vue';
import ServiceSetting from './components/ServiceSetting';
import SchemeEditor from './components/SchemeEditor/index'

const expandedKeys = ref<string[]>(['0-0-0']);
const selectedKeys = ref<string[]>([]);

const activeKey = ref('3');

</script>

<style scoped lang="less">
.container{
  margin: 20px;
  background: #ffffff;
  min-height: 100vh;
}
.content{
  margin: 20px;
  .header{
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;
  }
}
</style>
