<template>
  <div class="datapool-main">
    <!-- header -->
    <div class="header">
      <a-form layout="inline">
        <a-form-item>
          <a-button class="editable-add-btn" type="primary" html-type="submit" style="margin-bottom: 8px"
                    @click="onCreate">
            新建
          </a-button>
        </a-form-item>
      </a-form>

      <a-input-search placeholder="输入数据池名称搜索" style="width: 200px"
                      v-model:value="keyword"
                      @search="handleSearch"/>
    </div>

    <!-- content -->
    <EmptyCom>
      <template #content>
        <a-table :data-source="dataSource" :columns="datapoolColumns" :rowKey="(_record, index) => index">
          <template #name="{ text, record }">
            <div class="serve-name">
              <EditAndShowField :custom-class="'custom-serve show-on-hover'" placeholder="请输入服务名称"
                                :value="text || ''"
                                @update="(e: string) => handleUpdateName(e, record)" @edit="onEdit(record)"/>
            </div>
          </template>

          <template #customStatus="{ text, record }">
            <a-tag :color="record.statusTag">{{ text }}</a-tag>
          </template>

          <template #operation="{ record }">
            <a-dropdown>
              <MoreOutlined/>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="onDisable(record)">禁用</a>
                  </a-menu-item>
                  <a-menu-item key="2">
                    <a class="operation-a" href="javascript:void (0)" @click="onDelete(record)">删除</a>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>
        </a-table>
      </template>
    </EmptyCom>

    <!-- 抽屉 -->
    <Drawer :edit-key="editKey"
            :drawer-visible="drawerVisible"
            @onClose="onClose"/>

  </div>
</template>
<script setup lang="ts">

import {computed, createVNode, ref, watch,} from 'vue';
import {useStore} from "vuex";
import {useRouter} from 'vue-router';
import {Modal} from 'ant-design-vue';
import {ExclamationCircleOutlined, MoreOutlined} from '@ant-design/icons-vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EmptyCom from '@/components/Empty/index.vue';
import Drawer from './Drawer.vue';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {datapoolColumns} from '../../config';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSetting.serviceOptions);
const route = useRouter();

const drawerVisible = ref(false);
const editKey = ref(0);
const keyword = ref('')

let formConfig = ref([
  {
    type: 'tooltip',
    title: '一个产品服务端通常对应一个或多个服务(微服务)，服务可以有多个版本并行，新的服务默认起始版本为v0.1.0。',
    text: '新建服务'
  },
  {
    type: 'input',
    modelName: 'name',
    placeholder: '服务名称',
    valueType: 'string'
  },
  {
    type: 'button',
    text: '新建'
  }
]);

const rules = {
  name: [
    {
      required: true,
      message: '服务名称不能为空'
    }
  ]
};

watch((currProject), async (newVal) => {
  await getList()
}, {
  immediate: true
})

async function getList(name = '') {
  await store.dispatch('ProjectSetting/getServersList', {
    projectId: currProject.value.id,
    page: 0,
    pageSize: 100,
    name
  })
}

function onClose() {
  drawerVisible.value = false;
}

function handleSearch(value: any) {
  getList(value);
}

function handleUpdateName(value: string, record: any) {
  const serviceInfo = {name: value, description: record.description, id: record.id};
  store.dispatch('ProjectSetting/saveStoreServe', {
    "projectId": currProject.value.id,
    formState: {...serviceInfo},
    action: 'update'
  });
}

const onCreate = () => {
  console.log('onCreate')

  editKey.value++;
  drawerVisible.value = true;
}

const onEdit = () => {
  console.log('onEdit')

  editKey.value++;
  drawerVisible.value = true;
}

async function onDelete(record: any) {
  Modal.confirm({
    title: '确认要删除该服务吗',
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      store.dispatch('ProjectSetting/deleteStoreServe', {id: record.id, projectId: currProject.value.id});
    }
  })
}

async function onDisable(record: any) {
  store.dispatch('ProjectSetting/disabledStoreServe', {id: record.id, projectId: currProject.value.id});
}

</script>

<style scoped lang="less">
.datapool-main {
  margin: 20px;

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;
  }
}

.drawer-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.operation-a {
  text-align: center;
  display: inline-block;
  width: 80px;
}

.serve-name {
  width: 120px;
}

.serve-description {
  max-width: 140px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
}
</style>
