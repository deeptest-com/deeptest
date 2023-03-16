<template>
  <div class="content">
    <div class="left-content">
      <div class="global">
        <div class="header">全局</div>
        <a-button :class="{'env-item':true,'env-item-active':isShowGlobalVars}"
                  @click="showGlobalVars"
                  :type="isShowGlobalVars ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">V</i>
          </template>
          全局变量
        </a-button>
        <a-button :class="{'env-item':true,'env-item-active':isShowGlobalParams}"
                  @click="showGlobalParams"
                  :type="isShowGlobalParams ? 'primary' : 'text'">
          <template #icon>
            <i class="var-icon">P</i>
          </template>
          全局参数
        </a-button>
      </div>
      <div style="margin: 0 16px;">
        <a-divider class="divider"/>
      </div>
      <div class="env">
        <div class="header">环境</div>
        <draggable
            tag="div"
            :list="envList"
            class="list-group"
            handle=".handle"
            item-key="name">
          <template #item="{ element, index }">
            <a-button
                :class="{'env-item':true,'env-item-active':activeEnvDetail?.name === element.name}"
                :type="activeEnvDetail?.name === element.name ? 'primary' : 'text'"
                @click="showEnvDetail(element)"
                class="env-item" :key="index">
              <MenuOutlined class="handle"/>
              <span class="text"> {{ element.displayName }} </span>
            </a-button>
          </template>
        </draggable>
        <div style="margin: 0 16px;">
          <a-divider class="divider"/>
        </div>
        <a-button
            :type="isShowAddEnv ? 'primary' : 'text'"
            :class="{
            'env-item':true,
            'env-item-footer':true,
            'env-item-active':isShowAddEnv}"
            @click="showEnvDetail(null,true)">
          <template #icon>
            <PlusOutlined/>
          </template>
          新建环境
        </a-button>
      </div>
    </div>
    <div class="right-content">
      <!-- ::::全局变量 -->
      <div class="globalVars" v-if="isShowGlobalVars">
        <div class="title">全局变量</div>
        <a-button
            class="envDetail-btn"
            @click="addGlobalVar"
        >
          <template #icon>
            <PlusOutlined/>
          </template>
          添加
        </a-button>

        <a-table bordered size="small" :pagination="false"
                 :columns="globalVarsColumns"
                 :data-source="globalVarsData">
          <template #customName="{ text,index }">
            <a-input @change="(e) => {
                handleGlobalVarsChange('name',index,e);
                }" :value="text" placeholder="请输入参数名"/>
          </template>
          <template #customLocalValue="{ text,index }">
            <a-input :value="text" @change="(e) => {
              handleGlobalVarsChange('localValue',index,e);
              }" placeholder="请输入本地值"/>
          </template>
          <template #customRemoteValue="{ text,index }">
            <a-input :value="text" @change="(e) => {
            handleGlobalVarsChange('remoteValue',index,e);
            }" placeholder="请输入远程值"/>
          </template>
          <template #customDescription="{ text,index }">
            <a-input :value="text" @change="(e) => {
              handleGlobalVarsChange('description',index,e);
            }" placeholder="说明"/>
          </template>
          <template #customAction="{index }">
            <a-button danger
                      type="text"
                      @click="handleGlobalVarsChange('description',index,'','delete');"
                      :size="'small'">删除
            </a-button>
          </template>
        </a-table>

        <div class="envDetail-footer">
          <a-button class="save-btn" @click="handleSaveGlobalVars" type="primary">保存</a-button>
        </div>
      </div>
      <!-- ::::全局参数 -->
      <div class="globalParams" v-if="isShowGlobalParams">
        <div class="title">全局参数</div>

        <a-tabs :pagination="false" v-model:activeKey="globalParamsActiveKey">
          <a-tab-pane key="header" tab="Header">

            <a-button
                class="envDetail-btn"
                @click="addGlobalParams"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>

            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                     :data-source="globalParamsData?.header || []">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('header','name',index,e);
              }" placeholder="请输入参数名"/>
              </template>
              <template #customType="{ text,index }">
                <a-select
                    class="custom-select"
                    :value="text"
                    style="width: 120px"
                    @change="(e) => {
                    handleGlobalParamsChange('header','type',index,e)
                    }">
                  <a-select-option value="string">string</a-select-option>
                  <a-select-option value="number">number</a-select-option>
                  <a-select-option value="integer">integer</a-select-option>
                </a-select>
              </template>
              <template #customRequired="{ text,index }">
                <a-switch :checked="text" @change="(checked) => {
                    handleGlobalParamsChange('header','required',index,checked)
                    }"/>
              </template>
              <template #customDefaultValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('header','defaultValue',index,e);
                }" placeholder="默认值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('header','description',index,e);
              }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleGlobalParamsChange('header','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="cookie" tab="Cookie">
            <a-button
                class="envDetail-btn"
                @click="addGlobalParams"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>
            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                     :data-source="globalParamsData?.cookie || []">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('cookie','name',index,e);
              }" placeholder="请输入参数名"/>
              </template>
              <template #customType="{ text,index }">
                <a-select
                    ref="select"
                    class="custom-select"
                    :value="text"
                    style="width: 120px"
                    @change="(e) => {
                    handleGlobalParamsChange('cookie','type',index,e)
                    }">
                  <a-select-option value="string">string</a-select-option>
                  <a-select-option value="number">number</a-select-option>
                  <a-select-option value="integer">integer</a-select-option>
                </a-select>
              </template>
              <template #customRequired="{ text,index }">
                <a-switch :checked="text" @change="(checked) => {
                    handleGlobalParamsChange('cookie','required',index,checked)
                    }"/>
              </template>
              <template #customDefaultValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('cookie','defaultValue',index,e);
                }" placeholder="默认值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('cookie','description',index,e);
              }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleGlobalParamsChange('cookie','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="query" tab="Query">
            <a-button
                class="envDetail-btn"
                @click="addGlobalParams"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>
            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                     :data-source="globalParamsData?.query || []">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('query','name',index,e);
              }" placeholder="请输入参数名"/>
              </template>
              <template #customType="{ text,index }">
                <a-select
                    ref="select"
                    :value="text"
                    class="custom-select"
                    style="width: 120px"
                    @change="(e) => {
                    handleGlobalParamsChange('query','type',index,e)
                    }">
                  <a-select-option value="string">string</a-select-option>
                  <a-select-option value="number">number</a-select-option>
                  <a-select-option value="integer">integer</a-select-option>
                </a-select>
              </template>
              <template #customRequired="{ text,index }">
                <a-switch :checked="text" @change="(checked) => {
                    handleGlobalParamsChange('query','required',index,checked)
                    }"/>
              </template>
              <template #customDefaultValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('query','defaultValue',index,e);
                }" placeholder="默认值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('query','description',index,e);
              }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleGlobalParamsChange('cookie','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </a-tab-pane>
          <a-tab-pane key="body" tab="Body">
            <a-button
                class="envDetail-btn"
                @click="addGlobalParams"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>
            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                     :data-source="globalParamsData?.body || []">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('body','name',index,e);
              }" placeholder="请输入参数名"/>
              </template>
              <template #customType="{ text,index }">
                <a-select
                    ref="select"
                    :value="text"
                    class="custom-select"
                    style="width: 120px"
                    @change="(e) => {
                    handleGlobalParamsChange('cookie','type',index,e)
                    }">
                  <a-select-option value="string">string</a-select-option>
                  <a-select-option value="number">number</a-select-option>
                  <a-select-option value="integer">integer</a-select-option>
                </a-select>
              </template>
              <template #customRequired="{ text,index }">
                <a-switch :checked="text" @change="(checked) => {
                    handleGlobalParamsChange('body','required',index,checked)
                    }"/>
              </template>
              <template #customDefaultValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('body','defaultValue',index,e);
                }" placeholder="默认值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
                  handleGlobalParamsChange('body','description',index,e);
              }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleGlobalParamsChange('cookie','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
              </template>
            </a-table>
          </a-tab-pane>
        </a-tabs>

        <div class="envDetail-footer">
          <a-button class="save-btn" @click="handleSaveGlobalParams" type="primary">保存</a-button>
        </div>
      </div>
      <!-- ::::环境详情 -->
      <div class="envDetail" v-if="isShowEnvDetail && activeEnvDetail">
        <div class="title">{{ activeEnvDetail.displayName }}</div>
        <div class="envDetail-content">
          <a-form-item :labelCol="{span: 2}" :wrapperCol="{span: 10}" label="环境名称">
            <a-input class="env-name" :value="activeEnvDetail.name || ''" @change="handleEnvNameChange" placeholder="请输入环境名称"/>
          </a-form-item>
          <div class="serveServers">
            <div class="serveServers-header">服务 (前置URL)</div>
            <a-button
                class="envDetail-btn"
                @click="addService"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              关联服务
            </a-button>
            <a-table v-if="activeEnvDetail.serveServers.length > 0"
                     size="small"
                     bordered :pagination="false"
                     :columns="serveServersColumns"
                     :data-source="activeEnvDetail.serveServers">
              <template #customName="{ text,index }">
                <a-input :value="text" @change="(e) => {
                handleEnvChange('serveServers','serveName',index,e);
                }" placeholder="请输入参数名"/>
              </template>

              <template #customUrl="{ text,index }">
                <a-input :value="text"
                         @change="(e) => {
                            handleEnvChange('serveServers','url',index,e);
                         }"
                         placeholder="http 或 https 起始的合法 URL"/>
              </template>


            </a-table>

          </div>
          <div class="vars">
            <div class="vars-header">环境变量</div>
            <a-button
                class="envDetail-btn"
                @click="addVar"
            >
              <template #icon>
                <PlusOutlined/>
              </template>
              添加
            </a-button>
            <a-table
                v-if="activeEnvDetail.vars.length > 0"
                bordered size="small"
                :pagination="false"
                :columns="globalVarsColumns"
                :data-source="activeEnvDetail.vars">
              <template #customName="{ text,index }">
                <a-input @change="(e) => {
                handleEnvChange('vars','name',index,e);
                }" :value="text" placeholder="请输入参数名"/>
              </template>
              <template #customLocalValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
              handleEnvChange('vars','localValue',index,e);
              }" placeholder="请输入本地值"/>
              </template>
              <template #customRemoteValue="{ text,index }">
                <a-input :value="text" @change="(e) => {
            handleEnvChange('vars','remoteValue',index,e);
            }" placeholder="请输入远程值"/>
              </template>
              <template #customDescription="{ text,index }">
                <a-input :value="text" @change="(e) => {
              handleEnvChange('vars','description',index,e);
            }" placeholder="说明"/>
              </template>
              <template #customAction="{index }">
                <a-button danger
                          type="text"
                          @click="handleEnvChange('vars','',index,'','delete');"
                          :size="'small'">删除
                </a-button>
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
  <a-modal
      v-model:visible="addServiceModalVisible"
      title="关联服务"
      @ok="handleAddServiceOk">
    <a-form-item class="select-service" :labelCol="{span: 6}" :wrapperCol="{span: 16}" label="请选择服务">
      <a-select
          v-model:value="selectedService"
          :options="serviceOptions"
          placeholder="请选择服务"
          style="width: 200px"/>
    </a-form-item>

  </a-modal>
