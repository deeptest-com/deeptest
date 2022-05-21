<template>
  <div id="main" class-interface-main>
    <div id="left-panel">
      <InterfaceTree />
    </div>

    <div id="splitter-h"></div>

    <div id="right-panel">
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
      resizeWidth('main', 'left-panel', 'splitter-h', 'right-panel',
          260, 800, collapsed.value ? 55 - 15 : 100 - 25)
    }

    return {

    }
  }

})
</script>

<style lang="less" scoped>
#main {
  display: flex;
  height: 100%;

  #left-panel {
    width: 260px;
    height: 100%;
  }

  #right-panel {
    flex: 1;
    height: 100%;
    overflow-y: auto;
    overflow-x: hidden;
  }

  #splitter-h {
    width: 1px;
    height: 100%;
    background-color: #e6e9ec;
    cursor: ew-resize;

    &:hover {
      width: 1px;
      background-color: #D0D7DE;
    }

    &.active {
      width: 1px;
      background-color: #a9aeb4;
    }
  }
}

</style>