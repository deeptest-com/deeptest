<template>
  <div class="content">
    <div class="left-content">
      <div class="global">
        <div class="header">全局</div>
        <a-button :class="{'env-item':true,'env-item-active':isShowGlobalVars}"
                  @click="showGlobalVars"
                  :type="isShowGlobalVars ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">V</i>
          </template>
          全局变量
        </a-button>

        <a-button :class="{'env-item':true,'env-item-active':isShowGlobalParams}"
                  @click="showGlobalParams"
                  :type="isShowGlobalParams ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">P</i>
          </template>
          全局参数
        </a-button>
      </div>

      <div style="margin: 0 16px;">
        <a-divider class="divider"/>
      </div>
      <div class="env">
        <div class="header">环境</div>
        <draggable
            tag="div"
            :list="envList"
            class="list-group"
            handle=".handle"
            item-key="name">
          <template #item="{ element, index }">
            <a-button
                :class="{'env-item':true,'env-item-active':activeEnvDetail?.id === element.id}"
                :type="activeEnvDetail?.id === element.id ? 'primary' : 'text'"
                @click="showEnvDetail(element)"
                class="env-item" :key="index">
              <MenuOutlined class="handle"/>
              <span class="text"> {{ element.displayName }} </span>
            </a-button>
          </template>
        </draggable>

        <div style="margin: 0 16px;">
          <a-divider class="divider"/>
        </div>
        <a-button
            :type="isShowAddEnv ? 'primary' : 'text'"
            :class="{
            'env-item':true,
            'env-item-footer':true,
            'env-item-active':isShowAddEnv}"
            @click="showEnvDetail(null,true)">
          <template #icon>
            <PlusOutlined/>
          </template>
          新建环境
        </a-button>
      </div>
    </div>
    <div class="right-content">
      <!-- ::::全局变量 -->
      <div class="globalVars" v-if="isShowGlobalVars">
        <GlobalVarCom
          @handle-global-vars-change="handleGlobalVarsChange"
          @handle-save-global-vars="handleSaveGlobalVars"
          @add-global-var="addGlobalVar"
        />
      </div>
      <!-- ::::全局参数 -->
      <div class="globalParams" v-if="isShowGlobalParams">
        <GlobalParamsCom
          @handle-global-params-change="handleGlobalParamsChange"
          @handle-save-global-params="handleSaveGlobalParams"
          @add-global-params="addGlobalParams"
        />
      </div>

      <div class="envDetail" v-if="isShowEnvDetail && activeEnvDetail">
        <EnvDetail
          :activeEnvDetail="activeEnvDetail"
          @deleteEnvData="deleteEnvData"
          @copyEnvData="copyEnvData"
          @addEnvData="addEnvData"
          @handleEnvChange="handleEnvChange"
          @handleEnvNameChange="handleEnvNameChange"
          @addVar="addVar"
          @addService="addService"
        />
      </div>

    </div>
  </div>

<!--  <a-modal
      v-model:visible="addServiceModalVisible"
      title="关联服务"
      @ok="handleAddServiceOk">
    <a-form-item class="select-service" :labelCol="{span: 6}" :wrapperCol="{span: 16}" label="请选择服务">
      <a-select
          v-model:value="selectedService"
          :options="serviceOptions"
          placeholder="请选择服务"
          style="width: 200px"/>
    </a-form-item>

  </a-modal>-->
</template>
<script setup lang="ts">

import {
  computed,
  ref,
  watch
} from 'vue';
import {MenuOutlined, PlusOutlined} from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import EnvDetail from './EnvDetail.vue';
import GlobalParamsCom from './GlobalParams.vue';
import GlobalVarCom from './GlobalVar.vue';
import { useGlobalEnv } from '../../hooks/globalEnv';
import { useGlobalVarAndParams } from '../../hooks/globalVar';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from "@/views/project-setting/store";
import {useStore} from "vuex";

// store 相关
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting:  ProjectSettingStateType}>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const serviceOptions = computed<any>(() => store.state.ProjectSetting.serviceOptions);
const envList = computed<any>(() => store.state.ProjectSetting.envList);

