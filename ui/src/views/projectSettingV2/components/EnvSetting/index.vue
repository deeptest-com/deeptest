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
        <div class="title">全局变量</div>
        <a-button
            class="envDetail-btn"
            @click="addGlobalVar"
        >
          <template #icon>
            <PlusOutlined/>
          </template>
          添加
        </a-button>

        <a-table bordered size="small" :pagination="false"
          :columns="globalVarsColumns"
          :data-source="globalVarsData">
          <template #customName="{ text,index }">
            <a-input @change="(e) => {
                handleGlobalVarsChange('name',index,e);
                }" :value="text" placeholder="请输入参数名"/>
          </template>
          <template #customLocalValue="{ text,index }">
            <a-input :value="text" @change="(e) => {
              handleGlobalVarsChange('localValue',index,e);
              }" placeholder="请输入本地值"/>
          </template>
          <template #customRemoteValue="{ text,index }">
            <a-input :value="text" @change="(e) => {
            handleGlobalVarsChange('remoteValue',index,e);
            }" placeholder="请输入远程值"/>
          </template>
          <template #customDescription="{ text,index }">
            <a-input :value="text" @change="(e) => {
              handleGlobalVarsChange('description',index,e);
            }" placeholder="请输入描述信息"/>
          </template>
          <template #customAction="{index }">
            <a-button danger
                      type="text"
                      @click="handleGlobalVarsChange('description',index,'','delete');"
                      :size="'small'">删除
            </a-button>
          </template>
        </a-table>

        <div class="envDetail-footer">
          <a-button class="save-btn" @click="handleSaveGlobalVars" type="primary">保存</a-button>
        </div>
      </div>
      <!-- ::::全局参数 -->
      <div class="globalParams" v-if="isShowGlobalParams">
        <div class="title">全局参数</div>

        <a-tabs :pagination="false" v-model:activeKey="globalParamsActiveKey">
          <a-tab-pane v-for="(tabItem) in tabPaneList" :key="tabItem.type" :tab="tabItem.name">

            <a-button
                class="envDetail-btn"
                @click="addGlobalParams"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>

            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                     :data-source="globalParamsData?.[tabItem.name] || []">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange(tabItem.name,'name',index,e);
              }" placeholder="请输入参数名"/>
              </template>
              <template #customType="{ text,index }">
                <a-select
                    class="custom-select"
                    :value="text"
                    style="width: 120px"
                    @change="(e) => {
                    handleGlobalParamsChange(tabItem.name,'type',index,e)
                    }">
                  <a-select-option value="string">string</a-select-option>
                  <a-select-option value="number">number</a-select-option>
                  <a-select-option value="integer">integer</a-select-option>
                </a-select>
              </template>
              <template #customRequired="{ text,index }">
                <a-switch :checked="text" @change="(checked) => {
                    handleGlobalParamsChange(tabItem.name,'required',index,checked)
                    }"/>
              </template>
              <template #customDefaultValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange(tabItem.name,'defaultValue',index,e);
                }" placeholder="默认值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange(tabItem.name,'description',index,e);
              }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleGlobalParamsChange(tabItem.name,'',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </a-tab-pane>
        </a-tabs>

        <div class="envDetail-footer">
          <a-button class="save-btn" @click="handleSaveGlobalParams" type="primary">保存</a-button>
        </div>
      </div>
      <!-- ::::环境详情 -->
      <div class="envDetail" v-if="isShowEnvDetail && activeEnvDetail">
        <div class="title">{{ activeEnvDetail.displayName }}</div>
        <div class="envDetail-content">
          <a-form-item :labelCol="{span: 2}" :wrapperCol="{span: 10}" label="环境名称">
            <a-input class="env-name" :value="activeEnvDetail.name || ''" @change="handleEnvNameChange"
                     placeholder="请输入环境名称"/>
          </a-form-item>
          <div class="serveServers">
            <div class="serveServers-header">服务 (前置URL)</div>
            <a-button
                class="envDetail-btn"
                @click="addService"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              关联服务
            </a-button>
            <a-table v-if="activeEnvDetail.serveServers.length > 0"
                     size="small"
                     bordered :pagination="false"
                     :columns="serveServersColumns"
                     :data-source="activeEnvDetail.serveServers">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                handleEnvChange('serveServers','serveName',index,e);
                }" placeholder="请输入参数名"/>
              </template>

              <template #customUrl="{ text,index }">
                <a-input :value="text"
                         @change="(e) => {
                            handleEnvChange('serveServers','url',index,e);
                         }"
                         placeholder="http 或 https 起始的合法 URL"/>
              </template>
            </a-table>
          </div>
          <div class="vars">
            <div class="vars-header">环境变量</div>
            <a-button
                class="envDetail-btn"
                @click="addVar"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>
            <a-table
                v-if="activeEnvDetail.vars.length > 0"
                bordered size="small"
                :pagination="false"
                :columns="globalVarsColumns"
                :data-source="activeEnvDetail.vars">
              <template #customName="{ text,index }">
                <a-input @change="(e) => {
                handleEnvChange('vars','name',index,e);
                }" :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customLocalValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
              handleEnvChange('vars','localValue',index,e);
              }" placeholder="请输入本地值"/>
              </template>
              <template #customRemoteValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
            handleEnvChange('vars','remoteValue',index,e);
            }" placeholder="请输入远程值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
              handleEnvChange('vars','description',index,e);
            }" placeholder="请输入描述信息"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleEnvChange('vars','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </div>
        </div>
        <div class="envDetail-footer">
          <a-button v-if="activeEnvDetail.id" class="save-btn" @click="deleteEnvData" type="danger">删除</a-button>
          <a-button v-if="activeEnvDetail.id" class="save-btn" @click="copyEnvData" type="primary">复制</a-button>
          <a-button class="save-btn" @click="addEnvData" type="primary">保存</a-button>
        </div>
      </div>
    </div>
  </div>
  <a-modal
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

  </a-modal>
