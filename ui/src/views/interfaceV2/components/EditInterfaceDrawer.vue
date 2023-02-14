<template>
  <a-drawer
      :placement="'right'"
      :width="1000"
      :closable="true"
      :visible="visible"
      class="drawer"
      wrapClassName="drawer-1"
      :bodyStyle="{padding:0,marginBottom:'60px'}"
      @close="onCloseDrawer">
    <!-- 头部信息  -->
    <template #title>
      <div class="title">
        <a-space :size="8">
          <contenteditable
              tag="div"
              :ref="interfaceDetailNameRef"
              class="interfaceName"
              :contenteditable="true"
              v-model="interfaceDetail.title"
              :no-nl="true"
              :no-html="true"
              @returned="enterPressed"/>
          <EditOutlined @click="editInterfaceDetailName"/>
        </a-space>
      </div>
    </template>
    <!-- 基本信息  -->
    <a-card
        class="card-baseInfo"
        :bordered="false"
        title="基本信息">
      <template #extra><a href="#">more</a></template>
      <a-descriptions :title="null">
        <a-descriptions-item label="创建人">Zhou Maomao</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ interfaceDetail.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="状态">Hangzhou, Zhejiang</a-descriptions-item>
        <a-descriptions-item label="服务版本">empty</a-descriptions-item>
        <a-descriptions-item label="最近更新">{{ interfaceDetail.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-card>
    <!-- 接口定义 -->
    <a-card
        style="width: 100%"
        title="接口定义"
        :tab-list="tabList"
        :active-tab-key="key"
        @tabChange="key => onTabChange(key, 'key')"
    >
      <div v-if="key === 'request'">
        <!-- ::::路径定义方式 -->
        <a-form-item class="request-module" label="路径">
          <a-input v-model:value="interfaceDetail.path">
            <template #addonBefore>
              <a-select :value="'http://localhost:3000'" style="width: 200px">
                <a-select-option value="http://localhost:3000">http://localhost:3000</a-select-option>
                <a-select-option value="http://localhost:3001">http://localhost:3001</a-select-option>
              </a-select>
            </template>
            <template #addonAfter>
              <a-button type="primary" @click="addPathParams">
                <template #icon>
                  <PlusOutlined/>
                </template>
                路径参数
              </a-button>
            </template>
          </a-input>
        </a-form-item>
        <!-- ::::路径参数 -->
        <div class="pathParam">
          <div v-for="item in interfaceDetail.pathParams" :key="item.id">
            <FieldItem :fieldData="item"
                       @del="deletePathParams"
                       @paramsNameChange="paramsNameChange"
                       @settingOther="settingOtherForPathParams"
                       @setRef="setRefForPathParams"
                       @setRequire="setPathParamsRequire"/>
          </div>
        </div>
        <!-- ::::请求方式定义 -->
        <a-form-item class="request-module" label="请求方式">
          <a-radio-group
              @change="selectedMethodChange"
              v-model:value="selectedMethod" button-style="solid">
            <a-radio-button :key="method.value" v-for="method in requestMethodOpts" :value="method.value">
              {{ method.label }}
            </a-radio-button>
          </a-radio-group>
        </a-form-item>

        <div class="request-module-method-defined">
          <div v-if="selectedMethodDetail">
            <a-form-item label="OPERATION ID">
              <a-input v-model:value="selectedMethodDetail.name"/>
            </a-form-item>
            <a-form-item label="DESCRIPTION">
              <a-input v-model:value="selectedMethodDetail.desc"/>
            </a-form-item>
            <div class="params-defined">
              <div class="params-defined-btns">
                <a-button type="primary" @click="setSecurity">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `Security` }}
                </a-button>
                <a-button type="primary" @click="addHeader">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `Header` }}
                </a-button>
                <a-button type="primary" @click="addQueryParams">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `Query Params` }}
                </a-button>
                <a-button type="primary" @click="addCookie">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `Cookie` }}
                </a-button>
              </div>
              <div class="params-defined-content">
                <div class="params-defined-item" v-if="selectedMethodDetail?.headers?.length">
                  <div>Header</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.headers" :key="item.id">
                      <FieldItem
                          :fieldData="item"
                          @del="deletePathParams"
                          @paramsNameChange="paramsNameChange"
                          @settingOther="settingOtherForPathParams"
                          @setRef="setRefForPathParams"
                          @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
                <div class="params-defined-item" v-if="selectedMethodDetail?.params?.length">
                  <div>Query Params</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.params" :key="item.id">
                      <FieldItem
                          :fieldData="item"
                          @del="deletePathParams"
                          @paramsNameChange="paramsNameChange"
                          @settingOther="settingOtherForPathParams"
                          @setRef="setRefForPathParams"
                          @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
                <div class="params-defined-item" v-if="selectedMethodDetail?.cookies?.length">
                  <div>Cookie</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.cookies" :key="item.id">
                      <FieldItem :fieldData="item"
                                 @del="deletePathParams"
                                 @paramsNameChange="paramsNameChange"
                                 @settingOther="settingOtherForPathParams"
                                 @setRef="setRefForPathParams"
                                 @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="body-defined">
              <div class="body-defined-btns">
                <a-button type="primary" @click="addBody">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `addBody` }}
                </a-button>
              </div>
              <div class="body-defined-content">
                <div class="params-defined-item" v-if="selectedMethodDetail?.headers?.length">
                  <div>Header</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.headers" :key="item.id">
                      <FieldItem
                          :fieldData="item"
                          @del="deletePathParams"
                          @paramsNameChange="paramsNameChange"
                          @settingOther="settingOtherForPathParams"
                          @setRef="setRefForPathParams"
                          @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
                <div class="params-defined-item" v-if="selectedMethodDetail?.params?.length">
                  <div>Query Params</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.params" :key="item.id">
                      <FieldItem
                          :fieldData="item"
                          @del="deletePathParams"
                          @paramsNameChange="paramsNameChange"
                          @settingOther="settingOtherForPathParams"
                          @setRef="setRefForPathParams"
                          @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
                <div class="params-defined-item" v-if="selectedMethodDetail?.cookies?.length">
                  <div>Cookie</div>
                  <div class="header-defined ">
                    <div v-for="item in selectedMethodDetail.cookies" :key="item.id">
                      <FieldItem :fieldData="item"
                                 @del="deletePathParams"
                                 @paramsNameChange="paramsNameChange"
                                 @settingOther="settingOtherForPathParams"
                                 @setRef="setRefForPathParams"
                                 @setRequire="setPathParamsRequire"/>
                    </div>
                  </div>
                </div>
              </div>
            </div>

          </div>
          <div class="no-defined" v-else>
            <a-button type="primary">
              <template #icon>
                <PlusOutlined/>
              </template>
              {{ `${selectedMethod} Operation` }}
            </a-button>
          </div>
        </div>

      </div>


      <div v-else-if="key === 'response'">response content</div>
      <div v-else-if="key === 'run'">run content</div>
      <div v-else-if="key === 'mock'">mock content</div>

      <template #extra>
        <a href="#">More</a>
      </template>

    </a-card>

    <!-- ::::接口提交按钮 -->
    <div class="drawer-btns">
      <a-space>
        <a-button type="primary">保存</a-button>
        <a-button>取消</a-button>
      </a-space>
    </div>

  </a-drawer>
