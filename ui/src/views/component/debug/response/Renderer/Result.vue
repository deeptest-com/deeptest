<template>
  <div class="response-result">
    <div class="row status">
      <span class="col">
        状态：{{ responseData.statusContent }}
      </span>
      <span class="col">
        耗时: {{ responseData.time }}毫秒
      </span>
      <span class="col">
        大小：{{ responseData.contentLength }}字节
      </span>
    </div>

    <template v-if="resultData?.length > 0">
      <ResponseDefine v-if="entityData && entityData.codes.length > 0"
        :codes="entityData.codes"
        :code="entityData.code?entityData.code:'200'"
        :open="!entityData.disabled"
        @change="change"
        />
      <ResultMsg :responseData="responseDataForDefine"/>
      <div class="title" v-if="responseDataForAssert.length > 0">断言结果</div>
      <ResultMsg :responseData="responseDataForAssert"/>
    </template>
  </div>
</template>

<script setup lang="ts">
import {computed, watch,ref} from "vue";
import {useStore} from "vuex";

import {StateType as Debug} from "@/views/component/debug/store";
import {useI18n} from "vue-i18n";
import ResponseDefine from "./ResponseDefine.vue";
import ResultMsg from "./ResultMsg.vue";
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const responseData = computed<any>(() => store.state.Debug.responseData);
const resultData = computed<any>(() => store.state.Debug.resultData);
const entityData = computed<any>(()=>store.state.Debug.debugData.responseDefine?.entityData)
const responseDataForDefine = computed(()=>resultData.value.filter((item:any)=>item.conditionEntityType=="responseDefine"))
const responseDataForAssert = computed(()=>resultData.value.filter((item:any)=>item.conditionEntityType=="checkpoint"))


watch(responseData, (newVal) => {
  console.log('responseData', responseData.value.invokeId)
  if (responseData.value.invokeId)
    store.dispatch("Debug/getInvocationResult", responseData.value.invokeId)
}, {immediate: true, deep: true})



const change = async (formState:any)=>{
  console.log(formState)
  await store.dispatch("Debug/saveResponseDefine",{id:entityData.value.id,disabled:!formState.open,code:formState.code} )
}



</script>

<style lang="less" scoped>
.response-result {
  height: 100%;
  overflow-y: auto;
  padding: 0px 6px;

  .status {
    padding: 12px 0 8px 0;
    .col {
      margin-right: 20px;
    }
  }

  .title {
    padding-left: 2px;
    font-weight: bold;
  }

}

</style>