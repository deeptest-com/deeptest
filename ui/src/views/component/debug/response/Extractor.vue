<template>
  <div class="response-extractor-main">
    <div class="head">
      <a-row type="flex" class="extractor item">
        <a-col flex="50px">编号</a-col>
        <a-col flex="70px">来源</a-col>
        <a-col flex="90px">提取类型</a-col>
        <a-col flex="1">表达式</a-col>
        <a-col flex="100px" style="padding-left: 10px;">变量</a-col>
        <a-col flex="1">结果</a-col>

        <a-col flex="100px" class="dp-right">
          <PlusOutlined @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="items">
      <a-row v-for="(item, idx) in extractorsData" :key="idx" type="flex" class="item">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="70px">{{ t(item.src) }}</a-col>
        <a-col flex="90px">{{ item.type ? t(item.type) : '' }}</a-col>
        <a-col flex="1"  style="width: 0; word-break: break-word;">
          <span v-if="item.src === ExtractorSrc.header">
            {{ item.key }}
          </span>
          <span v-if="item.src === ExtractorSrc.body">
            {{ item.type === ExtractorType.boundary ?
                `${item.boundaryStart}-${item.boundaryEnd}[${item.boundaryIndex}] ${item.boundaryIncluded}` :
                item.expression }}
          </span>
        </a-col>
        <a-col flex="100px" style="padding-left: 10px;">{{ item.variable }}</a-col>
        <a-col flex="1" :class="[item.result==='extractor_err'? 'dp-color-fail': '']"  style="width: 0; word-break: break-word;">
          {{item.result==='extractor_err'? t(item.result) : item.result}}
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip v-if="!item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip v-if="item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
          </a-tooltip>

          <EditOutlined @click.stop="edit(item)" class="dp-icon-btn dp-trans-80" />
          <DeleteOutlined @click.stop="remove(item)" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <a-modal
        :title="modelRef.id ? '编辑' : '创建' + '提取器'"
        :destroy-on-close="true"
        :mask-closable="false"
        :visible="editVisible"
        :onCancel="cancel"
        :footer="null"
        width="700px">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="数据来源" v-bind="validateInfos.src">
            <a-radio-group name="srcGroup" @change="selectSrc" v-model:value="modelRef.src"
                           @blur="validate('src', { trigger: 'change' }).catch(() => {})">
              <a-radio v-for="(item, idx) in srcOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-radio>
            </a-radio-group>

          </a-form-item>

          <!-- for body -->
          <a-form-item v-if="modelRef.src === 'body'" label="提取方法" v-bind="validateInfos.type">
            <a-select v-model:value="modelRef.type" @change="selectType"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in typeOptions" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <!-- for header -->
          <a-form-item v-if="modelRef.src === 'header'" label="键值" v-bind="validateInfos.key">
            <a-input v-model:value="modelRef.key"
                     @blur="validate('key', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <template v-if="modelRef.src === 'body' && modelRef.type === 'boundary'">
            <a-form-item label="边界开始" v-bind="validateInfos.boundaryStart">
              <a-input v-model:value="modelRef.boundaryStart"
                       @blur="validate('boundaryStart', { trigger: 'blur' }).catch(() => {})" />
            </a-form-item>
            <a-form-item  label="边界结束" v-bind="validateInfos.boundaryEnd">
              <a-input v-model:value="modelRef.boundaryEnd"
                       @blur="validate('boundaryEnd', { trigger: 'blur' }).catch(() => {})" />
            </a-form-item>
            <a-form-item  label="索引值">
              <a-input-number v-model:value="modelRef.boundaryIndex" />
            </a-form-item>
            <a-form-item  label="是否包含边界">
              <a-switch v-model:checked="modelRef.boundaryIncluded" />
            </a-form-item>
          </template>

          <a-form-item v-if="modelRef.src === 'body' && modelRef.type !== 'boundary'"
                       :label="modelRef.type==='regx'?'表达式':'XPath'" v-bind="validateInfos.expression">
            <a-input v-model:value="modelRef.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="变量名称" v-bind="validateInfos.variable">
            <a-input-group compact>
              <a-input v-model:value="modelRef.variable"
                       @change="onVarChanged"
                       @blur="validate('variable', { trigger: 'blur' }).catch(() => {})"
                       style="width: 72%"/>

              <a-select v-model:value="modelRef.code"
                        @change="onVarSelected"
                        style="width: 28%">
                <a-select-option value="">
                  选择变量
                </a-select-option>

                <a-select-option v-for="(item, idx) in debugData.shareVars"
                                 :key="idx"
                                 :value="item.id + '-' + item.name">
                  {{item.name}}
                </a-select-option>
              </a-select>
            </a-input-group>
          </a-form-item>

          <a-form-item label="变量作用域">
            <a-radio-group v-model:value="modelRef.scope">
              <a-radio value="public">公有</a-radio>
              <a-radio value="private">私有</a-radio>
            </a-radio-group>
            <div class="dp-input-tip">
              公有变量在接口所在服务及场景下有效。
            </div>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <a-button type="primary" @click="save" class="dp-btn-gap">保存</a-button> &nbsp;
            <a-button @click="cancel" class="dp-btn-gap">取消</a-button>
          </a-form-item>

        </a-form>

      </div>

    </a-modal>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, reactive, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Form} from 'ant-design-vue';
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  DeleteOutlined,
  EditOutlined,
  PlusOutlined
} from '@ant-design/icons-vue';
import {VarScope} from "@/utils/enum";
import {ExtractorSrc, ExtractorType, UsedBy} from "@/utils/enum";

