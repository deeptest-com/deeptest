<template>
  <a-modal width="1000px"
           :visible="visible"
           @ok="finish"
           @cancel="cancel"
           :title="'备选用例'"
           class="case-generate-main">

    <div class="toolbar">
      <div class="left"></div>
      <div class="right">
        <a-button @click="saveAsCase">
          另存为用例
        </a-button>
      </div>
    </div>

    <a-form :label-col="{ span: 3 }"
            :wrapper-col="{ span: 20 }">

<!--      <a-form-item label="请求方法" v-bind="validateInfos.method">
        <a-select class="select-method"
                  v-model:value="modelRef.method"
                  @change="onMethodChanged">
          <template v-for="method in Methods">
            <a-select-option v-if="hasDefinedMethod(method)" :value="method" :key="method">
              {{ method }}
            </a-select-option>
          </template>
        </a-select>
      </a-form-item>-->

      <a-form-item label="名称前缀" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.prefix"
                 @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
        <div class="dp-input-tip">
          {{`生成的用例会以"${modelRef.prefix}-"开头`}}
        </div>
      </a-form-item>

      <a-form-item label="备选用例">
        <a-tree
            :replaceFields="replaceFields"
            :tree-data="alternativeCases"
            :expandedKeys="expandedKeys"
            :checkable="true"
            v-model:checkedKeys="checkedKeys"
            :show-icon="true">
          <template #title="nodeProps">
            <span class="tree-title">
              <span>{{ nodeProps.title}}</span>
              <template v-if="nodeProps.category==='case'">
                <span>: &nbsp;&nbsp;&nbsp;</span>

                <span v-if="treeDataMap[nodeProps.key]?.isEdit">
                  <a-input size="small"
                           :style="{width: '160px'}"
                           v-model:value="sampleRef" />
                  &nbsp;
                  <CheckOutlined @click="editFinish(nodeProps.key)" class="dp-icon-btn2 dp-trans-80" />
                  <CloseOutlined @click="editCancel(nodeProps.key)"  class="dp-icon-btn2 dp-trans-80" />
                </span>

                <span v-else>
                  {{ nodeProps.sample ? nodeProps.sample : '空' }}
                  &nbsp;
                  <EditOutlined @click="editStart(nodeProps.key)" />
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
  </a-modal>
</template>

<script lang="ts" setup>
import {computed, defineProps, inject, reactive, ref, watch} from 'vue';
import {Methods, UsedBy} from "@/utils/enum";
import {Form} from "ant-design-vue";
import {useStore} from "vuex";
import {FolderOutlined, FolderOpenOutlined, FileOutlined, CheckOutlined, EditOutlined, CloseOutlined} from '@ant-design/icons-vue';

import {Endpoint} from "@/views/endpoint/data";
import {StateType as EndpointStateType} from "@/views/endpoint/store";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy

const store = useStore<{ Endpoint: EndpointStateType }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const alternativeCases = computed<any>(() => store.state.Endpoint.alternativeCases);

const sampleRef = ref('')
const treeDataMap = ref({})

watch(alternativeCases, (newVal) => {
  getNodeMap({key: '', children: newVal}, treeDataMap.value)
}, {deep: true, immediate: true});

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  model: {
    required: true,
    type: Object,
  },
  onFinish: {
    type: Function,
    required: true,
  },
  onCancel: {
    type: Function,
    required: true,
  },
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

// const onMethodChanged = () => {
//   loadCaseTree()
// }

watch(() => props.visible, () => {
  console.log('watch props.visible', props?.visible)
  modelRef.value = {
    baseId: props?.model.baseId,
    prefix: props?.model?.prefix || '异常路径',
  }

  loadCaseTree()
}, {immediate: true, deep: true})

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
  method: [
    {required: true, message: '请选择请求方法', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const saveAsCase = () => {
  console.log('saveAsCase', checkedKeys.value)
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
  props.onCancel()
}

function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
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

</script>

<style lang="less">
.case-generate-main {
  .ant-modal-content {
    .ant-modal-body {
      padding-top: 10px;
      height: calc(100vh - 266px);
      overflow-y: auto;

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

      .modal-btns {
        display: flex;
        justify-content: flex-end;

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
    }
  }

}
</style>
