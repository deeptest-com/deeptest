<template>
  <div class="case-generate-main">
    <div class="toolbar">
      <div class="left">
        <a-button @click="back" size="small" class="btn">
          <template #icon>
            <icon-svg class="" type="back"/>
          </template>
          返回
        </a-button>
      </div>
      <div class="right">
        <a-button @click="saveAsCase">
          另存为用例
        </a-button>
      </div>
    </div>

    <a-form :label-col="{ span: 3 }"
            :wrapper-col="{ span: 20 }">

      <a-form-item label="备选路径">
        <a-tree
            :replaceFields="replaceFields"
            :tree-data="alternativeCases"
            :expandedKeys="expandedKeys"
            :checkable="true"
            v-model:checkedKeys="checkedKeys"
            :show-icon="true">
          <template #title="nodeProps">
            <span class="tree-title">
              <span>{{ nodeProps.title }}</span>
              <template v-if="nodeProps.category==='case'">
                <span>: &nbsp;&nbsp;&nbsp;</span>

                <span v-if="treeDataMap[nodeProps.key]?.isEdit">
                  <a-input size="small"
                           :style="{width: '160px'}"
                           v-model:value="sampleRef"/>
                  &nbsp;
                  <CheckOutlined @click="editFinish(nodeProps.key)" class="dp-icon-btn2 dp-trans-80"/>
                  <CloseOutlined @click="editCancel(nodeProps.key)" class="dp-icon-btn2 dp-trans-80"/>
                </span>

                <span v-else>
                  {{ nodeProps.sample ? nodeProps.sample : '空' }}
                  &nbsp;
                  <EditOutlined @click="editStart(nodeProps.key)"/>
                </span>

              </template>
            </span>
          </template>

          <template #icon="slotProps">
            <FolderOutlined v-if="slotProps.isDir && !slotProps.expanded"/>
            <FolderOpenOutlined v-if="slotProps.isDir && slotProps.expanded"/>
            <FileOutlined v-if="!slotProps.isDir"/>
          </template>
        </a-tree>
      </a-form-item>
    </a-form>

    <SaveAlternative
        :visible="saveAsVisible"
        :onClose="saveAsClosed"
        :model="saveAsModel"/>
  </div>
</template>

<script lang="ts" setup>
import {computed, defineProps, inject, reactive, ref, watch} from 'vue';
import {UsedBy} from "@/utils/enum";
import {Form} from "ant-design-vue";
import {useStore} from "vuex";
import IconSvg from "@/components/IconSvg";
import {
  CheckOutlined,
  CloseOutlined,
  EditOutlined,
  FileOutlined,
  FolderOpenOutlined,
  FolderOutlined
} from '@ant-design/icons-vue';

import {Endpoint} from "@/views/endpoint/data";
import {StateType as EndpointStateType} from "@/views/endpoint/store";
import SaveAlternative from "./saveAlternative.vue";

const useForm = Form.useForm;

const store = useStore<{ Endpoint: EndpointStateType }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const alternativeCases = computed<any>(() => store.state.Endpoint.alternativeCases);

const sampleRef = ref('')
const treeDataMap = ref({})

watch(alternativeCases, (newVal) => {
  getNodeMap({key: '', children: newVal}, treeDataMap.value)
}, {deep: true, immediate: true});

const props = defineProps({
  model: {
    required: true,
    type: Object,
  },
  onBack: {
    type: Function,
    required: true,
  }
})

const modelRef = ref({
  baseId: 0,
  prefix: '异常路径',
});

const replaceFields = {key: 'key'};
const expandedKeys = ref<string[]>([]);
const checkedKeys = ref<string[]>([])

const loadCaseTree = () => {
  store.dispatch('Endpoint/loadAlternativeCases', modelRef.value.baseId).then((result) => {
    console.log('loadCaseTree', result)
    expandAll()
  })
  store.dispatch('Endpoint/loadAlternativeCasesSaved', modelRef.value.baseId)
}

function expandAll() {
  const keys: any = [];
  const data = alternativeCases.value;

  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    arr.forEach((item, index) => {
      keys.push(item.key);
      if (Array.isArray(item.children)) {
        fn(item.children)
      }
    });
  }

  fn(data);
  expandedKeys.value = keys;
}

watch(() => props.model, () => {
  console.log('watch props.visible', props.model)
  modelRef.value = {
    baseId: props.model.baseId,
    prefix: props.model?.prefix || '异常路径',
  }

  loadCaseTree()
}, {immediate: true, deep: true})

const rulesRef = reactive({

});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const saveAsVisible = ref(false)
const saveAsModel = ref({} as any)
const saveAsCase = () => {
  console.log('saveAsCase', checkedKeys.value)
  saveAsVisible.value = true

  const paths = ref([] as any[])
  checkedKeys.value.forEach((key) => {
    if (treeDataMap.value[key]) {
      paths.value.push(treeDataMap.value[key].path)
    }
  })
  const baseId = modelRef.value.baseId
  saveAsModel.value = {paths, baseId}
}
const saveAsClosed = () => {
  saveAsVisible.value = false
  saveAsModel.value = {}
}

const finish = () => {
  console.log('finish', modelRef.value)
  validate().then(() => {
    props.onFinish(modelRef.value)
    resetFields();
  }).catch((error) => console.log('error', error))
}

const cancel = () => {
  console.log('cancel')
  resetFields()
}

const editStart = (key) => {
  console.log('editStart', key)
  resetEdit()
  treeDataMap.value[key].isEdit = true
  sampleRef.value = treeDataMap.value[key].sample
}
const editFinish = async (key) => {
  console.log('editFinish', key, treeDataMap.value[key])

  const item = treeDataMap.value[key]
  const data = {baseId: modelRef.value.baseId, path: item.path}

  const result = await store.dispatch('Endpoint/saveAlternativeCase', data)
  if (result) {
    treeDataMap.value[key].isEdit = false
    treeDataMap.value[key].sample = sampleRef.value
  }
}
const editCancel = (key) => {
  console.log('editCancel', key)
  treeDataMap.value[key].isEdit = false
}

function resetEdit() {
  Object.keys(treeDataMap.value).forEach((key) => {
    treeDataMap.value[key].isEdit = false
  })
}

function getNodeMap(treeNode: any, mp: any) {
  if (!treeNode) return

  treeNode.entity = null
  mp[treeNode.key] = treeNode

  if (treeNode.children) {
    treeNode.children.forEach((item, index) => {
      getNodeMap(item, mp)
    })
  }

  return
}


// const generateCasesFinish = async (model) => {
//   console.log('generateCasesFinish', model, debugData.value.url)
//
//   const data = Object.assign({...model}, debugInfo.value)
//
//   store.commit("Global/setSpinning",true)
//   const res = await store.dispatch('Debug/generateCases', data)
//   store.commit("Global/setSpinning",false)
//
//   if (res === true) {
//     generateCasesVisible.value = false
//
//     notifySuccess(`自动生成用例成功`);
//   } else {
//     notifyError(`自动生成用例保存失败`);
//   }
// }
// const generateCasesCancel = () => {
//   console.log('generateCasesCancel')
//   generateCasesVisible.value = false
// }

const back = () => {
  console.log('back')
  props.onBack()
}

</script>

<style lang="less">
.case-generate-main {
  .toolbar {
    display: flex;
    margin-bottom: 10px;

    .left {
      flex: 1;
    }

    .right {
      width: 100px;
      text-align: right;
    }
  }

  .ant-tree {
    .ant-tree-title {
      height: 24px;

      input {
        height: 24px;
        background-color: white;
      }
    }
  }

}
</style>
