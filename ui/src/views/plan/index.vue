<template>
  <div id="plan-index-main" class="dp-splits-v">
    <div id="plan-index-left">
      <PlanTree />
    </div>

    <div id="plan-index-splitter" class="splitter"></div>

    <div id="plan-index-right">
        <PlanList />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";

import PlanTree from './list/tree.vue';
import PlanList from './list/list.vue';

const useForm = Form.useForm;

const {t} = useI18n();

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('plan-index-main',
      'plan-index-left', 'plan-index-splitter', 'plan-index-right', 320, 600)
}

</script>

<style lang="less" scoped>
#plan-index-main {
  display: flex;
  height: 100%;

  #plan-index-left {
    width: 260px;
    height: 100%;
  }

  #plan-index-right {
    flex: 1;
    width: 0;
    height: 100%;
  }
}

</style>