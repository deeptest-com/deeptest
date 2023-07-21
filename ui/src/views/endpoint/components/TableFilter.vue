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
            style="width: 200px;"
            allowClear
            @change="(e) => {
              handleFilterChange('tagNames',e);
            }"
            :value="formState?.tagNames"
            placeholder="请选择标签"
            max-tag-count="responsive"
            :options="tagList"/>
            <template v-slot:max-tag-placeholder>
        数据同步
        <a-tooltip placement="topLeft" arrow-point-at-center overlayClassName="memo-tooltip">
          <template v-slot:title>
            <span class="title">完全覆盖</span><br>
            通过swagger导入/同步的接口定义，同步更新时使用接口方法和路径进行匹配。<br>
            匹配到的相同接口同步时不保留平台中的旧数据，完全使用swagger文档中的新数据进行覆盖。<br>
            通过平台创建的接口定义不会被覆盖。<br>
         </template>
        <QuestionCircleOutlined class="icon" style=" font-size: 14px;transform: scale(0.9)" />
        </a-tooltip>
      </template>
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
const maxTagPlaceholder = (num) => {
  console.log(num,"++++")
                return 'more ';
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
