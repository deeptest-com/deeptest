<template>
  <div class="config-main">
    <a-tabs v-model:activeKey="activeKey" :animated="false">
      <a-tab-pane key="1" tab="查询参数">
        <QueryParameters v-if="activeKey === '1'"></QueryParameters>
      </a-tab-pane>

      <a-tab-pane key="2" tab="路径参数">
        <PathParameters v-if="activeKey === '2'"></PathParameters>
      </a-tab-pane>

      <a-tab-pane key="3" tab="请求体">
        <RequestBody v-if="activeKey === '3'"></RequestBody>
      </a-tab-pane>

      <a-tab-pane key="4" tab="请求头">
        <RequestHeaders v-if="activeKey === '4'"></RequestHeaders>
      </a-tab-pane>

      <a-tab-pane key="5" tab="授权">
        <Authorization v-if="activeKey === '5'"></Authorization>
      </a-tab-pane>

      <a-tab-pane key="6" tab="预处理">
        <PreRequestScript v-if="activeKey === '6'"></PreRequestScript>
      </a-tab-pane>

      <!-- <a-tab-pane key="7" tab="验证">
              <ValidationScript v-if="activeKey === '7'"></ValidationScript>
            </a-tab-pane>-->
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
import PreRequestScript from "./config/PreRequestScript.vue";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const activeKey = ref('1');

</script>

<style lang="less">
.config-main {
  height: calc(100% - 32px);
  min-height: 200px;

  .ant-tabs-line {
    height: 100%;

    .ant-tabs-top-content {
      height: calc(100% - 61px);
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
