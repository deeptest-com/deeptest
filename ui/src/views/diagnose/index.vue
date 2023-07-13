<template>
  <div id="diagnose-interface-main" class="dp-splits-v">
    <div id="diagnose-interface-left" v-if="!collapsed">
      <DiagnoseInterfaceTree />
    </div>

    <CollapsedIcon
        :style="{left:'294px',top:'300px'}"
        :collapsedStyle="{left:'-9px', top:'300px'}"
        @click="collapsed = !collapsed" :collapsed="collapsed" />

    <div id="diagnose-interface-right">
      <DiagnoseInterfaceDesign />
    </div>
  </div>
</template>


<script setup lang="ts">
import {computed, ref} from "vue";
import {useStore} from "vuex";
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"

import DiagnoseInterfaceTree from './components/tree.vue';
import DiagnoseInterfaceDesign from './design/index.vue';

import {StateType as DiagnoseInterfaceStateType} from "@/views/diagnose/store";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ServeStateType} from "@/store/serve";

const collapsed = ref(false);

const store = useStore<{ DiagnoseInterface: DiagnoseInterfaceStateType, ProjectGlobal: ProjectStateType, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

</script>

<style lang="less" scoped>
#diagnose-interface-main {
  position: relative;
  display: flex;
  height: 100%;
  padding: 16px;

  #diagnose-interface-left {
    width: 300px;
    height: 100%;
    background-color: #ffffff;
    border-right: 1px solid #f0f0f0;
  }
  #diagnose-interface-right {
    flex: 1;
    height: 100%;
    width: 0;
    background-color: #ffffff;
  }
}
</style>
