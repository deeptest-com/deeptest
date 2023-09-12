<template>
  <a-modal 
    :visible="true"
    :title="mockExpectDetail.id ? '编辑Mock期望' : '新建Mock期望'" 
    width="1000px"
    @cancel="handleCancel">
    <template #footer>
      <a-button key="back" @click="handleCancel">取消</a-button>
      <a-button :loading="loading" key="submit" html-type="submit" type="primary" @click="handleOk">确定</a-button>
    </template>
    <div style="max-height: 650px;overflow-y: scroll;">
      <a-form 
        layout="vertical"
        :model="formState" 
        class="mock-detail-form"
        ref="mockFormRef">
        <a-form-item class="expect-name" label="期望名称" :rules="rules.name" name="name">
          <a-input style="width: 600px" v-model:value="formState.name" placeholder="请填写mock期望名称" />
        </a-form-item>
        <a-form-item class="expect-method" label="请求方法" :rules="rules.method" name="method">
          <a-select style="width: 200px;" v-model:value="formState.method" placeholder="请选择请求方法" :options="methods" />
        </a-form-item>
        <a-form-item label="期望条件">
          <a-tabs v-model:activeKey="requestActiveKey">
            <a-tab-pane v-for="(item) in requestTabs" :key="item.type" :tab="item.title">
              <div v-if="item.type === 'requestBodies'" class="requestbodies-select-type">
                <span>匹配对象: &nbsp;</span>
                <a-radio-group v-model:value="requestBodyType" @change="handleRequestBodyTypeChanged">
                  <a-radio value="keyValue">键值</a-radio>
                  <a-radio value="XPath">XPath元素</a-radio>
                  <a-radio value="fullText">全文本</a-radio>
                </a-radio-group>
              </div>
              <MockData 
                :type="item.type" 
                :data="formState" 
                :optionsMap="dropdownOptions"
                :selectType="requestBodyType"
                @columnChange="handleChange" 
                @delete="handleDelete" />
            </a-tab-pane>
          </a-tabs>
        </a-form-item>
        <a-form-item label="响应数据">
          <a-row class="mock-response-data">
            <a-form-item label="返回HTTP状态码" name="code" style="margin-right: 20px;" :rules="rules.reponseCode">
              <a-auto-complete
                v-model:value="formState.code"
                :options="defaultResponseCodes.map(e => ({ value: e, label: e }))"
                style="width: 200px"
                placeholder="请输入或选择http状态码"
                :filter-option="false"
              />
            </a-form-item>
            <a-form-item label="返回延迟" name="delayTime" :rules="rules.responseDelay">
              <a-input-number v-model:value="formState.delayTime" />
              ms
            </a-form-item>
          </a-row>
          <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1" tab="响应体">
              <div class="form-response">
                <div class="top">
                  <div class="response-left">
                    <a-radio-group v-model:value="language" @change="handleLanguageChange">
                      <a-radio-button value="json">pretty</a-radio-button>
                      <a-radio-button value="raw">raw</a-radio-button>
                    </a-radio-group>
                  </div>
                  <div class="response-right">
                    <a-button type="link" @click="generateJsonExample">自动生成</a-button>
                  </div>
                </div>
                <div class="bottom">
                  <MonacoEditor
                    ref="monacoEditor"
                    customId="request-body-main"
                    class="editor"
                    v-model:value="formState.responseBody.value"
                    :language="language"
                    theme="vs"
                    height="300"
                    :timestamp="timestamp"
                    @change="handleEditorChange"
                    :options="editorOptions"/>
                </div>
              </div>
            </a-tab-pane>
            <a-tab-pane key="2" tab="响应头">
              <MockData type="responseHeaders" :data="formState"  @columnChange="handleChange" @delete="handleDelete"/>
            </a-tab-pane>
          </a-tabs>
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
</template>
<script lang="tsx" setup>
/**
 * ::todo 后续在做代码层面的拆分优化
 */
import { ref, defineProps, defineEmits, reactive, onMounted, computed, watch, unref, } from 'vue';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';
import cloneDeep from "lodash/cloneDeep";

import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import { MockData } from './components/index';

import { requestTabs } from './index';
import { MonacoOptions } from "@/utils/const";
import { defaultResponseCodes, responseCodes } from '@/config/constant';

