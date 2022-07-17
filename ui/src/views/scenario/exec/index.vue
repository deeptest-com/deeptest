<template>
  <div id="scenario-exec-main" class="scenario-exec-main dp-splits-v">
    <div id="scenario-exec-left" class="left">
      <ScenarioTree :scenarioId="id" />
    </div>

    <div id="scenario-exec-splitter" class="splitter"></div>

    <div id="scenario-exec-right" class="right">
      RIGHT
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref, watch} from "vue";
import {Scenario} from '../data.d';
import {useStore} from "vuex";

import {StateType as ScenarioStateType, StateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal} from "ant-design-vue";
import {resizeWidth} from "@/utils/dom";
import ScenarioTree from "./components/Tree.vue"
import {StateType as GlobalStateType} from "@/store/global";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const list = computed<Scenario[]>(() => store.state.Scenario.listResult.list);

const id = ref(+router.currentRoute.value.params.id)

onMounted(() => {
  console.log('onMounted')
})

onMounted(() => {
  console.log('onMounted')
})

const resize = () => {
  resizeWidth('scenario-exec-main',
      'scenario-exec-left', 'scenario-exec-splitter', 'scenario-exec-right',
      260, 800, collapsed.value ? 55 - 15 : 100 - 25)
}

</script>

<style lang="less" scoped>
.scenario-exec-main {

}
</style>