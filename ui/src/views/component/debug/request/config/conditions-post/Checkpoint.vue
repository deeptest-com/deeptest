<template>
  <div class="response-checkpoint-main">
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

        <div class="dp-input-tip">{{t('tips_expression_bool', {name: '{name}'})}}</div>
      </a-form-item>

      <a-form-item v-if="model.type !== 'judgement'" label="数值" v-bind="validateInfos.value">
        <a-input v-model:value="model.value"
                 @blur="validate('value', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onBeforeUnmount, onMounted, PropType, reactive, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {message, Form, notification} from 'ant-design-vue';
import { PlusOutlined, EditOutlined, DeleteOutlined, CloseCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';

import {
  listExtractorVariable
} from "@/views/component/debug/service";
import {ComparisonOperator, CheckpointType, UsedBy} from "@/utils/enum";
import {isInArray} from "@/utils/array";
import {getResultCls} from "@/utils/dom"
import {getCompareOptsForRespCode, getCompareOptsForString} from "@/utils/compare";
import {StateType as Debug} from "@/views/component/debug/store";
import {Checkpoint} from "@/views/component/debug/data";
import {getEnumSelectItems} from "@/utils/comm";
import {NotificationKeyCommon} from "@/utils/const";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const model = computed<any>(() => store.state.Debug.checkpointData);

const props = defineProps({
  condition: {
    type: Object,
    required: true,
  },
  finish: {
    type: Function,
    required: false,
  },
})

const types = getEnumSelectItems(CheckpointType)
const operators = getEnumSelectItems(ComparisonOperator)
const operatorsForString = getCompareOptsForString()
const operatorsForCode = getCompareOptsForRespCode()

const load = () => {
  console.log('load', props.condition)
  store.dispatch('Debug/getCheckpoint', props.condition.entityId)
}
load()

const variables = ref([])

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

const save = () => {
  console.log('save', model.value)
  validate().then(() => {
    model.value.debugInterfaceId = debugInfo.value.debugInterfaceId
    model.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
    model.value.projectId = debugData.value.projectId

    store.dispatch('Debug/saveCheckpoint', model.value).then((result) => {
      if (result) {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存成功`,
        });
        if (props.finish) {
          props.finish()
        }
      } else {
        notification.error({
          key: NotificationKeyCommon,
          message: `保存失败`,
        });
      }
    })
  })
}
const cancel = () => {
  console.log('cancel')
  if (props.finish) {
    props.finish()
  }
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.eventConditionSave, save);
})
onBeforeUnmount( () => {
  console.log('onBeforeUnmount')
  bus.off(settings.eventConditionSave, save);
})

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
    if (!model.value.operator) model.value.operator = ComparisonOperator.equal
  }

  if (model.value.type === CheckpointType.extractor) {
    rules.extractorVariable = [extractorVariableRequired]

    listExtractorVariable(debugInfo.value).then((jsn) => {
      variables.value = jsn.data
    })
  } else {
    rules.extractorVariable = []
  }
}
loadExtractorVariable()

const labelCol = { span: 4 }
const wrapperCol = { span: 18 }

</script>

<style lang="less" scoped>
.response-checkpoint-main {
  height: 100%;
  width: 100%;
}
</style>