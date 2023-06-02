<template>
  <div v-if="drawerVisible">
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }"
              :visible="drawerVisible"
              @close="onClose">
      <template #title>
        <div class="drawer-header">
          <div>{{ '测试计划执行详情' }}</div>
        </div>
      </template>
      <div class="drawer-content">
        <ReportBasicInfo :items="basicInfoList || []" :showBtn="false"/>
        <StatisticTable :data="statisticData" :value="statInfo"/>
        <Progress :exec-status="progressStatus"
                  :percent="progressValue"
                  @exec-cancel="execCancel"/>
        <LogTreeView :treeData="scenarioReports" :expandKeys="expandKeys"/>
      </div>
    </a-drawer>
  </div>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, computed, watch, onMounted} from 'vue';
import {useStore} from 'vuex';

import {
  ReportBasicInfo,
  StatisticTable,
  Progress,
  LogTreeView
} from '@/views/component/Report/components';

import {StateType as ReportStateType, StateType as PlanStateType} from "../store";
import settings from "@/config/settings";
import bus from "@/utils/eventBus";
import {getToken} from "@/utils/localToken";
import {WebSocket} from "@/services/websocket";
import {momentUtc} from "@/utils/datetime";
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ProjectSettingStateType} from "@/views/projectSetting/store";
import {StateType as UserStateType} from "@/store/user";
import {getDivision, getPercent, getPercentStr} from '@/utils/number';
import {
  scenarioReports, expandKeys,
  clearLog,
  execLogs, execRes, updateExecLogs, updateExecRes,statInfo
  , statisticData, initData, progressStatus, progressValue, updatePlanRes,
} from '@/composables/useExecLogs';

const props = defineProps<{
  drawerVisible: boolean
  title: string
  scenarioExpandActive: boolean
  showScenarioInfo: boolean
  scene: string // 查看详情的场景 【执行测试计划 exec_plan， 执行测试场景 exec_scenario， 查看报告详情 query_detail】
  reportId?: number
}>();

const emits = defineEmits(['onClose']);

const store = useStore<{
  Plan: PlanStateType,
  Global: GlobalStateType,
  Exec: ExecStatus,
  Report: ReportStateType,
  ProjectSetting: ProjectSettingStateType,
  User: UserStateType
  CurrentUser
}>();

const currPlan = computed<any>(() => store.state.Plan.currPlan);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
// TODO： 这里的envList是从ProjectSetting中获取的，需要修改下，会污染其他作用域下的数据
const envList = computed(() => store.state.ProjectSetting.envList);
const currentUser = computed(() => store.state.User.currentUser);
const currUser = computed(() => store.state.User.currentUser);

// 执行计划的基本信息
const basicInfoList = computed(() => {
  const curEnv = envList.value.find((item: any) => item.id === currEnvId.value)
  return [
    {
      label: '测试计划',
      value: currPlan.value.name || '-'
    },
    {
      label: '开始时间',
      value: momentUtc(new Date())
    },
    {
      label: '执行环境',
      value: curEnv ? curEnv.name : '--'
    },
    {
      label: '创建人',
      value: currUser.value.username || '--'
    },
  ]
})

const execStart = async () => {
  clearLog();
  const token = await getToken();
  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: token,
    planId: currPlan.value && currPlan.value.id,
    environmentId: currEnvId.value
  }
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execPlan', planExecReq: data}));
};

const execCancel = () => {
  progressStatus.value = 'exception';
  const msg = {act: 'stop', execReq: {planId: currPlan.value && currPlan.value.id}};
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
};


const OnWebSocketMsg = (data: any) => {
  console.log('832 websocket msg', data)

  if (!data.msg) return;
  if (progressStatus.value === 'cancel') return;
  const wsMsg = JSON.parse(data.msg);
  const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
  // 开始执行，初始化数据
  if (wsMsg.category == 'initialize') {
    initData(log);
    progressStatus.value = 'in_progress';
  }
  // 执行中
  else if (wsMsg.category == 'in_progress') {
    progressStatus.value = 'in_progress';
  }
  // 更新【计划】的执行结果
  else if (wsMsg.category == 'result' && log.planId) {
    updatePlanRes(log);
    console.log('计划的结果', log)
  }
  //  更新【场景】的执行结果
  else if (wsMsg.category == 'result' && log.scenarioId) {
    updateExecRes(log);
    console.log('场景的结果', log)
  }
  // 更新【场景里每条编排】的执行记录
  else if (wsMsg.category === "processor" && log.scenarioId) {
    console.log('场景里每条编排的执行记录', log)
    updateExecLogs(log);
  }
  // 执行完毕
  else if (wsMsg.category == 'end') {
    progressStatus.value = 'end';
  } else {
    console.log('其他情况：严格来说，不能执行到这儿');
  }
};

// websocket 连接状态 查询
const onWebSocketConnStatusMsg = (data: any) => {
  if (!data.msg) {
    return;
  }
  const {conn}: any = JSON.parse(data.msg);
  progressStatus.value = conn === 'success' ? 'in_progress' : 'failed';
}

function onClose() {
  emits('onClose');
}


watch(() => {
  return props.drawerVisible
}, (newVal: any) => {
  if (newVal) {
    execStart();
    bus.on(settings.eventWebSocketMsg, OnWebSocketMsg);
    bus.on(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
  } else {
    execCancel();
    bus.off(settings.eventWebSocketMsg, OnWebSocketMsg);
    bus.off(settings.eventWebSocketConnStatus, onWebSocketConnStatusMsg);
  }
}, {
  immediate: true,
});

onMounted(() => {
  execCancel();
});

</script>
<style scoped lang="less">
.report-drawer {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>


