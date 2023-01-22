<template>
  <div id="scenario-index-main" class="dp-splits-v">
    <div id="scenario-index-left">
      <ScenarioTree />
    </div>

    <div id="scenario-index-splitter" class="splitter"></div>

    <div id="scenario-index-right">
        <ScenarioList />
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";

import ScenarioTree from './list/tree.vue';
import ScenarioList from './list/list.vue';

provide('usedBy', UsedBy.interface)
const useForm = Form.useForm;

const {t} = useI18n();

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('scenario-index-main',
      'scenario-index-left', 'scenario-index-splitter', 'scenario-index-right', 320, 600)
}

</script>

<style lang="less" scoped>
#scenario-index-main {
  display: flex;
  height: 100%;

  #scenario-index-left {
    width: 320px;
    height: 100%;
  }

  #scenario-index-right {
    flex: 1;
    width: 0;
    height: 100%;
  }
}

</style>