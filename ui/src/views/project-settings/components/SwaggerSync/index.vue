<template>
  <div class="content">
<a-form :model="formState" :label-col="{ style: { width: '300px' } }" :wrapper-col="{ span: 14 }" :rules="rules">
  <a-form-item class="desc">
    开启Swagger自动同步，系统将从指定的Swagger地址中定时自动同步接口定义到项目当前中
    </a-form-item>
    <a-form-item label="是否开启自动同步">
      <a-switch v-model:checked="formState.switch" />
    </a-form-item>
    <a-form-item name="syncType" label="数据同步" v-if="formState.switch">
      <a-select v-model:value="formState.syncType" :options="syncTypes" />
      完全覆盖会导致通过平台上的接口定义更新被覆盖，请谨慎使用
    </a-form-item>
    <a-form-item label="所属分类" name="categoryId" v-if="formState.switch">
        <a-tree-select
            @change="selectedCategory"
            :value="formState.categoryId"
            show-search
            :multiple="false"
            :treeData="treeData"
            style="width: 100%"
            :treeDefaultExpandAll="true"
            :replaceFields="{ title: 'name',value:'id'}"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            placeholder="请选择所属分类"
            allow-clear/>
            所有接口都将同步到该分类目录下
      </a-form-item>
      
    <a-form-item name="url" label="项目的swagger文档 URL地址"  v-if="formState.switch">
      <a-input v-model:value="formState.url" type="textarea" placeholder="请输入swagger url地址"/>
    </a-form-item>
    <a-form-item name="cron" label="类cron风格表达式(默认每天更新一次)" v-if="formState.switch">
      <a-input v-model:value="formState.cron" type="textarea" placeholder="请输入Linux定时任务表达式"/>
      <a>点此查看</a>cron表达式格式说明
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }" v-if="formState.switch">
      <a-button type="primary" @click="onSubmit">保存</a-button>
      <a-button style="margin-left: 10px">取消</a-button>
    </a-form-item>
  </a-form>
  </div>
</template>

<script setup lang="ts">
import { reactive, toRaw,computed,watch,onMounted } from 'vue';
import type { UnwrapRef } from 'vue';
import {SwaggerSync} from './data';
import {useStore} from "vuex";
const store = useStore<{ Endpoint,ProjectGlobal }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);
  const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return  data?.[0]?.children || [];
});


const formState: UnwrapRef<SwaggerSync> = reactive({
    switch: false,
    syncType: 1,
    categoryId:'',
    url: '',
    cron: '23 * * *',
});

const rules = {
  syncType: [{required: true}],
  categoryId: [{required: true}],
  url: [{required: true}]
};

const onSubmit = () => {
    console.log('submit!', toRaw(formState));
};

const syncTypes = [
      { label: '完全覆盖', value: 1 },
    ];
  

function selectedCategory(value) {
  formState.categoryId = value;
}

async function loadCategories() {
  await store.dispatch('Endpoint/loadCategory');
  //expandAll();
}

onMounted(async () => {
  await loadCategories();
 // expandAll();
})

watch(() => {
  return currProject.value;
}, async (newVal) => {
  if (newVal?.id) {
    await loadCategories();
  }
}, {
  immediate: true
})

</script>

<style scoped lang="less">
.content {
  margin: 20px;
 
 }
</style>
