<template>
  <a-modal width="600px"
           :visible="visible"
           @ok="finish"
           @cancel="cancel"
           :title="(!model.id ? '新建' : '修改') + '用例'">
    <a-form :label-col="{ span: 6 }"
            :wrapper-col="{ span: 14 }">

      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.name"
                 @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import {defineProps, reactive, ref, watch} from 'vue';
import {Form} from "ant-design-vue";

const useForm = Form.useForm;

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
  id: 0,
  name: '',
});

watch(() => props.visible, () => {
  console.log('watch props.visible', props?.visible)
  modelRef.value = {
    id: props?.model?.id,
    name: props?.model?.name,
  }
}, {immediate: true, deep: true})

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
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

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
