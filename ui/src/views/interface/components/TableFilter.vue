<template>
  <a-row type="flex" :gutter="16" justify="space-between" style="width: 100%">
    <a-col :span="8">
      <a-form-item label="创建人" style="margin-bottom: 0">
        <a-select
            @change="(e) => {
              handleFilterChange('createUser',e);
            }"
            placeholder="请选择创建人"
            :options="userList"
            option-label-prop="username"
            :value="formState?.createUser"/>
      </a-form-item>
    </a-col>
    <a-col :span="8">
      <a-form-item label="状态" style="margin-bottom: 0;">
        <a-select
            @change="(e) => {
              handleFilterChange('status',e);
            }"
            :value="formState?.status"
            placeholder="请选择状态"
            :options="interfaceStatusOpts"/>
      </a-form-item>
    </a-col>
    <a-col :span="8">
      <a-input-search
          style="display: flex;justify-content: end;"
          placeholder="请输入关键词"
          enter-button
          :value="formState?.title"
          @change="(e) => {
              handleFilterChange('title',e);
            }"
          @search="async () => {
            await handleFilter()
          }"/>
    </a-col>
  </a-row>
</template>

<script lang="ts" setup>
import {interfaceStatusOpts} from '@/config/constant';
import {filterFormState} from "../data";
import {
  defineEmits, ref,
  onMounted, computed, watch, Ref
} from 'vue';

const store = useStore<{ Interface, ProjectGlobal, Project }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
let userList = computed<any>(() => store.state.Project.userList);
let filterState = computed<any>(() => store.state.Interface.filterState);

import {useStore} from "vuex";

const emit = defineEmits(['filter']);

const formState: Ref<filterFormState> = ref({
  "status": "",
  "createUser": "",
  "title": ""
});

async function handleFilterChange(type, e) {
  if (type === 'status') {
    formState.value.status = e;
    await handleFilter();
  }
  if (type === 'createUser') {
    formState.value.createUser = e;
    await handleFilter();
  }
  if (type === 'title') {
    formState.value.title = e.target.value;
    // await handleFilter();
  }
}

async function handleFilter() {
  emit('filter', {
    ...filterState.value,
    ...formState.value
  });
}

watch(() => {
  return filterState.value
}, (newVal) => {
  formState.value = {...newVal}
}, {
  immediate: true,
})

onMounted(async () => {
  await store.dispatch('Project/getUserList');
})

</script>

<style lang="less" scoped>
.requireActived {
  color: #0000cc;
}
</style>
