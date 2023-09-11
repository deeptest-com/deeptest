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
              <MockData 
                :type="item.type" 
                :data="formState" 
                :optionsMap="dropdownOptions"
                @columnChange="handleChange" 
                @delete="handleDelete" />
            </a-tab-pane>
          </a-tabs>
        </a-form-item>
        <a-form-item label="响应数据">
          <a-row class="mock-response-data">
            <a-form-item label="返回HTTP状态码" name="code" style="margin-right: 20px;" :rules="rules.reponseCode">
              <a-input placeholder="请输入http状态码" v-model:value="formState.code" />
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
                    v-model:value="formState.responseBody.value"
                    :language="'javascript'"
                    theme="vs"
                    height="300"
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
import { requestMethodOpts } from '@/config/constant';

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
}
/**
 * form 表单信息
 */
const formState: any = reactive({
  name: '', // 期望名称
  method: requestMethodOpts[0].value, // 请求方法
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

/**
 * 页面绑定data
 */
const mockFormRef = ref();
const activeKey = ref('1');
const requestActiveKey = ref('requestHeaders');
const editorOptions = ref(Object.assign( { usedWith: 'response', readOnly:false }, MonacoOptions ) );
const language = ref('pretty');
const dropdownOptions = computed(() => {
  return store.state.Endpoint.mockExpectOptions;
});
const mockExpectDetail = computed(() => {
  return store.state.Endpoint.mockExpectDetail;
});
const loading = ref(false);

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

const labelCol = { style: { width: '150px' } };

const setDataList = (data: any[], type?: string) => {
  return cloneDeep(data).filter(e => e.name !== '').map(e => {
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

const handleChange = (...args) => {
  const [type] = args;
  try {
    if (!formState[type].some(e => e.name === '')) {
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
  formState.responseBody.value = e;
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
  }
}, {
  immediate: true,
  deep: true
})

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
      white-space: pre;
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