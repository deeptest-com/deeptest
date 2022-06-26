<template>
  <div class="scenario-node-edit-main">
<!--    <ProcessorThreadDefault v-if="selectedNode.processorType === 'processor_thread_default'" />-->
    <ProcessorGroupDefault v-if="selectedNode.processorType === 'processor_group_default'" />
    <ProcessorTimerDefault  v-if="selectedNode.processorType === 'processor_time_default'" />

    <ProcessorLogicIf   v-if="selectedNode.processorType === 'processor_logic_if'" />
    <ProcessorLogicElse v-if="selectedNode.processorType === 'processor_logic_else'" />

    <ProcessorLoopRepeatTime  v-if="selectedNode.processorType === 'processor_loop_repeat_time'" />
    <ProcessorLoopRepeatUntil v-if="selectedNode.processorType === 'processor_loop_repeat_until'" />
    <ProcessorLoopRepeatIn    v-if="selectedNode.processorType === 'processor_loop_repeat_in'" />
    <ProcessorLoopRepeatRange v-if="selectedNode.processorType === 'processor_loop_range'" />
    <ProcessorLoopRepeatBreak v-if="selectedNode.processorType === 'processor_loop_break'" />

    <ProcessorVariableSet   v-if="selectedNode.processorType === 'processor_variable_set'" />
    <ProcessorVariableGet   v-if="selectedNode.processorType === 'processor_variable_get'" />
    <ProcessorVariableClear v-if="selectedNode.processorType === 'processor_variable_clear'" />

    <ProcessorAssertionEqual      v-if="selectedNode.processorType === 'processor_assertion_equal'" />
    <ProcessorAssertionNotEqual   v-if="selectedNode.processorType === 'processor_assertion_not_equal'" />
    <ProcessorAssertionContain    v-if="selectedNode.processorType === 'processor_assertion_contain'" />
    <ProcessorAssertionNotContain v-if="selectedNode.processorType === 'processor_assertion_not_contain'" />

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
import ProcessorTimerDefault from "./proccessors/timer/default.vue"

import ProcessorLogicIf  from "./proccessors/logic/if.vue"
import ProcessorLogicElse from "./proccessors/logic/else.vue"

import ProcessorLoopRepeatTime from "./proccessors/loop/repeat_time.vue"
import ProcessorLoopRepeatUntil from "./proccessors/loop/repeat_until.vue"
import ProcessorLoopRepeatIn  from "./proccessors/loop/repeat_in.vue"
import ProcessorLoopRepeatRange from "./proccessors/loop/repeat_range.vue"
import ProcessorLoopRepeatBreak from "./proccessors/loop/repeat_break.vue"

import ProcessorVariableSet  from "./proccessors/variable/set.vue"
import ProcessorVariableGet  from "./proccessors/variable/get.vue"
import ProcessorVariableClear from "./proccessors/variable/clear.vue"

import ProcessorAssertionEqual     from "./proccessors/assertion/equal.vue"
import ProcessorAssertionNotEqual  from "./proccessors/assertion/not_equal.vue"
import ProcessorAssertionContain   from "./proccessors/assertion/contain.vue"
import ProcessorAssertionNotContain from "./proccessors/assertion/not_contain.vue"

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
const selectedNode = computed<boolean>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.scenario-node-edit-main {
  padding: 16px;
}

</style>