</template>

<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {defineComponent, reactive, ref, toRaw, UnwrapRef, defineProps, defineEmits, watch} from 'vue';
import {requestMethodOpts} from '@/config/constant';
import {getInterfaceDetail} from '../service';
import {PlusOutlined, EditOutlined} from '@ant-design/icons-vue';
import contenteditable from 'vue-contenteditable';
import FieldItem from './FieldItem.vue'
import {momentUtc} from '@/utils/datetime';

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  interfaceId: {
    required: true,
    // type: an,
  }
})


const emit = defineEmits(['ok', 'close']);

const collapseActiveKey = ref(['1']);
const activeKey = ref('1');
const selectedMethod = ref('GET');
const interfaceNameEditable = ref(false);

function onCloseDrawer() {
  emit('close');
}

const selectedMethodDetail:any = ref(null);

function selectedMethodChange(e) {
  console.log(e.target.value);
  let curInterface = interfaceDetail.value.interfaces.find((item) => {
    return item.method === e.target.value;
  })
  selectedMethodDetail.value = curInterface;
}


function setSecurity() {
  console.log('setSecurity')
}

function addCookie() {
  selectedMethodDetail.value.cookes.push({
    name: '',
    value: '',
    desc: ''
  })
}

function addQueryParams() {
  selectedMethodDetail.value.params.push({
    name: '',
    value: '',
    desc: ''
  })
}

