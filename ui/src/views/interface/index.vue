<template>
  <div id="interface-index-main" class="dp-splits-v">
    <div id="interface-index-left">
      <InterfaceTree />
    </div>

    <div id="interface-index-splitter" class="splitter"></div>

    <div id="interface-index-right">
        <InterfaceList />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {resizeWidth} from "@/utils/dom";
import {UsedBy} from "@/utils/enum";

import InterfaceTree from './list/tree.vue';
import InterfaceList from './list/list.vue';

const useForm = Form.useForm;

const {t} = useI18n();

onMounted(() => {
  console.log('onMounted')
  resize()
})

const resize = () => {
  resizeWidth('interface-index-main',
      'interface-index-left', 'interface-index-splitter', 'interface-index-right', 320, 600)
}

</script>

<style lang="less" scoped>
#interface-index-main {
  display: flex;
  height: 100%;

  #interface-index-left {
    width: 260px;
    height: 100%;
  }

  #interface-index-right {
    flex: 1;
    width: 0;
    height: 100%;
  }
}

</style>