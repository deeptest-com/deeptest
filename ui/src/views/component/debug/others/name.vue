<template>
  <div class="interface-name-main">
    <icon-svg v-if="showBind" type="link" />&nbsp;
    <span>{{bind}}</span>
    <span>{{processorNode.name }}</span>
  </div>
</template>

<script setup lang="ts">
import {computed, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";

import IconSvg from "@/components/IconSvg";

import {StateType as Debug} from "@/views/component/debug/store";
import {ProcessorInterfaceSrc, UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";

const usedBy = inject('usedBy') as UsedBy

const {t} = useI18n();
const store = useStore<{  Debug: Debug, Scenario: ScenarioStateType; }>();
const debugData = computed<any>(() => store.state.Debug.debugData);
const processorNode = computed<any>(()=> store.state.Scenario.nodeData);

const bind = computed(() => {
  let ret = ''

  const src = debugData.value.processorInterfaceSrc

  if (usedBy !== UsedBy.ScenarioDebug) {
    ret = ''
  } else if (src && src !== ProcessorInterfaceSrc.Custom && src !== ProcessorInterfaceSrc.Curl) {
    ret = '绑定'+t(`processor_interface_src_${debugData.value.processorInterfaceSrc}`) + '：'
  } else {
    ret = ''
  }

  return ret
})

const showBind = computed(() => {
  const src = debugData.value.processorInterfaceSrc
  return  src && src !== ProcessorInterfaceSrc.Custom && src !== ProcessorInterfaceSrc.Curl
})

</script>

<style lang="less">
.interface-name-main {
    padding-top: 16px;
}
</style>
