<template>

  <a-modal
      width="600px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      title="新建分类/修改分类">

    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 4 }"
        :wrapper-col="{ span: 14 }">

      <a-form-item label="接口名称" name="name">
        <a-input placeholder="接口名称" v-model:value="formState.name"/>
      </a-form-item>

      <a-form-item label="备注" name="remark">
        <a-input placeholder="备注" v-model:value="formState.remark"/>
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
  }
})

const emit = defineEmits(['ok', 'cancal']);


watch(() => {
  return props.visible
}, () => {
  console.log('832', props.visible)
})

/**
 *
 *   tag: '接口类型1',
 *   value: 1,
 *   name: '用户详情信息',
 *   method:'GET',
 *   path:'/api/user/:userId'
 * */
interface FormState {
  name: string;
  remark: string | undefined;
}


function ok() {
  emit('ok');
}

function cancal() {
  emit('cancal');
}

const formRef = ref();

const formState: UnwrapRef<FormState> = reactive({
  name: '接口类型1',
  remark: '用户信息相关',
});

const rules = {
  name: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 3, max: 50, message: '最长多少个字符', trigger: 'blur'},
  ],
  path: [{required: true, message: 'Please select Activity zone', trigger: 'change'}],
  tag: [{required: true, message: 'Please select activity resource', trigger: 'change'}],
};

const onSubmit = () => {
  formRef.value
      .validate()
      .then(() => {
        console.log('values', formState, toRaw(formState));
      })
      .catch((error: ValidateErrorEntity<FormState>) => {
        console.log('error', error);
      });
};
const resetForm = () => {
  formRef.value.resetFields();
};
</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