const usedBy = inject('usedBy') as UsedBy
const useForm = Form.useForm;
const {t} = useI18n();

const srcOptions = getEnumSelectItems(ExtractorSrc)
const typeOptions = getEnumSelectItems(ExtractorType)

import {Extractor, Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/utils/comm";
const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);
const extractorsData = computed(() => store.state.Debug.extractorsData);

const listExtractor = () => {
  store.dispatch('Debug/listExtractor')
}

watch(debugData, () => {
  console.log('watch debugData', debugData.value.id, usedBy)
  listExtractor()
}, {immediate: true, deep: true})

const modelRef = ref({variable: '', code: '', scope: VarScope.ScopePublic, usedBy: usedBy} as any)
const results = ref({})
const editVisible = ref(false)

const typeRequired = { required: true, message: '请选择类型', trigger: 'change' }
const expressionRequired = { required: true, message: '请输入元素路径', trigger: 'blur' }
const keyRequired = { required: true, message: '请输入键值', trigger: 'blur' }
const boundaryStartRequired = { required: true, message: '请输入边界开始字符串', trigger: 'blur' }
const boundaryEndRequired = { required: true, message: '请输入边界结束字符串', trigger: 'blur' }

const rules = reactive({
  src: [
    { required: true, message: '请选择来源', trigger: 'change' },
  ],
  type: [
    typeRequired,
  ],
  expression: [
    expressionRequired,
  ],
  key: [
    keyRequired,
  ],
  boundaryStart: [
    boundaryStartRequired,
  ],
  boundaryEnd: [
    boundaryEndRequired,
  ],
  variable: [
    { required: true, message: '请输入变量名', trigger: 'blur' },
  ],
});

const { resetFields, validate, validateInfos } = useForm(modelRef, rules);

const add = () => {
  editVisible.value = true
  modelRef.value = {
    src: ExtractorSrc.body,
    type: ExtractorType.boundary,
    expression: '',
    variable: '',
    code: '',
    scope: VarScope.ScopePublic,
    usedBy: usedBy} as Extractor

  selectSrc()
  if (responseData.value.contentLang === 'json') {
    modelRef.value.type = ExtractorType.jsonquery
  } else if (responseData.value.contentLang === 'xml') {
    modelRef.value.type = ExtractorType.xmlquery
  } else if (responseData.value.contentLang === 'html') {
    modelRef.value.type = ExtractorType.htmlquery
  }
}

const edit = (item) => {
  console.log('edit')
  modelRef.value = item
  editVisible.value = true

  selectSrc()
}

const save = () => {
  console.log('save')
  validate().then(() => {
    modelRef.value.debugInterfaceId = debugInfo.value.debugInterfaceId
    modelRef.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
    modelRef.value.projectId = debugData.value.projectId

    store.dispatch('Debug/saveExtractor', modelRef.value).then((result) => {
      if (result) {
        editVisible.value = false
      }
    })
  })
}

const cancel = () => {
  console.log('cancel')
  editVisible.value = false
}

const remove = (item) => {
  console.log('remove')
  store.dispatch('Debug/removeExtractor', {id: item.id})
}

const disable = (item) => {
  console.log('disabled')
  item.disabled = !item.disabled
  store.dispatch('Debug/saveExtractor', item)
}

const selectSrc = () => {
  console.log('selectSrc', modelRef.value.src)

  if (modelRef.value.src === ExtractorSrc.header) {
    rules.key = [keyRequired]
    rules.expression = []
    rules.type = []
  } else {
    rules.key = []
    rules.expression = [expressionRequired]
    rules.type = [typeRequired]
  }

  selectType()
}
const selectType = () => {
  console.log('selectType', modelRef.value.type)

  if (modelRef.value.type === ExtractorType.boundary) {
    rules.boundaryStart = [boundaryStartRequired]
    rules.boundaryEnd = [boundaryEndRequired]
    rules.expression = []
  } else {
    rules.boundaryStart = []
    rules.boundaryEnd = []
    rules.expression = [expressionRequired]
  }
}

const onVarChanged = (e) => {
  console.log('onVarChanged', e)

  const value = e.target.value.trim()

  if (!value) {
    modelRef.value.code = ''
    return
  }

  let found = false
  for (let i in debugData.value.shareVars) {
    const item = debugData.value.shareVars[i]

    if (value === item.name) {
      modelRef.value.code = item.id + '-' + item.name
      found = true
      break
    }
  }

  if (!found) {
    modelRef.value.code = ''
  }
};

const onVarSelected = (value) => {
  console.log('onVarSelected')

  const arr = value.split('-')
  modelRef.value.variable = arr[1]
};

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>

<style lang="less">
</style>

<style lang="less" scoped>
.response-extractor-main {
  height: 100%;

  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .items {
    padding: 6px;
    height: calc(100% - 30px);
    overflow-y: auto;

    .item {
      .ant-col {
        padding: 0 3px;
        word-break: break-all;
      }
    }
  }
}
</style>