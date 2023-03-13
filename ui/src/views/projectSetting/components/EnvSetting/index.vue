<template>
  <div class="content">
    <div class="left-content">
      <ul class="global">
        <li class="header">全局</li>
        <ul>
          <li :class="{'env-item':true,'env-item-active':isShowGlobalVars}" @click="showGlobalVars"><i class="var-icon">V</i>全局变量
          </li>
          <li :class="{'env-item':true,'env-item-active':isShowGlobalParams}" @click="showGlobalParams"><i
              class="param-icon">P</i>全局参数
          </li>
        </ul>
      </ul>
      <div style="margin: 0 16px;">
        <a-divider class="divider"/>
      </div>
      <ul class="env">
        <li class="header">环境</li>
        <draggable
            tag="ul"
            :list="envList"
            class="list-group"
            handle=".handle"
            item-key="name">
          <template #item="{ element, index }">
            <li
                :class="{'env-item':true,'env-item-active':activeEnvDetail?.name === element.name}"
                @click="showEnvDetail(element)"
                class="env-item" :key="index">
              <MenuOutlined class="handle"/>
              <span class="text">{{ element.name }} </span>
            </li>
          </template>
        </draggable>
        <div style="margin: 0 16px;">
          <a-divider class="divider"/>
        </div>
        <li
            :class="{
          'env-item':true,
          'env-item-footer':true,
          'env-item-active':isShowAddEnv}"
            @click="showEnvDetail(null,true)">
          <a-button
              class="btn"
              @click="addEnv"
              type="text">
            <template #icon>
              <PlusOutlined/>
            </template>
            新建环境
          </a-button>
        </li>
      </ul>
    </div>
    <div class="right-content">
      <!-- ::::全局变量 -->
      <div class="globalVars" v-if="isShowGlobalVars">
        <div class="title">全局变量</div>
        <a-table :columns="globalVarsColumns" :data-source="globalVarsData">
          <template #customName="{ text }">
            <a-input :value="text" placeholder="请输入参数名"/>
          </template>
        </a-table>
      </div>
      <!-- ::::全局参数 -->
      <div class="globalParams" v-if="isShowGlobalParams">
        <div class="title">全局参数</div>
        <a-tabs v-model:activeKey="activeKey">
          <a-tab-pane key="1" tab="Header">
            <a-table :columns="globalParamscolumns" :data-source="globalParamsData?.header || []">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customRequired="{ text }">
                {{ text ? '是' : '否' }}
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="2" tab="Cookie">
            <a-table :columns="globalParamscolumns" :data-source="globalParamsData?.cookie || []">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customRequired="{ text }">
                {{ text ? '是' : '否' }}
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="3" tab="Query">
            <a-table :columns="globalParamscolumns" :data-source="globalParamsData?.query || []">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customRequired="{ text }">
                {{ text ? '是' : '否' }}
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="4" tab="Body">
            <a-table :columns="globalParamscolumns" :data-source="globalParamsData?.body || []">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customRequired="{ text }">
                {{ text ? '是' : '否' }}
              </template>
            </a-table>
          </a-tab-pane>
        </a-tabs>

      </div>
      <!-- ::::环境详情 -->
      <div class="envDetail" v-if="isShowEnvDetail && activeEnvDetail">
        <div class="title">{{ activeEnvDetail.name }}</div>
        <div class="envDetail-content">
          <a-form-item :labelCol="{span: 4}" :wrapperCol="{span: 14}" label="环境名称">
            <a-input :value="activeEnvDetail.name"/>
          </a-form-item>
          <div class="serveServers">
            <div class="serveServers-header">服务 (前置URL)</div>
            <a-table :columns="serveServersColumns" :data-source="activeEnvDetail.serveServers">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customUrl="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
            </a-table>
          </div>
          <div class="vars">
            <div class="vars-header">环境变量</div>
            <a-table :columns="globalVarsColumns" :data-source="activeEnvDetail.vars">
              <template #customName="{ text }">
                <a-input :value="text" placeholder="请输入参数名"/>
              </template>
            </a-table>
          </div>

        </div>
        <div class="envDetail-footer">
          <a-button class="save-btn" @click="addEnvData" type="primary">保存</a-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">

import {computed, defineComponent, defineEmits, defineProps, onMounted, reactive, Ref, ref, UnwrapRef} from 'vue';
import {CheckOutlined, EditOutlined, MenuOutlined, BorderlessTableOutlined, PlusOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';
import {
  getServeList,
  deleteServe,
  copyServe,
  disableServe,
  getGlobalVarsList,
  getEnvironmentsParamList,
  getEnvList,
} from '../../service';
import {momentUtc} from '@/utils/datetime';
import {message} from "ant-design-vue";
import {serveStatus} from "@/config/constant";
import draggable from 'vuedraggable'
import Combobox from "ant-design-vue/es/vc-select/examples/combobox";
import mounted = Combobox.mounted;

const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);


/*************************************************
 * ::::环境列表 管理 start
 ************************************************/
const envList = ref([]);
const isShowEnvDetail = ref(false);
const isShowAddEnv = ref(false);
const activeEnvDetail: any = ref(null);
const serveServersColumns: any = [
  {
    title: '服务名',
    dataIndex: 'name',
    key: 'name',
    slots: {customRender: 'customName'},
  },
  {
    title: '前置 URL ',
    dataIndex: 'url',
    key: 'url',
    slots: {customRender: 'customUrl'},
  },
];
// 请求环境列表
onMounted(async () => {
  const res = await getEnvList({
    projectId: 1
  });
  if (res.code === 0) {
    envList.value = res.data;
  }
})

