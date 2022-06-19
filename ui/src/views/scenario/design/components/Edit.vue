<template>
  <div class="scenario-node-edit-main">
<!--    <ProcessorThreadDefault v-if="selectedNode.entityType === 'processor_thread_default'" />-->
    <ProcessorSimpleDefault v-if="selectedNode.entityType === 'processor_simple_default'" />
    <ProcessorTimerDefault  v-if="selectedNode.entityType === 'processor_time_default'" />

    <ProcessorLogicIf   v-if="selectedNode.entityType === 'processor_logic_if'" />
    <ProcessorLogicElse v-if="selectedNode.entityType === 'processor_logic_else'" />

    <ProcessorLoopRepeatTime  v-if="selectedNode.entityType === 'processor_loop_repeat_time'" />
    <ProcessorLoopRepeatUntil v-if="selectedNode.entityType === 'processor_loop_repeat_until'" />
    <ProcessorLoopRepeatIn    v-if="selectedNode.entityType === 'processor_loop_repeat_in'" />
    <ProcessorLoopRepeatRange v-if="selectedNode.entityType === 'processor_loop_range'" />
    <ProcessorLoopRepeatBreak v-if="selectedNode.entityType === 'processor_loop_break'" />

    <ProcessorVariableSet   v-if="selectedNode.entityType === 'processor_variable_set'" />
    <ProcessorVariableGet   v-if="selectedNode.entityType === 'processor_variable_get'" />
    <ProcessorVariableClear v-if="selectedNode.entityType === 'processor_variable_clear'" />

    <ProcessorAssertionEqual      v-if="selectedNode.entityType === 'processor_assertion_equal'" />
    <ProcessorAssertionNotEqual   v-if="selectedNode.entityType === 'processor_assertion_not_equal'" />
    <ProcessorAssertionContain    v-if="selectedNode.entityType === 'processor_assertion_contain'" />
    <ProcessorAssertionNotContain v-if="selectedNode.entityType === 'processor_assertion_not_contain'" />

    <ProcessorExtractorBoundary  v-if="selectedNode.entityType === 'processor_extractor_boundary'" />
    <ProcessorExtractorJsonQuery v-if="selectedNode.entityType === 'processor_extractor_jsonquery'" />
    <ProcessorExtractorHtmlQuery v-if="selectedNode.entityType === 'processor_extractor_htmlquery'" />
    <ProcessorExtractorXmlQuery  v-if="selectedNode.entityType === 'processor_extractor_xmlquery'" />

    <ProcessorCookieGet   v-if="selectedNode.entityType === 'processor_cookie_get'" />
    <ProcessorCookieSet   v-if="selectedNode.entityType === 'processor_cookie_set'" />
    <ProcessorCookieClear v-if="selectedNode.entityType === 'processor_cookie_clear'" />

    <ProcessorDataText    v-if="selectedNode.entityType === 'processor_data_text'" />
    <ProcessorDataExcel   v-if="selectedNode.entityType === 'processor_data_excel'" />
    <ProcessorDataZenData v-if="selectedNode.entityType === 'processor_data_zendata'" />

  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import {StateType as ScenarioStateType} from "../../store";

import ProcessorSimpleDefault from "./proccessors/simple/default.vue";
import ProcessorTimerDefault from "./proccessors/timer/default"

import ProcessorLogicIf  from "./proccessors/logic/if"
import ProcessorLogicElse from "./proccessors/logic/else"

import ProcessorLoopRepeatTime from "./proccessors/loop/repeat_time"
import ProcessorLoopRepeatUntil from "./proccessors/loop/repeat_until"
import ProcessorLoopRepeatIn  from "./proccessors/loop/repeat_in"
import ProcessorLoopRepeatRange from "./proccessors/loop/repeat_range"
import ProcessorLoopRepeatBreak from "./proccessors/loop/repeat_break"

import ProcessorVariableSet  from "./proccessors/variable/set"
import ProcessorVariableGet  from "./proccessors/variable/get"
import ProcessorVariableClear from "./proccessors/variable/clear"

import ProcessorAssertionEqual     from "./proccessors/assertion/equal"
import ProcessorAssertionNotEqual  from "./proccessors/assertion/not_equal"
import ProcessorAssertionContain   from "./proccessors/assertion/contain"
import ProcessorAssertionNotContain from "./proccessors/assertion/not_contain"

import ProcessorExtractorBoundary from "./proccessors/extractor/boundary"
import ProcessorExtractorJsonQuery from "./proccessors/extractor/jsonquery"
import ProcessorExtractorHtmlQuery from "./proccessors/extractor/htmlquery"
import ProcessorExtractorXmlQuery from "./proccessors/extractor/xmlquery"

import ProcessorCookieGet  from "./proccessors/cookie/get"
import ProcessorCookieSet  from "./proccessors/cookie/set"
import ProcessorCookieClear from "./proccessors/cookie/clear"

import ProcessorDataText   from "./proccessors/data/text"
import ProcessorDataExcel  from "./proccessors/data/excel"
import ProcessorDataZenData from "./proccessors/data/zendata"

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const selectedNode = computed<boolean>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.scenario-node-edit-main {
  padding: 16px;
}

</style>