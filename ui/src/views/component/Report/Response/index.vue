<template>
  <div v-if="responseDrawerVisible">
    <a-drawer
      :visible="responseDrawerVisible"
      :closable="true"
      :width="1000"
      :bodyStyle="{ padding: '24px', height: '100%' }"
      @close="onClose"
    >
      <template #title>
        <span>接口运行结果</span>
      </template>

      <div class="drawer-content">
        <a-tabs v-model:activeKey="activeKey" class="dp-tabs-full-height">
          <a-tab-pane key="body" :tab="title" class="uppercase">
            <BodyInfo :data="interfaceResDetail.bodyInfo" />
          </a-tab-pane>
          <a-tab-pane key="header" tab="响应头">
            <ParamGrid :list="interfaceResDetail.headers" />
          </a-tab-pane>
          <a-tab-pane key="cookie" tab="Cookie">
            <ParamGrid :list="interfaceResDetail.cookies" />
          </a-tab-pane>
          <a-tab-pane key="console" tab="控制台">
            <ResponseConsole :data="interfaceResDetail" />
          </a-tab-pane>
          <a-tab-pane key="info" tab="实际请求">
            <ResponseInfo :data="interfaceResDetail.requestData" />
          </a-tab-pane>
        </a-tabs>
      </div>
    </a-drawer>
  </div>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, watch, computed, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { BodyInfo } from "./Components";
import ParamGrid from "@/views/component/debug/comp/param-grid.vue";
import ResponseConsole from "@/views/component/debug/response/Renderer/Console.vue";
import ResponseInfo from "@/views/component/debug/response/Renderer/Info.vue";
import { InterfaceDetail } from "./data";

const props = defineProps({
  responseDrawerVisible: {
    type: Boolean,
  },
  data: {
    type: Object,
  },
});

const { t } = useI18n();

const emits = defineEmits(["onClose"]);

const activeKey = ref("body");
const interfaceResDetail = reactive(<InterfaceDetail>{
  contentLang: "",
  headers: [],
  contentType: "",
  cookies: [],
});

const title = computed(() => t(interfaceResDetail.contentLang || "empty"));

function onClose() {
  emits("onClose");
}

watch(
  () => {
    return props.responseDrawerVisible;
  },
  (newVal) => {
    if (!newVal) return;
    const { resContent = {}, reqContent = {} }: any = props.data;
    Object.assign(interfaceResDetail, {
      ...resContent,
      bodyInfo: {
        content: resContent.content || "",
        contentLang: resContent.contentLang || "",
        contentType: resContent.contentType || "",
      },
      headers: resContent.headers || [],
      cookies: resContent.cookies || [],
      requestData: reqContent,
    });
  },
  {
    immediate: true,
  }
);
</script>

<style scoped lang="less">
.drawer-content {
  height: calc(100% - 46px);
}
</style>
