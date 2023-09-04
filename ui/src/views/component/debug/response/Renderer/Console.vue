<template>
  <div class="response-console-main" v-if="consoleData && consoleData.length">
    <div v-for="(item, index) in consoleData"
         :key="index"
         :class="getResultClass(item)" class="item">

      <span v-if="item.resultStatus===ResultStatus.Pass">
        <CheckCircleOutlined />
      </span>
      <span v-if="item.resultStatus===ResultStatus.Fail">
        <CloseCircleOutlined />
      </span>&nbsp;

      <span>
        <icon-svg v-if="item.conditionEntityType === ConditionType.extractor"
                  type="variable"
                  class="icon variable" />
        <icon-svg v-if="item.conditionEntityType === ConditionType.checkpoint"
                  type="checkpoint"
                  class="icon"  />
        <icon-svg v-if="item.conditionEntityType === ConditionType.script"
                  type="script"
                  class="icon"  />
      </span>
      &nbsp;
      <span v-html="item.resultMsg"></span>
    </div>
  </div>
  <Empty :desc="'暂无数据'" style="margin-top: 100px" v-else/>
</template>

<script setup lang="ts">
import {computed, inject, watch, defineProps} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {useI18n} from "vue-i18n";
import {ConditionType, ResultStatus} from "@/utils/enum";
import { CheckCircleOutlined, CloseCircleOutlined} from '@ant-design/icons-vue';
import IconSvg from "@/components/IconSvg";
import Empty from "@/components/others/empty.vue";

const {t} = useI18n();

const props = defineProps<{
  data?: any;
}>();

const store = useStore<{  Debug: Debug }>();
const responseData = computed<any>(() => props.data || store.state.Debug.responseData);
const consoleData = computed<any>(() => store.state.Debug.consoleData);

watch(responseData, (newVal) => {
  console.log('watch responseData', responseData.value.invokeId)
  if (responseData.value.invokeId)
    store.dispatch("Debug/getInvocationLog", responseData.value.invokeId)
}, {deep: true, immediate: true})

const getResultClass = (item) => {
  return item.resultStatus===ResultStatus.Pass? 'pass':
      item.resultStatus===ResultStatus.Fail ? 'fail' : ''
}

</script>

<style lang="less">
.response-console-main {
}
</style>

<style lang="less" scoped>
.response-console-main {
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