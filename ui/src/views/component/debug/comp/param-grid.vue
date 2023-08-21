<template>
  <div class="param-grid-main">
    <div v-if="title"  class="head">
      <ConBoxTitle :title="title" />
    </div>

    <div class="items">
      <a-table v-if="list?.length > 0" :dataSource="list || []" :columns="columns"
               :rowKey="(record, index) => {return index}"
               :tableLayout="'fixed'"
               class="dp-small-table"
               :rowClassName="(record, index) => {return record.name==='' ? 'hidden' : ''}"
               :pagination="false"/>
      <div v-else>空</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {defineProps, PropType} from "vue";
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';

const {t} = useI18n();

const props = defineProps({
  title: {
    type: String,
    required: false
  },
  list: {
    type: Array,
    required: true
  },
})

const columns = [
  {
    title: '参数名',
    dataIndex: 'name',
  },
  {
    title: '参数值',
    dataIndex: 'value',
  }
];

</script>

<style lang="less" scoped>
.param-grid-main {
  margin-bottom: 16px;
  .head {
  }
  .items {

  }
  .empty {
    padding: 10px 10px;
  }
}

</style>