function addHeader() {
  selectedMethodDetail.value.headers.push({
    name: '',
    value: '',
    desc: ''
  })
}

const interfaceDetailNameRef: any = ref(null)

function editInterfaceDetailName() {
  // todo 为什么没有拿到 ref ，回头再看看，先用不好的方式来实现
  // interfaceDetailNameRef.value.focus();
  document.getElementsByClassName('interfaceName')[0].focus()
  // interfaceNameEditable.value = true;
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

  //同步替换删除path中的param参数
  // let path = interfaceDetail.value.path;
  // interfaceDetail.value.path = path.replace(`{${data.name}}`, '');

}

// var a = 'api/user/{id}/{detailID}'
// var parsePathReg = /\{(\w+)(?:=\})/g
// // var parsePathReg = /\{(\w+)\}/g
// var s = a.match(parsePathReg);
// console.log(s)
// console.log(parsePathReg.exec(a));
// console.log(a.match(parsePathReg.$1))

/**
 * 删除路径参数
 * */
function deletePathParams(data) {
  let index = interfaceDetail.value.pathParams.find((item) => {
    return item.id === data.id;
  })
  interfaceDetail.value.pathParams.splice(index, 1);

  //同步替换删除path中的param参数
  let path = interfaceDetail.value.path;
  interfaceDetail.value.path = path.replace(`{${data.name}}`, '');

}


function paramsNameChange(val) {
  // var a = 'api/user/{id}/{detailID}'
  // 解析path 中的参数
  let parsePathReg = /\{(\w+)\}/g
  let path = interfaceDetail.value.path;
  let params = path.match(parsePathReg);
  if (val) {
    params.push(`{${val}}`)
  }
  // todo 需要处理，几个表单项的联动场景
  console.log(832, params, val);
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

function enterPressed() {
  console.log(1212)
}

interface FormState {
  name: string;
  remark: string | undefined;
}


const tabList = [
  {
    key: 'request',
    tab: '请求定义',
    slots: {tab: 'customRenderRequest'},
  },
  {
    key: 'response',
    tab: '响应定义',
    slots: {tab: 'customRenderResponse'},
  },
  {
    key: 'run',
    tab: '运行调试',
    slots: {tab: 'customRenderRun'},

  },
  {
    key: 'mock',
    tab: 'Mock',
    slots: {tab: 'customRenderMock'},
  },
];

const key = ref('request');

const interfaceDetail = ref({
  title: '接口名称'
});

const onTabChange = (value: string, type: string) => {
  console.log(value, type);
  if (type === 'key') {
    key.value = value;
  }
};


const formRef = ref();

const formState: UnwrapRef<FormState> = reactive({
  name: '接口类型1',
  remark: '用户信息相关',
});

const rules = {
  name: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 3, max: 50, message: '最长多少个字符', trigger: 'blur'},
  ],
  path: [{required: true, message: 'Please select Activity zone', trigger: 'change'}],
  tag: [{required: true, message: 'Please select activity resource', trigger: 'change'}],
};


/**
 * 监控接口ID，如果变化，需要重新请求接口详情信息
 * */
watch(() => {
  return props.interfaceId;
}, async (newVal) => {
  if (newVal) {
    let res = await getInterfaceDetail(newVal);
    let data = res.data;
    data.createdAt = momentUtc(data.createdAt);
    data.updatedAt = momentUtc(data.updatedAt);
    console.log(832, res)
    interfaceDetail.value = {...res.data}
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

.card-baseInfo {
  width: 100%;

  :deep(.ant-card-body) {
    padding: 12px 24px;
  }
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

.request-module {
  //margin-bottom: 0;
  //pointer-events: none;
}

.pathParam {
  margin-left: 40px;
  margin-bottom: 32px;
}

.request-module-method-defined {
  .no-defined {
    margin: 100px auto;
    text-align: center;
  }
}

.params-defined-item{
  
}
</style>
