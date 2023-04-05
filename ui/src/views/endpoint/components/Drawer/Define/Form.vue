<template>
  <div class="content">
    <!-- ::::路径定义方式 -->
    <a-row class="form-item">
      <a-col :span="2" class="form-label">路径</a-col>
      <a-col :span="16">
        <a-input :value="endpointDetail.path" @change="updatePath" placeholder="请输入路径">
          <template #addonBefore>
            <a-select
                :options="serveServers"
                :value="serveServers?.[0]?.value"
                placeholder="请选择服务器"
                style="width: 200px;text-align: left" />
          </template>
          <template #addonAfter>
            <a-button @click="addPathParams">
              <template #icon>
                <PlusOutlined/>
              </template>
              路径参数
            </a-button>
          </template>
        </a-input>
        <!-- ::::路径参数 -->
        <div class="path-param-list">
            <Field
                v-for="(item,index) in endpointDetail.pathParams"
                :key="item.id"
                :fieldData="{...item,index:index}"
                :showRequire="true"
                :refsOptions="[
                    {
                    label: '组件 1',
                    value: 'COM1'
                  }, {
                    label: '组件 2',
                    value: 'COM2'
                  }]"
                @del="deletePathParams(index)"
                @change="pathParamsNameChange"/>
        </div>
      </a-col>
    </a-row>

    <!-- ::::请求方式定义 -->
    <a-row class="form-item">
      <a-col :span="2" class="form-label">请求方式</a-col>
      <a-col :span="22">
        <!-- ::::请求方法定义 -->
        <a-radio-group v-model:value="selectedMethod" button-style="solid">
          <a-radio-button
              :class="{'has-defined': hasDefinedMethod(method.value)}"
              :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
            {{ method.label }}
          </a-radio-button>
        </a-radio-group>

        <div class="form-item-request">
          <div v-if="selectedMethodDetail">
            <!-- ::::Operation ID -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label">
                Operation ID
              </a-col>
              <a-col :span="12">
                <a-input v-model:value="selectedMethodDetail.operationId"/>
              </a-col>
            </a-row>
            <!-- ::::Description -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label">
                Description
              </a-col>
              <a-col :span="12">
                <a-input v-model:value="selectedMethodDetail.description"/>
              </a-col>
            </a-row>

            <!-- ::::增加请求参数 -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label">
                增加请求参数
              </a-col>
              <a-col :span="15">
                <div class="params-defined-btns">
                  <a-button @click="setSecurity">
                    <template #icon>
                      <PlusOutlined/>
                    </template>
                    {{ `Security` }}
                  </a-button>
                  <a-button @click="addHeader">
                    <template #icon>
                      <PlusOutlined/>
                    </template>
                    {{ `Header` }}
                  </a-button>
                  <a-button @click="addQueryParams">
                    <template #icon>
                      <PlusOutlined/>
                    </template>
                    {{ `Query Params` }}
                  </a-button>
                  <a-button @click="addCookie">
                    <template #icon>
                      <PlusOutlined/>
                    </template>
                    {{ `Cookie` }}
                  </a-button>
                </div>
              </a-col>
            </a-row>

            <!-- ::::请求参数展示：headers、cookies、query params等 -->
            <a-row class="form-item-request-item">
              <a-col :span="3"></a-col>
              <a-col :span="21">
                <div class="params-defined">
                  <div class="params-defined-content">
                    <div class="params-defined-item" v-if="showSecurity">
                      <div class="params-defined-item-header">
                        <span>Security</span>
                      </div>
                      <div class="header-defined header-defined-items">
                        <a-select @change="securityChange"
                                  allowClear
                                  :value="selectedMethodDetail.security"
                                  :options="securityOpts" style="width: 300px;"/>
                        <a-tooltip placement="topLeft" arrow-point-at-center title="删除 Security">
                          <a-button @click="delSecurity">
                            <template #icon>
                              <DeleteOutlined/>
                            </template>
                          </a-button>
                        </a-tooltip>
                        <a-tooltip placement="topLeft" arrow-point-at-center title="去添加或编辑 Security">
                          <a-button @click="goEditSecurity">
                            <template #icon>
                              <PlusOutlined/>
                            </template>
                            Security
                          </a-button>
                        </a-tooltip>
                      </div>
                    </div>
                    <div class="params-defined-item" v-if="selectedMethodDetail?.headers?.length">
                      <div class="params-defined-item-header">
                        <span>Header</span>
                      </div>
                      <div class="header-defined header-defined-items">
                        <div v-for="(item,index) in selectedMethodDetail.headers" :key="item.id">
                          <Field
                              :fieldData="{...item,index:index}"
                              :showRequire="true"
                              @del="deleteParams('headers',index)"
                              @change="(val) => {
                                handleParamsChange('headers',val);
                              }"/>
                        </div>
                      </div>
                    </div>
                    <div class="params-defined-item" v-if="selectedMethodDetail?.params?.length">
                      <div class="params-defined-item-header">
                        <span>Query Params</span>
                      </div>
                      <div class="header-defined ">
                        <div v-for="(item,index) in selectedMethodDetail.params" :key="item.id">
                          <Field
                              :fieldData="{...item,index:index}"
                              :showRequire="true"
                              @del="deleteParams('params',index)"
                              @change="(val) => {
                                handleParamsChange('params',val);
                              }"/>
                        </div>
                      </div>
                    </div>
                    <div class="params-defined-item" v-if="selectedMethodDetail?.cookies?.length">
                      <div class="params-defined-item-header">
                        <span>Cookie</span>
                      </div>
                      <div class="header-defined ">
                        <div v-for="(item,index) in selectedMethodDetail.cookies" :key="item.id">
                          <Field
                              :fieldData="{...item,index:index}"
                              :showRequire="true"
                              @del="deleteParams('cookies',index)"
                              @change="(val) => {
                                handleParamsChange('cookies',val);
                              }"/>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </a-col>
            </a-row>
            <!-- ::::增加请求体 -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label">
                增加请求体
              </a-col>
              <a-col :span="18">
                <a-select
                    placeholder="请选择格式"
                    v-if="selectedMethodDetail.requestBody"
                    v-model:value="selectedMethodDetail.requestBody.mediaType"
                    style="width: 300px"
                    :options="mediaTypesOpts"
                ></a-select>
              </a-col>
            </a-row>
            <!-- ::::增加请求体 - 描述  -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label"></a-col>
              <a-col :span="20">
                <a-input placeholder="请输入描述" v-model:value="selectedMethodDetail.requestBody.description"/>
              </a-col>
            </a-row>
            <!-- ::::增加请求体 - scheme定义 -->
            <a-row class="form-item-request-item">
              <a-col :span="3" class="form-label"></a-col>
              <a-col :span="21">
                <SchemaEditor
                    @generateFromJSON="(data) => {
                      generateFromJSON('req',data);
                    }"
                    @generateExample="(data) => {
                       handleGenerateExample('req',data);
                    }"
                    @change="(data) => {
                      handleChange('req',data);
                    }"
                    :tab-content-style="{width:'100%'}"
                    :value="activeReqBodySchema"/>
              </a-col>
            </a-row>
            <!-- ::::响应定义  -->
            <a-row class="form-item-response">
              <a-col :span="3" class="form-label">
                选择响应代码
              </a-col>
              <a-col :span="21">
                <a-radio-group v-model:value="selectedCode" button-style="solid">
                  <a-radio-button
                      :class="{'has-defined': hasDefinedCode(code.value)}"
                      :key="code.value" v-for="code in repCodeOpts"
                      :value="code.value">
                    {{ code.label }}
                  </a-radio-button>
                </a-radio-group>
                <div class="form-item-response">
                  <div v-if="selectedCodeDetail">
                    <!-- ::::Description -->
                    <a-row class="form-item-response-item">
                      <a-col :span="4" class="form-label">
                        Description
                      </a-col>
                      <a-col :span="18">
                        <a-input v-model:value="selectedCodeDetail.desc"/>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应头 -->
                    <a-row class="form-item-response-item">
                      <a-col :span="4" class="form-label">
                        增加响应头
                      </a-col>
                      <a-col :span="18">
                        <div class="params-defined-btns">
                          <a-button type="primary" @click="addResponseHeader">
                            <template #icon>
                              <PlusOutlined/>
                            </template>
                            {{ `添加` }}
                          </a-button>
                        </div>
                      </a-col>
                    </a-row>
                    <!-- ::::响应头展示-->
                    <a-row class="form-item-response-item">
                      <a-col :span="4"></a-col>
                      <a-col :span="20">
                        <div class="params-defined">
                          <div class="params-defined-content">
                            <div class="params-defined-item" v-if="selectedCodeDetail?.headers?.length">
                              <div class="header-defined header-defined-items">
                                <div v-for="(item,index) in selectedCodeDetail.headers" :key="item.id">
                                  <Field
                                      :fieldData="{...item,index:index}"
                                      :showRequire="false"
                                      @del="deleteResHeader(index)"
                                      @change="(val) => {
                                        handleResHeaderChange(val);
                                      }"/>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应体体 -->
                    <a-row class="form-item-response-item">
                      <a-col :span="4" class="form-label">
                        增加响应体
                      </a-col>
                      <a-col :span="18">
                        <a-select
                            placeholder="请选择格式"
                            v-model:value="selectedCodeDetail.mediaType"
                            style="width: 300px"
                            :options="mediaTypesOpts"
                        ></a-select>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应体 - 描述  -->
                    <a-row class="form-item-response-item">
                      <a-col :span="4" class="form-label"></a-col>
                      <a-col :span="18">
                        <a-input placeholder="请输入描述" v-model:value="selectedCodeDetail.description"/>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应体 - scheme定义 -->
                    <a-row class="form-item-response-item">
                      <a-col :span="4" class="form-label"></a-col>
                      <a-col :span="20">
                        <SchemaEditor
                            @generateFromJSON="(data) => {
                                 generateFromJSON('res',data);
                            }"
                            @change="(data) => {
                               handleChange('res',data);
                            }"
                            @generateExample="(data) => {
                               handleGenerateExample('res',data);
                            }"
                            :tab-content-style="{width:'600px'}"
                            :value="activeResBodySchema"/>
                      </a-col>
                    </a-row>
                  </div>
                  <div v-if="!selectedCodeDetail">
                    <a-button type="primary" @click="addCodeResponse">
                      <template #icon>
                        <PlusOutlined/>
                      </template>
                      {{ `Add Response` }}
                    </a-button>
                  </div>
                </div>
              </a-col>
            </a-row>
          </div>
          <div class="no-defined" v-else>
            <a-button type="primary" @click="addEndpoint">
              <template #icon>
                <PlusOutlined/>
              </template>
              {{ `${selectedMethod} Operation` }}
            </a-button>
          </div>
        </div>
      </a-col>
    </a-row>
  </div>
