<!-- 响应定义 -->
<template>
  <!-- 响应定义  -->
  <a-row class="form-item-response">
    <a-col :span="3" class="form-label form-label-first">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">选择响应代码</span>
    </a-col>
    <a-col :span="21">
      <div class="select-code-con">
        <a-radio-group class="select-code-group" v-model:value="selectedCode" button-style="outline">
          <a-radio-button
              :class="{'has-defined': hasDefinedCode(code.value),'res-code-btn':true}"
              :style="{ color: hasDefinedCode(code.value) ? code.color : '',
                      'box-shadow': `none` ,
                      background: code.value !== selectedCode ? '#f5f5f5' : '#fff',
                     'border-color': '#d9d9d9'}"
              :key="code.value" v-for="code in responseCodes.filter(item => selectedCodes.includes(item.value))"
              :value="code.value">
            <span class="res-code-btn-content">
              <span class="text">{{ code.label }}</span>
              </span>
          </a-radio-button>
        </a-radio-group>
        <a-dropdown class="select-code-options"
                    :overlayClassName="'deeptest-res-code-dropdown'">
          <template #overlay>
            <a-menu class="menu-con">
              <a-menu-item v-for="opt in responseCodes"
                           @click="() => {handleCodeOptionsChange(opt.value) }"
                           :disabled="selectedCodes.includes(opt.value)"
                           :key="opt.value">
                <span style="font-size: 13px;">{{ opt.description }}</span>
              </a-menu-item>
            </a-menu>
          </template>
          <a-button>
            <PlusOutlined/>
            添加响应码
          </a-button>
        </a-dropdown>
        <span class="delete-icon">
                  <a-popconfirm title="确定删除当前选中状态码吗？"
                                @confirm="confirmDeleteCode"
                                @cancel="cancelDeleteCode">
                    <template #icon><question-circle-outlined style="color: red"/></template>
                     <a-button v-if="selectedCode && selectedCode !== '200'"><DeleteOutlined/>删除 </a-button>
                  </a-popconfirm>
              </span>
      </div>
      <div :class="{'form-item-response-body':true,'has-data':selectedCodeDetail}"  v-if="collapse">
        <div v-if="selectedCodeDetail">
          <!-- Description -->
          <a-row class="form-item-response-item">
            <a-col :span="3" class="form-label">
              描述
            </a-col>
            <a-col :span="21">
              <a-input @change="handleResDescriptionChange" :placeholder="'请输入描述信息'"
                       :value="selectedCodeDetail.description"/>
            </a-col>
          </a-row>
          <!-- 增加响应头 -->
          <a-row class="form-item-response-item">
            <a-col :span="3" class="form-label form-label-header">
              <RightOutlined v-if="!resHeaderCollapse" @click="resHeaderCollapse = !resHeaderCollapse"/>
              <DownOutlined v-if="resHeaderCollapse" @click="resHeaderCollapse = !resHeaderCollapse"/>
              <span class="label-name">增加响应头</span>
            </a-col>
            <a-col :span="21">
              <div class="params-defined-btns">
                <a-button @click="addResponseHeader">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `添加` }}
                </a-button>
              </div>
            </a-col>
          </a-row>
          <!-- 响应头展示-->
          <a-row class="form-item-response-item form-item-response-item-header" style="margin-top: 0px;"
                 v-if="resHeaderCollapse">
            <a-col :span="3"></a-col>
            <a-col :span="21">
              <div class="params-defined">
                <div class="params-defined-content">
                  <div class="params-defined-item" v-if="selectedCodeDetail?.headers?.length">
                    <div class="header-defined header-defined-items">
                      <div v-for="(item,index) in selectedCodeDetail.headers" :key="item.id">
                        <Field
                            :fieldData="{...item,index:index}"
                            :showRequire="false"
                            @del="deleteResHeader(index)"
                            @change="(val) => { handleResHeaderChange(val);}"/>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </a-col>
          </a-row>
          <!-- 响应体 -->
          <ResponseBody/>
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
  responseCodes,
  defaultHeaderParams,
  defaultCodeResponse,
  defaultResponseCodes,
} from '@/config/constant';
import {PlusOutlined, DownOutlined, RightOutlined, DeleteOutlined, QuestionCircleOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import ResponseBody from './ResponseBody.vue'
import {Endpoint} from "@/views/endpoint/data";
import {cloneByJSON} from "@/utils/object";
import cloneDeep from "lodash/cloneDeep";
import {message} from "ant-design-vue";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const selectedMethodDetail: any = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const selectedCodeDetail = computed<any>(() => store.state.Endpoint.selectedCodeDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
// 是否折叠,默认展开
const collapse = ref(true);
const resHeaderCollapse = ref(true);
const props = defineProps({});
const emit = defineEmits([]);
const selectedCode = ref('200');

// 是否定义了请求方法的响应体
function hasDefinedCode(code: string) {
  return selectedMethodDetail?.value?.responseBodies?.some((item) => {
    return item.code === code;
  })
}

const selectedCodes: any = computed(() => {
  const codes = (selectedMethodDetail?.value?.responseCodes.split(',') || []).filter((item) => {
    return !!item;
  });
  // 如果没有定义响应码,则默认返回默认的 codes 枚举
  if (!codes.length) {
    return cloneDeep(defaultResponseCodes);
  }
  return codes;
})

watch(() => {
  return selectedCode.value
}, (newVal, oldVal) => {
  if (newVal !== oldVal) {
    collapse.value = true;
  }
  const detail = selectedMethodDetail?.value?.responseBodies?.find((item) => {
    return item.code === newVal;
  })
  store.commit('Endpoint/setSelectedCodeDetail', detail);
}, {immediate: true});


// 添加可选的响应码
function handleCodeOptionsChange(code: string) {
  if (selectedCodes.value.includes(code)) {
    return;
  }
  selectedCodes.value.push(code);
  selectedMethodDetail.value.responseCodes = selectedCodes.value.toString();
  store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
}

function confirmDeleteCode() {
  if (!selectedCodes.value.includes(selectedCode.value)) {
    return;
  }
  if (selectedCode.value === '200') {
    message.warning('200状态码不可删除');
    return;
  }
  const index = selectedCodes.value.findIndex((item) => {
    return item === selectedCode.value;
  })
  selectedCodes.value.splice(index, 1);
  selectedCode.value = selectedCodes.value?.[index - 1] || '200';
  selectedMethodDetail.value.responseCodes = [...selectedCodes.value].toString();
  store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
}

function cancelDeleteCode(code: string) {
  console.log('cancelDeleteCode', code);
}

function addResponseHeader() {
  resHeaderCollapse.value = true;
  selectedCodeDetail.value.headers.push(cloneByJSON({
    ...defaultHeaderParams,
    name: "header" + (selectedCodeDetail.value.headers.length + 1)
  }));
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail.value);
}

function addCodeResponse() {
  const item = {
    ...cloneByJSON(defaultCodeResponse),
    "code": selectedCode.value,
    "endpointId": selectedMethodDetail.value.id,
  }
  store.commit('Endpoint/setSelectedCodeDetail', item);
}

function handleResDescriptionChange(e) {
  selectedCodeDetail.value.description = e.target.value;
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail.value);
}

