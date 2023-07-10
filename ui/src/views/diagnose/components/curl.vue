<template>
  <a-modal width="600px"
           title="导入cURL"
           :visible="visible"
           @ok="finish"
           @cancel="cancel">
    <a-form :label-col="{ span: 0 }" :wrapper-col="{ span: 24 }" >

      <a-form-item label="" v-bind="validateInfos.content">
        <a-textarea :rows="8"
                    v-model:value="modelRef.content"/>
      </a-form-item>

    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import {defineProps, reactive, ref, watch} from 'vue';
import {Form} from 'ant-design-vue';

const useForm = Form.useForm;

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
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
/*
const test = `curl -X POST 'http://127.0.0.1:8085/api/v1/mock?reqType=file'` +
    `  -H 'Content-Type: application/json'` +
    `  -H 'Cookie: BIDUPSID=88B7FC40D50C2F811E57590167144216;'` +
    `  -F 'name=aaron -F myFile=@/Users/aaron/rd/project/github/gcurl/tests/file.txt;type=text/plain'`
    */

const test = ""    
const modelRef = ref({content: test} as any);

const rulesRef = reactive({
  content: [
    {required: true, message: '请输入cURL内容', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

watch(() => {return props.visible}, () => {
  console.log('watch props.visible', props.visible)
  modelRef.value = {
    content: test,
  }
}, {deep: true})

const finish = () => {
  validate().then(() => {
    props.onFinish(modelRef.value)
    resetFields();
  }).catch((error) => console.log('error', error))
}

const cancel = () => {
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
