<template>
  <div>
    <DrawerLayout :visible="visible" @close="onCloseDrawer" :stickyKey="stickyKey"
                  class="scenario-interface-design">
      <!-- 头部信息  -->
      <template #header>
        <div class="header-text">
          <span class="serialNumber">[{{ detailResult.serialNumber }}]</span>
          <EditAndShowField placeholder="修改标题"
                            :value="detailResult?.title || ''"
                            @update="updateTitle"/>
        </div>
      </template>

      <!-- 基本信息 -->
      <template #basicInfo>
        <BasicInfo @change="changeBasicInfo"/>
      </template>
      <template #tabHeader>
        <div class="tab-header-items">
          <div class="tab-header-item"
               :class="{'active':tab.key === activeKey}" v-for="tab in tabsList"
               :key="tab.key"
               @click="changeTab(tab.key)">
            <span>{{ tab.label }}</span>
          </div>
        </div>

        <div class="tab-header-btns">
          <div v-if="activeKey==='1'"
               :style="{right: isShowSync ? '200px' : '110px'}"
               class="exec-scenario-btn">
            <a-button @click="exec" type="primary">
              <span>执行场景</span>
            </a-button>
          </div>
        </div>
      </template>

      <template #tabContent>
        <div class="tab-pane">
          <div v-if="activeKey==='1'" >
            <Design :id="detailResult?.id"/>
          </div>
          <div v-if="activeKey==='2'" style="padding: 16px">
            <ExecList @showDetail="showDetail"/>
          </div>
          <div style="padding: 16px" v-if="activeKey==='3'">
            <PlanList :linked="true"/>
          </div>
        </div>
      </template>
    </DrawerLayout>

    <!-- 动态场景执行抽屉 -->
    <a-drawer
        :placement="'right'"
        :width="1000"
        :closable="true"
        :visible="execDrawerVisible"
        :title="'执行场景'"
        class="drawer"
        wrapClassName="drawer-exec"
        :bodyStyle="{padding:'16px',marginBottom:'56px'}"
        @close="onCloseExecDrawer">
      <ExecInfo v-if="execDrawerVisible"/>
    </a-drawer>

    <EnvSelector
        :env-select-drawer-visible="selectEnvVisible"
        @on-cancel="cancelSelectExecEnv"
        @on-ok="selectExecEnv" />

    <!-- ::::静态数据：查看执行历史的详情 -->
    <a-drawer
        :placement="'right'"
        :width="1000"
        :title="'执行详情'"
        :closable="true"
        :visible="execListDetailVisible"
        class="drawer"
        wrapClassName="drawer-exec-history-detail"
        :bodyStyle="{padding:'16px',marginBottom:'56px'}"
        @close="execListDetailVisible = false">
      <template #title>
        <div class="drawer-header">
          <div>{{ '测试报告详情' }}</div>
        </div>
      </template>
      <ExecListDetail/>
    </a-drawer>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, reactive, watch,
} from 'vue';
import BasicInfo from './BasicInfo.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';

import {useStore} from "vuex";
import {Scenario} from "@/views/Scenario/data";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as ScenarioStateType} from "../../store";
import Design from "../Design/index.vue"
import PlanList from "./PlanList.vue";
import ExecList from "./ExecList.vue";
import ExecInfo  from "../Exec/index.vue";
import EnvSelector from "@/views/component/EnvSelector/index.vue";
import ExecListDetail from "./ExecListDetail.vue";
import DrawerLayout from "@/views/component/DrawerLayout/index.vue";
import {ProcessorInterfaceSrc, UsedBy} from "@/utils/enum";

const store = useStore<{ Debug: Debug, Scenario: ScenarioStateType, ProjectGlobal, ServeGlobal, Report }>();
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);
const debugData = computed<any>(() => store.state.Debug.debugData);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  drawerTabKey: {
    required: true,
    type: String
  },
  execVisible: {
    required: true,
    type: Boolean,
  }
});

