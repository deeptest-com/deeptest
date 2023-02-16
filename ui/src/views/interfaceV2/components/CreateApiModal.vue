<template>

  <a-modal
      width="600px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      title="新建接口">

    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
    >
      <a-form-item label="接口分类" name="tag">
        <a-select placeholder="请选择接口分类" v-model:value="formState.tag">
          <a-select-option value="shanghai">接口类型1</a-select-option>
          <a-select-option value="beijing">接口类型2</a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item label="接口名称" name="title">
        <a-input placeholder="接口名称" v-model:value="formState.title"/>
      </a-form-item>

      <a-form-item label="接口路径" name="path">
        <a-input v-model:value="formState.path" class="form-item-con" placeholder="请输入接口路径">
          <template #addonBefore>
            <a-select
                style="width: 120px"
                v-model:value="formState.method"
                :options="requestMethodOpts"
                mode="tags"
                placeholder="请选择请求方法">
            </a-select>
          </template>
        </a-input>
      </a-form-item>


      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <span class="">
           注：接口请求方法可以通过详情页添加
        </span>
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
//
// watch(props.visible, () => {
//   console.log('832', props.visible)
// })

/**
 *
 *   tag: '接口类型1',
 *   value: 1,
 *   title: '用户详情信息',
 *   method:'GET',
 *   path:'/api/user/:userId'
 * */
interface FormState {
  title: string;
  tag: string | undefined;
  method: string | undefined;
  path: string | undefined;
  value: string | undefined;
}


function ok() {
  emit('ok',formState);
}

function cancal() {
  emit('cancal',formState);
}

const formRef = ref();

const formState: UnwrapRef<FormState> = reactive({
  tag: '接口类型1',
  value: '1',
  title: '',
  method: 'GET',
  path: ''
});

const rules = {
  title: [
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
