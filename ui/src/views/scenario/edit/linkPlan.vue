<template>
  <div class="select-plan-table">
    <a-table :dataSource="dataSource" :columns="columns" />
  </div>
</template>

<script setup lang="ts">
import {defineComponent, computed, ref, reactive, ComputedRef, defineProps, PropType} from "vue";
import {useRouter} from "vue-router";

import { useI18n } from "vue-i18n";
import {Form} from 'ant-design-vue';
const useForm = Form.useForm;


const router = useRouter();
const { t } = useI18n();

const dataSource = [
  {
    key: '1',
    name: '胡彦斌',
    age: 32,
    address: '西湖区湖底公园1号',
  },
  {
    key: '2',
    name: '胡彦祖',
    age: 42,
    address: '西湖区湖底公园1号',
  },
];
    const columns = [
  {
    title: '姓名',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '年龄',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: '住址',
    dataIndex: 'address',
    key: 'address',
  },
];

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },
  categoryId: {
    type: Number,
    required: true
  },
  onFinish: {
    type: Function as PropType<() => void>,
    required: true
  }
})


</script>

<style lang="less" scoped>
.select-plan-table{

}
</style>
