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

<script lang="ts">
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

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    InterfaceTree, InterfaceDesigner,
  },
  setup() {
    const store = useStore<{ Global: GlobalStateType, User: UserStateType, Interface: StateType }>();
    const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    onMounted(() => {
      console.log('onMounted')
      resize()
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

    return {
      interfaceData
    }
  }

})
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