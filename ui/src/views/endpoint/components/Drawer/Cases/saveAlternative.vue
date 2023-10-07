<template>
  <a-modal width="600px"
           :visible="visible"
           @ok="submit"
           @cancel="cancel"
           title="另存为用例">
    <a-form :label-col="{ span: 6 }" :wrapper-col="{ span: 14 }">

      <a-form-item label="名称前缀" v-bind="validateInfos.prefix">
        <a-input v-model:value="modelRef.prefix"
                 @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
        <div class="dp-input-tip">
          {{ `生成的用例会以"${modelRef.prefix}-"开头` }}
        </div>
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
  onClose: {
    type: Function,
    required: true,
  },
})

const modelRef = ref({
  prefix: '',
});

watch(() => props.visible, () => {
  console.log('watch props.visible', props?.visible)
  modelRef.value = {
    prefix: '备选用例-',
  }
}, {immediate: true, deep: true})

const rulesRef = reactive({
  prefix: [
    {required: true, message: '请输入用例前缀', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submit = () => {
  console.log('submit', modelRef.value, props.model)
  validate().then(() => {
    // submit

    resetFields();
    props.onClose()
  }).catch((error) => console.log('error', error))
}

const cancel = () => {
  console.log('cancel')
  resetFields()
  props.onClose()
}

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
