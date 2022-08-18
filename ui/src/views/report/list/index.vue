<template>
  <div class="report-main">
    <a-card :bordered="false">
      <template #title>
      </template>
      <template #extra>
        <a-select @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr" class="status-select">
        </a-select>
        <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                        placeholder="输入关键字搜索" style="width:270px;margin-left: 16px;"/>
      </template>

      <div>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {SelectTypes} from 'ant-design-vue/es/select';
import {useStore} from "vuex";

import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal} from "ant-design-vue";

const statusArr = ref<SelectTypes['options']>([
  {
    label: '所有状态',
    value: '',
  },
  {
    label: '启用',
    value: '1',
  },
  {
    label: '禁用',
    value: '0',
  },
]);

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
.report-main {

}
</style>