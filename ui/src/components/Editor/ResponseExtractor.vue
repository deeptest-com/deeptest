<template>
  <a-modal
      title="提取变量"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="true"
      :onCancel="onCancel"
      :footer="null"
  >
    <div>
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="XPath" v-bind="validateInfos.xpath">
          <a-input v-model:value="modelRef.xpath"
                   @blur="validate('xpath', { trigger: 'blur' }).catch(() => {})" />
        </a-form-item>

        <a-form-item label="变量" v-bind="validateInfos.varName">
          <a-input-group compact>
            <a-input v-model:value="modelRef.varName" style="width: 72%"
                     @blur="validate('varName', { trigger: 'blur' }).catch(() => {})"  />
            <a-select v-model:value="modelRef.varId" style="width: 28%">
              <a-select-option v-for="(item, idx) in environmentData.vars" :key="idx"
                               :value="item.value">{{item.name}}</a-select-option>

              <a-select-option v-for="(item, idx) in validExtractorVariablesData" :key="idx"
                               :value="item.value">{{item.name}}</a-select-option>
            </a-select>
          </a-input-group>
        </a-form-item>

        <a-form-item label="">
          <a-button type="link">测试</a-button>
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
  content:{
    type: String,
    required: true
  },
  selection:{
    type: Object,
    required: true
  },

  onCancel:{
    type: Function,
    required: true
  },
  onFinish:{
    type: Function,
    required: true
  }
});

const { t } = useI18n();

const store = useStore<{ Interface: InterfaceStateType, Environment: EnvironmentStateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const environmentData = computed<any>(() => store.state.Environment.environmentData); // environmentData.vars
const validExtractorVariablesData = computed(() => store.state.Interface.validExtractorVariablesData);

const rulesRef = reactive({
  xpath: [
    { required: true, message: '请输入XPath表达式', trigger: 'blur' },
  ],
  varName: [
    { required: true, message: '请输入变量名', trigger: 'blur' },
  ],
});

const modelRef = ref<any>({
  xpath: '',
  varName: '',
})

const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

const onSubmit = async () => {
  console.log('onSubmit', modelRef)

  validate().then(async () => {
    store.dispatch('Interface/saveExtractor', modelRef.value).then(() => {
      props.onFinish();
    })
  }).catch(err => { console.log('') })
}

onMounted(()=> {
  console.log('onMounted')

})

// const emit = defineEmits(["update:dialogVisible"]);
// const commit = ()=>{
//   emit('update:dialogVisible',false)
// };

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>