<template>
  <div id="interface-design-main" class="interface-design-main dp-splits-v">
    <div id="interface-design-left" class="left">
      <InterfaceTree />
    </div>

    <div id="interface-design-splitter" class="splitter"></div>

    <div id="interface-design-right" class="right">
      <InterfaceDesigner v-if="!interfaceData.isDir"></InterfaceDesigner>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";

import {resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";
import {Interface} from "@/views/interface/data";
import {StateType} from "@/views/interface/store";

import InterfaceDesigner from './Designer.vue';
import InterfaceTree from "./Tree.vue"

    const store = useStore<{ Global: GlobalStateType, User: UserStateType, Interface: StateType }>();
    const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    onMounted(() => {
      console.log('onMounted')
      resize()
      store.dispatch('Interface/getInterface', null)
    })
    onUnmounted(() => {
      console.log('onUnmounted')
    })

    watch(collapsed, () => {
      console.log('watch collapsed')
      resize()
    }, {deep: true})

    const resize = () => {
      resizeWidth('interface-design-main',
            'interface-design-left', 'interface-design-splitter', 'interface-design-right',
          260, 800)
    }

</script>

<style lang="less" scoped>
.interface-design-main {
  .left {
    width: 260px;
  }
  .right {
    flex: 1;
  }
}
</style>