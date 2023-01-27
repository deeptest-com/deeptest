<template>
  <div class="handsontable-main">
    <hot-table :settings="settings"
               ref="wrapperRef"
    ></hot-table>
  </div>
</template>

<script setup lang="ts">
import {ref, watch} from "vue";
import { HotTable } from '@handsontable/vue3';
import {registerAllModules} from "handsontable/registry";
import {
  registerLanguageDictionary,
  zhCN,
} from 'handsontable/i18n';
import {defineProps, onMounted, onUnmounted} from "vue";
import {useI18n} from "vue-i18n";
import {resizeHandler} from "@/utils/dom";

registerAllModules();
registerLanguageDictionary(zhCN);

const { t } = useI18n();

const props = defineProps({
  data: {
    type: Object,
    required: true
  },
})

const settings = {
  licenseKey: 'non-commercial-and-evaluation',
  language: t('lang'),
  height: 'auto',
  contextMenu: true,
  rowHeaders: false,
  colHeaders: false,
  fixedRowsTop: 1,
  data: props.data
}

const wrapperRef = ref()
const handsontableRef = ref()

watch(() => props.data, () => {
  console.log('watch data')

  settings.data = props.data

  handsontableRef.value.updateSettings({
    colHeaders: settings.colHeaders,
    data: settings.data,
  })
  handsontableRef.value.render()
}, {deep: false})

onMounted(() => {
  console.log('onMounted')
  handsontableRef.value = wrapperRef.value.hotInstance;
})

onUnmounted(() => {
  console.log('onUnmounted')
})

</script>

<style lang="less">
.handsontable-main {
  height: 100%;
  overflow-y: auto;

  .wtHolder {
    min-width: 100%;

    tbody {
      tr:first-child td  {
        background-color: #f0f0f0 !important;
      }
    }
  }
}

</style>

<style src="../../../node_modules/handsontable/dist/handsontable.css"></style>