</template>
<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  watch,
  computed,
} from 'vue';
import {useStore} from "vuex";
import {
  requestMethodOpts,
  mediaTypesOpts,
  repCodeOpts,
  defaultCookieParams,
  defaultHeaderParams,
  defaultQueryParams,
  defaultPathParams,
  defaultEndpointDetail,
  defaultCodeResponse,
} from '@/config/constant';
import {PlusOutlined, DeleteOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import {Endpoint} from "@/views/endpoint/data";
import {StateType as Debug} from "@/views/component/debug/store";
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {cloneByJSON} from "@/utils/object";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);

const currInterface = computed<any>(() => store.state.Debug?.currInterface);

const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
const serveServers: any = computed<Endpoint>(() => store.state.Endpoint.serveServers);
const securityOpts: any = computed<any>(() => store.state.Endpoint.securityOpts);

const props = defineProps({});
const emit = defineEmits([]);

const selectedMethod = ref(currInterface.value?.method ? currInterface.value?.method : 'GET');
const selectedCode = ref('200');

// 是否定义了请求方法
function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

// 是否定义了请求方法的响应体
function hasDefinedCode(code: string) {
  return selectedMethodDetail?.value?.responseBodies?.some((item) => {
    return item.code === code;
  })
}

// 当前选中的请求方法详情
const selectedMethodDetail: any = ref(null);
// 当前选中的请求方法的响应体详情
const selectedCodeDetail: any = ref(null);
// 是否展示安全定义
const showSecurity = ref(false);

