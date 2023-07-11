<!-- 请求参数定义：包括 header、query 、 security、cookie  -->
<template>
  <!-- 增加请求参数 -->
  <a-row class="form-item-request-item" style="margin-top: 16px;">
    <a-col :span="3" class="form-label form-label-first">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">增加请求参数</span>
    </a-col>
    <a-col :span="15">
      <div class="params-defined-btns">

        <a-radio-group v-model:value="selectedParamType">
          <a-radio-button
              class="request-query-btn"
              :style="{ color: selectedParamType !== item.value ? '#999999' : '#1890ff',
                      'box-shadow': `none` ,
                      background:  selectedParamType !== item.value ? '#f5f5f5' : '#fff',
                     'border-color': '#d9d9d9'}"
              v-for="item in paramTypeOpts" :key="item.value" :value="item.value">{{item.label}}</a-radio-button>
        </a-radio-group>

        <span style="width: 12px;display: inline-block"></span>

        <a-button @click="setSecurity" v-show="selectedParamType === 'security'">
          <template #icon>
            <PlusOutlined/>
          </template>
          {{ `设置 Security` }}
        </a-button>
        <a-button @click="addHeader"  v-show="selectedParamType === 'header'">
          <template #icon>
            <PlusOutlined/>
          </template>
          {{ `添加 Header` }}
        </a-button>
        <a-button @click="addQueryParams"  v-show="selectedParamType === 'query'">
          <template #icon>
            <PlusOutlined/>
          </template>
          {{ `添加 Params 参数` }}
        </a-button>
        <a-button @click="addCookie"  v-show="selectedParamType === 'cookie'">
          <template #icon>
            <PlusOutlined/>
          </template>
          {{ `添加 Cookie` }}
        </a-button>
      </div>
    </a-col>
  </a-row>
  <!-- 请求参数展示：headers、cookies、query params等 -->
  <a-row class="form-item-request-item-con" v-if="collapse">
    <a-col :span="3"></a-col>
    <a-col :span="21">
      <div class="params-defined">
        <div class="params-defined-content">
          <div class="params-defined-item" v-if="showSecurity && selectedParamType === 'security'">
            <div class="params-defined-item-header" style="margin-top: 16px;margin-bottom: 8px;"/>
            <div class="header-defined header-defined-items">
              <a-select @change="securityChange"
                        allowClear
                        placeholder="请选择 Security"
                        @focus="getSecurityOpts"
                        :value="selectedMethodDetail.security || null"
                        :options="securityOpts" style="width: 300px;"/>
              <a-tooltip placement="topLeft" arrow-point-at-center title="去添加或编辑 Security">
                <a-button @click="goEditSecurity">
                  <template #icon>
                    <FormOutlined/>
                  </template>
                  Security
                </a-button>
              </a-tooltip>
              <a-tooltip placement="topLeft" arrow-point-at-center title="删除 Security">
                <a-button @click="delSecurity">
                  <template #icon>
                    <DeleteOutlined/>
                  </template>
                </a-button>
              </a-tooltip>
            </div>
          </div>
          <div class="params-defined-item" v-if="selectedMethodDetail?.headers?.length &&  selectedParamType === 'header'">
            <div class="header-defined header-defined-items">
              <div v-for="(item,index) in selectedMethodDetail.headers" :key="item.id">
                <RequestHeader
                    :fieldData="{...item,index:index}"
                    :showRequire="true"
                    @del="deleteParams('headers',index)"
                    @change="(val) => {handleParamsChange('headers',val);}"/>
              </div>
            </div>
          </div>
          <div class="params-defined-item" v-if="selectedMethodDetail?.params?.length  &&  selectedParamType === 'query'">
            <div class="header-defined ">
              <div v-for="(item,index) in selectedMethodDetail.params" :key="item.id">
                <Field
                    :fieldData="{...item,index:index}"
                    :showRequire="true"
                    @del="deleteParams('params',index)"
                    @change="(val) => {handleParamsChange('params',val);}"/>
              </div>
            </div>
          </div>
          <div class="params-defined-item" v-if="selectedMethodDetail?.cookies?.length &&  selectedParamType === 'cookie'">
            <div class="header-defined ">
              <div v-for="(item,index) in selectedMethodDetail.cookies" :key="item.id">
                <Field
                    :fieldData="{...item,index:index}"
                    :showRequire="true"
                    @del="deleteParams('cookies',index)"
                    @change="(val) => {handleParamsChange('cookies',val);}"/>
              </div>
            </div>
          </div>
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
  defaultCookieParams,
  defaultHeaderParams,
  defaultQueryParams,
} from '@/config/constant';
import {PlusOutlined, DeleteOutlined, RightOutlined, DownOutlined, FormOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import RequestHeader from './RequestHeader.vue'
import {Endpoint} from "@/views/endpoint/data";
import {cloneByJSON} from "@/utils/object";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const selectedMethodDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const securityOpts: any = computed<any>(() => store.state.Endpoint.securityOpts);
const props = defineProps({});
const emit = defineEmits([]);
// 是否折叠,默认展开
const collapse = ref(true);
// 是否展示安全定义
const showSecurity = ref(!!selectedMethodDetail.value?.security);

const paramTypeOpts = [
  {
    label: 'Security',
    value: 'security'
  },
  {
    label: 'Header',
    value: 'header'
  },
  {
    label: 'Query Params',
    value: 'query'
  },
  {
    label: 'Cookie',
    value: 'cookie'
  },
]

// 当前选中的参数类型
const selectedParamType = ref('security');

function goEditSecurity() {
  window.open(`/#/project-setting/service-setting?sectab=service-security&serveId=${endpointDetail.value.serveId}`, '_blank')
}

function delSecurity() {
  showSecurity.value = false;
  selectedMethodDetail.value.security = null;
}

function securityChange(val) {
  selectedMethodDetail.value.security = val || null;
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function setSecurity() {
  collapse.value = true;
  showSecurity.value = true;
}

async function getSecurityOpts() {
    await store.dispatch('Endpoint/getSecurityList', {
      id: endpointDetail.value.serveId
    });
}

function addCookie() {
  collapse.value = true;
  selectedMethodDetail.value.cookies.push(cloneByJSON({
    ...defaultCookieParams,
    name: 'cookie' + (selectedMethodDetail.value.cookies.length + 1)
  }));
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function addQueryParams() {
  collapse.value = true;
  selectedMethodDetail.value.params.push(cloneByJSON({
    ...defaultQueryParams,
    name: 'param' + (selectedMethodDetail.value.params.length + 1)
  }));
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function addHeader() {
  collapse.value = true;
  selectedMethodDetail.value.headers.push(cloneByJSON({
    ...defaultHeaderParams,
    name: 'header' + (selectedMethodDetail.value.headers.length + 1)
  }));
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function deleteParams(type, index) {
  selectedMethodDetail.value[type].splice(index, 1);
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function handleParamsChange(type, data) {
  selectedMethodDetail.value[type][data.index] = {...data};
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

</script>
<style lang="less" scoped>
.form-label {
  font-weight: bold;
  position: relative;
  left: -18px;
}

.form-label-first {
  font-weight: bold;
  position: relative;
  left: -18px;
}

.params-defined-item-header {
  font-weight: bold;
  //margin-bottom: 8px;
  margin-top: 16px;
}
.form-item-request-item-con{
  position: relative;
  &:before{
    content:"";
    position: absolute;
    left: -12px;
    top: 0px;
    width: 2px;
    background: #E5E5E5;
    height: 100%;
  }
}

.label-name {
  display: inline-block;
  margin-left: 4px;
  margin-top: 4px;
}

.ant-radio-button-wrapper-checked.request-query-btn{
  &:before{
    display: none;
  }
}
</style>
