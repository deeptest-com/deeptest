<template>
  <a-modal
      title="提取变量"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="true"
      :onCancel="onCancel"
      :footer="null"
      width="800px"
      height="600px"
  >
    <div>
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item :label="exprType === 'regx'?'正则表达式':'XPath'" v-bind="validateInfos.expression">
          <a-input v-model:value="modelRef.expression"
                   @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" >
            <template #addonAfter>
              <div @click="test" :class="{'dp-disabled':!modelRef.expression}" class="dp-link">测试</div>
            </template>
          </a-input>
        </a-form-item>

        <a-form-item label="变量" v-bind="validateInfos.variable">
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

              <a-select-option v-for="(item, idx) in validExtractorVariablesData"
                               :key="idx"
                               :value="item.id + '-' + item.name">
                {{item.name}}
              </a-select-option>
            </a-select>
          </a-input-group>
        </a-form-item>

        <a-form-item label="变量作用域">
          <a-radio-group v-model:value="modelRef.scope">
            <!--              <a-radio value="private">私有</a-radio>-->
            <a-radio value="local">局部</a-radio>
            <a-radio value="global">全局</a-radio>
          </a-radio-group>
          <div class="dp-input-tip">
            局部变量在整个接口设计器及其诞生的场景目录下有效。
          </div>
        </a-form-item>

        <a-form-item v-if="result" label="结果">
          {{result}}
        </a-form-item>

        <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
          <a-button type="primary" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
          <a-button @click="() => onCancel()" class="dp-btn-gap">取消</a-button>
        </a-form-item>

      </a-form>

    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, defineEmits, onMounted, reactive, ref, Ref, computed} from "vue";
import {message, Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getEnvironment, saveEnvironment} from "@/views/interface/service";
import {useStore} from "vuex";
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as EnvironmentStateType} from "@/store/environment";
import {Interface} from "@/views/interface/data";
const useForm = Form.useForm;

const props = defineProps({
  interfaceId:{
    type: Number,
    required: true
  },
  exprType: {
    String,
    required: true
  },
  expr:{
    type: String,
    required: true
  },
  result:{
    type: String,
    required: true
  },

  onCancel:{
    type: Function,
    required: true
  },
  onFinish:{
    type: Function,
    required: true
  },
  onTest:{
    type: Function,
    required: true
  }
});

const { t } = useI18n();

const store = useStore<{ Interface: InterfaceStateType, Environment: EnvironmentStateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const validExtractorVariablesData = computed(() => store.state.Interface.validExtractorVariablesData);

const modelRef = ref<any>({
  expression: props.expr,
  variable: '',
  scope: 'local',
  code: '',
})

const onVarChanged = (e) => {
  console.log('onVarChanged', e)

  const value = e.target.value.trim()

  if (!value) {
    modelRef.value.code = ''
    return
  }

  let found = false
  for (let i in validExtractorVariablesData.value) {
    const item = validExtractorVariablesData.value[i]

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

const onSubmit = async () => {
  console.log('onSubmit', modelRef.value)
  validate().then(() => {
    props.onFinish(modelRef.value);
  })
}

const test  = async () => {
  console.log('test', modelRef.value)
  if (!modelRef.value.expression) return
  props.onTest(modelRef.value.expression, props.exprType);
}

onMounted(()=> {
  console.log('onMounted')
})

const rulesRef = reactive({
  xpath: [
    { required: true, message: '请输入XPath表达式', trigger: 'blur' },
  ],
  variable: [
    { required: true, message: '请输入变量名', trigger: 'blur' },
  ],
});
const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

// const emit = defineEmits(["update:dialogVisible"]);
// const commit = ()=>{
//   emit('update:dialogVisible',false)
// };

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>