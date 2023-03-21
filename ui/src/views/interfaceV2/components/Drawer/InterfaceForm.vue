<template>
  <div class="interface-form">
    <!-- ::::路径定义方式 -->
    <a-row class="request-module">
      <a-col :span="2" class="path-defined-label">路径</a-col>
      <a-col :span="18">
        <a-input v-model:value="interfaceDetail.path">
          <template #addonBefore>
            <a-select :value="'http://localhost:3000'" style="width: 200px">
              <a-select-option value="http://localhost:3000">http://localhost:3000</a-select-option>
              <a-select-option value="http://localhost:3001">http://localhost:3001</a-select-option>
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
        <div class="pathParam">
          <div v-for="(item,index) in interfaceDetail.pathParams" :key="item.id">
            <FieldItem :fieldData="item"
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
    <a-row class="request-module">
      <a-col :span="2" class="path-defined-label">请求方式</a-col>
      <a-col :span="22">
        <!-- ::::请求方法定义 -->
        <a-radio-group
            @change="selectedMethodChange"
            v-model:value="selectedMethod" button-style="solid">
          <a-radio-button :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
            {{ method.label }}
          </a-radio-button>
        </a-radio-group>
        <div class="request-module-method-defined">
          <div v-if="selectedMethodDetail">
            <!-- ::::Operation ID -->
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label">
                Operation ID
              </a-col>
              <a-col :span="12">
                <a-input v-model:value="selectedMethodDetail.operationId"/>
              </a-col>
            </a-row>
            <!-- ::::Description -->
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label">
                Description
              </a-col>
              <a-col :span="12">
                <a-input v-model:value="selectedMethodDetail.description"/>
              </a-col>
            </a-row>
            <!-- ::::增加请求参数 -->
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label">
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
            <a-row class="method-item">
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
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label">
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
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label"></a-col>
              <a-col :span="21">
                <a-input placeholder="请输入描述" v-model:value="selectedMethodDetail.requestBody.description"/>
              </a-col>
            </a-row>
            <!-- ::::增加请求体 - scheme定义 -->
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label"></a-col>
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
            <a-row class="method-item">
              <a-col :span="3" class="method-item-label">
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
                <div class="request-module-method-defined">
                  <div v-if="selectedCodeDetail">
                    <!-- ::::Description -->
                    <a-row class="method-item">
                      <a-col :span="4" class="method-item-label">
                        Description
                      </a-col>
                      <a-col :span="18">
                        <a-input v-model:value="selectedCodeDetail.desc"/>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应头 -->
                    <a-row class="method-item">
                      <a-col :span="4" class="method-item-label">
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
                    <a-row class="method-item">
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
                    <a-row class="method-item">
                      <a-col :span="4" class="method-item-label">
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
                    <a-row class="method-item">
                      <a-col :span="4" class="method-item-label"></a-col>
                      <a-col :span="18">
                        <a-input placeholder="请输入描述" v-model:value="selectedCodeDetail.description"/>
                      </a-col>
                    </a-row>
                    <!-- ::::增加响应体 - scheme定义 -->
                    <a-row class="method-item">
                      <a-col :span="4" class="method-item-label"></a-col>
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
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {
  defineComponent,
  reactive,
  ref,
  toRaw,
  UnwrapRef,
  defineProps,
  defineEmits,
  watch,
  computed,
  onUnmounted
} from 'vue';
import {useStore} from "vuex";
import {requestMethodOpts, interfaceStatus, mediaTypesOpts, repCodeOpts} from '@/config/constant';
import {getInterfaceDetail, saveInterface, getYaml} from '../../service';
import {PlusOutlined, EditOutlined, CodeOutlined, BarsOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import FieldItem from './FieldItem.vue'
import {momentUtc} from '@/utils/datetime';
import {Interface} from "@/views/interface/data";

const store = useStore<{ InterfaceV2, ProjectGlobal }>();
const interfaceDetail = computed<Interface[]>(() => store.state.InterfaceV2.interfaceDetail);
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
const collapseActiveKey = ref(['1']);
const activeKey = ref('1');
const activeResCodeKey = ref('1');
const selectedMethod = ref('GET');
const selectedCode = ref('200');
const interfaceNameEditable = ref(false);
const checked1 = ref(false);

const selectedMethodDetail: any = ref(null);
const selectedCodeDetail: any = ref(null);

function editorChange(newScriptCode) {
  console.log(832, interfaceDetail.value)
  // if (selectedMethodDetail.value?.requestBody?.schemaItem?.content) {
  //   selectedMethodDetail.value?.requestBody?.schemaItem?.content = newScriptCode
  // }
}

function selectedMethodChange(e) {
  let curInterface = interfaceDetail.value.interfaces.find((item) => {
    return item.method === e.target.value;
  })
  selectedMethodDetail.value = curInterface;
}

const showMode = ref('form');

const yamlCode = ref('');


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

// const interfaceDetailNameRef: any = ref(null)


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


//同步替换删除path中的param参数
// let path = interfaceDetail.value.path;
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

function updateTitle(title) {
  interfaceDetail.value.title = title
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
// interfaceDetail.value = null;
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


// onUnmounted(() => {
//
// })
const loading = ref(false);


/**
 * 打开窗口时，需要重新获取
 * */
watch(() => {
  return props.visible;
}, async (newVal) => {
  if (newVal) {
    let res = await getInterfaceDetail(props.interfaceId);
    let data = res.data;
    data.createdAt = momentUtc(data.createdAt);
    data.updatedAt = momentUtc(data.updatedAt);
    interfaceDetail.value = {...res.data};


// todo 默认选中第一个有值的method ，临时方案，应该高亮展示一些场景
    if (interfaceDetail.value.interfaces[0]?.method) {
      selectedMethod.value = interfaceDetail.value.interfaces[0].method;
      selectedMethodDetail.value = interfaceDetail.value.interfaces[0];
    }
    if (selectedMethodDetail.value?.responseBodies[0]?.code) {
      selectedCode.value = selectedMethodDetail.value?.responseBodies[0]?.code;
      selectedCodeDetail.value = selectedMethodDetail.value?.responseBodies[0];
    }


  } else {
// interfaceDetail.value = null;
// selectedMethodDetail.value = null;
// selectedCodeDetail.value = null;
  }
}, {
  immediate: true
})

</script>

<style lang="less" scoped>

.drawer {
  margin-bottom: 60px;

  .title {
    width: auto;

    .ant-input-affix-wrapper {
      width: auto;
      border: none;

      &:focus {
        border: none;
        outline: none;
        box-shadow: none;
      }
    }

    input {
      width: auto;
      border: none;

      &:focus {
        border: none;
        border: none;
        outline: none;
        box-shadow: none;
      }
    }
  }


}

.drawer-btns {
  background: #ffffff;
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  position: absolute;
  bottom: 0;
  right: 0;
  width: 100%;
  height: 60px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 16px;
}

:deep(.ant-drawer-body) {
  padding: 0 !important;
  border: 1px solid red;
  margin-bottom: 60px;
}

.drawer {

}


.interfaceName {
  min-width: 1em;

  &:focus {
    outline: none;
  }

  &:hover,
  &:focus {
    outline: none;
    //border-bottom: 1px solid rgba(0, 0, 0, 0.65);;
  }
}

.schema {
  margin-left: 20px;
  width: 50%;
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  border: 1px solid rgba(0, 0, 0, .1);
  border-radius: 8px;
  padding: 12px;
}

.method-item {
  margin-bottom: 16px;
}

.request-module {
  margin-top: 16px;
}

.request-module-method-defined {
  margin-top: 16px;
}

</style>
