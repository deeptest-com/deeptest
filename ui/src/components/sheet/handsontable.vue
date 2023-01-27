<template>
  <div class="handsontable-main">
    <hot-table :settings="settings"
               :data="data"
    ></hot-table>
  </div>
</template>

<script setup lang="ts">
import { HotTable } from '@handsontable/vue3';
import {registerAllModules} from "handsontable/registry";
import {
  registerLanguageDictionary,
  zhCN,
} from 'handsontable/i18n';
import {defineProps} from "vue";
import {useI18n} from "vue-i18n";

registerAllModules();
registerLanguageDictionary(zhCN);

const { t } = useI18n();

const props = defineProps({
  headers: {
    required: true
  },
  data: {
    required: true
  }
})

const settings = {
  licenseKey: 'non-commercial-and-evaluation',
  rowHeaders: [],
  colHeaders: props.headers,
  height: 'auto',
  contextMenu: true,
  language: t('lang'),
}

</script>

<style lang="less">
.handsontable-main {
  height: 100%;
  overflow-y: auto;

  .wtHolder {
    min-width: 100%;
  }
}

</style>

<style src="../../../node_modules/handsontable/dist/handsontable.css"></style>