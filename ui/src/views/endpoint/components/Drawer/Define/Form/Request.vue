<template>
  <!-- 请求方式定义 -->
  <a-row class="form-item">
    <a-col :span="2" class="form-label">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">请求方式</span>
    </a-col>
    <a-col :span="22">
      <!-- 请求方法定义 -->
      <a-radio-group v-model:value="selectedMethod" button-style="solid">
        <a-radio-button
            :class="{'has-defined': hasDefinedMethod(method.value)}"
            :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
          {{ method.label }}
        </a-radio-button>
      </a-radio-group>
      <div class="form-item-request" v-if="collapse">
        <div v-if="selectedMethodDetail">
          <!-- Operation ID -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label">
              Operation ID
            </a-col>
            <a-col :span="12">
              <a-input v-model:value="selectedMethodDetail.operationId"/>
            </a-col>
          </a-row>
          <!-- Description -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label">
              Description
            </a-col>
            <a-col :span="12">
              <a-input v-model:value="selectedMethodDetail.description"/>
            </a-col>
          </a-row>
          <!-- 增加请求参数 -->
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
          <!-- 请求参数展示：headers、cookies、query params等 -->
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
          <!-- 增加请求体 -->
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
          <!-- 增加请求体 - 描述  -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label"></a-col>
            <a-col :span="20">
              <a-input placeholder="请输入描述" v-model:value="selectedMethodDetail.requestBody.description"/>
            </a-col>
          </a-row>
          <!-- 增加请求体 - scheme定义 -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label"></a-col>
            <a-col :span="21">
              <SchemaEditor
                  @generateFromJSON="generateFromJSON"
                  @generateExample="handleGenerateExample"
                  @change="handleChange"
                  :tab-content-style="{width:'100%'}"
                  :value="activeReqBodySchema"/>
            </a-col>
          </a-row>
          <!-- 响应定义  -->
          <Response :method="selectedMethod"/>
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
  defaultCookieParams,
  defaultHeaderParams,
  defaultQueryParams,
  defaultEndpointDetail,
} from '@/config/constant';
import {PlusOutlined, DeleteOutlined, RightOutlined, DownOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import Response from './Response.vue';
import {Endpoint} from "@/views/endpoint/data";
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {cloneByJSON} from "@/utils/object";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);
const currInterface = computed<any>(() => store.state.Debug?.currInterface);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
const securityOpts: any = computed<any>(() => store.state.Endpoint.securityOpts);
const props = defineProps({});
const emit = defineEmits([]);
const selectedMethod = ref(currInterface.value?.method ? currInterface.value?.method : 'GET');
// 是否折叠,默认展开
const collapse = ref(true);

// 是否定义了请求方法
function hasDefinedMethod(method: string) {
  return endpointDetail?.value?.interfaces?.some((item) => {
    return item.method === method;
  })
}

// 当前选中的请求方法详情
const selectedMethodDetail: any = ref(null);
// 是否展示安全定义
const showSecurity = ref(false);

watch(() => {
  return selectedMethod.value
}, (newVal, oldVal) => {
  selectedMethodDetail.value = interfaceMethodToObjMap.value[newVal];
  if (selectedMethodDetail.value) {
    store.dispatch('Debug/setInterface', selectedMethodDetail.value);
    store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
    showSecurity.value = !!selectedMethodDetail.value.security;
  } else {
    store.dispatch('Debug/setInterface', {});
    store.commit('Endpoint/setSelectedMethodDetail', {});
  }
}, {immediate: true});


function goEditSecurity() {
  window.open(`/#/projectSetting/index?firtab=3&sectab=3&serveId=${endpointDetail.value.serveId}`, '_blank')
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

function addEndpoint() {
  const item = {
    ...cloneByJSON(defaultEndpointDetail),
    "projectId": endpointDetail.value.projectId,
    "serveId": endpointDetail.value.serveId,
    "useId": currentUser.value.id,
    "method": selectedMethod.value,
  }
  selectedMethodDetail.value = item;
  store.dispatch('Debug/setInterface', selectedMethodDetail.value);
  store.commit('Endpoint/setInterfaceMethodToObjMap', {
    method: item.method,
    value: item,
  });
  store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    interfaces: [...endpointDetail.value.interfaces, item],
  })
}


function deleteParams(type, index) {
  selectedMethodDetail.value[type].splice(index, 1);
}

function handleParamsChange(type, data) {
  selectedMethodDetail.value[type][data.index] = {...data};
}

const activeReqBodySchema: any = ref({
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


async function generateFromJSON(type: string, JSONStr: string) {
  const res = await store.dispatch('Endpoint/example2schema', {data: JSONStr});
  activeReqBodySchema.value.content = res;
}

async function handleGenerateExample(examples: any) {
  const res = await store.dispatch('Endpoint/schema2example', {data: JSON.stringify(activeReqBodySchema.value.content)});
  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res),
  };
  activeReqBodySchema.value.examples.push(example);
}

function handleChange(json: any) {
  const {content, examples} = json;
  // activeReqBodySchema.value.content = content;
  // activeReqBodySchema.value.type = content.type;
  if (selectedMethodDetail?.value?.requestBody) {
    selectedMethodDetail.value.requestBody.schemaItem.content = JSON.stringify(content);
    selectedMethodDetail.value.requestBody.examples = JSON.stringify(examples);
    selectedMethodDetail.value.requestBody.schemaItem.type = content.type;
  }
  // store.commit('Endpoint/setInterfaceDetail', {
  //   ...interfaceDetail.value,
  // });
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

.label-name {
  margin-left: 4px;
}
</style>
