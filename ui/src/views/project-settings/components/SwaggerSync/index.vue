<template>
  <div class="content">
<a-form ref="" :model="formState" :label-col="{ style: { width: '300px' } }" :wrapper-col="{ span: 14 }" :rules="rules">
  <a-form-item class="desc">
    开启Swagger自动同步，系统将从指定的Swagger地址中定时自动同步接口定义到当前项目中
    </a-form-item>
    <a-form-item label="是否开启自动同步">
      <a-switch v-model:checked="formState.switch" :checkedValue="1" :unCheckedValue="2" />
    </a-form-item>
    <a-form-item class="execTime" v-if="formState.switch==1"> 上次更新时间：{{formState.execTime || '-'}}</a-form-item>
    <a-form-item name="syncType" v-if="formState.switch==1">
      <template v-slot:label>
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
      <a-select v-model:value="formState.syncType" :options="syncTypes" />
      完全覆盖会导致通过平台上的接口定义更新被覆盖，请谨慎使用
    </a-form-item>
    <a-form-item label="所属分类" name="categoryId" v-if="formState.switch==1">
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

    <a-form-item v-bind="validateInfos.url" label="项目的swagger文档 URL地址"  v-if="formState.switch==1">
      <a-input  v-model:value="formState.url" type="textarea" placeholder="请输入swagger url地址"/>
    </a-form-item>
    <a-form-item v-bind="validateInfos.cron" label="类cron风格表达式(默认每天更新一次)" v-if="formState.switch==1">
      <a-input  v-model:value="formState.cron" type="textarea" placeholder="请输入Linux定时任务表达式"/>
      <a href="https://wiki.nancalcloud.com/pages/viewpage.action?pageId=5969821" target="_blank">点此查看</a>cron表达式格式说明
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
      <a-button type="primary" @click="onSubmit" :disabled="disbaled">保存</a-button>
    </a-form-item>
  </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed,watch,onMounted,ref } from 'vue';
import type { UnwrapRef } from 'vue';
import {SwaggerSync} from './data';
import {useStore} from "vuex";
import {message} from "ant-design-vue";
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import {Form} from 'ant-design-vue';
import {pattern} from "@/utils/const";
import debounce from "lodash.debounce";
const useForm = Form.useForm;
const store = useStore<{ Endpoint,ProjectGlobal,ProjectSetting }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);

  const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return  data?.[0]?.children || [];
});

let formState: UnwrapRef<SwaggerSync>  = computed<SwaggerSync>(()=>store.state.ProjectSetting.swaggerSyncDetail)

const rules = {
  syncType: [{required: true}],
  categoryId: [{required: true}],
  url: [{required: true,message: '请输入swagger url' ,trigger: 'blur'}],
  cron: [{required: true,pattern:pattern.cron,message: '请正确的linux定时任务表达',trigger: 'blur'}]
};

const {validate,validateInfos  } = useForm(formState, rules);

const onSubmit = () => {

    validate().then(()=>{
      saveSwaggerSync(formState.value)

      message.success('保存成功');
    }).catch(()=>{
      console.log('error:', formState.value);
    })

};

async function saveSwaggerSync(data:SwaggerSync) {
  console.log(data)
  await store.dispatch('ProjectSetting/saveSwaggerSync', data);
}

const syncTypes = [
      { label: '完全覆盖', value: 1 },
    ];


function selectedCategory(value) {
  formState.value.categoryId = value;
}

async function loadCategories() {
  await store.dispatch('Endpoint/loadCategory');
}

onMounted(async () => {
  await loadCategories();
  await store.dispatch('ProjectSetting/getSwaggerSync');
  formState.value.projectId = currProject.value.id
  watch(() => {
     return formState.value;
    }, async () => {
    disbaled.value = false
  }, {
     immediate: false,
     deep:true
  })
})

const disbaled = ref(true)
watch(() => {
  return currProject.value;
}, async (newVal) => {
  if (newVal?.id) {
    await loadCategories();
    await store.dispatch('ProjectSetting/getSwaggerSync');
  }
}, {
  immediate: true
})




</script>

<style scoped lang="less">
.content {
  margin: 20px;
 }
 .title {
    font-size: 16px;
  }
  .execTime {
    padding-left: 300px;
    margin-top: -25px;
}
</style>
