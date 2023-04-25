<template>
  <a-modal
      width="600px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      :title="mode === 'new' ? '新建分类' : '修改分类'">
    <a-form
        ref="formRef"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }">
      <a-form-item label="接口分类名称" name="name">
        <a-input placeholder="请输入接口分类名称" v-model:value="formState.name"/>
      </a-form-item>
      <a-form-item label="备注信息" name="description">
        <a-input placeholder="请输入备注信息" v-model:value="formState.description"/>
      </a-form-item>
    </a-form>
  </a-modal>


</template>
<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {defineComponent, reactive, ref, toRaw, UnwrapRef, defineProps, defineEmits, watch} from 'vue';
import {requestMethodOpts} from '@/config/constant';

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  nodeInfo: {
    required: true,
    type: Object,
  },
  mode: {
    required: true,
    type: String,
  }
})

const emit = defineEmits(['ok', 'cancal']);


const formState: any = ref(null);

watch(() => {
  return props.visible
}, (newVal) => {
  if(newVal){
    formState.value = props.nodeInfo;
    if (props.mode === 'new') {
      formState.value.name = '';
      formState.value.description = '';
    }
  }
})

function ok() {
  emit('ok', formState.value);
}

function cancal() {
  emit('cancal', formState.value);
}

// const formRef = ref();


const rules = {
  name: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 3, max: 50, message: '最长多少个字符', trigger: 'blur'},
  ],
  path: [{required: true, message: 'Please select Activity zone', trigger: 'change'}],
  tag: [{required: true, message: 'Please select activity resource', trigger: 'change'}],
};

// const onSubmit = () => {
//   formRef.value
//       .validate()
//       .then(() => {
//         console.log('values', formState, toRaw(formState));
//       })
//       .catch((error: ValidateErrorEntity<FormState>) => {
//         console.log('error', error);
//       });
// };

// const resetForm = () => {
//   formRef.value.resetFields();
// };

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>