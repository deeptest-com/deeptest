<template>
  <div class="response-meta">
    <div class="row">
      <span class="col" :class="[responseData.statusCode===200? 'dp-color-pass': 'dp-color-fail']">
        状态：{{ responseData.statusContent }}
      </span>
      <span class="col">
        耗时: {{ responseData.time }}毫秒
      </span>
      <span class="col">
        大小：{{ responseData.contentLength }}字节
      </span>
    </div>

    <div v-for="(item, index) in resultData" :key="index" class="item">
      {{item.id}}
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, watch} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {useI18n} from "vue-i18n";
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const responseData = computed<any>(() => store.state.Debug.responseData);
const resultData = computed<any>(() => store.state.Debug.resultData);

watch(responseData, (newVal) => {
  console.log('responseData', responseData.value.invokeId)
  if (responseData.value.invokeId)
    store.dispatch("Debug/getInvocationResult", responseData.value.invokeId)
}, {deep: true, immediate: true})

</script>

<style lang="less" scoped>
.response-meta {
  padding: 0 6px;

  .row {
    padding: 2px 0;
    .col {
      margin-right: 20px;
    }
  }
}

</style>