</template>
<script setup lang="ts">

import {
  computed,
  defineEmits,
  defineProps,
  ref,
  watch
} from 'vue';
import {MenuOutlined, PlusOutlined} from '@ant-design/icons-vue';

import draggable from 'vuedraggable'
import Combobox from "ant-design-vue/es/vc-select/examples/combobox";
import mounted = Combobox.mounted;
import { useGlobalEnv } from '../../hooks/globalEnv';
import { useGlobalVarAndParams } from '../../hooks/globalVar';
import {globalParamscolumns, globalVarsColumns, serveServersColumns, tabPaneList} from '@/views/projectSettingV2/config';

const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);


import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from "@/views/projectSettingV2/store";
import {useStore} from "vuex";

// store 相关
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSettingV2:  ProjectSettingStateType}>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const globalParamsData = computed<any>(() => store.state.ProjectSettingV2.globalParamsData);
const globalVarsData = computed<any>(() => store.state.ProjectSettingV2.globalVarsData);
const serviceOptions = computed<any>(() => store.state.ProjectSettingV2.serviceOptions);
const envList = computed<any>(() => store.state.ProjectSettingV2.envList);

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
  await store.dispatch('ProjectSettingV2/getServersList', {
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
    //padding: 0 16px;
    padding-left: 8px;
    //display: inline-block;
    //&:hover{
    //  color: #1677ff;
    //}

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

  //.env-item-footer {
  //  padding: 0;
  //}
}

.right-content {
  flex: 1;
  height: calc(100vh - 138px);
  overflow-y: scroll;
  position: relative;
  //margin: 16px;
  padding: 16px;

  .globalVars, .globalParams, .envDetail {
    padding: 8px;

    .title {
      font-weight: bold;
      font-size: 18px;
      margin-bottom: 16px;
    }
  }

  .envDetail-content {
    position: relative;
  }

  .envDetail-footer {
    height: 60px;
    position: absolute;
    top: 8px;
    right: 8px;
    width: 300px;
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .save-btn {
      margin-right: 16px;
    }
  }
}

.vars-header, .serveServers-header {
  padding: 0 0 8px;
  line-height: 1.4;
  white-space: normal;
  text-align: left;
  margin-bottom: 8px;
}

.var-icon {
  display: inline-block;
}

.envDetail-btn {
  margin-top: 16px;
  margin-bottom: 16px;
}

.serveServers-header, .vars-header {
  font-weight: bold;
  margin-bottom: 0;
  margin-top: 16px;
}


.select-service {
  .ant-select-selector {
    border: 1px solid #d9d9d9;
  }
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