watch(() => {
  return selectedMethod.value
}, (newVal, oldVal) => {
  console.log('watch selectedMethod', newVal)

  selectedMethodDetail.value = interfaceMethodToObjMap.value[newVal]

  if (selectedMethodDetail.value) {
    store.dispatch('Debug/setInterface', selectedMethodDetail.value);

    showSecurity.value = !!selectedMethodDetail.value.security;
    selectedCodeDetail.value = selectedMethodDetail?.value?.responseBodies?.find((item) => {
      return item.code === selectedCode.value;
    })
  } else {
    store.dispatch('Debug/setInterface', {});
  }

}, {immediate: true});

watch(() => {
  return selectedCode.value
}, (newVal, oldVal) => {
  selectedCodeDetail.value = selectedMethodDetail?.value?.responseBodies?.find((item) => {
    return item.code === newVal;
  })
}, {immediate: true});

const selectedMethodIndex: any = computed(() => {
  return endpointDetail?.value?.interfaces?.findIndex((item) => {
    return item.method === selectedMethod.value;
  })
});
const selectedCodeIndex: any = computed(() => {
  return selectedMethodDetail?.value?.responseBodies?.findIndex((item) => {
    return item.code === selectedCode?.value;
  })
});

function goEditSecurity() {
  window.open(`/#/projectSetting/index?firtab=3&sectab=3&serveId=${endpointDetail.value.serveId}`,'_blank')
}

