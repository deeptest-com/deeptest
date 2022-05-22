<template>
  <div id="main" class="scenario-design-main">
    <div id="left-panel">
      <ScenarioTree />
    </div>

    <div id="splitter-h"></div>

    <div id="right-panel">

    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, watch} from "vue";
import {useRouter} from "vue-router";

import {resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as GlobalStateType} from "@/store/global";
import {StateType as ScenarioStateType} from "../store";

import ScenarioTree from "./components/Tree.vue"

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();

const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const detailResult = computed<boolean>(() => store.state.Scenario.detailResult);

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