</template>
<script setup lang="ts">

import {
  computed,
  defineComponent,
  defineEmits,
  defineProps,
  onMounted,
  reactive,
  Ref,
  ref,
  UnwrapRef,
  watch
} from 'vue';
import {CheckOutlined, EditOutlined, MenuOutlined, BorderlessTableOutlined, PlusOutlined} from '@ant-design/icons-vue';

import {
  getServeList,
  deleteServe,
  copyServe,
  disableServe,
  getGlobalVarsList,
  getEnvironmentsParamList,
  getEnvList, saveEnv,
  saveGlobalVars,
  saveEnvironmentsParam,
} from '../../service';
import {momentUtc} from '@/utils/datetime';
import {message} from "ant-design-vue";
import {serveStatus} from "@/config/constant";
import draggable from 'vuedraggable'
import Combobox from "ant-design-vue/es/vc-select/examples/combobox";
import mounted = Combobox.mounted;

const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);


import {StateType as ProjectStateType} from "@/store/project";
import {useStore} from "vuex";

const store = useStore<{ ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

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
    dataIndex: 'serveName',
    key: 'serveName',
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
async function getEnvsList() {
  const res = await getEnvList({
    projectId: currProject.value.id
  });
  res.data.forEach((item) => {
    item.displayName = item.name;
  })
  if (res.code === 0) {
    envList.value = res.data;
  }
}

function showEnvDetail(item, isAdd?: boolean) {
  if (isAdd) {
    isShowAddEnv.value = true;
    isShowEnvDetail.value = true;
    activeEnvDetail.value = {
      displayName: "新建环境",
      name: "",
      serveServers: [],
      vars: [],
    };
  } else {
    isShowEnvDetail.value = true;
    isShowAddEnv.value = false;
    activeEnvDetail.value = item;
    activeEnvDetail.value.name = item.name || '';
    activeEnvDetail.value.displayName = item.name || '';
  }
  isShowGlobalParams.value = false;
  isShowGlobalVars.value = false;
}

const addServiceModalVisible = ref(false);
const serviceOptions = ref([]);
const selectedService = ref('');

async function addService() {
  addServiceModalVisible.value = true;
}

function addVar() {
  activeEnvDetail.value.vars.push({
    "name": "",
    "rightValue": "",
    "localValue": "",
    "remoteValue": "",
    // "environmentId": 7
  })
}


async function getServersList() {
  // 请求服务列表
  let res = await getServeList({
    projectId: currProject.value.id,
    "page": 0,
    "pageSize": 100,
  });
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.label = item.name;
      item.value = item.id;
    })
    serviceOptions.value = res.data.result;
  }
}

