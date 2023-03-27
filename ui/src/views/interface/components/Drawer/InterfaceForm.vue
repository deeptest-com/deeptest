<template>
  <div class="content">
    <!-- ::::路径定义方式 -->
    <a-row class="form-item">
      <a-col :span="2" class="form-label">路径</a-col>
      <a-col :span="16">
        <a-input :value="interfaceDetail.path" @change="updatePath" placeholder="请输入路径">
          <template #addonBefore>
            <a-select
                :options="serveServers"
                :value="serveServers?.[0]?.value"
                placeholder="请选择服务器"
                style="width: 200px;text-align: left"/>
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
          <div v-for="(item,index) in interfaceDetail.pathParams" :key="item.id">
            <FieldItem
                :fieldData="{...item,index:index}"
                :showRequire="true"
                @del="deletePathParams"
                @paramsNameChange="paramsNameChange"
                @setRequire="setPathParamsRequire"/>
          </div>
        </div>
      </a-col>
    </a-row>
    <!-- ::::请求方式定义 -->
    <a-row class="form-item">
      <a-col :span="2" class="form-label">请求方式</a-col>
      <a-col :span="22">
        <!-- ::::请求方法定义 -->
        <a-radio-group v-model:value="selectedMethod" button-style="solid">
          <a-radio-button :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
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
                    <div class="params-defined-item" v-if="selectedMethodDetail?.headers?.length">
                      <div class="params-defined-item-header">
                        <span>Header</span>
                      </div>
                      <div class="header-defined header-defined-items">
                        <div v-for="(item,index) in selectedMethodDetail.headers" :key="item.id">
                          <FieldItem
                              :fieldData="item"
                              @del="deleteParams('headers',index)"
                          />
                        </div>
                      </div>
                    </div>
                    <div class="params-defined-item" v-if="selectedMethodDetail?.params?.length">
                      <div class="params-defined-item-header">
                        <span>Query Params</span>
                      </div>
                      <div class="header-defined ">
                        <div v-for="(item,index) in selectedMethodDetail.params" :key="item.id">
                          <FieldItem
                              :fieldData="item"
                              @del="deleteParams('params',index)"/>
                        </div>
                      </div>
                    </div>
                    <div class="params-defined-item" v-if="selectedMethodDetail?.cookies?.length">
                      <div class="params-defined-item-header">
                        <span>Cookie</span>
                      </div>
                      <div class="header-defined ">
                        <div v-for="(item,index) in selectedMethodDetail.cookies" :key="item.id">
                          <FieldItem :fieldData="item"
                                     @del="deleteParams('cookies',index)"/>
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
                <a-button
                    v-if="!selectedMethodDetail.requestBody"
                    type="primary" @click="addReqBody">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `添加` }}
                </a-button>

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
                    @generateFromJSON="generateFromJSON"
                    @exampleChange="handleExampleChange"
                    @generateExample="handleGenerateExample"
                    @schemaTypeChange="handleSchemaTypeChange"
                    @contentChange="handleContentChange"
                    :tab-content-style="{width:'700px'}"
                    :value="selectedMethodDetail.requestBody.schemaItem.content"/>
              </a-col>
            </a-row>
            <!-- ::::响应定义  -->
            <a-row class="form-item-response">
              <a-col :span="3" class="form-label">
                选择响应代码
              </a-col>
              <a-col :span="21">
                <a-radio-group v-model:value="selectedCode" button-style="solid">
                  <a-radio-button :key="code.value" v-for="code in repCodeOpts" :value="code.value">
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
                                  <FieldItem
                                      :fieldData="item"
                                      @del="deleteResHeader(index)"/>
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
                        <!--                              <a-button-->
                        <!--                                  v-if="!selectedCodeDetail.mediaType"-->
                        <!--                                  type="primary" @click="addResBody">-->
                        <!--                                <template #icon>-->
                        <!--                                  <PlusOutlined/>-->
                        <!--                                </template>-->
                        <!--                                {{ `添加` }}-->
                        <!--                              </a-button>-->
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
                            @generateFromJSON="generateFromJSON"
                            @exampleChange="handleExampleChange"
                            @generateExample="handleGenerateExample"
                            @schemaTypeChange="handleSchemaTypeChange"
                            @contentChange="handleContentChange"
                            :tab-content-style="{width:'600px'}"
                            :value="activeSchema"/>
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
            <a-button type="primary" @click="addInterface">
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
  defaultPathParams
} from '@/config/constant';
import {PlusOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import FieldItem from './FieldItem.vue'
import {Interface} from "@/views/interface/data";

const store = useStore<{ Interface, ProjectGlobal, User }>();
const interfaceDetail: any = computed<Interface>(() => store.state.Interface.interfaceDetail);
const currentUser: any = computed<Interface>(() => store.state.User.currentUser);
const serveServers: any = computed<Interface>(() => store.state.Interface.serveServers);
import SchemaEditor from '@/components/SchemaEditor/index.vue';

const props = defineProps({});
const emit = defineEmits([]);

const selectedMethod = ref('GET');
const selectedCode = ref('200');
// 当前选中的请求方法详情
const selectedMethodDetail: any = computed(() => {
  return interfaceDetail?.value?.interfaces?.find((item) => {
    return item.method === selectedMethod.value;
  })
});
const selectedMethodIndex: any = computed(() => {
  return interfaceDetail?.value?.interfaces?.findIndex((item) => {
    return item.method === selectedMethod.value;
  })
});
// 当前选中的请求方法的响应体详情
const selectedCodeDetail: any = computed(() => {
  return selectedMethodDetail?.value?.responseBodies?.find((item) => {
    return item.code === selectedCode?.value;
  })
});

const selectedCodeIndex: any = computed(() => {
  return selectedMethodDetail?.value?.responseBodies?.findIndex((item) => {
    return item.code === selectedCode?.value;
  })
});

function setSecurity() {
  console.log('setSecurity');
}

function addCookie() {
  selectedMethodDetail.value.cookies.push(defaultCookieParams);
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
  })
}

