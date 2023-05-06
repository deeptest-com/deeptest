<template>
  <div id="scenario-design-main">
    <div class="toolbar">
      <a-button @click="exec" type="link" class="dp-btn-group">执行</a-button>
      <a-button href="#/scenario/index" type="link" class="dp-btn-group">返回</a-button>
    </div>

    <div class="scenario-design-content">
      <Processor></Processor>
    </div>
  </div>

</template>

<script setup lang="ts">
import {onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {useStore} from "vuex";

import {useRouter} from "vue-router";
import {UsedBy} from "@/utils/enum";

import Processor from './components/Processor.vue';

const router = useRouter();
provide('usedBy', UsedBy.ScenarioDebug)
const useForm = Form.useForm;

const {t} = useI18n();

const id = ref(+router.currentRoute.value.params.id)

const exec = () => {
  router.push(`/scenario/exec/${id.value}`)
}

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
#scenario-design-main {
  height: 100%;

  .toolbar {
    height: 32px;
    padding-left: 12px;
    padding-right: 12px;
    text-align: right;
  }

  .scenario-design-content {
    height: calc(100% - 32px);
    width: 100%;
  }
}

</style>