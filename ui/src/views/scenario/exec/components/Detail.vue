<template>
  <div class="scenario-exec-detail-main">
    DETAIL
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import {StateType as ScenarioStateType} from "../../store";

import ProcessorGroupDefault from "./proccessors/group/default.vue";
import ProcessorInterfaceDefault from "./proccessors/interface/default.vue"
import ProcessorTimerDefault from "./proccessors/timer/default.vue"

import ProcessorLogicIf  from "./proccessors/logic/if.vue"
import ProcessorLogicElse from "./proccessors/logic/else.vue"

import ProcessorLoopTime from "./proccessors/loop/time.vue"
import ProcessorLoopUntil from "./proccessors/loop/until.vue"
import ProcessorLoopIn  from "./proccessors/loop/in.vue"
import ProcessorLoopRange from "./proccessors/loop/range.vue"
import ProcessorLoopBreak from "./proccessors/loop/break.vue"

import ProcessorVariableSet  from "./proccessors/variable/set.vue"
import ProcessorVariableClear from "./proccessors/variable/clear.vue"

import ProcessorAssertionDefault     from "./proccessors/assertion/default.vue"

import ProcessorExtractorBoundary from "./proccessors/extractor/boundary.vue"
import ProcessorExtractorJsonQuery from "./proccessors/extractor/jsonquery.vue"
import ProcessorExtractorHtmlQuery from "./proccessors/extractor/htmlquery.vue"
import ProcessorExtractorXmlQuery from "./proccessors/extractor/xmlquery.vue"

import ProcessorCookieGet  from "./proccessors/cookie/get.vue"
import ProcessorCookieSet  from "./proccessors/cookie/set.vue"
import ProcessorCookieClear from "./proccessors/cookie/clear.vue"

import ProcessorDataText   from "./proccessors/data/text.vue"
import ProcessorDataExcel  from "./proccessors/data/excel.vue"
import ProcessorDataZenData from "./proccessors/data/zendata.vue"
import {StateType as GlobalStateType} from "@/store/global";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();
const collapsed = computed<boolean>(()=> store.state.Global.collapsed);
const execData = computed<any>(()=> store.state.Scenario.execData);

const id = ref(+router.currentRoute.value.params.id)

// store.dispatch('Scenario/loadExecScenario', id.value);

</script>

<style lang="less" scoped>
.scenario-exec-detail-main {
  height: 100%;
  padding: 6px;
}

</style>