<template>
  <div id="test-interface-main" class="dp-splits-v">
    <div id="test-interface-left" v-if="!collapsed">
      <TestInterfaceTree />
    </div>

    <CollapsedIcon
        :style="{left:'294px',top:'300px'}"
        :collapsedStyle="{left:'-9px', top:'300px'}"
        @click="collapsed = !collapsed" :collapsed="collapsed" />

    <div id="test-interface-right">
      <TestInterfaceDesign />
    </div>
  </div>
</template>


<script setup lang="ts">
import {computed, ref} from "vue";
import {useStore} from "vuex";
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"

import TestInterfaceTree from './components/tree.vue';
import TestInterfaceDesign from './design/index.vue';

import {StateType as TestInterfaceStateType} from "@/views/debugger/store";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ServeStateType} from "@/store/serve";

const collapsed = ref(false);

const store = useStore<{ TestInterface: TestInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

</script>

<style lang="less" scoped>
#test-interface-main {
  display: flex;
  height: 100%;
  margin: 16px;
  position: relative;

  #test-interface-left {
    width: 300px;
    height: 100%;
    background-color: #ffffff;
    border-right: 1px solid #f0f0f0;
  }
  #test-interface-right {
    flex: 1;
    width: 0;
    height: 100%;
    background-color: #ffffff;
  }
}
</style>