// 页面state相关
const isShowGlobalVars = ref(false);
const isShowGlobalParams: any = ref(false);
const addServiceModalVisible = ref(false);
const selectedService = ref('');

const globalParamsActiveKey = ref('header');


// 环境设置相关
const {
  isShowAddEnv,
  isShowEnvDetail,
  activeEnvDetail,
  getEnvsList,
  showEnvDetail,
  addVar,
  addEnvData,
  deleteEnvData,
  copyEnvData,
  handleEnvChange,
  handleEnvNameChange
} = useGlobalEnv({ isShowGlobalParams, isShowGlobalVars});

// 全局变量和全局参数相关
const {
  showGlobalParams,
  showGlobalVars,
  addGlobalVar,
  addGlobalParams,
  handleSaveGlobalParams,
  handleSaveGlobalVars,
  handleGlobalVarsChange,
  handleGlobalParamsChange
} = useGlobalVarAndParams({
  isShowAddEnv,
  isShowEnvDetail,
  activeEnvDetail,
  isShowGlobalParams,
  isShowGlobalVars,
  globalParamsActiveKey
});

/**
 * 请求服务列表
 */
async function getServersList() {
  await store.dispatch('ProjectSetting/getServersList', {
    projectId: currProject.value.id,
    page: 0,
    pageSize: 100
  })
}

// 添加服务弹窗操作
async function addService() {
  addServiceModalVisible.value = true;
}

/**
 * modal弹窗确认选择服务后操作
 */
function handleAddServiceOk() {
  addServiceModalVisible.value = false;
  const selectServe: any = serviceOptions.value.find((item: any) => {
    return selectedService.value === item.id;
  })
  activeEnvDetail.value.serveServers.push({
    // "environmentId": 7,
    "url": "",
    "serveName": selectServe.name,
    "serveId": selectServe.id,
  })
}

// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await getServersList();
  await getEnvsList();
  // 默认展示全局变量
  await showGlobalVars();
}, {
  immediate: true
})


</script>
<style scoped lang="less">
.content {
  display: flex;
}

.left-content {
  height: calc(100vh - 138px);
  overflow-y: scroll;
  background-color: #F9F9F9;
  width: 240px;

  div {
    list-style: none;
    margin: 0;
    padding: 0;

    li {
      padding: 0;
      margin: 0;
      height: 36px;
      line-height: 36px;
      list-style: none;
      cursor: pointer;
    }

    .header {
      margin-left: 16px;
      margin-top: 16px;
      margin-bottom: 16px;
      opacity: 0.5;
    }
  }

  button.env-item {
    width: 200px;
    text-align: left;
    margin-bottom: 8px;
    border-radius: 6px;
  }

  .env-item {
    margin: 0 16px;
    padding-left: 8px;

    i {
      width: 18px;
      height: 18px;
      background-color: #515152;
      border-radius: 100%;
      color: #ffffff;
      font-size: 10px;
      text-align: center;
      line-height: 18px;
      margin-right: 8px;
    }


    .handle {
      margin-right: 8px;
      cursor: pointer;
    }
  }

  .divider {
    margin: 16px 0;
  }
}

.right-content {
  flex: 1;
  height: calc(100vh - 138px);
  overflow-y: scroll;
  position: relative;
  padding: 16px;

  .globalVars, .globalParams, .envDetail {
    padding: 8px;
  }
}

.var-icon {
  display: inline-block;
}

::v-deep {

  .custom-select {
    .ant-select-selector {
      border: 1px solid transparent !important;
    }

    .ant-select-selector:hover, .ant-select-selector:active, .ant-select-selector:focus {
      border: 1px solid #4096ff !important
    }
  }


  .ant-input:not(.env-name) {
    border: 1px solid transparent !important;
  }

  .ant-input:not(.env-name):hover, .ant-input:active, .ant-input:focus {
    border: 1px solid #4096ff !important
  }
}

</style>
