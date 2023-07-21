<template>
  <a-descriptions :size="'small'" :title="null" :column="4">
    <a-descriptions-item label="创建人">{{ endpointDetail?.createUser }}</a-descriptions-item>
    <a-descriptions-item label="状态">
      <EditAndShowSelect
          :label="endpointStatus.get(endpointDetail?.status || 0 )"
          :value="endpointDetail?.status"
          :options="endpointStatusOpts"
          @update="handleChangeStatus"/>
    </a-descriptions-item>
    <a-descriptions-item label="标签">
      <Tags
       :options="tagList"
       :size="'small'"
       :values="endpointDetail.tags"
       @updateTags = "(values:[])=>{
          updateTags(values,endpointDetail.id,endpointDetail.projectId)
        }"
      />
    </a-descriptions-item>
    <a-descriptions-item label="描述">
      <EditAndShowField :placeholder="'请输入描述'" :value="endpointDetail?.description || '暂无'"
                        @update="updateDescription"/>
    </a-descriptions-item>
    <a-descriptions-item label="分类">
      <EditAndShowTreeSelect
          :label="categoryLabel"
          :value="endpointDetail?.categoryId"
          :treeData="treeData"
          @update="handleChangeCategory"/>
    </a-descriptions-item>
    <a-descriptions-item label="创建时间">{{ endpointDetail?.createdAt }}</a-descriptions-item>
    <a-descriptions-item label="最近更新">{{ endpointDetail?.updatedAt }}</a-descriptions-item>
  </a-descriptions>
</template>
<script lang="ts" setup>

import {
  defineProps,
  ref,
  defineEmits,
  computed,
onMounted,
} from 'vue';
import {endpointStatusOpts, endpointStatus} from '@/config/constant';
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EditAndShowTreeSelect from '@/components/EditAndShowTreeSelect/index.vue';
import Tags from '../Tags/index.vue';

const store = useStore<{ Endpoint }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const tagList: any = computed(()=>store.state.Endpoint.tagList);
//  const tagList = ref(["aabdd","sddsd"])
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);
const treeData: any = computed(() => {
  return treeDataCategory.value?.[0]?.children || [];
});
const categoryLabel = computed(() => {
  if (!endpointDetail.value?.categoryId) {
    return '未分类'
  }
  const data = treeDataCategory.value?.[0]?.children || [];
  let label = "";
  let hasFind = false;

  // 递归查找目录树的文案
  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    for (let i = 0; i < arr.length; i++) {
      const item = arr[i];
      if (item.id === endpointDetail.value?.categoryId) {
        label = item.name;
        hasFind = true;
      }
      if (Array.isArray(item.children) && !hasFind) {
        fn(item.children)
      }
    }
  }

  fn(data);
  return label;
});

const emit = defineEmits(['changeStatus', 'changeDescription', 'changeCategory']);

function handleChangeStatus(val) {
  emit('changeStatus', val);
}

function handleChangeCategory(val) {
  emit('changeCategory', val);
}

function updateDescription(val: string) {
  emit('changeDescription', val);
}

const updateTags = async (tags :[],id:number,projectId:number)=>{  
   await store.dispatch('Endpoint/updateEndpointTag', {
      id:id,tagNames:tags
    });

  await store.dispatch('Endpoint/loadList', {projectId: projectId});
    
}

</script>

