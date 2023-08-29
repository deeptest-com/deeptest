<template>
  <a-modal width="600px"
           :visible="visible"
           @ok="finish"
           @cancel="cancel"
           :title="(!model.id ? '新建' : '修改') + '用例'">
    <a-form :label-col="{ span: 6 }" :wrapper-col="{ span: 14 }">

      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.name"
                 @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

      <a-form-item label="请求方法" v-bind="validateInfos.method">
        <a-select class="select-method"
                  v-model:value="modelRef.method"
                  :disabled="modelRef.id > 0">
          <template v-for="method in Methods">
            <a-select-option v-if="hasDefinedMethod(method)" :value="method" :key="method">
              {{ method }}
            </a-select-option>
          </template>
        </a-select>
      </a-form-item>

    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import {computed, defineProps, inject, reactive, ref, watch} from 'vue';
import {Methods, UsedBy} from "@/utils/enum";
import {Form} from "ant-design-vue";
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {StateType as EndpointStateType} from "@/views/endpoint/store";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy

const store = useStore<{ Endpoint: EndpointStateType }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

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

const validMethods = listDefinedMethod()

const modelRef = ref({
  id: 0,
  name: '',
  method: validMethods.length > 0 ? validMethods[0] : '',
});

watch(() => props.visible, () => {
  console.log('watch props.visible', props?.visible)
  modelRef.value = {
    id: props?.model?.id,
    name: props?.model?.name,
    method: props?.model?.method ? props?.model?.method : validMethods.length > 0 ? validMethods[0] : '',
  }
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

function listDefinedMethod() {
  const methods = [] as string[]
  endpointDetail?.value?.interfaces?.forEach((item) => {
    methods.push(item.method)
  })

  return methods
}

function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
