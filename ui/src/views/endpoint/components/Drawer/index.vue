<template>
  <a-drawer class="drawer"
      :placement="'right'"
      :width="1200"
      :closable="true"
      :visible="visible"
      wrapClassName="drawer-1"
      :headerStyle="{position:'sticky',top:0,zIndex:9999}"
      :bodyStyle="{padding:0,minHeight:'100vh'}"
      @close="onCloseDrawer">

    <!-- 头部信息  -->
    <template #title>
      <div class="header-text">
        <span class="serialNumber">[{{ endpointDetail.serialNumber }}]</span>
        <EditAndShowField :custom-class="'show-on-hover'" placeholder="修改标题" :value="endpointDetail?.title || ''"
                          @update="updateTitle"/>
      </div>
    </template>

    <!-- 基本信息 -->
    <EndpointBasicInfo @changeStatus="changeStatus"
                       @change-description="changeDescription"
                       @changeCategory="changeCategory"/>

    <!-- 接口设计区域 -->
    <a-card
        style="width: 100%"
        :bordered="false"
        :size="'small'"
        :bodyStyle="{padding:'0 16px'}"
        :headStyle="{padding:'0 16px',borderBottom:'none'}">
      <template #title>
        <div>
          <ConBoxTitle :show-arrow="true" @expand="expandInfo" :backgroundStyle="'background: #FBFBFB;'"
                       :title="'接口设计'"/>
        </div>
      </template>

      <a-tabs class="tabs"
          :tabBarStyle="{marginBottom: 0}"
          v-show="expand" :activeKey="key" :animated="false" @change="changeTab">
        <template #tabBarExtraContent>
          <a-button v-if="key === 'request' && showFooter" type="primary" @click="save">
            <template #icon>
              <SaveOutlined/>
            </template>
            保存
          </a-button>
        </template>

        <a-tab-pane key="request" tab="定义" class="tab">
          <div class="tab-container">
            <EndpointDefine v-if="key === 'request'" @switchMode="switchMode"/> <!-- use v-if to force page reload-->
          </div>
        </a-tab-pane>

        <a-tab-pane key="run" tab="调试">
          <div class="tab-container">
            <!-- use v-if to force page reload -->
            <EndpointDebug v-if="key === 'run'" @switchToDefineTab="switchToDefineTab"/>
          </div>
        </a-tab-pane>

        <a-tab-pane key="cases" tab="用例">
          <div class="tab-container">
            <!-- use v-if to force page reload -->
            <EndpointCases v-if="key === 'cases'" @switchToDefineTab="switchToDefineTab"/>
          </div>
        </a-tab-pane>

        <a-tab-pane key="docs" tab="文档">
          <Docs :onlyShowDocs="true"
                :showHeader="false"
                v-if="key === 'docs' && docsData"
                :data="docsData"
                @switchToDefineTab="switchToDefineTab"
                :show-menu="true"/> <!-- use v-if to force page reload-->
        </a-tab-pane>
      </a-tabs>
    </a-card>
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
import EndpointCases from './Cases/index.vue';
import Docs from '@/components/Docs/index.vue';

import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {message} from "ant-design-vue";
import {SaveOutlined} from '@ant-design/icons-vue';

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

const docsData = ref(null);

async function changeTab(value) {
  console.log('changeTab', value)

  key.value = value;
  // 切换到调试页面时，需要先保存
  if (value === 'run') {
    // Comment out since it cause a issue in ./Debug/method @chenqi
    // await store.dispatch('Endpoint/updateEndpointDetail',
    //     {...endpointDetail.value}
    // );
    // 获取最新的接口详情,比如新增的 接口的id可能会变化，所以需要重新获取
    // await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});

  } else if (value === 'docs') {
    const res = await store.dispatch('Endpoint/getDocs', {
      endpointIds: [endpointDetail.value.id],
    });
    docsData.value = res;
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

const expand = ref(true)

function expandInfo(val) {
  expand.value = val
}

</script>
<style lang="less" scoped>
.drawer {
  margin-bottom: 60px;

  .tabs {
    .tab {
      .tab-container {
        margin-top: 16px;
      }
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
  padding-right: 24px;
  height: 56px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 16px;
  z-index: 99;
}

.header-text {
  display: flex;
  max-width: 80%;

  .serialNumber {
    margin-right: 6px;
  }
}
</style>