function delSecurity() {
  showSecurity.value = false;
  selectedMethodDetail.value.security = null;
}

function securityChange(val) {
  selectedMethodDetail.value.security = val || null;
}

function setSecurity() {
  showSecurity.value = true;
}

function addCookie() {
  selectedMethodDetail.value.cookies.push(cloneByJSON(defaultCookieParams));
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
  })
}

function addQueryParams() {
  selectedMethodDetail.value.params.push(cloneByJSON(defaultQueryParams));
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
  })
}

function addHeader() {
  selectedMethodDetail.value.headers.push(cloneByJSON(defaultHeaderParams));
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
  })
}

function addResponseHeader() {
  selectedCodeDetail.value.headers.push(cloneByJSON(defaultHeaderParams));
}

function addCodeResponse() {
  const item = {
    ...cloneByJSON(defaultCodeResponse),
    "code": selectedCode.value,
    "endpointId": selectedMethodDetail.value.id,
  }
  store.commit('Endpoint/setEndpointDetailByIndex', {
    methodIndex: selectedMethodIndex.value,
    codeIndex: selectedCodeIndex.value,
    value: item
  })
  selectedCodeDetail.value = item;
}

function addEndpoint() {
  const item = {
    ...cloneByJSON(defaultEndpointDetail),
    "projectId": endpointDetail.value.projectId,
    "serveId": endpointDetail.value.serveId,
    "useId": currentUser.value.id,
    "method": selectedMethod.value,
  }
  selectedMethodDetail.value = item;
  selectedCode.value = '200';
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    interfaces: [...endpointDetail.value.interfaces, item],
  })
}

/**
 * 添加路径参数
 * */
function addPathParams() {
  endpointDetail.value.pathParams.push(cloneByJSON(defaultPathParams));
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
}

/**
 * 删除路径参数
 * */
function deletePathParams(data) {
  endpointDetail.value.pathParams.splice(data.index, 1);
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
}


/**
 * 更新参数名称
 * */
function pathParamsNameChange(data) {
  console.log(data);
  endpointDetail.value.pathParams[data.index] = data;
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
}

/**
 * 处理 path 与 pathParams 字段联动的情况
 * */
function handlePathLink() {
  // ::::todo 待补充
  // let parsePathReg = /\{(\w+)\}/g
  // let path = endpointDetail.value.path;
  // let params = path.match(parsePathReg);
  // if (data.name) {
  //   params.push(data.name)
  // }
  console.log('handlePathLink');
}

function deleteParams(type, index) {
  selectedMethodDetail.value[type].splice(index, 1);
}

function handleParamsChange(type, data) {
  selectedMethodDetail.value[type][data.index] = {...data};
}

