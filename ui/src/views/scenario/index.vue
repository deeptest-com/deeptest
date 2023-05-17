<template>
  <div id="scenario-index-main" class="dp-splits-v">
    <div id="scenario-index-left" v-if="!collapsed">
      <Tree/>
    </div>
    <CollapsedIcon
        :style="{left:'294px',top:'300px'}"
        :collapsedStyle="{left:'-9px', top:'300px'}"
        @click="collapsed = !collapsed" :collapsed="collapsed"/>
    <div id="scenario-index-right">
        <ScenarioList />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';
import {UsedBy} from "@/utils/enum";
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"
import Tree from './components/Tree/index.vue';
import ScenarioList from './components/List/index.vue';

provide('usedBy', UsedBy.InterfaceDebug)
const useForm = Form.useForm;
const collapsed = ref(false);
const {t} = useI18n();

</script>

<style lang="less" scoped>
#scenario-index-main {
  display: flex;
  height: 100%;
  margin: 16px;
  position: relative;

  #scenario-index-left {
    width: 300px;
    height: 100%;
    background-color: #ffffff;
    border-right: 1px solid #f0f0f0;
  }

  #scenario-index-right {
    flex: 1;
    width: 0;
    height: 100%;
    background-color: #ffffff;
  }
}

</style>