const emit = defineEmits(['ok', 'close', 'refreshList', 'closeExecDrawer']);
const activeKey = ref('1');
const stickyKey = ref(0);
const tabsList = [
  {
    "key": "1",
    "label": "测试开发"
  },
  {
    "key": "2",
    "label": "执行历史"
  },
  {
    "key": "3",
    "label": "关联测试计划"
  },
]
async function changeTab(value) {
  activeKey.value = value;
  stickyKey.value ++;
}


const execDrawerVisible = ref(false);
const selectEnvVisible = ref(false);

// 执行历史详情
const execListDetailVisible = ref(false);
async function cancelSelectExecEnv(record: any) {
  selectEnvVisible.value = false;
}

async function showDetail(record:any) {
  execListDetailVisible.value = true;
  await store.dispatch('Scenario/getScenariosReportsDetail', {id: record.id});
}

async function selectExecEnv() {
  selectEnvVisible.value = false;
  execDrawerVisible.value = true;
}

function getScenarioList() {
  console.log('get')
}

function onCloseDrawer() {
  emit('close');
}

function onCloseExecDrawer() {
  execDrawerVisible.value = false;
  emit('closeExecDrawer');
}

function exec() {
  selectEnvVisible.value = true;
}

watch(() => {
  return props.drawerTabKey;
}, (val) => {
  activeKey.value = val;
});

watch(() => {
  return props.execVisible;
}, (val) => {
  execDrawerVisible.value = val;
});

// 更新标题
async function updateTitle(title) {
  await store.dispatch('Scenario/saveScenario',
      {id: detailResult.value.id, name: title}
  );
  emit('refreshList');
}

async function changeBasicInfo(type, value) {
  if(type==='status') {
    await store.dispatch('Scenario/updateStatus',
        {id: detailResult.value.id, status: value}
    );
    emit('refreshList');
  }
  if(type==='priority') {
    await store.dispatch('Scenario/updatePriority',
        {id: detailResult.value.id, priority: value}
    );
    emit('refreshList');
  }
  if(type==='desc') {
    await store.dispatch('Scenario/saveScenario',
        {id: detailResult.value.id, desc: value}
    );
    emit('refreshList');
  }
  if(type==='categoryId') {
    await store.dispatch('Scenario/updateCategoryId',
        {id: detailResult.value.id, categoryId: value}
    );
    emit('refreshList');
  }
  if(type==='type') {
    await store.dispatch('Scenario/saveScenario',
        {id: detailResult.value.id, type: value}
    );
    emit('refreshList');
  }
}

async function cancel() {
  emit('close');
}

const isShowSync = computed(() => {
  const ret = debugData.value.processorInterfaceSrc !== ProcessorInterfaceSrc.Custom  &&
      debugData.value.processorInterfaceSrc !== ProcessorInterfaceSrc.Curl

  return ret
})

</script>

<style lang="less" scoped>
.scenario-interface-design {
  .tab-header-btns {
    position: relative;
    .exec-scenario-btn {
      position: absolute;
      //right: 200px;
      top: -16px;
    }
  }
}

.drawer {
  margin-bottom: 60px;

  .dp-tabs-full-height {
    height: calc(100% - 161px);

    .test-developer {
      height: 100%;
      width: 100%;
      position: relative;
    }
  }

  .title {
    width: auto;

    .ant-input-affix-wrapper {
      width: auto;
      border: none;

      &:focus {
        border: none;
        outline: none;
        box-shadow: none;
      }
    }

    input {
      width: auto;
      border: none;

      &:focus {
        border: none;
        border: none;
        outline: none;
        box-shadow: none;
      }
    }
  }
}

.drawer-btns {
  background: #ffffff;
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  position: absolute;
  bottom: 0;
  //right: 0;
  width: 100%;
  height: 56px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 16px;
  z-index: 99;
}



.drawer-exec-history-detail {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>
