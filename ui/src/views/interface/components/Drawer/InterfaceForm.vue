<template>
  <div class="content">
    <!-- ::::路径定义方式 -->
    <a-row class="form-item">
      <a-col :span="2" class="form-label">路径</a-col>
      <a-col :span="16">
        <a-input :value="interfaceDetail.path" @change="updatePath">
          <template #addonBefore>
            <a-select :value="'http://localhost:3000'" style="width: 200px">
              <a-select-option value="http://localhost:3000">http://localhost:3000</a-select-option>
            </a-select>
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
                :fieldData="item"
                @del="deletePathParams(index)"
                @paramsNameChange="paramsNameChange"
                @settingOther="settingOtherForPathParams"
                @setRef="setRefForPathParams"
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
        <a-radio-group
            @change="selectedMethodChange"
            v-model:value="selectedMethod" button-style="solid">
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
                    :value="activeSchema"/>
              </a-col>
            </a-row>
            <!-- ::::响应定义  -->
            <a-row class="form-item-response">
              <a-col :span="3" class="form-label">
                选择响应代码
              </a-col>
              <a-col :span="21">
                <a-radio-group
                    @change="selectedCodeChange"
                    v-model:value="selectedCode" button-style="solid">
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
import {requestMethodOpts, interfaceStatus, mediaTypesOpts, repCodeOpts} from '@/config/constant';
import {saveInterface} from '../../service';
import {PlusOutlined, EditOutlined, CodeOutlined, BarsOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import FieldItem from './FieldItem.vue'
import {momentUtc} from '@/utils/datetime';
import {Interface} from "@/views/interface/data";

const store = useStore<{ Interface, ProjectGlobal }>();
const interfaceDetail: any = computed<Interface[]>(() => store.state.Interface.interfaceDetail);
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {example2schema, schema2example} from "@/views/projectSetting/service";

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  interfaceId: {
    required: true,
  }
});
const emit = defineEmits(['ok', 'close', 'refreshList']);
const activeKey = ref('1');
const selectedMethod = ref('GET');
const selectedCode = ref('200');
const selectedMethodDetail: any = ref(null);
const selectedCodeDetail: any = ref(null);

function selectedMethodChange(e) {
  let curInterface = interfaceDetail.value.interfaces.find((item) => {
    return item.method === e.target.value;
  })
  selectedMethodDetail.value = curInterface;
}

function selectedCodeChange(e) {
  let curCode = selectedMethodDetail.value.responseBodies.find((item) => {
    return item.code == e.target.value;
  })
  selectedCodeDetail.value = curCode;
}

function setSecurity() {
  console.log('setSecurity')
}

function addCookie() {
  selectedMethodDetail.value.cookies.push({
    name: '',
    value: '',
    desc: '',
    type: 'string',
  })
}

function addQueryParams() {
  selectedMethodDetail.value.params.push({
    name: '',
    value: '',
    desc: '',
    type: 'string',
  })
}

