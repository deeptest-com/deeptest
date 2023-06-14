<template>
  <div id="test-interface-index-main" class="dp-splits-v">
    <div id="test-interface-index-left" v-if="!collapsed">
      <TestInterfaceTree />
    </div>

    <CollapsedIcon
        :style="{left:'294px',top:'300px'}"
        :collapsedStyle="{left:'-9px', top:'300px'}"
        @click="collapsed = !collapsed" :collapsed="collapsed" />

    <div id="test-interface-index-right">

    </div>
  </div>
</template>


<script setup lang="ts">
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"
import TestInterfaceTree from './components/tree.vue';
import {ref} from "vue";

const collapsed = ref(false);

</script>

<style lang="less" scoped>
.debugger-main {
  display: flex;
  height: 100%;
  margin: 16px;
  position: relative;

  #test-interface-index-left {
    width: 300px;
    height: 100%;
    background-color: #ffffff;
    border-right: 1px solid #f0f0f0;
  }
  #test-interface-index-right {
    flex: 1;
    width: 0;
    height: 100%;
    background-color: #ffffff;
  }
}
</style>