function handleAddServiceOk() {
  addServiceModalVisible.value = false;
  const selectServe: any = serviceOptions.value.find((item: any) => {
    return selectedService.value === item.id;
  })
  activeEnvDetail.value.serveServers.push({
    // "environmentId": 7,
    "url": "",
    "serveName":selectServe.name,
    "serveId":selectServe.id,
  })
}

async function addEnvData() {
  // const serveServers = activeEnvDetail.value?.serveServers.
  let res = await saveEnv({
    id: activeEnvDetail.value?.id,
    projectId: currProject.value.id,
    name: activeEnvDetail.value?.name,
    "serveServers": activeEnvDetail.value?.serveServers || [],
    "vars": activeEnvDetail.value?.vars || [],
  });
  if (res.code === 0) {
    message.success('保存环境成功')
  }
}

function handleEnvChange(type, field, index, e, action) {
  if (action === 'delete') {
    activeEnvDetail.value[type].splice(index, 1);
  } else {
    activeEnvDetail.value[type][index][field] = e.target.value;
  }

}

function handleEnvNameChange(e) {
  activeEnvDetail.value.name = e.target.value;
}

/*************************************************
 * ::::环境列表 管理 end
 ************************************************/


/*************************************************
 * ::::全局参数+全局变量  逻辑 start
 ************************************************/
