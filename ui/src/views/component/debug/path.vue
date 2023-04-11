<!-- 路径定义方式 -->
<template>
  <a-row class="form-item">
    <a-col :span="2" class="form-label">
      <span class="label-name">路径</span>
    </a-col>

    <a-col :span="17">
      <div class="path-param-header">
        <a-input class="path-param-header-input" :value="endpointDetail.path" @change="updatePath" placeholder="请输入路径">
          <template #addonBefore>
            <a-select
                :options="serveServers"
                v-model:value="currentEnvURL"
                placeholder="请选择环境"
                style="width: 120px;text-align: left">
              <template #notFoundContent>
                <a-button type="link" @click="addEnv" class="add-env-btn">
                  <PlusOutlined/>&nbsp;去新建
                </a-button>
              </template>
            </a-select>
            <span v-if="currentEnvURL" style="width: 150px;display: inline-block"
                  class="currentEnvURL">{{ currentEnvURL || '---' }}</span>
          </template>
        </a-input>
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

const currentEnvURL = ref(serveServers?.value[0]?.value);

/**
 * path 变动，联动 pathParams
 * */
function handlePathLinkParams() {
  // 支持字母下划线及中划线
  let reg = /\{([\w-]+)\}/g
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
  .form-label {
    line-height: 26px;
  }
}

.label-name {
  margin-left: 4px;
}

.path-param-list {
  margin-top: 16px;
  //padding-top: 16px;
}

.path-param-header {
  display: inline-block;
  overflow: hidden;
  width: 100%;
}

.path-param-header-input {
  width: 85%;
}

.path-param-header-btn {
  width: 15%;
}

.form-label {
  font-weight: bold;
}

.add-env-btn {

}

.currentEnvURL {

}

</style>
