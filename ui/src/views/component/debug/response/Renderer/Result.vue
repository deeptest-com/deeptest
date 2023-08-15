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
      <ResponseDefine v-if="entityData.codes.length > 0" 
        :codes="entityData.codes" 
        :code="entityData.code?entityData.code:'200'"
        :open="!entityData.disabled"
        @change="change"
        />
      <div class="title">断言结果</div>
      <div v-for="(item, index) in resultData"
           :key="index"
           :class="getResultClass(item)" class="item">

        <span v-if="item.resultStatus===ResultStatus.Pass">
          <CheckCircleOutlined />
        </span>

        <span v-if="item.resultStatus===ResultStatus.Fail">
          <CloseCircleOutlined />
        </span>&nbsp;

        <span>{{item.resultMsg}}</span>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import {computed, watch,ref} from "vue";
import {useStore} from "vuex";
import { CheckCircleOutlined, CloseCircleOutlined} from '@ant-design/icons-vue';

import {ResultStatus} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {useI18n} from "vue-i18n";
import ResponseDefine from "./ResponseDefine.vue";
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const responseData = computed<any>(() => store.state.Debug.responseData);
const resultData = computed<any>(() => store.state.Debug.resultData);
const entityData = computed<any>(()=>store.state.Debug.debugData.responseDefine?.entityData)


watch(responseData, (newVal) => {
  console.log('responseData', responseData.value.invokeId)
  if (responseData.value.invokeId)
    store.dispatch("Debug/getInvocationResult", responseData.value.invokeId)
}, {immediate: true, deep: true})

const getResultClass = (item) => {
  return item.resultStatus===ResultStatus.Pass? 'pass':
      item.resultStatus===ResultStatus.Fail ? 'fail' : ''
}

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

  .item {
    margin: 3px;
    padding: 5px;
    &.pass {
      color: #14945a;
      background-color: #F1FAF4;
    }
    &.fail {
      color: #D8021A;
      background-color: #FFECEE;
    }
  }
}

</style>