function deleteResHeader(index) {
  selectedCodeDetail.value.headers.splice(index, 1);
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail.value);
}

function handleResHeaderChange(data) {
  selectedCodeDetail.value.headers[data.index] = {...data};
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail.value);
}

</script>
<style lang="less" scoped>

.form-label {
  font-weight: bold;
}

.form-label-first, .form-label-header {
  font-weight: bold;
  position: relative;
  left: -18px;
}


.form-item-response {
  margin-top: 16px;

  .form-item-response-item {
    margin-top: 16px;
    align-items: center;
  }
}

.form-item-response-body {
  margin-top: 8px;
  position: relative;
  &.has-data{
    &:before {
      content: "";
      position: absolute;
      left: -121px;
      top: -12px;
      width: 2px;
      background: #E5E5E5;
      min-height: 80vh;
      height: calc(100% + 36px);
    }
  }

}

.form-item-response-item-header {
  position: relative;

  &:before {
    content: "";
    position: absolute;
    left: -12px;
    top: 0;
    width: 2px;
    background: #E5E5E5;
    height: 100%;
  }
}

.has-defined {
  color: #1890ff;
  //font-weight: bold;
}

.res-code-btn-content {
  //display: flex;
  //width: 40px;
  //justify-content: center;
  //align-items: center;
}

.delete-icon {
  margin-left: 8px;
  display: block;
  color: #ff4d4f;
}

.ant-radio-button-wrapper-checked.res-code-btn {

  &:before {
    display: none;
  }
}

.label-name {
  margin-left: 4px;
}

.select-code-con {
  display: flex;
  width: 750px;
  overflow-x: scroll;

  .select-code-group {
    display: flex;
    height: 37px;
    max-width: 700px;
    overflow-x: scroll;
  }

  .select-code-options {
    width: 120px;
    margin-left: 8px;
  }
}
</style>
