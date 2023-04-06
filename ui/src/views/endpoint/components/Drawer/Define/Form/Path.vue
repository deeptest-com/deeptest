<!-- 路径定义方式 -->
<template>
  <a-row class="form-item">
    <a-col :span="2" class="form-label">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">路径</span>
    </a-col>
    <a-col :span="16">
      <a-input :value="endpointDetail.path" @change="updatePath" placeholder="请输入路径">
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
      <!-- 路径参数 -->
      <div class="path-param-list" v-if="collapse && endpointDetail?.pathParams?.length > 0">
        <Field
            v-for="(item,index) in endpointDetail.pathParams"
            :key="item.id +index"
            :fieldData="{...item,index:index}"
            :showRequire="true"
            @del="deletePathParams(index)"
            @change="handleChange"/>
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
  defaultPathParams,
} from '@/config/constant';
import {PlusOutlined, DeleteOutlined, RightOutlined, DownOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import {Endpoint} from "@/views/endpoint/data";
import {cloneByJSON} from "@/utils/object";

const props = defineProps({});
const emit = defineEmits([]);

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
const serveServers: any = computed<Endpoint>(() => store.state.Endpoint.serveServers);

// 是否折叠,默认展开
const collapse = ref(true);

/**
 * 添加路径参数
 * */
function addPathParams() {
  collapse.value = true;
  endpointDetail.value.pathParams.push(cloneByJSON(defaultPathParams));
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
  handleParamsLinkPath();
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
  handleParamsLinkPath();
}

/**
 * 更新路径参数
 * */
function handleChange(data) {
  endpointDetail.value.pathParams[data.index] = data;
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
  handleParamsLinkPath();
}

/**
 * path 变动，联动 pathParams
 * */
function handlePathLinkParams() {
  let reg = /\{(\w+)\}/g
  let path = endpointDetail.value.path;
  let pathParams = endpointDetail.value?.pathParams || [];
  const params: any = [];
  let param: any | Array<any> = reg.exec(path);
  while (param !== null) {
    params.push(param[1]);
    param = reg.exec(path)
  }
  if (params.length < pathParams.length) {
    pathParams.splice(params.length - 1);
  }
  params.forEach((item, index) => {
    if (pathParams[index]) {
      pathParams[index].name = item;
    } else {
      pathParams.push({
        ...cloneByJSON(defaultPathParams),
        name: item,
      })
    }
  })
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    pathParams: endpointDetail.value.pathParams
  })
}

/**
 * pathParams 变动，联动 Path
 * params = [a1,a2]
 * {a}/xxx{b} ===> {a1}/xxx{a2}
 * */
function handleParamsLinkPath() {
  let path = endpointDetail.value.path;
  let pathParams = endpointDetail.value.pathParams || [];
  let params = pathParams.map(item => item.name);
  let paths = path.split(/(\{\w*\})/g);

  let idx = 0;
  paths.forEach((item, index) => {
    if (item && item.startsWith('{') && item.endsWith('}')) {
      paths[index] = params[idx] ? `{${params[idx]}}` : '';
      idx++;
    }
  })
  if (params.length > idx) {
    params.slice(idx).forEach((item) => {
      paths.push(item ? `/{${item}}` : '')
    })
  }
  let newPath = paths.filter((item) => {
    return !!item
  }).join('').replace('//', '/');

  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    path: newPath,
  });
}

/**
 * 更新路径
 * */
function updatePath(e) {
  const path = e.target.value;
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
    path: path,
  });
  handlePathLinkParams();
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

.label-name {
  margin-left: 4px;
}

.path-param-list {
  margin-top: 16px;
  //padding-top: 16px;
}

.form-label {
  font-weight: bold;
}

</style>