function showEnvDetail(item, isAdd?: boolean) {
  if (isAdd) {
    isShowAddEnv.value = true;
    isShowEnvDetail.value = true;
    activeEnvDetail.value = {
      name: "",
      serveServers: [
        {
          "id": 2,
          // "createdAt": "2023-03-08T14:41:31+08:00",
          // "updatedAt": "2023-03-08T14:41:31+08:00",
          // "environmentId": 7,
          // "serveId": 1,
          "name": "",
          "url": ""
        }
      ],
      vars: [
        {
          "id": 1,
          "name": "",
          "rightValue": "",
          "localValue": "",
          "remoteValue": "",
          // "environmentId": 7,
          // "projectId": 0
        }
      ],
    };
  } else {
    isShowEnvDetail.value = true;
    isShowAddEnv.value = false;
    activeEnvDetail.value = item;
  }
  isShowGlobalParams.value = false;
  isShowGlobalVars.value = false;

}

/*************************************************
 * ::::环境列表 管理 end
 ************************************************/


/*************************************************
 * ::::新建环境 start
 ************************************************/
const addEnvVisible = ref(false);

function addEnv() {
  addEnvVisible.value = false;
}

function addEnvData() {
  console.log('保存数据');
}

function onCloseEnvDrawer() {
  addEnvVisible.value = false
}

/*************************************************
 * ::::新建环境  end
 ************************************************/




// ::::
/*************************************************
 * ::::全局参数 + 全局变量  逻辑 start
 ************************************************/
const activeKey = ref('1');
const isShowGlobalParams: any = ref(false);
const globalParamscolumns: any = [
  //    "id": 6,
  //   "createdAt": "2023-03-09T16:11:46+08:00",
  //   "updatedAt": "2023-03-09T16:11:46+08:00",
  //   "name": "var1",
  //   "type": "string",
  //   "required": false,
  //   "defaultValue": "1",
  //   "description": "var1",
  //   "in": "body",
  //   "projectId": 1
  {
    title: '参数名',
    dataIndex: 'name',
    key: 'name',
    slots: {customRender: 'customName'},
  },
  {
    title: '类型',
    dataIndex: 'type',
    key: 'type',
  },
  {
    title: '必须',
    dataIndex: 'required',
    key: 'required',
    slots: {customRender: 'customRequired'},
  },
  {
    title: '默认值',
    key: 'defaultValue',
    dataIndex: 'defaultValue',
  },
  // {
  //   title: '默认启用',
  //   key: 'description',
  //   dataIndex: 'description',
  // },
  {
    title: '说明',
    key: 'description',
    dataIndex: 'description',
  },
];
const globalParamsData: any = ref({
  header: [],
  cookie: [],
  body: [],
  query: []
});
const isShowGlobalVars = ref(false);
const globalVarsColumns = [
  {
    title: '变量名',
    dataIndex: 'name',
    key: 'name',
    slots: {customRender: 'customName'},
  },
  {
    title: '远程值',
    dataIndex: 'remoteValue',
    key: 'remoteValue',
  },
  {
    title: '本地值',
    dataIndex: 'localValue',
    key: 'localValue',
  },
  {
    title: '说明',
    key: 'description',
    dataIndex: 'description',
    slots: {customRender: 'description'},
  },
];
const globalVarsData = ref([]);

async function showGlobalParams() {
  isShowGlobalParams.value = true;
  isShowGlobalVars.value = false;
  isShowAddEnv.value = false;
  isShowEnvDetail.value = false;
  const res = await getEnvironmentsParamList({
    projectId: 1
  });
  if (res.code === 0) {
    globalParamsData.value = res.data;
  }
}

async function showGlobalVars() {
  isShowGlobalParams.value = false;
  isShowGlobalVars.value = true;
  isShowAddEnv.value = false;
  isShowEnvDetail.value = false;
  const res = await getGlobalVarsList({
    projectId: 1
  });
  if (res.code === 0) {
    globalVarsData.value = res.data;
  }
}

/*************************************************
 * ::::全局参数 + 全局变量 逻辑 end
 ************************************************/

</script>

<style scoped lang="less">
.content {
  display: flex;
}

.left-content {
  min-height: calc(100vh - 96px);
  height: 100%;
  background-color: #F9F9F9;
  width: 240px;

  ul {
    list-style: none;
    margin: 0;
    padding: 0;

    li {
      padding: 0;
      margin: 0;
      height: 36px;
      line-height: 36px;
      list-style: none;
      cursor: pointer;
    }

    .header {
      margin-left: 16px;
      margin-top: 16px;
      opacity: 0.5;
    }

    .footer {

    }
  }

  .env-item {
    margin: 0 16px;
    padding: 0 16px;

    i {
      width: 18px;
      height: 18px;
      background-color: #515152;
      border-radius: 100%;
      color: #ffffff;
      font-size: 10px;
      text-align: center;
      line-height: 18px;
      display: inline-block;
      margin-right: 8px;
    }

    &.env-item-active {
      background-color: #1aa391;
      border-radius: 6px;
    }

    &:hover {
      background-color: #1aa391;
      border-radius: 6px;
    }

    .handle {
      margin-right: 8px;
      cursor: pointer;
    }
  }

  .divider {
    margin: 16px 0;
  }

  .env-item-footer {
    padding: 0;
  }
}

.right-content {
  flex: 1;
  min-height: calc(100vh - 116px);
  margin: 16px;

  .globalVars, .globalParams, .envDetail {
    .title {
      font-weight: bold;
      font-size: 18px;
      margin-bottom: 16px;
    }
  }

  .envDetail-content {
    position: relative;
  }

  .envDetail-footer {
    height: 60px;
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .save-btn {
      margin-right: 16px;
    }
  }
}

.vars-header, .serveServers-header {
  padding: 0 0 8px;
  line-height: 1.4;
  white-space: normal;
  text-align: left;
  margin-bottom: 8px;
}


</style>
