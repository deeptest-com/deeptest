<template>
  <a-modal width="600px"
           title="导入cURL"
           :visible="visible"
           @ok="finish"
           @cancel="cancel">
    <a-form :label-col="{ span: 0 }" :wrapper-col="{ span: 24 }" >

      <a-form-item label="" v-bind="validateInfos.content">
        <a-textarea :rows="8"
                    v-model:value="modelRef.content"
                    />
      </a-form-item>
      <a-form-item>
      <a-alert message="不是合法的cURL请求，请重试。" type="error" show-icon v-if="showError"/>
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
const showError = ref(false)

let validateCurl = async (rule: any, value: string,callback: any) => {
      if (value === '') {
        showError.value = true
        return Promise.reject("请输Curl请求")
      } else {
        if (!rule.pattern.test(value)){
          showError.value = true
          return Promise.reject("不是合法的cURL请求，请重试。")
        }
        showError.value = false
        return Promise.resolve();
      }
    };

const rulesRef = reactive({
  content: [
    {required: true,  message: '',validator:validateCurl, trigger: 'change',pattern:/curl\s+.*\s+.*/},
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
  }).catch((error) => {
    console.log('error', error)
    //showError.value = true
  })
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
.ant-alert-error {
    background-color: #ffffff;
    border: 1px solid #ffffff;
}

.ant-form-item {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    color: rgba(0, 0, 0, 0.65);
    font-size: 14px;
    font-variant: tabular-nums;
    line-height: 1.5715;
    list-style: none;
    font-feature-settings: 'tnum';
    margin-bottom: 0px;
    vertical-align: top;
}

</style>
