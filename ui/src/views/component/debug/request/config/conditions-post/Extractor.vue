<template>
  <div class="response-extractor-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="数据来源" v-bind="validateInfos.src" required>
        <a-radio-group name="srcGroup" v-model:value="model.src"
                       @blur="validate('src', { trigger: 'change' }).catch(() => {})">
          <a-radio v-for="(item, idx) in srcOptions" :key="idx" :value="item.value">
            {{ t(item.label) }}
          </a-radio>
        </a-radio-group>
      </a-form-item>

      <!-- for body -->
      <a-form-item v-if="model.src === 'body'" label="提取方法" v-bind="validateInfos.type" required>
        <a-select v-model:value="model.type"
                  @blur="validate('type', { trigger: 'change' }).catch(() => {})">
          <a-select-option v-for="(item, idx) in typeOptions" :key="idx" :value="item.value">
            {{ t(item.label) }}
          </a-select-option>
        </a-select>
      </a-form-item>

      <!-- for header -->
      <a-form-item v-if="model.src === 'header'" label="键值" v-bind="validateInfos.key" required>
        <a-input v-model:value="model.key"
                 @blur="validate('key', { trigger: 'blur' }).catch(() => {})"/>
      </a-form-item>

      <template v-if="model.src === 'body' && model.type === 'boundary'">
        <a-form-item label="边界开始" v-bind="validateInfos.boundaryStart" required>
          <a-input v-model:value="model.boundaryStart"
                   @blur="validate('boundaryStart', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>
        <a-form-item label="边界结束" v-bind="validateInfos.boundaryEnd" required>
          <a-input v-model:value="model.boundaryEnd"
                   @blur="validate('boundaryEnd', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>
        <a-form-item label="索引值">
          <a-input-number v-model:value="model.boundaryIndex"/>
        </a-form-item>
        <a-form-item label="是否包含边界">
          <a-switch v-model:checked="model.boundaryIncluded"/>
        </a-form-item>
      </template>

      <a-form-item v-if="model.src === 'body' && model.type !== 'boundary'"
                   :label="model.type==='regx'?'表达式':'XPath'" v-bind="validateInfos.expression" required>
        <a-input v-model:value="model.expression"
                 @blur="validate('expression', { trigger: 'blur' }).catch(() => {})"/>
      </a-form-item>

      <a-form-item label="变量名称" v-bind="validateInfos.variable" required>
        <a-input-group compact>
          <a-input v-model:value="model.variable"
                   @change="onVarChanged"
                   @blur="validate('variable', { trigger: 'blur' }).catch(() => {})"
                   style="width: 72%"/>

          <a-select v-model:value="model.code"
                    @change="onVarSelected"
                    style="width: 28%">
            <a-select-option value="">
              选择变量
            </a-select-option>

            <a-select-option v-for="(item, idx) in debugData.shareVars"
                             :key="idx"
                             :value="item.id + '-' + item.name">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-input-group>
      </a-form-item>

      <a-form-item label="变量作用域">
        <a-radio-group v-model:value="model.scope">
          <a-radio value="public">公有</a-radio>
          <a-radio value="private">私有</a-radio>
        </a-radio-group>
        <div class="dp-input-tip">
          公有变量在接口所在服务及场景下有效。
        </div>
      </a-form-item>

    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onBeforeUnmount, onMounted, reactive, watch, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Form, notification} from 'ant-design-vue';
import {CheckpointType, ComparisonOperator, ExtractorSrc, ExtractorType, UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import {getEnumSelectItems} from "@/utils/comm";
import {NotificationKeyCommon} from "@/utils/const";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{ Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);
const model = computed<any>(() => store.state.Debug.extractorData);

const typeRequired = [{required: true, message: '请选择类型', trigger: 'change'}]
const expressionRequired = [{required: true, message: '请输入元素路径', trigger: 'blur'}]
const keyRequired = [{required: true, message: '请输入键值', trigger: 'blur'}]
const boundaryStartRequired = [{required: true, message: '请输入边界开始字符串', trigger: 'blur'}]
const boundaryEndRequired = [{required: true, message: '请输入边界结束字符串', trigger: 'blur'}]

const isInit = ref(true)
const rules = computed(() => { return {
  src: [
    {required: true, message: '请选择来源', trigger: 'change'},
  ],
  type: model.value.src === ExtractorSrc.header ? [] : typeRequired,
  key: model.value.src === ExtractorSrc.header ? keyRequired : [],

  expression: model.value.src === ExtractorSrc.header || model.value.type === ExtractorType.boundary ? [] : expressionRequired,
  boundaryStart: model.value.src !== ExtractorSrc.header && model.value.type === ExtractorType.boundary ? boundaryStartRequired : [],
  boundaryEnd: model.value.src !== ExtractorSrc.header && model.value.type === ExtractorType.boundary ? boundaryEndRequired : [],

  variable: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
}})


watch(model, (newVal) => {
  if (!isInit.value) return

  isInit.value = false

  if (responseData.value.contentLang === 'json') {
    model.value.type = ExtractorType.jsonquery
  } else if (responseData.value.contentLang === 'xml') {
    model.value.type = ExtractorType.xmlquery
  } else if (responseData.value.contentLang === 'html') {
    model.value.type = ExtractorType.htmlquery
  }
  }, {immediate: true, deep: true}
)

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
const srcOptions = getEnumSelectItems(ExtractorSrc)
const typeOptions = getEnumSelectItems(ExtractorType)

const load = () => {
  console.log('load', props.condition)
  store.dispatch('Debug/getExtractor', props.condition.entityId)
}
load()

let {resetFields, validate, validateInfos} = useForm(model, rules);

const save = () => {
  console.log('save')
  validate().then(() => {
    model.value.debugInterfaceId = debugInfo.value.debugInterfaceId
    model.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
    model.value.projectId = debugData.value.projectId

    store.dispatch('Debug/saveExtractor', model.value).then((result) => {
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

const onVarChanged = (e) => {
  console.log('onVarChanged', e)

  const value = e.target.value.trim()

  if (!value) {
    model.value.code = ''
    return
  }

  let found = false
  for (let i in debugData.value.shareVars) {
    const item = debugData.value.shareVars[i]

    if (value === item.name) {
      model.value.code = item.id + '-' + item.name
      found = true
      break
    }
  }

  if (!found) {
    model.value.code = ''
  }
};

const onVarSelected = (value) => {
  console.log('onVarSelected')

  const arr = value.split('-')
  model.value.variable = arr[1]
}

const labelCol = { span: 4 }
const wrapperCol = { span: 18 }

</script>

<style lang="less" scoped>
.response-extractor-main {
  height: 100%;
  width: 100%;

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