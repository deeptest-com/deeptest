<template>
  <div class="config-main">
    <a-tabs v-model:activeKey="activeKey" :animated="false" class="dp-tabs-full-height">
      <a-tab-pane key="query-param" tab="查询参数">
        <QueryParameters v-if="activeKey === 'query-param'" />
      </a-tab-pane>

      <a-tab-pane key="path-param" tab="路径参数">
        <PathParameters v-if="activeKey === 'path-param'" />
      </a-tab-pane>

      <a-tab-pane key="body" tab="请求体">
        <RequestBody v-if="activeKey === 'body'" />
      </a-tab-pane>

      <a-tab-pane key="header" tab="请求头">
        <RequestHeaders v-if="activeKey === 'header'" />
      </a-tab-pane>

      <a-tab-pane key="auth" tab="授权">
        <Authorization v-if="activeKey === 'auth'" />
      </a-tab-pane>

      <a-tab-pane key="pre-condition" tab="预处理">
        <PreCondition v-if="activeKey === 'pre-condition'" />
      </a-tab-pane>

      <a-tab-pane key="post-condition" tab="后处理">
        <PostCondition v-if="activeKey === 'post-condition'" />
      </a-tab-pane>

      <a-tab-pane key="assertion" tab="断言">
        <Assertion v-if="activeKey === 'assertion'" />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup lang="ts">
import {inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {UsedBy} from "@/utils/enum";

import QueryParameters from "./config/QueryParameters.vue";
import PathParameters from "./config/PathParameters.vue";

import RequestBody from "./config/Body.vue";
import RequestHeaders from "./config/Headers.vue";
import Authorization from "./config/Authorization.vue";
import PreCondition from "./config/ConditionPre.vue";
import PostCondition from "./config/ConditionPost.vue";
import Assertion from "./config/Assertion.vue";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const activeKey = ref('query-param');

</script>

<style lang="less">
.config-main {
  height: 100%;

  .ant-tabs-line {
    height: 100%;

    .ant-tabs-top-content {
      height: calc(100% - 45px);
      overflow-y: auto;
    }
  }
}
</style>

<style lang="less" scoped>
.config-main {
  padding: 3px;

  .ant-tabs-tabpane-active {
    height: 100%;
  }
}
</style>
