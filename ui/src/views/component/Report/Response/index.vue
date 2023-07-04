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

          <a-tab-pane key="extractor" tab="响应提取器" >
            <ExtractorInfo :data="extractors"/>
          </a-tab-pane>

          <a-tab-pane key="checkpoint" tab="响应验证点" >
            <CheckpointInfo :data="checkpoints"/>
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
import {BodyInfo, CookieInfo, RequestInfo, HeaderInfo, ExtractorInfo, CheckpointInfo} from './Components/index';

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

const extractors = ref([]);
const checkpoints = ref([]);

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

  extractors.value = resContent.cookies || [];
  checkpoints.value = resContent.checkpoints || [];

}, {
  immediate: true,
})
</script>

<style scoped lang="less">
.drawer-content {
  height: calc(100% - 46px);
}
</style>