const props = defineProps<{
  title?: String;
}>();

const emits = defineEmits(['cancel']);
const store = useStore();
const defaultData = {
  idx: 0,
  id: 0,
  name: '',
  value: '',
  compareWay: ''
};
/**
 * method 可选项
 */
const methods = computed(() => {
  return (store.state.Endpoint.endpointDetail.interfaces || []).map(e => e.method).map(e => ({ label: e, value: e }));
});
/**
 * form 表单信息
 */
const formState: any = reactive({
  name: '', // 期望名称
  method: methods.value[0]?.value, // 请求方法
  code: '',
  delayTime: 0,
  // 列表信息
  requestHeaders: [{...defaultData}], // 请求头
  requestBodies: [{...defaultData}], // 请求体
  requestQueryParams: [{...defaultData}], // 查询参数
  requestPathParams: [{...defaultData}], // 路径参数
  responseBody: {
    code: '',
    delayTime: '',
    value: ''
  },
  responseHeaders: [{...defaultData}], // 响应头
});

const requestBodyType = ref<any>('keyValue'); // 请求体选择类型： keyValue / XPath表达式 / fullText 对应的表格填写内容不同

const jsonContent = ref('');

/**
 * 页面绑定data
 */
const mockFormRef = ref();
const activeKey = ref('1'); //响应内容tab切换key
const requestActiveKey = ref('requestQueryParams'); //请求条件tab切换key
const editorOptions = ref(Object.assign( { usedWith: 'response', readOnly:false }, MonacoOptions ) );
const language = ref('json');
const dropdownOptions = computed(() => {
  return store.state.Endpoint.mockExpectOptions;
}); // 下拉选项
const mockExpectDetail = computed(() => {
  return store.state.Endpoint.mockExpectDetail;
}); // 期望详情
const loading = ref(false);

const timestamp = ref('');

/**
 * http响应码校验
 * @param args 校验信息
 */
const responseCodeValidator = (...args) => {
  const value = args[1];
  const exp = new RegExp('^[0-9]*[1-9][0-9]*$');
  if (!value) {
    return Promise.reject('请输入响应HTTP状态码');
  }
  if (!exp.test(value)) {
    return Promise.reject('请输入正确的响应HTTP状态码');
  }
  return Promise.resolve();
};

const responseDelayValidator = (...args) => {
  const value = args[1];
  if (['', undefined, null].includes(value)) {
    return Promise.reject('返回延迟不可为空');
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
    validator: responseDelayValidator
  }]
}

// const labelCol = { style: { width: '150px' } };

const setDataList = (data: any[], type?: string) => {
  const list = (type === 'body' && requestBodyType.value === 'fullText') ? cloneDeep(data).filter(e => e.value !== '') : cloneDeep(data).filter(e => e.name !== '');
  return list.map(e => {
    delete e.idx;
    if (!e.id) {
      delete e.id;
    }
    if (type) {
      e.source = type;
    }
    if (formState.id) {
      e.endpointMockExpectId = formState.id;
    }
    if (type === 'body') {
      e.selectType = requestBodyType.value;
    }
    return e;
  })
};

const handleOk = async (e: MouseEvent) => {
  try {
    await mockFormRef.value.validateFields();
    loading.value = true;
    const requestParams = { ...formState };
    requestParams.responseBody.code = formState.code;
    requestParams.responseBody.delayTime = formState.delayTime;
    requestParams.requestHeaders = setDataList(formState.requestHeaders, 'header');
    requestParams.requestBodies = setDataList(formState.requestBodies, 'body');
    requestParams.requestQueryParams = setDataList(formState.requestQueryParams, 'query');
    requestParams.requestPathParams = setDataList(formState.requestPathParams, 'path');
    requestParams.responseHeaders = setDataList(formState.responseHeaders);
    delete requestParams.code;
    delete requestParams.delayTime;
    const result = await store.dispatch('Endpoint/saveMockExpect', requestParams);
    loading.value = false;
    if (result) {
      message.success(`${formState.id ? '修改' : '新建'}Mock期望成功`);
    }
    emits('cancel');
  } catch (err: any) {
    console.log('saveGlobalParams validateFailed--', err);
  }
  // console.log(e);
  // emits('cancel', false);
};

