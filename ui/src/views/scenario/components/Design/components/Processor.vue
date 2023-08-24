<template>
  <div class="scenario-processor-edit-main dp-processors-container">
    <!--    <ProcessorThreadDefault v-if="selectedNode.processorType === 'processor_thread_default'" />-->

    <ProcessorInterfaceDefault v-if="selectedNode.processorType === 'processor_interface_default'" />
    <ProcessorGroupDefault v-else-if="selectedNode.processorType === 'processor_group_default'" />
    <ProcessorTimerDefault  v-else-if="selectedNode.processorType === 'processor_time_default'" />
    <ProcessorPrintDefault  v-else-if="selectedNode.processorType === 'processor_print_default'" />

    <ProcessorLogicIf   v-else-if="selectedNode.processorType === 'processor_logic_if'" />
    <ProcessorLogicElse v-else-if="selectedNode.processorType === 'processor_logic_else'" />

    <ProcessorLoopTime  v-else-if="selectedNode.processorType === 'processor_loop_time'" />

    <ProcessorLoopUntil v-else-if="selectedNode.processorType === 'processor_loop_until'" />
    <ProcessorLoopIn    v-else-if="selectedNode.processorType === 'processor_loop_in'" />
    <ProcessorLoopRange v-else-if="selectedNode.processorType === 'processor_loop_range'" />

    <ProcessorVariableSet   v-else-if="selectedNode.processorType === 'processor_variable_set'" />
    <ProcessorVariableClear v-else-if="selectedNode.processorType === 'processor_variable_clear'" />

    <ProcessorAssertionDefault      v-else-if="selectedNode.processorCategory === 'processor_assertion'" />

    <ProcessorExtractorBoundary  v-else-if="selectedNode.processorType === 'processor_extractor_boundary'" />
    <ProcessorExtractorJsonQuery v-else-if="selectedNode.processorType === 'processor_extractor_jsonquery'" />
    <ProcessorExtractorHtmlQuery v-else-if="selectedNode.processorType === 'processor_extractor_htmlquery'" />
    <ProcessorExtractorXmlQuery  v-else-if="selectedNode.processorType === 'processor_extractor_xmlquery'" />

    <ProcessorCookieGet   v-else-if="selectedNode.processorType === 'processor_cookie_get'" />
    <ProcessorCookieSet   v-else-if="selectedNode.processorType === 'processor_cookie_set'" />
    <ProcessorCookieClear v-else-if="selectedNode.processorType === 'processor_cookie_clear'" />

    <ProcessorDataDefault    v-else-if="selectedNode.processorType === 'processor_data_default'" />
    <ProcessorCustomCode v-else-if="selectedNode.processorType === 'processor_custom_code'" />

    <span v-else>
      <a-empty style="margin-top: 100px;" :description="'请先在左侧目录上选择编排场景'"/>
    </span>

  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import {StateType as ScenarioStateType} from "../../../store";

import ProcessorGroupDefault from "./proccessors/group/default.vue";
import ProcessorInterfaceDefault from "./proccessors/interface/default.vue"
import ProcessorTimerDefault from "./proccessors/timer/default.vue"
import ProcessorPrintDefault from "./proccessors/print/default.vue"

import ProcessorLogicIf  from "./proccessors/logic/if.vue"
import ProcessorLogicElse from "./proccessors/logic/else.vue"

import ProcessorLoopTime from "./proccessors/loop/time.vue"
import ProcessorLoopUntil from "./proccessors/loop/until.vue"
import ProcessorLoopIn  from "./proccessors/loop/in.vue"
import ProcessorLoopRange from "./proccessors/loop/range.vue"

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

import ProcessorDataDefault   from "./proccessors/data/default.vue"
import ProcessorCustomCode from "./proccessors/custom_code/default.vue"

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const selectedNode = computed<any>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less">

</style>

<style lang="less" scoped>
.scenario-processor-edit-main {
  height: 100%;

  :deep(.ant-form) {
    .ant-row.ant-form-item.processor-btn {
      .ant-form-item-control-input > .ant-form-item-control-input-content {
        display: flex;
        justify-content: flex-end;
      }
    }
  }
}
</style>
