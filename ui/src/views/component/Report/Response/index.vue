<template>
  <div v-if="responseDrawerVisible">
    <a-drawer
        :visible="responseDrawerVisible"
        :closable="true" :width="1000"
        :bodyStyle="{ padding: '24px', height: '100%' }"
        @close="onClose">
      <template #title>
        <span>接口运行结果</span>
      </template>

      <div class="drawer-content">
        <a-tabs v-model:activeKey="activeKey">
          <a-tab-pane key="Request" tab="请求">
            <RequestInfo :data="requestContent"/>
          </a-tab-pane>

          <a-tab-pane key="Body" tab="响应Body">
            <BodyInfo :data="bodyInfo"/>
          </a-tab-pane>

          <a-tab-pane key="Header" tab="响应Header" force-render>
            <HeaderInfo :data="headers"/>
          </a-tab-pane>

          <a-tab-pane key="Cookie" tab="响应Cookie" >
            <CookieInfo :data="cookies"/>
          </a-tab-pane>

<!--      <a-tab-pane key="Console" tab="控制台">
            <ConsoleInfo/>
          </a-tab-pane> -->
        </a-tabs>
      </div>
    </a-drawer>
  </div>

</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, watch, computed} from 'vue';
import {BodyInfo, ConsoleInfo, CookieInfo, RequestInfo, HeaderInfo} from './Components/index';

const props = defineProps({
  responseDrawerVisible: {
    type: Boolean
  },
  data: {
    type: Object
  }
});

const emits = defineEmits(['onClose']);

const activeKey = ref('Request');
const bodyInfo = ref({});

const cookies = ref([]);
const headers = ref([]);
const requestContent = ref({});

function onClose() {
  emits('onClose');
}

watch(() => {
  return props.responseDrawerVisible;
}, (newVal) => {
  if (!newVal) return

  const {resContent = {}, reqContent = {}}: any = props.data;

  bodyInfo.value = {
    content: resContent.content || '',
    contentLang: resContent.contentLang || ''
  };

  cookies.value = resContent.cookies || [];
  headers.value = resContent.headers || [];
  requestContent.value = reqContent;

}, {
  immediate: true,
})
</script>

<style scoped lang="less">
.drawer-content {
  height: calc(100% - 46px);
}
</style>
