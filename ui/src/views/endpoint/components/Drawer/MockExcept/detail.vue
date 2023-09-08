<template>
  <a-modal 
    :visible="true"
    :title="title || '新建Mock期望'" 
    width="1000px"
    @cancel="handleCancel">
    <template #footer>
      <a-button key="back" @click="handleCancel">取消</a-button>
      <a-button key="submit" html-type="submit" type="primary" @click="handleOk">确定</a-button>
    </template>
    <a-form 
      layout="vertical"
      :model="formState" 
      :label-col="labelCol" 
      class="mock-detail-form"
      ref="mockFormRef">
      <a-form-item label="期望名称" :rules="rules.name" name="name">
        <a-input style="width: 600px" v-model:value="formState.name" placeholder="请填写mock期望名称" />
      </a-form-item>
      <a-form-item label="请求方法" :rules="rules.method" name="method">
        <a-select style="width: 200px;" v-model:value="formState.method" placeholder="请选择请求方法" :options="requestMethodOpts" />
      </a-form-item>
      <a-form-item label="期望条件">
        <a-tabs v-model:activeKey="requestActiveKey">
          <a-tab-pane v-for="(item) in requestTabs" :key="item.type" :tab="item.title">
            <MockData :type="item.type" :data="formState.data" @change="handleChange" @delete="handleDelete" />
          </a-tab-pane>
        </a-tabs>
      </a-form-item>
      <a-form-item label="响应数据">
        <a-row class="mock-response-data">
          <a-form-item label="返回HTTP状态码" name="responseCode" style="margin-right: 20px;" :rules="rules.reponseCode">
            <a-input placeholder="请输入http状态码" v-model:value="formState.responseCode" />
          </a-form-item>
          <a-form-item label="返回延迟" name="responseDelay" :rules="formState.responseDelay">
            <a-input-number v-model:value="formState.responseDelay" />
            ms
          </a-form-item>
        </a-row>
        <a-tabs v-model:activeKey="activeKey">
          <a-tab-pane key="1" tab="响应体">
            <div class="form-response">
              <div class="top">
                <div class="response-left">
                  <a-radio-group v-model:value="language">
                    <a-radio-button value="pretty">pretty</a-radio-button>
                    <a-radio-button value="raw">raw</a-radio-button>
                  </a-radio-group>
                </div>
                <div class="response-right">
                  <a-button type="link">自动生成</a-button>
                </div>
              </div>
              <div class="bottom">
                <MonacoEditor
                  ref="monacoEditor"
                  customId="request-body-main"
                  class="editor"
                  v-model:value="formState.responseJson"
                  :language="'javascript'"
                  theme="vs"
                  height="300"
                  :options="editorOptions"/>
              </div>
            </div>
          </a-tab-pane>
          <a-tab-pane key="2" tab="响应头">
            <MockData type="responseHeader" :data="formState.data"  @change="handleChange" @delete="handleDelete"/>
          </a-tab-pane>
        </a-tabs>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="tsx" setup>
/**
 * ::todo 后续在做代码层面的拆分优化
 */
import { ref, defineProps, defineEmits, reactive, } from 'vue';
import { useStore } from 'vuex';

import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import { MockData } from './components/index';

import { requestTabs } from './index';
import { MonacoOptions } from "@/utils/const";
import { requestMethodOpts } from '@/config/constant';

const props = defineProps<{
  title?: String;
}>();

const emits = defineEmits(['cancel']);
const store = useStore();
const defaultData = {
  idx: 0,
  name: '',
  value: '',
  condition: ''
}
/**
 * form 表单信息
 */
const formState: any = reactive({
  name: '', // 期望名称
  method: requestMethodOpts[0].value, // 请求方法
  reponseCode: '', // 响应码
  responseDelay: 0, // 响应延迟
  reponseJson: '', // 响应内容
  // 列表信息
  data: {
    header: [{...defaultData}], // 请求头
    body: [{...defaultData}], // 请求体
    queryParams: [{...defaultData}], // 查询参数
    pathParams: [{...defaultData}], // 路径参数
    responseHeader: [{...defaultData}], // 响应头
  }
});

/**
 * 页面绑定data
 */
const mockFormRef = ref();
const activeKey = ref('1');
const requestActiveKey = ref('header');
const editorOptions = ref(Object.assign( { usedWith: 'response', readOnly:false }, MonacoOptions ) );
const language = ref('pretty');
/**
 * http响应码校验
 * @param args 校验信息
 */
const responseCodeValidator = (...args) => {
  const value = args[1];
  const exp = new RegExp('^[0-9]*[1-9][0-9]*$');
  console.log(value);
  if (!value) {
    return Promise.reject('请输入响应HTTP状态码');
  }
  if (!exp.test(value)) {
    return Promise.reject('请输入正确的响应HTTP状态码');
  }
  return Promise.resolve();
};

// 校验规则
const rules = {
  name: [{
    required: true,
    message: 'Mock期望名称不可为空'
  }],
  reponseCode: [{
    required: true,
    validator: responseCodeValidator
  }],
  method: [{
    required: true,
    message: '请求方法不可为空'
  }],
  responseDelay: [{
    required: true,
    message: '返回延迟不可为空'
  }]
}

const labelCol = { style: { width: '150px' } };

const handleOk = async (e: MouseEvent) => {
  try {
    await mockFormRef.value.validateFields();
    console.log(formState);
    store.commit("Global/setSpinning",true)
  } catch (err: any) {
    console.log('saveGlobalParams validateFailed--', err);
  }
  // console.log(e);
  // emits('cancel', false);
};

const handleCancel = () => {
  emits('cancel', false);
};

const handleChange = (...args) => {
  const [type] = args;
  console.log(`当前修改${type} --`);
  if (!formState.data[type].some(e => e.name === '')) {
    const lastElIdx = formState.data[type][formState.data[type].length - 1].idx;
    formState.data[type].push({ ...defaultData, idx: lastElIdx + 1 });
  }
};

const handleDelete = (record, type) => {
  if (formState.data[type].length === 1) {
    return;
  }
  const index = formState.data[type].findIndex(e => e.idx === record.idx);
  formState.data[type].splice(index, 1);
}

</script>
<style scoped lang="less">
.mock-response-data {

  &:has(.ant-form-item-has-error) {
    margin-bottom: 20px;
  }
  :deep(.ant-form-item) {
    display: flex;
    align-items: center;
    flex-direction: row;
    margin-bottom: 0 !important;

    .ant-form-item-label {
      padding: 0 !important;
      margin-right: 6px;
    }

    .ant-row.ant-form-item {
      margin-bottom: 0 !important;
    }

    .ant-form-item-explain.ant-form-item-explain-error {
      position: absolute;
      left: 0;
      bottom: -24px;
    }
  }
}

.form-response {
  .top {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;

    .response-left {

    }
  }

  .bottom {
    width: 100%;
    border: 1px solid #d9d9d9;
  }
}
</style>