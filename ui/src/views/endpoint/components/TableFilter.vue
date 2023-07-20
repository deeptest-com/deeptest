<template>
  <a-form :layout="'inline'">
    <a-space :size="16">
      <a-form-item label="创建人" style="margin-bottom: 0;">
        <a-select
            allowClear
            @change="(e) => {
              handleFilterChange('createUser',e);
            }"
            placeholder="请选择创建人"
            :options="userList"
            option-label-prop="name"
            style="width: 140px;"
            :value="formState?.createUser"/>
      </a-form-item>
      <a-form-item label="状态" style="margin-bottom: 0;">
        <a-select
            style="width: 120px;"
            allowClear
            @change="(e) => {
              handleFilterChange('status',e);
            }"
            :value="formState?.status"
            placeholder="请选择状态"
            :options="endpointStatusOpts"/>
      </a-form-item>
      <a-form-item label="标签" style="margin-bottom: 0;">
        <a-select
            mode="multiple"
            style="width: 140px;"
            allowClear
            @change="(e) => {
              handleFilterChange('tagNames',e);
            }"
            :value="formState?.tagNames"
            placeholder="请选择标签"
            :options="tagList"/>
      </a-form-item>
      <a-form-item :label="null">
        <a-input-search
            style="display: flex;justify-content: end;width: 250px;"
            placeholder="请输入关键词"
            enter-button
            :value="formState?.title"
            @change="(e) => {
              handleFilterChange('title',e);
            }"
            @search="async () => {
            await handleFilter()
          }"/>
      </a-form-item>
    </a-space>
  </a-form>
</template>

<script lang="ts" setup>
import {endpointStatusOpts} from '@/config/constant';
import {filterFormState} from "../data";
import {
  defineEmits, ref,
  onMounted, computed, watch, Ref
} from 'vue';

const store = useStore<{ Endpoint, ProjectGlobal, Project }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
let userList = computed<any>(() => store.state.Project.userList);
let filterState = computed<any>(() => store.state.Endpoint.filterState);
const tagList: any = computed(()=>store.state.Endpoint.tagList);

import {useStore} from "vuex";

const emit = defineEmits(['filter']);

const formState: Ref<filterFormState> = ref({
  "status": "",
  "createUser": "",
  "title": "",
  "tags":[],
});

async function handleFilterChange(type, e) {
  if (type === 'status') {
    formState.value.status = e;
    await handleFilter();
  }
  if (type === 'tagNames') {
    formState.value.tagNames = e;
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
