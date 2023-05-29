<template>
  <div v-if="drawerVisible">
    <a-drawer class="report-drawer" :closable="true" :width="1000" :bodyStyle="{ padding: '24px' }"
              :visible="drawerVisible"
              @close="onClose">
      <template #title>
        <div class="drawer-header">
          <div>{{ '测试报告详情' }}</div>
        </div>
      </template>
      <div class="drawer-content">
        <ReportBasicInfo :items="basicInfoList || []" :showBtn="false"/>
        <StatisticTable  :data="statisticData"/>
        <!--      <Progress :exec-status="execResult.progressStatus" v-if="scene !== ReportDetailType.QueryDetail"-->
        <!--                :percent="execResult.progressValue || 10"-->
        <!--                @exec-cancel="execCancel"/>-->
        <LogTreeView :treeData="scenarioReports"/>
      </div>
    </a-drawer>
  </div>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, computed, watch} from 'vue';
import {useStore} from 'vuex';

import {
  ReportBasicInfo,
  StatisticTable,
  ScenarioCollapsePanel,
  EndpointCollapsePanel,
  Progress,
  LogTreeView
} from '@/views/component/Report/components';

import {StateType as ReportStateType, StateType as PlanStateType} from "../store";
import {ReportDetailType, WsMsgCategory} from '@/utils/enum';
import settings from "@/config/settings";
import bus from "@/utils/eventBus";
import {getToken} from "@/utils/localToken";
import {WebSocket} from "@/services/websocket";
import {momentUtc} from "@/utils/datetime";
import {StateType as GlobalStateType} from "@/store/global";
import {ExecStatus} from "@/store/exec";
import {StateType as ProjectSettingStateType} from "@/views/projectSetting/store";
import {StateType as UserStateType} from "@/store/user";

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
}>();

const currPlan = computed<any>(() => store.state.Plan.currPlan);
const currEnvId = computed(() => store.state.ProjectSetting.selectEnvId);
const envList = computed(() => store.state.ProjectSetting.envList);

const currUser = computed(() => store.state.User.currentUser);
const processNum = ref(0);
const transformWithUndefined = (num: number | undefined) => {
  return num || 0;
}

const calcNum = (currNum, lastNum) => {
  return currNum + transformWithUndefined(lastNum);
}

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
    }
)

const statisticData = computed(() => {
  return [
    {
      label: '通过',
      value: currPlan.value.totalCaseNum || 0
    },
    {
      label: '总耗时',
      value: currPlan.value.execCaseNum || 0
    },
    {
      label: '失败',
      value: currPlan.value.passCaseNum || 0
    },
    {
      label: '平均接口耗时',
      value: currPlan.value.failCaseNum || 0
    },
    {
      label: '未测',
      value: currPlan.value.skipCaseNum || 0
    },
  ]
})

const scenarioReports = computed(() => {
  let res = [...genLogTreeView(execLogs.value, execRes.value)];
  console.log('832 scenarioReports', res);
  return res;
})

const execStart = async () => {
  console.log('execStart');
  // processNum.value = 0;
  execLogs.value = [];
  const data = {
    serverUrl: process.env.VUE_APP_API_SERVER, // used by agent to submit result to server
    token: await getToken(),
    planId: currPlan.value && currPlan.value.id,
    environmentId: currEnvId.value
  }
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify({act: 'execPlan', planExecReq: data}));
};

const execCancel = () => {
  console.log('execCancel');
  // flattenExecLogs.value = [];
  const msg = {act: 'stop', execReq: {planId: currPlan.value && currPlan.value.id}};
  WebSocket.sentMsg(settings.webSocketRoom, JSON.stringify(msg))
};

/**
 * @description 生成执行记录树
 * @param execLogs 场景的执行日志
 * @param execRes  场景的执行结果
 * */
function genLogTreeView(execLogs, execRes) {
  // 用于存储根节点，即场景节点,即 processorCategory为 processor_root的节点
  const scenarioReports: any = [];
  execLogs.forEach((item: any) => {
    if (item.processorCategory === "processor_root") {
      const res = execRes.find((log) => log.scenarioId === item.scenarioId) || {};
      item.logs = [{
        ...item,
        ...res,
        logs: []
      }];
      scenarioReports.push(item);
    }
  });
  scenarioReports.forEach((scenario) => {
    function fn(array, rootId) {
      const res: any = [];
      // 用于存储 树节点的map
      const map = {};
      array.forEach((item) => {
        map[item.id] = {...item}
      });
      array.forEach((item) => {
        const {id, parentId} = item;
        if (!map[id]) return;
        if (!parentId) return;
        if (parentId === rootId) {
          res.push(map[id]);
        } else {
          if (map[parentId]?.logs) {
            map[parentId].logs.push(map[id]);
          } else {
            map[parentId].logs = [map[id]];
          }
        }
      })
      return res;
    }

    scenario.logs[0].logs = fn(execLogs, scenario.id);
  });
  return scenarioReports;
}

// 打平的执行记录
const execLogs: any = ref([]);

// 更新场景的执行记录，不包括场景的执行结果
function updateExecLogs(log) {
  // 1. 更新执行记录
  if (execLogs.value.some((item: any) => item.id === log.id)) {
    execLogs.value.forEach((item: any) => {
      if (item.id === log.id) {
        item = {...item, ...log};
      }
    });
    // 2. 新增执行记录
  } else {
    execLogs.value.push(log);
  }
}

// 场景的执行结果记录
const execRes: any = ref([]);

// 更新场景的执行结果
function updateExecRes(res) {
  // 1. 更新执行结果
  if (execRes.value.some((item: any) => item.scenarioId === res.scenarioId)) {
    execRes.value.forEach((item: any) => {
      if (item.scenarioId === res.scenarioId) {
        item = {...item, ...res};
      }
    });
    // 2. 新增执行结果
  } else {
    execRes.value.push(res);
  }
}

// 【计划】的执行结果记录
function updatePlanRes(res) {
  console.log(res);
}

const OnWebSocketMsg = (data: any) => {
  if (!data.msg) return;
  const wsMsg = JSON.parse(data.msg);
  const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
  console.log('832222', wsMsg);
  // 开始执行
  if (wsMsg.category == 'in_progress') {
    console.log('计划开始执行');
    // 更新【计划】的执行结果
  } else if (wsMsg.category == 'result' && log.planId && !('scenarioId' in log)) {
    updatePlanRes(log);
    //  更新【场景】的执行结果
  } else if (wsMsg.category == 'result' && 'scenarioId' in log && !log.planId) {
    updateExecRes(log);
    // 更新【场景里每条编排】的执行记录
  } else if (wsMsg.category === "processor" && 'scenarioId' in log) {
    updateExecLogs(log);
  } else if (wsMsg.category == 'end') {
    // 执行完毕
    console.log('计划执行完成');
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
  // logDetailData.value.progressStatus = conn === 'success' ? WsMsgCategory.InProgress : 'failed';
  // store.dispatch('Plan/setExecResult', logDetailData.value);
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

</script>
<style scoped lang="less">
.report-drawer {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>
