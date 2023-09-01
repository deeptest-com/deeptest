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
      <a-radio-group v-model:value="selectedMethod" button-style="outline">
        <a-radio-button
            v-for="method in requestMethodOpts"
            :class="{'has-defined': hasDefinedMethod(method.value),'request-method-btn':true}"
            :style="{ color: hasDefinedMethod(method.value) ? method.color : '#999999',
                      'box-shadow': `none` ,
                      background: method.value !== selectedMethod ? '#f5f5f5' : '#fff',
                     'border-color': '#d9d9d9'}"
            :size="'small'"
            :key="method.value" :value="method.value">
          {{ method.label }}
        </a-radio-button>
      </a-radio-group>

      <div class="form-item-request" v-if="collapse">
        <div v-if="selectedMethodDetail">
          <!-- Operation ID -->
<!--          <a-row class="form-item-request-item">-->
<!--            <a-col :span="3" class="form-label">-->
<!--              Operation ID-->
<!--            </a-col>-->
<!--            <a-col :span="12">-->
<!--              <a-input placeholder="Operation ID" v-model:value="selectedMethodDetail.operationId"/>-->
<!--            </a-col>-->
<!--          </a-row>-->
          <!-- Description -->
          <a-row class="form-item-request-item">
            <a-col :span="3" class="form-label">
              描述信息
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
<!--            {{ `${selectedMethod} Operation` }}-->
            {{ `定义 ${selectedMethod} 方法` }}
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
  onMounted, onUnmounted,
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

const props = defineProps({});
const emit = defineEmits([]);

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const interfaceDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const interfaceMethodToObjMap = computed<any>(() => store.state.Endpoint.interfaceMethodToObjMap);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);

const selectedMethod = ref('')

onUnmounted(async () => {
  await store.dispatch('Endpoint/removeUnSavedMethods')
})

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
  console.log('selectedMethod', selectedMethod.value)
  if (!newVal) {
    if (endpointDetail.value?.interfaces?.length) {
      selectedMethod.value = endpointDetail.value?.interfaces[0].method;
    } else {
      selectedMethod.value = 'GET';
    }
    return;
  }
  selectedMethodDetail.value = interfaceMethodToObjMap.value[newVal];
  if (selectedMethodDetail.value) {
    store.commit('Endpoint/setSelectedMethodDetail', selectedMethodDetail.value);
  } else {
    store.commit('Endpoint/setSelectedMethodDetail', {});
  }
  // 根据选中的请求方法决定是否展示请求体设置，暂定以下三种方法是不需要请求体的
  showRequestBody.value = ['POST', 'PUT', 'PATCH'].includes(newVal);
}, {immediate: true, deep: true});

function addEndpoint() {
  const item = {
    ...cloneByJSON(defaultEndpointDetail),
    "projectId": endpointDetail.value.projectId,
    "serveId": endpointDetail.value.serveId,
    "useId": currentUser.value.id,
    "method": selectedMethod.value,
  }
  selectedMethodDetail.value = item;
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
    display: none;
    content:"";
    position: absolute;
    left: -72px;
    top: -12px;
    width: 2px;
    background: #E5E5E5;
    //min-height: 80vh;
    height: calc(100% + 24px);
  }
  .form-item-request-item {
    margin-top: 16px;
    align-items: center;
  }
}

.has-defined {
  color: #1890ff;
}
.ant-radio-button-wrapper-checked.request-method-btn{
  &:before{
    display: none;
  }
}
.label-name {
  margin-left: 4px;
}
</style>