function deleteResHeader(index) {
  selectedCodeDetail.value.headers.splice(index, 1);
}

function handleResHeaderChange(data) {
  selectedCodeDetail.value.headers[data.index] = {...data};
}

function updatePath(e) {
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    path: e.target.value,
  })
}

const activeReqBodySchema: any = ref({
  content: null,
  examples: [],
});

const activeResBodySchema: any = ref({
  content: null,
  examples: [],
});


watch(() => {
  return selectedMethodDetail?.value?.requestBody?.schemaItem?.content
}, (newVal, oldValue) => {
  activeReqBodySchema.value.content = JSON.parse(newVal || 'null')
}, {immediate: true});

watch(() => {
  return selectedMethodDetail?.value?.requestBody?.examples
}, (newVal, oldValue) => {
  activeReqBodySchema.value.examples = JSON.parse(newVal || '[]')
}, {immediate: true});

watch(() => {
  return selectedCodeDetail?.value?.schemaItem?.content
}, (newVal, oldValue) => {
  activeResBodySchema.value.content = JSON.parse(newVal || 'null')
}, {immediate: true});

watch(() => {
  return selectedCodeDetail?.value?.examples
}, (newVal, oldValue) => {
  activeResBodySchema.value.examples = JSON.parse(newVal || '[]')
}, {immediate: true});


async function generateFromJSON(type: string, JSONStr: string) {
  const res = await store.dispatch('Endpoint/example2schema',
      {data: JSONStr}
  );
  if (type === 'req') {
    activeReqBodySchema.value.content = res;
  }
  if (type === 'res') {
    activeResBodySchema.value.content = res;
  }
}

async function handleGenerateExample(type: string, examples: any) {
  const res = await store.dispatch('Endpoint/schema2example',
      {data: JSON.stringify(type === 'req' ? activeReqBodySchema.value.content : activeResBodySchema.value.content)}
  );
  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res),
  };
  if (type === 'req') {
    activeReqBodySchema.value.examples.push(example);
  }
  if (type === 'res') {
    activeResBodySchema.value.examples.push(example);
  }
}

function handleChange(type, json: any) {
  const {content, examples} = json;
  if (type === 'req') {
    // activeReqBodySchema.value.content = content;
    // activeReqBodySchema.value.type = content.type;
    if(selectedMethodDetail?.value?.requestBody){
      selectedMethodDetail.value.requestBody.schemaItem.content = JSON.stringify(content);
      selectedMethodDetail.value.requestBody.examples = JSON.stringify(examples);
      selectedMethodDetail.value.requestBody.schemaItem.type = content.type;
    }
    // store.commit('Endpoint/setInterfaceDetail', {
    //   ...interfaceDetail.value,
    // });
  }
  if (type === 'res') {
    // activeResBodySchema.value.content = content;
    // activeResBodySchema.value.type = content.type;
    // activeResBodySchema.value.examples = examples;
    if(selectedCodeDetail?.value){
      selectedCodeDetail.value.schemaItem.content = JSON.stringify(content);
      selectedCodeDetail.value.examples = JSON.stringify(examples);
      selectedCodeDetail.value.schemaItem.type = content.type;
    }
    // store.commit('Endpoint/setInterfaceDetailByIndex', {
    //   methodIndex: selectedMethodIndex.value,
    //   codeIndex: selectedCodeIndex.value,
    //   value: selectedCodeDetail.value
    // })
  }
}

</script>
<style lang="less" scoped>
.content {
  padding-top: 16px;
}

.form-item {
  margin-bottom: 16px;
  align-items: baseline;
}

.path-param-list {
  margin-top: 16px;
}

.form-label {
  font-weight: bold;
}

.form-item-request {
  margin-top: 16px;

  .form-item-request-item {
    margin-top: 16px;
    align-items: center;
  }

  .form-item-response {
    margin-top: 16px;

    .form-item-response-item {
      margin-top: 16px;
      align-items: center;
    }
  }
}

.params-defined-item-header {
  font-weight: bold;
  margin-bottom: 8px;
  margin-top: 8px;
}

.has-defined {
  color: #1890ff;
  //font-weight: bold;
}
</style>
