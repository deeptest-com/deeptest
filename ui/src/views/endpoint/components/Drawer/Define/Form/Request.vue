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
              <a-input placeholder="Operation ID" v-model:value="selectedMethodDetail.operationId"/>
            </a-col>
          </a-row>
          <!-- Description -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label">
              Description
            </a-col>
            <a-col :span="12">
              <a-input placeholder="描述信息" v-model:value="selectedMethodDetail.description"/>
            </a-col>
          </a-row>
          <RequestParams/>
          <!-- 增加请求体 -->
          <RequestBody v-if="showRequestBody"/>
          <!-- 响应定义  -->
          <Response/>
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
  defaultEndpointDetail,
} from '@/config/constant';
import {PlusOutlined, RightOutlined, DownOutlined} from '@ant-design/icons-vue';
import Response from './Response.vue';
import RequestParams from './RequestParams.vue';
import RequestBody from './RequestBody.vue';
import {Endpoint} from "@/views/endpoint/data";
import {cloneByJSON} from "@/utils/object";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);
const currInterface = computed<any>(() => store.state.Debug?.currInterface);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
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
// 是否展示请求体设置，比如 get 请求是不需要请求体的
const showRequestBody = ref(false);
watch(() => {
  return selectedMethod.value
}, (newVal, oldVal) => {
  selectedMethodDetail.value = interfaceMethodToObjMap.value[newVal];
  if (selectedMethodDetail.value) {
    store.dispatch('Debug/setDefineInterface', selectedMethodDetail.value);
    store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
  } else {
    store.dispatch('Debug/setDefineInterface', {});
    store.commit('Endpoint/setSelectedMethodDetail', {});
  }
  // 根据选中的请求方法决定是否展示请求体设置，暂定以下三种方法是不需要请求体的
  showRequestBody.value = ['POST', 'PUT', 'PATCH'].includes(newVal);
}, {immediate: true});

function addEndpoint() {
  const item = {
    ...cloneByJSON(defaultEndpointDetail),
    "projectId": endpointDetail.value.projectId,
    "serveId": endpointDetail.value.serveId,
    "useId": currentUser.value.id,
    "method": selectedMethod.value,
  }
  selectedMethodDetail.value = item;
  store.dispatch('Debug/setDefineInterface', selectedMethodDetail.value);
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

</script>
<style lang="less" scoped>

.form-item {
  //margin-bottom: 16px;
  align-items: baseline;

}

.form-label {
  font-weight: bold;
}

.form-item-request {
  margin-top: 16px;
  position: relative;
  &:before{
    content:"";
    position: absolute;
    left: -72px;
    top: -12px;
    width: 2px;
    background: #E5E5E5;
    height: calc(100% + 120px);
  }
  .form-item-request-item {
    margin-top: 16px;
    align-items: center;
  }
}

.has-defined {
  color: #1890ff;
}

.label-name {
  margin-left: 4px;
}
</style>
