<template>
  <div class="scenario-node-edit-main">
    <!--    <ProcessorThreadDefault v-if="selectedNode.processorType === 'processor_thread_default'" />-->

    <ProcessorInterfaceDefault v-if="selectedNode.processorType === 'processor_interface_default'" />
    <ProcessorGroupDefault v-if="selectedNode.processorType === 'processor_group_default'" />
    <ProcessorTimerDefault  v-if="selectedNode.processorType === 'processor_time_default'" />

    <ProcessorLogicIf   v-if="selectedNode.processorType === 'processor_logic_if'" />
    <ProcessorLogicElse v-if="selectedNode.processorType === 'processor_logic_else'" />

    <ProcessorLoopTime  v-if="selectedNode.processorType === 'processor_loop_time'" />
    <ProcessorLoopUntil v-if="selectedNode.processorType === 'processor_loop_until'" />
    <ProcessorLoopIn    v-if="selectedNode.processorType === 'processor_loop_in'" />
    <ProcessorLoopRange v-if="selectedNode.processorType === 'processor_loop_range'" />
    <ProcessorLoopBreak v-if="selectedNode.processorType === 'processor_loop_break'" />

    <ProcessorVariableSet   v-if="selectedNode.processorType === 'processor_variable_set'" />
    <ProcessorVariableClear v-if="selectedNode.processorType === 'processor_variable_clear'" />

    <ProcessorAssertionDefault      v-if="selectedNode.processorCategory === 'processor_assertion'" />

    <ProcessorExtractorBoundary  v-if="selectedNode.processorType === 'processor_extractor_boundary'" />
    <ProcessorExtractorJsonQuery v-if="selectedNode.processorType === 'processor_extractor_jsonquery'" />
    <ProcessorExtractorHtmlQuery v-if="selectedNode.processorType === 'processor_extractor_htmlquery'" />
    <ProcessorExtractorXmlQuery  v-if="selectedNode.processorType === 'processor_extractor_xmlquery'" />

    <ProcessorCookieGet   v-if="selectedNode.processorType === 'processor_cookie_get'" />
    <ProcessorCookieSet   v-if="selectedNode.processorType === 'processor_cookie_set'" />
    <ProcessorCookieClear v-if="selectedNode.processorType === 'processor_cookie_clear'" />

    <ProcessorDataText    v-if="selectedNode.processorType === 'processor_data_text'" />
    <ProcessorDataExcel   v-if="selectedNode.processorType === 'processor_data_excel'" />
    <ProcessorDataZenData v-if="selectedNode.processorType === 'processor_data_zendata'" />

  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
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

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const selectedNode = computed<any>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.scenario-node-edit-main {
  height: 100%;
  padding: 0;
}

</style>