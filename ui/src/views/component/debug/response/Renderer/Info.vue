<template>
  <div class="info-main">
    <div>
      <ConBoxTitle title="请求地址" />
      <div class="content">
        <span :style="{color: getMethodColor(requestData.method)}">
          {{requestData.method}}
        </span>&nbsp;
        <span>{{requestData.fullUrlToDisplay || requestData.url}}</span>
      </div>
    </div>

    <ParamGrid title="查询参数" :list="requestData.queryParams || []" />
    <ParamGrid title="路径参数" :list="requestData.pathParams || []" />
    <ParamGrid title="请求头" :list="requestData.headers || []" />

    <ParamGrid v-if="requestData.bodyType==='multipart/form-data' && requestData.bodyFormData"
               title="表单数据"
               :list="requestData.bodyFormData || []" />

    <ParamGrid v-else-if="requestData.bodyType==='application/x-www-form-urlencoded' && requestData.bodyFormUrlencoded"
               title="表单数据（UrlEncoded）"
               :list="requestData.bodyFormUrlencoded || []" />

    <ParamContent v-else
                  title="请求体"
                  :content="requestData.body || ''" />
  </div>
</template>

<script setup lang="ts">
import {computed, inject, defineProps} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import ParamGrid from "../../comp/param-grid.vue";
import ParamContent from "../../comp/param-content.vue";
import {StateType as Debug} from "@/views/component/debug/store";
import {getMethodColor} from "@/utils/dom";

const {t} = useI18n();
const props = defineProps<{
  data?: any;
}>();
const store = useStore<{  Debug: Debug }>();

const requestData = computed<any>(() => props.data || store.state.Debug.requestData);

</script>

<style lang="less" scoped>
.info-main {
  height: 100%;
  overflow-y: auto;
  //max-height: 480px;
  //max-width: 758px;
  .content {
    padding: 10px 10px;
  }
}

</style>
