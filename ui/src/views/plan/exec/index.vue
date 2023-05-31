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
import {getDivision, getPercent, getPercentStr} from '@/utils/number';

const progressStatus = ref('in_progress');
const progressValue = computed(() => {
  const {
    totalProcessorNum = 1,
    finishProcessorNum,
  } = statInfo.value;
  return Math.round(((finishProcessorNum || execLogs.value.length) / totalProcessorNum) * 100);
});

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
const currentUser = computed(()=> store.state.User.currentUser);
const currUser = computed(() => store.state.User.currentUser);
const processNum = ref(0);

// 统计聚合数据
const statInfo = ref({
  failAssertionNum: 0,
  failInterfaceNum: 0,
  failRequestNum: 0,
  failScenarioNum: 0,
  passAssertionNum: 0,
  passInterfaceNum: 0,
  passRequestNum: 0,
  passScenarioNum: 0,
  totalAssertionNum: 0,
  totalInterfaceNum: 0,
  totalRequestNum: 0,
  totalScenarioNum: 0,
  totalProcessorNum: 0,
  notTestNum: 0,
  finishProcessorNum: 0,
  duration: 0,
})

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
const statisticData = computed(() => {
  const {
    failAssertionNum,
    failInterfaceNum,
    failRequestNum,
    failScenarioNum,
    passAssertionNum,
    passInterfaceNum,
    passRequestNum,
    passScenarioNum,
    totalAssertionNum,
    totalInterfaceNum,
    totalRequestNum,
    totalScenarioNum,
    totalProcessorNum,
    duration,
    notTestNum,
  } = statInfo.value;
  // 计算平均接口耗时
  let interfaceDuration = 0;
  let interfaceNum = 0;
  execLogs.value.forEach((item: any) => {
    if (item.processorCategory === "processor_interface") {
      interfaceDuration += (item.cost || 0);
      interfaceNum++;
    }
  });
  const passRate = getPercentStr(passAssertionNum, totalAssertionNum);
  const notPassRate = getPercentStr(failAssertionNum, totalAssertionNum);
  const notTestNumRate = getPercentStr(notTestNum, totalAssertionNum);
  // 平均接口耗时
  const avgInterfaceDuration = getDivision(interfaceDuration, interfaceNum);
  return [
    {
      label: '通过',
      value: `${passRate} ${passAssertionNum} 个`,
    },
    {
      label: '接口总耗时',
      value: `${interfaceDuration} 毫秒`
    },
    {
      label: '失败',
      value: `${notPassRate}  ${failAssertionNum} 个`,
      class: 'fail',
    },
    {
      label: '平均接口耗时',
      value: `${avgInterfaceDuration} 毫秒`,
    },
    {
      label: '未测',
      value: `${notTestNumRate} ${notTestNum}个`,
      class: 'fail',
    },
    {
      label: '测试场景 (成功/失败)',
      value: `${totalScenarioNum} (${passScenarioNum}/${failScenarioNum})`,
    },

  ]
})

const scenarioReports = computed(() => {
  let res = [...genLogTreeView(execLogs.value, execRes.value)];
  return res;
})

const expandKeys = computed(() => {
  return scenarioReports.value.map((item: any) => item.key);
})

const execStart = async () => {
  processNum.value = 1;
  execLogs.value = [];
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

/**
 * @description 生成执行记录树
 * @param execLogs 场景的执行日志 数组
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
  console.log('832 scenarioReports', scenarioReports)
  return scenarioReports;
}

// 打平的执行记录
const execLogs: any = ref([]);

// 更新场景的执行记录，不包括场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
function updateExecLogs(log) {
  const isExist = execLogs.value.some((item: any) => {
    return item.logId === log.logId && item.scenarioId === log.scenarioId;
  });
  // 1. 更新执行记录
  if (isExist) {
    execLogs.value.forEach((item: any) => {
      if (item.logId === log.logId && item.scenarioId === log.scenarioId) {
        item = {...item, ...log};
      }
    });
    // 2. 新增执行记录
  } else {
    execLogs.value.push(log);
  }
}

// 场景的执行结果列表
const execRes: any = ref([]);

// 更新场景的执行结果
// todo 优化: 可以优化成算法，使用 hash
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

/**
 * 从每次返回的执行日志中更新统计数据
 * */
function updateStatFromLog(res: any) {
  const {
    failAssertionNum = 0,
    failInterfaceNum = 0,
    failRequestNum = 0,
    failScenarioNum = 0,
    passAssertionNum = 0,
    passInterfaceNum = 0,
    passRequestNum = 0,
    passScenarioNum = 0,
    totalAssertionNum = 0,
    totalInterfaceNum = 0,
    totalRequestNum = 0,
    totalScenarioNum = 0,
    totalProcessorNum = 0,
    finishProcessorNum = 0,
    duration = 0,
  }: any = res;
  console.log('updateStatFromLog', res);
  const notTestNum = totalAssertionNum - passAssertionNum - failAssertionNum;
  statInfo.value = {
    failAssertionNum,
    failInterfaceNum,
    failRequestNum,
    failScenarioNum,
    passAssertionNum,
    passInterfaceNum,
    passRequestNum,
    passScenarioNum,
    totalAssertionNum,
    totalInterfaceNum,
    totalRequestNum,
    totalScenarioNum,
    totalProcessorNum,
    finishProcessorNum,
    notTestNum: notTestNum >= 0 ? notTestNum : 0,
    duration
  }
}

// 更新统计聚合数据
function initData(res: any) {
  updateStatFromLog(res);
}

// 【计划】的执行最终结果 用于更新最终的执行结果
function updatePlanRes(res) {
  updateStatFromLog(res);
}

const OnWebSocketMsg = (data: any) => {
  if (!data.msg) return;
  if(progressStatus.value === 'cancel') return;
  const wsMsg = JSON.parse(data.msg);
  const log = wsMsg.data ? JSON.parse(JSON.stringify(wsMsg.data)) : {};
  console.log('832222', wsMsg);
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
  progressStatus.value = conn === 'success' ? WsMsgCategory.InProgress : 'failed';
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