const globalParamsActiveKey = ref('header');
const isShowGlobalParams: any = ref(false);
const globalParamscolumns: any = [
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
    slots: {customRender: 'customType'},
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
    slots: {customRender: 'customDefaultValue'},
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
    slots: {customRender: 'customDescription'},
  },
  {
    title: '操作',
    key: 'action',
    slots: {customRender: 'customAction'},
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
    slots: {customRender: 'customRemoteValue'},
  },
  {
    title: '本地值',
    dataIndex: 'localValue',
    key: 'localValue',
    slots: {customRender: 'customLocalValue'},
  },
  {
    title: '说明',
    key: 'description',
    dataIndex: 'description',
    slots: {customRender: 'customDescription'},
  },
  {
    title: '操作',
    key: 'action',
    slots: {customRender: 'customAction'},
  },
];
const globalVarsData: any = ref([]);

async function showGlobalParams() {
  isShowGlobalParams.value = true;
  isShowGlobalVars.value = false;
  isShowAddEnv.value = false;
  isShowEnvDetail.value = false;

  activeEnvDetail.value = null;
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
    projectId: currProject.value.id
  });
  if (res.code === 0) {
    globalVarsData.value = res.data;
  }
}

function addGlobalVar() {
  globalVarsData.value.push({
    "name": "",
    "rightValue": "",
    "localValue": "",
    "remoteValue": ""
  })
}

