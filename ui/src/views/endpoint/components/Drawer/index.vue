<template>
  <a-drawer
      :placement="'right'"
      :width="1200"
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
          <EditAndShowField placeholder="修改标题" :value="endpointDetail?.title || ''" @update="updateTitle"/>
        </a-col>
      </a-row>
    </template>
    <!-- 基本信息 -->
    <EndpointBasicInfo @changeStatus="changeStatus" @change-description="changeDescription"
                       @changeCategory="changeCategory"/>
    <!-- 接口设计区域 -->
    <a-card
        style="width: 100%"
        :bordered="false"
        :size="'small'"
        :headStyle="{padding:'0 24px',borderBottom:'none'}"
        :bodyStyle="{padding:'0 24px 0 24px'}">
      <template #title>
        <div style="margin-top: -12px;">
          <ConBoxTitle :backgroundStyle="'background: #FBFBFB;'" :title="'接口设计'"/>
        </div>
      </template>
      <a-tabs :activeKey="key" :animated="false" @change="changeTab">
        <a-tab-pane key="request" tab="定义">
          <EndpointDefine v-if="key === 'request'" @switchMode="switchMode"/> <!-- use v-if to force page reload-->
        </a-tab-pane>
        <a-tab-pane key="run" tab="调试">
          <!-- use v-if to force page reload -->
          <EndpointDebug v-if="key === 'run'" @switchToDefineTab="switchToDefineTab"/>
        </a-tab-pane>
        <a-tab-pane key="docs" tab="文档">
          <Docs v-if="key === 'docs'" @switchMode="switchMode"/> <!-- use v-if to force page reload-->
        </a-tab-pane>
      </a-tabs>
    </a-card>
    <div v-if="key === 'request' && showFooter" class="drawer-btns">
      <a-space>
        <a-button type="primary" @click="save">保存</a-button>
        <a-button @click="cancel">取消</a-button>
      </a-space>
    </div>
  </a-drawer>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';
import EndpointBasicInfo from './EndpointBasicInfo.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import EndpointDefine from './Define/index.vue';
import EndpointDebug from './Debug/index.vue';
import Docs from './Docs/index.vue';

import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {message} from "ant-design-vue";

const store = useStore<{ Endpoint, ProjectGlobal, ServeGlobal }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

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

async function changeTab(value) {
  key.value = value;
  // 切换到调试页面时，需要先保存
  if (value === 'run') {
    await store.dispatch('Endpoint/updateEndpointDetail',
        {...endpointDetail.value}
    );
    // 获取最新的接口详情,比如新增的 接口的id可能会变化，所以需要重新获取
    await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});
  }
}

function switchToDefineTab() {
    key.value = 'request';
}

const showFooter = ref(true);
function switchMode(val) {
  showFooter.value = (val === 'form');
}

async function changeStatus(status) {
  await store.dispatch('Endpoint/updateStatus',
      {id: endpointDetail.value.id, status: status}
  );
  await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});
}

async function updateTitle(title) {
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...endpointDetail.value, title: title}
  );
  await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});
}

async function changeDescription(description) {
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...endpointDetail.value, description}
  );
  await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});
}

async function changeCategory(value) {
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...endpointDetail.value, categoryId: value}
  );
  await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});
}


const key = ref('request');

async function cancel() {
  emit('close');
}

async function save() {
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...endpointDetail.value}
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

</style>
