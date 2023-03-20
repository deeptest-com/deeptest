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
            :value="fieldState.createUser"/>
      </a-form-item>
    </a-col>
    <a-col :span="8">
      <a-form-item label="状态" style="margin-bottom: 0;" :value="fieldState.status">
        <a-select
            @change="(e) => {
              handleFilterChange('status',e);
            }"
            :value="fieldState.status"
            placeholder="请选择状态"
            :options="interfaceStatusOpts"/>
      </a-form-item>
    </a-col>
    <a-col :span="8">
      <a-input-search
          style="display: flex;justify-content: end;"
          placeholder="请输入关键词"
          enter-button
          :value="fieldState.title"
          @change="(e) => {
              handleFilterChange('title',e);
            }"
          @search="() => {
            handleFilter()
          }"/>
    </a-col>
  </a-row>
</template>

<script lang="ts" setup>
import {interfaceStatusOpts} from '@/config/constant';
import {
  ref,
  defineEmits,
  onMounted, computed
} from 'vue';

const store = useStore<{ InterfaceV2, ProjectGlobal, Project }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
let userList = computed<any>(() => store.state.Project.userList);

import {useStore} from "vuex";

const emit = defineEmits(['filter']);


// todo 提到 store 中去
const fieldState = ref({
  "status": null,
  "createUser": null,
  "title": null
});


function handleFilter() {
  emit('filter', fieldState.value);
}

function handleFilterChange(type, e) {
  if (type === 'status') {
    fieldState.value.status = e;
  } else if (type === 'createUser') {
    fieldState.value.createUser = e;
  } else if (type === 'title') {
    fieldState.value.title = e.target.value;
  }
}

onMounted(async () => {
  await store.dispatch('Project/getUserList');
})


</script>

<style lang="less" scoped>
.requireActived {
  color: #0000cc;
}
</style>