function addGlobalParams() {
  globalParamsData.value[globalParamsActiveKey.value].push({
    "name": "",
    "type": "string",
    "defaultValue": "",
    "description": "",
    "required": false
  })
}

async function handleSaveGlobalParams() {
  let res = await saveEnvironmentsParam(globalParamsData.value);
  if (res.code === 0) {
    message.success('保存全局参数成功');
  }
}

async function handleSaveGlobalVars() {
  let res = await saveGlobalVars(globalVarsData.value);
  if (res.code === 0) {
    message.success('保存全局变量成功');
  }
}

function handleGlobalVarsChange(field, index, e, action?: string) {
  // 删除
  if (action === 'delete') {
    globalVarsData.value.splice(index, 1);
  } else {
    globalVarsData.value[index][field] = e.target.value;
  }
}

function handleGlobalParamsChange(type, field, index, e, action?: string) {
  if (action === 'delete') {
    globalParamsData.value[type].splice(index, 1);
  } else {
    globalParamsData.value[type][index][field] = ["string", "boolean"].includes(typeof e) ? e : e.target.value;
  }
}

/*************************************************
 * ::::::::全局参数+全局变量 + 全局变量 逻辑 end
 ************************************************/



// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await getServersList();
  await getEnvsList();
  // 默认展示全局变量
  await showGlobalVars();
}, {
  immediate: true
})


</script>
<style scoped lang="less">
.content {
  display: flex;
}

.left-content {
  height: calc(100vh - 138px);
  overflow-y: scroll;
  background-color: #F9F9F9;
  width: 240px;

  div {
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
      margin-bottom: 16px;
      opacity: 0.5;
    }

    .footer {

    }
  }

  button.env-item {
    width: 200px;
    text-align: left;
    margin-bottom: 8px;
    border-radius: 6px;
  }

  .env-item {
    margin: 0 16px;
    //padding: 0 16px;
    padding-left: 8px;
    //display: inline-block;
    //&:hover{
    //  color: #1677ff;
    //}

    i {
      width: 18px;
      height: 18px;
      background-color: #515152;
      border-radius: 100%;
      color: #ffffff;
      font-size: 10px;
      text-align: center;
      line-height: 18px;
      margin-right: 8px;
    }


    .handle {
      margin-right: 8px;
      cursor: pointer;
    }
  }

  .divider {
    margin: 16px 0;
  }

  //.env-item-footer {
  //  padding: 0;
  //}
}

.right-content {
  flex: 1;
  height: calc(100vh - 138px);
  overflow-y: scroll;
  position: relative;
  //margin: 16px;
  padding: 16px;

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
    top: 0;
    right: 0;
    width: 100px;
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

.var-icon {
  display: inline-block;
}

.envDetail-btn {
  margin-top: 16px;
  margin-bottom: 16px;
}

.serveServers-header, .vars-header {
  font-weight: bold;
  margin-bottom: 0;
  margin-top: 16px;
}


.select-service {
  .ant-select-selector {
    border: 1px solid #d9d9d9;
  }
}

::v-deep {

  .custom-select {
    .ant-select-selector {
      border: 1px solid transparent !important;
    }

    .ant-select-selector:hover, .ant-select-selector:active, .ant-select-selector:focus {
      border: 1px solid #4096ff !important
    }
  }


  .ant-input:not(.env-name) {
    border: 1px solid transparent !important;
  }

  .ant-input:not(.env-name):hover, .ant-input:active, .ant-input:focus {
    border: 1px solid #4096ff !important
  }
}

</style>
