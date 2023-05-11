<template>
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
    <BasicInfo @change="change" @change-description="changeDescription" @changeCategory="changeCategory"/>

    <!-- Tab 切换区域 -->
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane class="test-developer" key="1" tab="测试开发">
        <DesignContent :id="detailResult?.id"/>
      </a-tab-pane>
      <a-tab-pane key="2" tab="执行历史" force-render>
        <div style="padding: 16px">
          <ExecList
              :list="[]"
              :show-scenario-operation="true"
              :columns="columns"
              :loading="true"
              :pagination="null"
              @refresh-list="getScenarioList" />
        </div>
      </a-tab-pane>
      <a-tab-pane key="3" tab="关联测试计划" force-render>
        <div style="padding: 16px">
          <PlanList
              :list="[]"
              :show-scenario-operation="true"
              :columns="columns"
              :loading="true"
              :pagination="null"
              @refresh-list="getScenarioList" />
        </div>
      </a-tab-pane>
    </a-tabs>

  </a-drawer>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, reactive,
} from 'vue';
import BasicInfo from './BasicInfo.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';

import {useStore} from "vuex";
import {Scenario} from "@/views/Scenario/data";
import {message} from "ant-design-vue";
import DesignContent from "../../design/index1.vue"
import  PlanList from "./PlanList.vue";
import  ExecList from "./ExecList.vue";
import Associate from  "./Associate.vue"
const store = useStore<{ Scenario, ProjectGlobal, ServeGlobal }>();
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
  },
  {
    title: '摘要',
    dataIndex: 'name',
    slots: { customRender: 'name' }
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: { customRender: 'status' }
  },
  {
    title: '测试通过率',
    dataIndex: 'testPassRate',
  },
  {
    title: '负责人',
    dataIndex: 'adminName',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
    slots: { customRender: 'updatedAt' }
  },
];

function getScenarioList() {
  console.log('get')
}
function onCloseDrawer() {
  emit('close');
}

const activeKey = ref('1');

async function change(type,val) {
  console.log('832 change',type,val)
  // await store.dispatch('Scenario/updateStatus',
  //     {id:detailResult.value.id, status: status}
  // );
  // await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

async function updateTitle(title) {
  console.log('832 updateTitle',title);
  // await store.dispatch('Scenario/updateEndpointDetail',
  //     {...detailResult.value, title: title}
  // );
  // await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

async function changeDescription(description) {
  console.log('832 changeDescription',description);
  // await store.dispatch('Scenario/updateEndpointDetail',
  //     {...detailResult.value, description}
  // );
  // await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

async function changeCategory(value) {
  await store.dispatch('Scenario/updateEndpointDetail',
      {...detailResult.value, categoryId: value}
  );
  await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

const key = ref('request');

async function cancel() {
  emit('close');
}

async function save() {
  await store.dispatch('Scenario/updateEndpointDetail',
      {...detailResult.value}
  );
  message.success('保存成功');
  emit('refreshList');
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
.test-developer{
  height: 100%;
  width: 1000px;
}
</style>
