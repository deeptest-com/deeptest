<template>
  <a-drawer
      :placement="'right'"
      :width="1000"
      :closable="true"
      :visible="visible"
      class="drawer"
      wrapClassName="drawer-1"
      :bodyStyle="{padding:0,marginBottom:'56px'}"
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
    <BasicInfo @changeStatus="changeStatus" @change-description="changeDescription" @changeCategory="changeCategory"/>

    <!-- Tab 切换区域 -->
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane class="test-developer" key="1" tab="测试开发">
        <DesignContent :id="detailResult?.id"/>
      </a-tab-pane>
      <a-tab-pane key="2" tab="执行历史" force-render>Content of Tab Pane 2</a-tab-pane>
      <a-tab-pane key="3" tab="关联计划" force-render>Content of Tab Pane 2</a-tab-pane>
    </a-tabs>

<!--    <div v-if="key === 'request'" class="drawer-btns">-->
<!--      <a-space>-->
<!--        <a-button type="primary" @click="save">保存</a-button>-->
<!--        <a-button @click="cancel">取消</a-button>-->
<!--      </a-space>-->
<!--    </div>-->
  </a-drawer>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed,
} from 'vue';
import BasicInfo from './BasicInfo.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';

import {useStore} from "vuex";
import {Scenario} from "@/views/Scenario/data";
import {message} from "ant-design-vue";
import DesignContent from "../design/index1.vue"
const store = useStore<{ Scenario, ProjectGlobal, ServeGlobal }>();
const detailResult = computed<Scenario>(() => store.state.Scenario.detailResult);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

function onCloseDrawer() {
  emit('close');
}

const activeKey = ref('1');

async function changeStatus(status) {
  await store.dispatch('Scenario/updateStatus',
      {id:detailResult.value.id, status: status}
  );
  await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

async function updateTitle(title) {
  await store.dispatch('Scenario/updateEndpointDetail',
      {...detailResult.value, title: title}
  );
  await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
}

async function changeDescription(description) {
  await store.dispatch('Scenario/updateEndpointDetail',
      {...detailResult.value, description}
  );
  await store.dispatch('Scenario/getEndpointDetail', {id: detailResult.value.id});
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
