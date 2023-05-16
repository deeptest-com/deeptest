<template>
  <div id="scenario-design-main" class="scenario-design-main dp-splits-v">
    <div id="scenario-design-left" class="left">
      <ScenarioTree/>
    </div>
    <div id="scenario-design-splitter" class="splitter"></div>
    <div id="scenario-design-right" class="right">
      <Designer/>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";

import {resizeWidth} from "@/utils/dom";
import {useStore} from "vuex";

import {StateType as GlobalStateType} from "@/store/global";
import {StateType as ScenarioStateType} from "../store";

import ScenarioTree from "./Tree.vue"
import Designer from "./Designer.vue"
import {Scenario} from "@/views/scenario/data";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();
// const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);
// const id = ref(+router.currentRoute.value.params.id)
// const props = defineProps({
//   id: {
//     required: true,
//     type: Number,
//   },
// })

const collapsed = computed<boolean>(()=> store.state.Global.collapsed);

onMounted(() => {
  console.log('onMounted')
  resize()
  store.dispatch('Scenario/getNode', null)
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
      260, 800)
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
