<template>
  <div class="handsontable-main">
    <hot-table :data="data" :rowHeaders="true" :colHeaders="true"></hot-table>

  </div>
</template>

<script setup>
import { HotTable } from '@handsontable/vue3';
import {registerAllModules} from "handsontable/registry";

registerAllModules();

const data = [
  ['', 'Ford', 'Volvo', 'Toyota', 'Honda'],
  ['2016', 10, 11, 12, 13],
  ['2017', 20, 11, 14, 13],
  ['2018', 30, 15, 12, 13]
];
</script>

<style scoped>
.handsontable-main {

}
</style>

<style src="../../../node_modules/handsontable/dist/handsontable.full.css"></style>