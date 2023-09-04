<template>
  <a-typography-text type="secondary">
    {{typesLabel}}
    <LinkOutlined v-if="props?.value?.ref"/>
  </a-typography-text>
</template>
<script lang="ts" setup>
import {ref, defineProps, defineEmits, watch, reactive, toRaw, computed, onMounted} from 'vue';

import {
  LinkOutlined
} from '@ant-design/icons-vue';

const props = defineProps(['value'])
const emit = defineEmits([]);
// 返回，如何展示类型
const typesLabel: any = computed(() => {
  let {type, types} = props.value || {};
  type = props?.value?.name || type || '';
  if (!type) {
    return '';
  }
  // // 引用类型
  // if (props?.value?.ref) {
  //   return props?.value?.name
  // }
  const labels = Array.isArray(types) ? [...types, type] : [type];
  const result = labels.reduceRight((acc, cur, index) => {
    if (index === labels.length - 1) {
      return [cur];
    }
    return [cur, acc];
  }, []);
  return JSON.stringify(result).replace(/[",]/g, '').replace(/^\[/, '').replace(/\]$/, '');
});
</script>

<style lang="less" scoped>

</style>
