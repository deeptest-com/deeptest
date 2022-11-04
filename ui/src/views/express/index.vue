<template>
  <div id="express-main" class="express-main dp-splits-v">
    <div id="express-left" class="left">
      <InterfaceTree />
    </div>

    <div id="express-splitter" class="splitter"></div>

    <div id="express-right" class="right">
      <InterfaceDesigner></InterfaceDesigner>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";

import InterfaceDesigner from './components/Designer.vue';
import {resizeWidth} from "@/utils/dom";
import InterfaceTree from "./components/Tree.vue"
import {useStore} from "vuex";
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {Interface} from "@/views/interface/data";
import {StateType} from "@/views/interface/store";

const store = useStore<{ Global: GlobalStateType, User: UserStateType, ProjectData: ProjectStateType,
  Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

onMounted(() => {
  console.log('onMounted')
  resize()
})
onUnmounted(() => {
  console.log('onUnmounted')
})

watch(interfaceData, () => {
  console.log('watch interfaceData')

}, {deep: true})

const resize = () => {
  resizeWidth('express-main',
        'express-left', 'express-splitter', 'express-right',
      260, 800)
}

</script>

<style lang="less" scoped>
.express-main {
  .left {
    width: 260px;
  }
  .right {
    flex: 1;
  }
}
</style>