function addHeader() {
  selectedMethodDetail.value.headers.push({
    name: '',
    value: '',
    desc: '',
    type: 'string',
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
    "createdAt": "2023-02-10T10:30:30+08:00",
    "updatedAt": "2023-02-10T10:30:30+08:00",
    "code": selectedCode.value,
// "interfaceId": 49,
    "mediaType": "",
    "description": "",
// "schemaRefId": 1,
    "examples": "{\"user\":{\"value\":{\"id\":1,\"name\":\"王大锤\"}},\"product\":{\"value\":{\"id\":1,\"name\":\"服装\"}}}",
    "schemaItem": {
// "id": 3,
      "createdAt": "2023-02-10T10:30:30+08:00",
      "updatedAt": "2023-02-10T10:30:30+08:00",
      "name": "name",
      "type": "object",
      "content": "{\"id\":{\"type\":\"integer\",\"format\":\"string\"},\"name\":{\"type\":\"string\",\"format\":\"string\"}}",
// "ResponseBodyId": 3
    },
    "headers": []
  };
  selectedMethodDetail.value.responseBodies.push(tpl)
  selectedCodeDetail.value = tpl;
}

/**
 * 添加路径参数
 * */
function addPathParams() {
  interfaceDetail.value.pathParams.push({
    name: '',
    type: 'string',
    desc: ''
  })
  // if (interfaceDetail.value?.pathParams?.length > 0) {
  //
  // } else {
  //   interfaceDetail.value.pathParams = [
  //     {
  //       name: '',
  //       type: 'string',
  //       desc: ''
  //     }
  //   ]
  // }
  // 同步替换删除path中的param参数
  // interfaceDetail.value.path = path.replace(`{${data.name}}`, '');
}

function addInterface() {
  const tpl = {
// "id": 49,
    "createdAt": "2023-02-10T10:30:30+08:00",
    "updatedAt": "2023-02-10T10:30:30+08:00",
    "name": "",
    "desc": "",
    "endpoint_id": 34,
    "security": "token,api_key",
    "isLeaf": false,
    "parentId": 0,
    "projectId": 0,
    "useId": 0,
    "ordr": 0,
    "slots": null,
    "url": "",
    "method": selectedMethod.value,
    "body": "{}",
    "bodyType": "",
    "authorizationType": "",
    "preRequestScript": "",
    "validationScript": "",
    "children": null,
    "params": [],
    "headers": [],
    "cookies": [],
    "requestBody": {
      "mediaType": "application/json",
      "description": "",
      "examples": "{\"user\":{\"value\":{\"id\":1,\"name\":\"王大锤\"}},\"product\":{\"value\":{\"id\":1,\"name\":\"服装\"}}}",
      "schemaItem": {
        "name": "name",
        "type": "object",
        "content": "{\"id\":{\"type\":\"integer\",\"format\":\"string\"},\"name\":{\"type\":\"string\",\"format\":\"string\"}}"
      }
    },
    "responseBodies": []
  };
  interfaceDetail.value.interfaces.push(tpl);
  selectedMethodDetail.value = tpl;
  selectedCode.value = '200';
  selectedCodeDetail.value = null
}

/**
 * 删除路径参数
 * */
function deletePathParams(index) {
// let index = interfaceDetail.value.pathParams.find((item) => {
//   return item.id === data.id;
// })
  interfaceDetail.value.pathParams.splice(index, 1);

//同步替换删除path中的param参数
// let path = interfaceDetail.value.path;
// interfaceDetail.value.path = path.replace(`{${data.name}}`, '');

}

function deleteParams(type, index) {
// let index = selectedMethodDetail.value[type].find((item) => {
//   return item.id === id;
// })
  selectedMethodDetail.value[type].splice(index, 1);
}

function deleteResHeader(index) {
  selectedCodeDetail.value.headers.splice(index, 1);
}

function paramsNameChange(val) {
// todo 待解析，联动接口字段
// var a = 'api/user/{id}/{detailID}'
// 解析path 中的参数
// let parsePathReg = /\{(\w+)\}/g
// let path = interfaceDetail.value.path;
// let params = path.match(parsePathReg);
// if (val) {
//   params.push(`{${val}}`)
// }
// // todo 需要处理，几个表单项的联动场景
// console.log(832, params, val);
// interfaceDetail.value.path = path.replace(`{${data.name}}`, '');
}

function settingOtherForPathParams() {
  console.log('settingOtherForPathParams')
}

function setRefForPathParams() {
  console.log('setRefForPathParams')
}

function setPathParamsRequire() {
  console.log('setPathParamsRequire')
}

function updatePath(e) {
  store.commit('Interface/setInterfaceDetail', {
    ...interfaceDetail.value,
    path: e.target.value,
  })
}

const key = ref('request');

function addReqBody() {
  console.log('add request body');
}

function addResBody() {
  console.log('add request body');
}

// 取消
async function cancal() {
  emit('close');
}

// 保存
async function save() {
  let res = await saveInterface(interfaceDetail.value);
  if (res.code === 0) {
    message.success('保存成功');
    emit('close');
    emit('refreshList')
  }
}

const activeSchema: any = ref({
  content: {
    "type": "object",
    "properties": {
      "a": {
        "type": "string"
      },
      "b": {
        "properties": {
          "c": {
            "properties": {
              "d": {
                "type": "number"
              }
            },
            "type": "object"
          }
        },
        "type": "object"
      }
    }
  },
  examples: [
    {
      name: 'example 1',
      content: '{"a":"string","b":{"c":{"d":0}}}'
    },
    {
      name: 'example 2',
      content: '{"a":"string","b":{"c":{"d":0}}}'
    }
  ],
  type: 'object'
});
const contentStr = ref('');
const schemaType = ref('object');
const exampleStr = ref('');

async function generateFromJSON(JSONStr: string) {
  const res = await example2schema({
    data: JSONStr
  });
  if (res.code === 0) {
// activeSchema.value.content = res.data;
// contentStr.value = JSON.stringify(res.data);
// schemaType.value = res.data.type;
  }
}

async function handleGenerateExample(examples: any) {
  const res = await schema2example({
    data: contentStr.value
  });

  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res.data),
  };

  if (res.code === 0) {
    activeSchema.value.examples.push(example);
  }
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
</style>