const handleCancel = () => {
  emits('cancel', false);
};

const handleLanguageChange = (e) => {
  const val = e.target.value;
  if (val == 'raw') {
    console.log(1);
    formState.responseBody.value = jsonContent.value.replace(/\r\n/g,'').replace(/\n/g,'').replace(/\s+/g,'')
  }
}

const handleRequestBodyTypeChanged = (e) => {
  formState.requestBodies = [{ ...defaultData }].map((e, index) => ({
    ...e,
    idx: index + 1,
  }));
}

const handleChange = (...args) => {
  const [type] = args;
  try {
    if (!formState[type].some(e => e.name === '')) {
      const lastElIdx = formState[type][formState[type].length - 1].idx;
      formState[type].push({ ...defaultData, idx: lastElIdx + 1 });
    }
    /**
     * 请求体类型为fullText时，name为空。
     */
    if (requestBodyType.value === 'fullText' && type === 'requestBodies' && formState[type].length === 1) {
      const lastElIdx = formState[type][formState[type].length - 1].idx;
      formState[type].push({ ...defaultData, idx: lastElIdx + 1 });
    }
  } catch (error) {
    console.log(error);
  }
};

const handleDelete = (record, type) => {
  if (formState[type].length === 1) {
    return;
  }
  const index = formState[type].findIndex(e => e.idx === record.idx);
  formState[type].splice(index, 1);
}

const handleEditorChange = (e) => {
  console.log(e);
  formState.responseBody.value = e;
  jsonContent.value = e;
}

const generateJsonExample = async () => {
  if (!formState.code) {
    message.error('请先填写http响应状态码');
    return;
  }
  const result = await store.dispatch('Endpoint/generateJsonExample', {
    code: formState.code,
    method: formState.method,
  });
  console.log(result);
  if (result !== false) {
    formState.responseBody.value = result === null ? '' : JSON.stringify(result);
  }
}

onMounted(() => {
  store.dispatch('Endpoint/getMockExpectOptions');
})

const initListData = (array: any) => {
  let result: any[] = [];
  if (array) {
    result = array.concat([defaultData]).map((e, index) => ({
      idx: index + 1,
      ...e,
    }));
  } else {
    result = [{ ...defaultData }];
  }
  return result;
};

watch(() => {
  return unref(mockExpectDetail);
}, val => {
  console.log('获取当前查看的mockExpect详情', val);
  if (val.id) {
    // 设置当前formState
    Object.assign(formState, {
      ...val,
      requestBodies: initListData(val.requestBodies),
      requestHeaders: initListData(val.requestHeaders),
      requestQueryParams: initListData(val.requestQueryParams),
      requestPathParams: initListData(val.requestPathParams),
      responseHeaders: initListData(val.responseHeaders),
      code: val.responseBody.code,
      delayTime: val.responseBody.delayTime,
    })
    requestBodyType.value = [...new Set((val.requestBodies || []).map(e => e.selectType))]?.[0] || 'keyValue';
  }
}, {
  immediate: true,
  deep: true
});

watch(() => {
  return formState.responseBody.value;
}, () => {
  timestamp.value = Date.now() + '';
}, {
  immediate: true,
  deep: true,
})

</script>
<style scoped lang="less">
.expect-name, .expect-method {
  flex-direction: row !important;

  :deep(.ant-form-item-label) {
    padding: 0;
    display: flex;
    align-items: center;
    margin-right: 10px;
    height: 32px;
  }
}
.mock-response-data {

  :deep(.ant-form-item) {
    flex-direction: row;
    margin-bottom: 0 !important;
  
    .ant-form-item-label {
      margin-right: 6px;
      padding: 0;
      display: flex;
      align-items: center;
      height: 32px;
    }

    .ant-row.ant-form-item {
      margin-bottom: 0 !important;
    }

    // .ant-form-item-explain.ant-form-item-explain-error {
    //   position: absolute;
    //   left: 0;
    //   bottom: -24px;
    //   white-space: pre;
    // }
  }
}

.form-response {
  .top {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
  }

  .bottom {
    width: 100%;
    border: 1px solid #d9d9d9;
  }
}

.requestbodies-select-type {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  margin-bottom: 18px;
}
</style>