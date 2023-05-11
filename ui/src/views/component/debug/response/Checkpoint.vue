<template>
  <div class="response-checkpoint-main">
    <div class="head">
      <a-row type="flex" class="item">
        <a-col flex="50px">编号</a-col>
        <a-col flex="60px">类型</a-col>
        <a-col flex="150px">变量 / 键值 / 表达式</a-col>
        <a-col flex="60px">运算符</a-col>
        <a-col flex="100px">数值</a-col>
        <a-col flex="80px">实际结果</a-col>
        <a-col flex="60px">状态</a-col>

        <a-col flex="80px" class="dp-right">
          <PlusOutlined v-if="usedBy===UsedBy.InterfaceDebug" @click.stop="add" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <a-row v-for="(item, idx) in checkpointsData" :key="idx" type="flex" class="item">
        <a-col flex="50px">{{idx + 1}}</a-col>
        <a-col flex="60px">{{ t(item.type) }}</a-col>
        <a-col flex="150px">{{ item.type === CheckpointType.extractor ? item.extractorVariable : item.expression }} </a-col>
        <a-col flex="60px">{{ t(item.operator) }}</a-col>
        <a-col flex="100px">{{ item.value }}</a-col>
        <a-col flex="80px" style="width: 0; word-break: break-word;">
          {{ item.actualResult }}
        </a-col>

        <a-col flex="60px" :class="getResultCls(item.resultStatus)">
          {{ item.resultStatus ? t(item.resultStatus) : '' }}
        </a-col>

        <a-col flex="80px" class="dp-right">
          <a-tooltip v-if="!item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>禁用</template>
            <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip v-if="item.disabled" @click="disable(item)" overlayClassName="dp-tip-small">
            <template #title>启用</template>
            <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
          </a-tooltip>

          <EditOutlined v-if="usedBy===UsedBy.InterfaceDebug" @click.stop="edit(item)" class="dp-icon-btn dp-trans-80" />
          <DeleteOutlined v-if="usedBy===UsedBy.InterfaceDebug" @click.stop="remove(item)" class="dp-icon-btn dp-trans-80" />
        </a-col>
      </a-row>
    </div>

    <a-modal
        :title="model.id ? '编辑' : '创建' + '检查点'"
        :destroy-on-close="true"
        :mask-closable="false"
        :visible="editVisible"
        :onCancel="cancel"
        :footer="null"
    >
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="类型" v-bind="validateInfos.type">
            <a-select v-model:value="model.type"
                      @change="selectType"
                      @blur="validate('type', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in types" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.type === 'responseHeader'"
                       :label="model.type === 'responseHeader' ? '键值' : ''"
                       v-bind="validateInfos.expression">
            <a-input v-model:value="model.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item v-if="model.type === 'extractor'" label="变量名称" v-bind="validateInfos.extractorVariable">
            <a-select v-model:value="model.extractorVariable"
                      @blur="validate('extractorVariable', { trigger: 'blur' }).catch(() => {})">
              <a-select-option key="" value=""></a-select-option>
              <a-select-option v-for="(item, idx) in variables" :key="idx" :value="item.name">
                {{ item.name }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item v-if="model.type !== 'judgement'" label="运算符" v-bind="validateInfos.operator">
            {{ void (options = model.type === 'responseStatus' ? operatorsForCode :
              isInArray(model.type, ['responseHeader', 'responseBody']) ? operatorsForString : operators) }}

            <a-select v-model:value="model.operator"
                      @blur="validate('operator', { trigger: 'change' }).catch(() => {})">

              <a-select-option v-for="(item, idx) in options" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>

            </a-select>
          </a-form-item>

          <a-form-item v-if="model.type === 'judgement'" label="判断表达式" v-bind="validateInfos.expression">
            <a-textarea v-model:value="model.expression" :auto-size="{ minRows: 2, maxRows: 5 }"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />

            <div class="dp-input-tip">{{t('tips_expression_bool')}}</div>
          </a-form-item>

          <a-form-item v-if="model.type !== 'judgement'" label="数值" v-bind="validateInfos.value">
            <a-input v-model:value="model.value"
                     @blur="validate('value', { trigger: 'blur' }).catch(() => {})" />
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
import {computed, inject, PropType, reactive, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {message, Form} from 'ant-design-vue';
import { PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';

import {
  getEnumSelectItems,
  listExtractorVariable
} from "@/views/component/debug/service";
import {ComparisonOperator, CheckpointType, UsedBy} from "@/utils/enum";
import {isInArray} from "@/utils/array";
import {getResultCls} from "@/utils/dom"
import {getCompareOptsForRespCode, getCompareOptsForString} from "@/utils/compare";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
import {Checkpoint} from "@/views/component/debug/data";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const checkpointsData = computed(() => store.state.Debug.checkpointsData);

const types = getEnumSelectItems(CheckpointType)
const operators = getEnumSelectItems(ComparisonOperator)
const operatorsForString = getCompareOptsForString()
const operatorsForCode = getCompareOptsForRespCode()

watch(debugData, () => {
  console.log('watch debugData')
  listCheckPoint()
}, {deep: true})

const listCheckPoint = () => {
  store.dispatch('Debug/listCheckpoint')
}
listCheckPoint()

const model = ref({
  type: CheckpointType.responseStatus,
  expression: '',
  usedBy: UsedBy.InterfaceDebug,
  extractorVariable: '',
  operator: ComparisonOperator.equal,
  value: ''} as Checkpoint)

const variables = ref([])
const editVisible = ref(false)

const extractorVariableRequired = { required: true, message: '请选择变量', trigger: 'change' }
const expressionRequired = { required: true, message: '请输入表达式', trigger: 'blur' }
const operatorRequired = { required: true, message: '请选择操作', trigger: 'change' }
const valueRequired = { required: true, message: '请输入数值', trigger: 'blur' }
const rules = reactive({
  type: [
    { required: true, message: '请选择类型', trigger: 'blur' },
  ],
  extractorVariable: [
    extractorVariableRequired
  ],
  expression: [
    expressionRequired
  ],
  operator: [
    operatorRequired,
  ],
  value: [
    valueRequired,
  ],
} as any);

let { resetFields, validate, validateInfos } = useForm(model, rules);

const add = () => {
  console.log('add')
  editVisible.value = true
  model.value = {
    type: CheckpointType.responseStatus,
    expression: '',
    usedBy: UsedBy.InterfaceDebug,
    extractorVariable: '',
    operator: ComparisonOperator.equal,
    value: ''} as Checkpoint
  loadExtractorVariable()
}

const edit = (item) => {
  console.log('edit')
  model.value = item
  editVisible.value = true

  loadExtractorVariable()
}

const save = () => {
  console.log('save')
  validate().then(() => {
    model.value.endpointInterfaceId = debugData.value.endpointInterfaceId

    store.dispatch('Debug/saveCheckpoint', model.value).then((result) => {
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
  store.dispatch('Debug/removeCheckpoint', item.id)
}

const disable = (item) => {
  console.log('disabled')
  item.disabled = !item.disabled
  store.dispatch('Debug/saveCheckpoint', item)
}

const selectType = () => {
  console.log('selectType')

  if (model.value.type === CheckpointType.responseBody) {
    model.value.operator = ComparisonOperator.contain
  } else {
    model.value.operator = ComparisonOperator.equal
  }

  loadExtractorVariable()
}

const loadExtractorVariable = () => {
  if (model.value.type === CheckpointType.responseHeader || model.value.type === CheckpointType.judgement) {
    rules.expression = [expressionRequired]
  } else {
    rules.expression = []
  }

  if (model.value.type === CheckpointType.judgement) {
    rules.operator = []
    rules.value = []

    model.value.operator = ComparisonOperator.empty
    model.value.value = ''
  } else {
    rules.operator = [operatorRequired]
    rules.value = [valueRequired]
    model.value.operator = ComparisonOperator.equal
  }

  if (model.value.type === CheckpointType.extractor) {
    rules.extractorVariable = [extractorVariableRequired]

    listExtractorVariable(debugData.value.id).then((jsn) => {
      variables.value = jsn.data
    })
  } else {
    rules.extractorVariable = []
  }
}

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>

<style lang="less">
.response-checkpoint-main {
}
</style>

<style lang="less" scoped>
.response-checkpoint-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
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