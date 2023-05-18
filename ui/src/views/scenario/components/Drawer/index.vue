<template>
  <div>
    <a-drawer
        :placement="'right'"
        :width="1000"
        :closable="true"
        :visible="visible"
        class="drawer"
        wrapClassName="drawer-1"
        :bodyStyle="{padding:'16px',marginBottom:'56px'}"
        @close="onCloseDrawer">
      <!-- 头部信息  -->
      <template #title>
        <a-row type="flex" style="align-items: center;width: 100%">
          <a-col :span="8">
            <EditAndShowField placeholder="修改标题" :value="detailResult.name" @update="updateTitle"/>
          </a-col>
        </a-row>
      </template>
      <!-- 基本信息 -->
      <BasicInfo @change="changeBasicInfo"/>
      <!-- Tab 切换区域 -->
      <a-tabs v-model:activeKey="activeKey" force-render>
        <a-tab-pane class="test-developer" key="1" tab="测试开发">
          <div v-if="activeKey==='1'">
            <div class="exec-btn">
              <a-button @click="exec" :size="'small'" type="primary"><span>&nbsp;执行&nbsp;</span></a-button>
            </div>
            <Design :id="detailResult?.id"/>
          </div>
        </a-tab-pane>
        <a-tab-pane key="2" tab="执行历史" force-render>
          <div style="padding: 16px" v-if="activeKey==='2'">
            <ExecList @showDetail="showDetail"/>
          </div>
        </a-tab-pane>
        <!-- :::: 关联测试计划Tab-->
        <a-tab-pane key="3" tab="关联测试计划" force-render>
          <div style="padding: 16px" v-if="activeKey==='3'">
            <PlanList :linked="true"/>
          </div>
        </a-tab-pane>
      </a-tabs>
    </a-drawer>
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
        @on-ok="selectExecEnv"/>
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
import {PaginationConfig, Scenario} from "@/views/Scenario/data";
import {message} from "ant-design-vue";
import Design from "../Design/index.vue"
import PlanList from "./PlanList.vue";
import ExecList from "./ExecList.vue";
import ExecInfo  from "../Exec/index.vue";
import EnvSelector from "@/views/component/EnvSelector/index.vue";
import ExecListDetail from "./ExecListDetail.vue";
const store = useStore<{ Scenario, ProjectGlobal, ServeGlobal,Report }>();
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);
const pagination = computed<PaginationConfig>(() => store.state.Scenario.listResult.pagination);

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
    await store.dispatch('Scenario/saveScenario',
        {id: detailResult.value.id, categoryId: value}
    );
    emit('refreshList');
  }
}

async function cancel() {
  emit('close');
}


</script>
<style lang="less" scoped>
.drawer {
  margin-bottom: 60px;

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

.test-developer {
  height: 100%;
  width: 1000px;
  position: relative;

  .exec-btn {
    position: absolute;
    right: 4px;
    top: -48px;
  }
}


.drawer-exec-history-detail {
  :deep(.ant-drawer-header) {
    box-shadow: 0px 1px 0px rgba(0, 0, 0, 0.06);
  }
}
</style>
