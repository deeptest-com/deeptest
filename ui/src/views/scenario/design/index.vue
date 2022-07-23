<template>
  <div id="scenario-design-main" class="scenario-design-main dp-splits-v">
    <div id="scenario-design-left" class="left">
      <ScenarioTree :scenarioId="id" />
    </div>

    <div id="scenario-design-splitter" class="splitter"></div>

    <div id="scenario-design-right" class="right">
      <Edit></Edit>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";

import bus from "@/utils/eventBus";
import {resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as GlobalStateType} from "@/store/global";
import {StateType as ScenarioStateType} from "../store";

import ScenarioTree from "./components/Tree.vue"
import Edit from "./components/Edit.vue"
import settings from "@/config/settings";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();

const collapsed = computed<boolean>(()=> store.state.Global.collapsed);

const id = ref(+router.currentRoute.value.params.id)

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
  resizeWidth('scenario-design-main',
      'scenario-design-left', 'scenario-design-splitter', 'scenario-design-right',
      260, 800, collapsed.value ? 55 - 15 : 100 - 25)
}

</script>

<style lang="less" scoped>
.scenario-design-main {
  .left {
    width: 260px;
  }
  .right {
    flex: 1;
  }
}
</style>