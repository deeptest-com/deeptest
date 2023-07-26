<template>
  <DrawerLayout :visible="visible" @close="emit('close');" :stickyKey="stickyKey">
    <!-- 头部信息  -->
    <template #header>
      <div class="header-text">
        <span class="serialNumber">[{{ endpointDetail.serialNumber }}]</span>
        <EditAndShowField :custom-class="'show-on-hover'" placeholder="修改标题" :value="endpointDetail?.title || ''"
                          @update="updateTitle"/>
      </div>
    </template>
    <template #basicInfo>
      <!-- 基本信息 -->
      <EndpointBasicInfo @changeStatus="changeStatus"
                         @change-description="changeDescription"
                         @changeCategory="changeCategory"/>
    </template>

    <template #tabHeader>
      <div class="tab-header-items">
        <div class="tab-header-item"
             v-for="tab in tabsList"
             :key="tab.key"
             :class="{'active':tab.key === activeTabKey}"
             @click="changeTab(tab.key)">
          <span>{{ tab.label }}</span>
        </div>
      </div>
      <div class="tab-header-btns">
        <a-button v-if="activeTabKey === 'request' && showFooter" type="primary" @click="save">
          <template #icon>
            <SaveOutlined/>
          </template>
          保存
        </a-button>
      </div>
    </template>

    <template #tabContent>
      <div class="tab-pane">
        <EndpointDefine v-if="activeTabKey === 'request'"
                        @switchMode="switchMode"/>

        <EndpointDebug v-if="activeTabKey === 'run'"
                       @switchToDefineTab="switchToDefineTab"/>

        <EndpointCases v-if="activeTabKey === 'cases'"
                       v-model:showList="showList"
                       @switchToDefineTab="switchToDefineTab"/>

        <Docs :onlyShowDocs="true"
              :showHeader="false"
              v-if="activeTabKey === 'docs' && docsData"
              :data="docsData"
              @switchToDefineTab="switchToDefineTab"
              :show-menu="true"/> <!-- use v-if to force page reload-->
      </div>
    </template>
  </DrawerLayout>
</template>

<script lang="ts" setup>
import {computed, defineEmits, defineProps, ref,} from 'vue';
import EndpointBasicInfo from './EndpointBasicInfo.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EndpointDefine from './Define/index.vue';
import EndpointDebug from './Debug/index.vue';
import EndpointCases from './Cases/index.vue';
import Docs from '@/components/Docs/index.vue';
import DrawerLayout from "@/views/component/DrawerLayout/index.vue";
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {message} from "ant-design-vue";
import {SaveOutlined} from '@ant-design/icons-vue';

const store = useStore<{ Endpoint, ProjectGlobal, ServeGlobal,Global }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  }
})

const emit = defineEmits(['ok', 'close', 'refreshList']);

const showList = ref(true)
const docsData = ref(null);

const tabsList = [
  {
    "key": "request",
    "label": "定义"
  },
  {
    "key": "run",
    "label": "调试"
  },
  {
    "key": "cases",
    "label": "用例"
  },
  {
    "key": "docs",
    "label": "文档"
  },
]

const stickyKey = ref(0);
async function changeTab(value) {
  console.log('changeTab', value)

  // click cases tab again, will cause EndpointCases component back to case list page
  if (activeTabKey.value === 'cases' && activeTabKey.value === value) {
    showList.value = true // back to list
    return
  }

  activeTabKey.value = value;
  stickyKey.value ++;
  // 切换到调试页面时，需要先保存
  if (value === 'run') {
    // Comment out since it cause a issue in ./Debug/method @chenqi
    // await store.dispatch('Endpoint/updateEndpointDetail',
    //     {...endpointDetail.value}
    // );
    // 获取最新的接口详情,比如新增的 接口的id可能会变化，所以需要重新获取
    // await store.dispatch('Endpoint/getEndpointDetail', {id: endpointDetail.value.id});

  } else if (value === 'docs') {
    docsData.value = await store.dispatch('Endpoint/getDocs', {
      endpointIds: [endpointDetail.value.id],
      needDetail: true,
    });
  }
}

function switchToDefineTab() {
  activeTabKey.value = 'request';
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


const activeTabKey = ref('request');

async function cancel() {
  emit('close');
}

async function save() {

  store.commit("Global/setSpinning",true)
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...endpointDetail.value}
  ).finally( ()=>{
        store.commit("Global/setSpinning",false)
      }
  );
  store.commit("Global/setSpinning",false)
  message.success('保存成功');
  emit('refreshList');
}


</script>

<style lang="less" scoped>
.header-text {
  display: flex;
  max-width: 80%;

  .serialNumber {
    margin-right: 6px;
  }
}
</style>