function addQueryParams() {
  selectedMethodDetail.value.params.push(defaultQueryParams);
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
  })
}

function addHeader() {
  selectedMethodDetail.value.headers.push(defaultHeaderParams);
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
  })
}

function addResponseHeader() {
  selectedCodeDetail.value.headers.push({
    name: '',
    desc: '',
    type: 'string',
  })
}

function addCodeResponse() {
  const tpl = {
    "code": selectedCode.value,
    "interfaceId": selectedMethodDetail.value.id,
    "mediaType": "application/json",
    "description": "",
    "schemaRefId": null,
    "examples": "",
    "schemaItem": {
      "id": null,
      "name": "",
      "type": "object",
      "content": "",
      "ResponseBodyId": null
    },
    "headers": []
  };
  store.commit('Interface/setInterfaceDetailByIndex', {
    methodIndex:selectedMethodIndex.value,
    codeIndex:selectedCodeIndex.value,
    value:tpl
  })
}

/**
 * 添加路径参数
 * */
function addPathParams() {
  interfaceDetail.value.pathParams.push(defaultPathParams);
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    pathParams: interfaceDetail.value.pathParams
  })
}

/**
 * 删除路径参数
 * */
function deletePathParams(data) {
  interfaceDetail.value.pathParams.splice(data.index, 1);
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    pathParams: interfaceDetail.value.pathParams
  })
}

/**
 * 更新路径参数的 require 为 true
 * */
function setPathParamsRequire(data) {
  interfaceDetail.value.pathParams[data.index] = {...data};
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    pathParams: interfaceDetail.value.pathParams
  })
}

/**
 * 更新参数名称
 * */
function paramsNameChange(data) {
  interfaceDetail.value.pathParams[data.index] = data;
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    pathParams: interfaceDetail.value.pathParams
  })
}

/**
 * 处理 path 与 pathParams 字段联动的情况
 * */
function handlePathLink() {
  // ::::todo 待补充
  // let parsePathReg = /\{(\w+)\}/g
  // let path = interfaceDetail.value.path;
  // let params = path.match(parsePathReg);
  // if (data.name) {
  //   params.push(data.name)
  // }
  console.log('handlePathLink');
}

function deleteParams(type, index) {
  selectedMethodDetail.value[type].splice(index, 1);
}

function addInterface() {
  const defaultInterfaceDetail = {
    "name": "",
    "projectId": interfaceDetail.value.projectId,
    "serveId": interfaceDetail.value.serveId,
    "useId": currentUser.value.id,
    "method": selectedMethod.value,
    "description": "",
    "operationId": "",
    "security": "token,api_key",
    "requestBody": {
      "id": null,
      "interfaceId": null,
      "mediaType": "",
      "description": "",
      "schemaRefId": null,
      "examples": "",
      "schemaItem": {
        "id": null,
        "name": "",
        "type": "object",
        "content": "",
        "requestBodyId": null
      }
    },
    "responseBodies": [],
    "bodyType": "application/json", // todo 确定 UI 交互
    "params": [],
    "headers": [],
    "cookies": []
  }
  interfaceDetail.value.interfaces.push(defaultInterfaceDetail);

  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    interfaces: [...interfaceDetail.value.interfaces],
  })

}

function deleteResHeader(index) {
  selectedCodeDetail.value.headers.splice(index, 1);
}


function updatePath(e) {
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    path: e.target.value,
  })
}


function addReqBody() {
  console.log('add request body');
}

function addResBody() {
  console.log('add request body');
}


const contentStr = ref('');
const schemaType = ref('object');
const exampleStr = ref('');

async function generateFromJSON(JSONStr: string) {
  console.log('generateFromJSON');
}

async function handleGenerateExample(examples: any) {
  console.log('handleGenerateExample');
}

function handleContentChange(str: string) {
  contentStr.value = str;
}

function handleSchemaTypeChange(str: string) {
  schemaType.value = str;
}

function handleExampleChange(str: string) {
  exampleStr.value = str;
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
}
</style>
