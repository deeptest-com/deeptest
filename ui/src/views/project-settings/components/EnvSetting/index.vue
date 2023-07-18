<template>
  <div class="content">
    <div class="left-content">
      <div class="global">
        <div class="header">全局</div>
        <a-button :class="{ 'env-item': true, 'env-item-active': isShowGlobalVars }" @click="toVarPage"
          :type="isShowGlobalVars ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">V</i>
          </template>
          全局变量
        </a-button>
        <a-button :class="{ 'env-item': true, 'env-item-active': isShowGlobalParams }" @click="toParamsPage"
          :type="isShowGlobalParams ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">P</i>
          </template>
          全局参数
        </a-button>
      </div>
      <div style="margin: 0 16px;">
        <a-divider class="divider" />
      </div>
      <div class="env">
        <div class="header">环境</div>

        <draggable tag="div" :list="envList" class="list-group" handle=".handle" item-key="name" @end="handleDragEnd">
          <template #item="{ element, index }">
            <a-button :class="{ 'env-item': true, 'env-item-active': activeEnvDetail?.id === element.id }"
              :type="activeEnvDetail?.id === element.id ? 'primary' : 'text'" @click="toEnvDetail(element)" :key="index">
              <MenuOutlined class="handle dp-drag" />

              <span class="text"> {{ element.displayName }} </span>
              <span class="action">
                <copy-outlined class="copy" @click.stop="copyEnvData(element)" />
                <delete-outlined class="delete" @click.stop="deleteEnvData(element)" />
              </span>
            </a-button>
          </template>
        </draggable>

        <div style="margin: 0 16px;">
          <a-divider class="divider" />
        </div>
        <a-button :type="isShowAddEnv ? 'primary' : 'text'" :class="{
          'env-item': true,
          'env-item-footer': true,
          'env-item-active': isShowAddEnv
        }" @click="toEnvDetail(null)">
          <template #icon>
            <PlusOutlined />
          </template>
          新建环境
        </a-button>
      </div>
    </div>
    <div class="right-content">
      <router-view></router-view>
    </div>
  </div>
</template>
<script setup lang="ts">

import {
  computed,
  ref,
  watch,
} from 'vue';
import { MenuOutlined, PlusOutlined, CopyOutlined, DeleteOutlined } from '@ant-design/icons-vue';
import draggable from 'vuedraggable'
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from "@/views/ProjectSetting/store";
import { EnvDataItem } from '../../data';
import { useGlobalEnv } from '../../hooks/useGlobalEnv';

enum RouteNameMap {
  Var = 'enviroment.var',
  Params = 'enviroment.params',
  Detail = 'enviroment.envdetail'
}

const { activeEnvDetail, deleteEnvData, copyEnvData } = useGlobalEnv();
// store 相关
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const envList = computed<any>(() => store.state.ProjectSetting.envList);
const router = useRouter();

const params = router.currentRoute.value.params;
const routeName = router.currentRoute.value.name;

// 页面state相关
const isShowGlobalVars = ref(routeName === RouteNameMap.Var);
const isShowGlobalParams: any = ref(routeName === RouteNameMap.Params);
const isShowAddEnv = ref(routeName === RouteNameMap.Detail && !params.id);


function handleDragEnd(_e: any) {
  const envIdList = envList.value.map((e: EnvDataItem) => {
    return e.id;
  })
  store.dispatch('ProjectSetting/sortEnvList', {
    data: envIdList,
    projectId: currProject.value.id,
  })
}

async function toVarPage() {
  isShowGlobalVars.value = true;
  isShowGlobalParams.value = false;
  store.dispatch('ProjectSetting/setEnvDetail', null);
  router.push('/project-setting/enviroment/var');
}

async function toParamsPage() {
  isShowGlobalParams.value = true;
  isShowGlobalVars.value = false;
  store.dispatch('ProjectSetting/setEnvDetail', null);
  router.push('/project-setting/enviroment/params');
}

async function toEnvDetail(env: any) {
  isShowGlobalVars.value = false;
  isShowGlobalParams.value = false;
  isShowAddEnv.value = !env;
  await store.dispatch('ProjectSetting/setEnvDetail', env);
  router.push(`/project-setting/enviroment/envdetail${env ? `/${env.id}` : ''}`);
}

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

// 请求环境列表
async function getEnvsList() {
  console.log('%c[GET ENV LIST] --  currProject [globalEnv.ts -- 16]', 'color: red', currProject.value);
  await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id });
  const params: any = router.currentRoute.value.params;
  const { id } = params;
  if (id && id !== -1) {
    const selectedEnv: any = envList.value.find((item: any) => {
      return Number(id) === item.id;
    })
    store.dispatch('ProjectSetting/setEnvDetail', selectedEnv);
  }
}

// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  if (Object.keys(newVal).length > 0) {
    await getServersList();
    await getEnvsList();
  }

  // 默认展示全局变量
  // await showGlobalVars();
}, {
  immediate: true
})

watch(() => {
  return router.currentRoute.value
}, (val) => {
  const { params: { id }, name } = val;
  isShowAddEnv.value = name === RouteNameMap.Detail && !id;
  isShowGlobalVars.value = name === RouteNameMap.Var;
  isShowGlobalParams.value = name === RouteNameMap.Params;
  if (name === RouteNameMap.Var || name === RouteNameMap.Params) {
    store.dispatch('ProjectSetting/setEnvDetail', null);
  }
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
    display: flex;
    align-items: center;

    &:hover {
      .action {
        display: block;
      }
    }

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

    .text {
      flex: 1;
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
      padding-right: 6px;
    }

    .action {
      display: none;
      .copy {
        margin-right: 6px;
      }
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

  .globalVars,
  .globalParams,
  .envDetail {
    padding: 8px;
  }
}

.var-icon {
  display: inline-block;
}

:deep(.ant-input:not(.env-name):hover),
:deep(.ant-input:active),
:deep(.ant-input:focus) {
  border: 1px solid #4096ff !important
}

:deep(.ant-input:not(.env-name)) {
  border: 1px solid transparent !important
}

:deep(.custom-select .ant-select-selector) {
  border: 1px solid transparent !important;
}

:deep(.custom-select .ant-select-selector:hover),
:deep(.custom-select .ant-select-selector:active),
:deep(.ant-select-selector:focus) {
  border: 1px solid #4096ff !important
}
</style>
