<template>
  <div class="handsontable-main">
    <hot-table :settings="settings"
               :data="data"
               :rowHeaders="true"
               :colHeaders="true"
    ></hot-table>
  </div>
</template>

<script setup>
import { HotTable } from '@handsontable/vue3';
import {registerAllModules} from "handsontable/registry";
import {defineProps} from "vue";

registerAllModules();

const props = defineProps({
  data: {
    required: true
  }
})

const settings = {
  licenseKey: 'non-commercial-and-evaluation',
  height: 'auto'
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