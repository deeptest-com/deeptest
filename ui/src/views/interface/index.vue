<template>
  <div id="interface-design-main" class="interface-design-main dp-splits-v">
    <div id="interface-design-left" class="left">
      <InterfaceTree />
    </div>

    <div id="interface-design-splitter" class="splitter"></div>

    <div id="interface-design-right" class="right">
      <InterfaceDesigner></InterfaceDesigner>
    </div>
  </div>
</template>

<script lang="ts">
import {computed, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import {useRouter} from "vue-router";

import {Form} from "ant-design-vue";

import InterfaceDesigner from './components/Designer.vue';
import {resizeWidth} from "@/utils/dom";
import InterfaceTree from "./components/Tree.vue"
import {useStore} from "vuex";
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";

export default defineComponent({
  name: 'InterfaceIndexPage',
  components: {
    InterfaceTree, InterfaceDesigner,
  },
  setup() {
    const router = useRouter();
    const store = useStore<{ Global: GlobalStateType; User: UserStateType; }>();

    const collapsed = computed<boolean>(()=> store.state.Global.collapsed);

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
          260, 800, collapsed.value ? 55 - 15 : 100 - 25)
    }

    return {

    }
  }

})
</script>

<style lang="less" scoped>
.interface-design-main {

}
</style>