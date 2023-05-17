<template>
  <div id="plan-index-main" class="dp-splits-v">
    <div id="plan-index-left"  v-if="!collapsed">
      <PlanTree />
    </div>
    <CollapsedIcon
        :style="{left:'294px',top:'300px'}"
        :collapsedStyle="{left:'-9px', top:'300px'}"
        @click="collapsed = !collapsed" :collapsed="collapsed"/>
    <div id="plan-index-right">
        <PlanList />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, provide, ref} from "vue";
import {useI18n} from "vue-i18n";
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"
import PlanTree from './list/tree.vue';
import PlanList from './list/list.vue';

const collapsed = ref(false);

const {t} = useI18n();


</script>

<style lang="less" scoped>
#plan-index-main {
  display: flex;
  height: 100%;
  margin: 16px;
  position: relative;
  #plan-index-left {
    width: 300px;
    height: 100%;
    background-color: #ffffff;
    border-right: 1px solid #f0f0f0;
  }
  #plan-index-right {
    flex: 1;
    width: 0;
    height: 100%;
    background-color: #ffffff;
  }
}

</style>
