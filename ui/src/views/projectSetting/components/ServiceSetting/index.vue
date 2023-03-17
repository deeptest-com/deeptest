<template>
  <div class="content">
    <div class="header">
      <a-form layout="inline" :model="formState">
        <a-form-item>
            <span style="cursor: pointer;font-weight: bold">新建服务
              <a-tooltip placement="topLeft" arrow-point-at-center
                         title="一个产品服务端通常对应一个或多个服务(微服务)，服务可以有多个版本并行，新的服务默认起始版本为v0.1.0。">
              <QuestionCircleOutlined class="icon"
                                      style="position: relative;top:-8px; font-size: 12px;transform: scale(0.9)"/>
              </a-tooltip>
            </span>
        </a-form-item>
        <a-form-item>
          <a-input v-model:value="formState.name" placeholder="服务名称"/>
        </a-form-item>
        <a-form-item>
          <a-select
              v-model:value="formState.userId"
              show-search
              placeholder="负责人(默认创建人)"
              style="width: 200px"
              :options="userListOptions"
              @focus="selectUserFocus">
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-input v-if="editServiceDesc"
                   v-model:value="formState.description"
                   placeholder="输入描述"/>
        </a-form-item>
        <a-form-item>
          <a-button class="editable-add-btn"
                    @click="handleOk"
                    type="primary"
                    style="margin-bottom: 8px">
            确定
          </a-button>
        </a-form-item>
      </a-form>


      <a-input-search
          v-model:value="keyword"
          placeholder="输入服务名称搜索"
          style="width: 300px"
          @search="onSearch"
      />
    </div>
    <a-table  :data-source="dataSource" :columns="columns" rowKey="id">

      <template #name="{ text, record }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            {{ text || ' ' }}
            <edit-outlined class="editable-cell-icon" @click="edit(record)"/>
          </div>
        </div>
      </template>
      <template #customServers="{ record }">
        <span v-if="record?.servers.length > 0">
               <span  v-for="server in record.servers" :key="server.id">
          {{server.name || server.description}}
        </span>
        </span>
        <span v-else>暂无关联服务</span>
      </template>
      <template #customStatus="{ text,record }">
        <a-tag :color="record.statusTag">{{ text }}</a-tag>
      </template>
      <template #operation="{ record }">
        <a-dropdown>
          <MoreOutlined/>
          <template #overlay>
            <a-menu>
              <a-menu-item key="1">
                <a class="operation-a"  href="javascript:void (0)" @click="onDisabled(record)">禁用</a>
              </a-menu-item>
              <a-menu-item key="2">
                <a  class="operation-a"  href="javascript:void (0)" @click="onCopy(record)">复制</a>
              </a-menu-item>
              <a-menu-item key="3">
                <a  class="operation-a"  href="javascript:void (0)" @click="onDelete(record)">删除</a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>
    </a-table>
    <!-- ::::新建服务弹框 -->
    <a-modal v-model:visible="visible"
             @cancel="handleCancel"
             title="新建服务"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="服务名称">
          <a-input v-model:value="formState.name" placeholder="请输入内容"/>
        </a-form-item>
        <a-form-item label="描述">
          <a-input v-if="editServiceDesc" v-model:value="formState.description" placeholder="请输入内容"/>
        </a-form-item>
      </a-form>
    </a-modal>

    <a-drawer
        :closable="true"
        :width="1000"
        :key="editKey"
        :visible="drawerVisible"
        @close="onClose"
    >
      <template #title>
        <div class="drawer-header">
          <div>服务编辑</div>
        </div>
      </template>
      <div class="drawer-content">
        <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col=" { span: 15 }">
          <a-form-item label="服务名称">
            <a-input
                v-if="isEditServiceName"
                @focusout="changeServiceInfo"
                @pressEnter="changeServiceInfo"
                v-model:value="editFormState.name"
                placeholder="请输入内容"/>
            <span v-else>{{ editFormState.name }}
              <edit-outlined class="editable-cell-icon" @click="editServiceName"/></span>
          </a-form-item>
          <a-form-item label="描述">
            <a-input
                @focusout="changeServiceInfo"
                @pressEnter="changeServiceInfo"
                v-if="isEditServiceDesc"
                v-model:value="editFormState.description"
                placeholder="请输入内容"/>
            <span v-if="!isEditServiceDesc">{{ editFormState.description }}
              <edit-outlined class="editable-cell-icon" @click="editServiceDesc"/> </span>
          </a-form-item>
          <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1" tab="服务版本">
              <ServiceVersion :serveId="editFormState.serveId"/>
            </a-tab-pane>
            <a-tab-pane key="2" tab="服务组件">
              <ServiceComponent :serveId="editFormState.serveId"/>
            </a-tab-pane>
          </a-tabs>
        </a-form>
      </div>
    </a-drawer>

  </div>
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
import {CheckOutlined, EditOutlined, ExclamationOutlined, QuestionCircleOutlined,MoreOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';

import {getServeList, deleteServe, copyServe, disableServe, saveServe, getUserList,} from '../../service';
import {momentUtc} from '@/utils/datetime';
import {message} from "ant-design-vue";
import {serveStatus,serveStatusTagColor} from "@/config/constant";
import {StateType as ProjectStateType} from "@/store/project";
import {useStore} from "vuex";
import {SelectTypes} from "ant-design-vue/es/select";

const store = useStore<{ ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  description: string;
  serveId?: string,


}

interface DataItem {
  key: string;
  name: string;
  age: number;
  address: string;
}


const formState: UnwrapRef<any> = reactive({
  name: '',
  description: '',
  userId: null,
});

const editFormState: UnwrapRef<any> = reactive({
  name: '',
  description: '',
  serveId: '',
});


const visible = ref(false);
const drawerVisible = ref(false);

const columns = [
  {
    title: '服务名称',
    dataIndex: 'name',
    slots: {customRender: 'name',title:'fdshfh'},
  },
  {
    title: '描述',
    dataIndex: 'description',
  },
  {
    title: '关联服务',
    dataIndex: 'servers',
    slots: {customRender: 'customServers'},
  },
  {
    title: '状态',
    dataIndex: 'statusDesc',
    slots: {customRender: 'customStatus'},
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
  },
  {
    title: '最近更新时间',
    dataIndex: 'updatedAt',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];
const dataSource: Ref<DataItem[]> = ref([]);
const count = computed(() => dataSource.value.length + 1);
const editKey = ref(0);

const keyword = ref('');

const activeKey = ref('1');

function onClose() {
  drawerVisible.value = false;

}


/*************************************************
 * ::::新建服务逻辑 start
 ************************************************/
const userListOptions = ref<SelectTypes['options']>([]);

async function setUserListOptions() {
  const res = await getUserList('');
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.label = item.name;
      item.value = item.id
    })
    console.log(832,res.data.result);
    userListOptions.value = res.data.result;
  }
}

async function selectUserFocus(e) {
  await setUserListOptions();
}

async function onSearch(e) {
  await getList();
}

/*************************************************
 * ::::新建服务逻辑 end
 ************************************************/


const edit = (record: any) => {
  editKey.value++;
  drawerVisible.value = true;
  editFormState.name = record.name;
  editFormState.description = record.description;
  editFormState.serveId = record.id;
};

const isEditServiceDesc = ref(false);
const isEditServiceName = ref(false);

function editServiceDesc() {
  isEditServiceDesc.value = true;
}

function editServiceName() {
  isEditServiceName.value = true;
}

async function changeServiceInfo(e) {
  isEditServiceDesc.value = false;
  isEditServiceName.value = false;
  if (editFormState.name && editFormState.description) {
    const res = await saveServe({
      "projectId": currProject.value.id,
      "name": editFormState.name,
      "description": editFormState.description,
      "id": editFormState.serveId,
    });
    if (res.code === 0) {
      // message.success('修改服务描述成功');
      await getList();
    } else {
      // message.error('修改服务描述失败');
    }
  }
}

async function onDelete(record: any) {
  const res = await deleteServe(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await getList();
  } else {
    message.error('删除失败');
  }
}

async function onDisabled(record: any) {
  const res = await disableServe(record.id);
  if (res.code === 0) {
    message.success('禁用服务成功');
    await getList();
  } else {
    message.error('禁用服务失败');
  }
}

async function onCopy(record: any) {
  const res = await copyServe(record.id);
  if (res.code === 0) {
    message.success('复制服务成功');
    await getList();
  } else {
    message.error('复制服务失败');
  }
}

const handleAdd = () => {
  visible.value = true;
};

// 确定
async function handleOk() {
  visible.value = false;
  // :::: todo 需要更换数据
  const res = await saveServe({
    "projectId": currProject.value.id,
    "name": formState.name,
    "description": formState.description,
    "userId": formState.userId,
  });
  if (res.code === 0) {
    message.success('新建服务成功');
    await getList();
  } else {
    message.error('新建服务失败');
  }
}

// 取消
function handleCancel() {
  visible.value = false;
}


async function getList() {
  let res = await getServeList({
    "projectId": currProject.value.id,
    "page": 0,
    "pageSize": 100,
    "name": keyword.value,
  });
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.statusDesc = serveStatus.get(item.status);
      item.statusTag = serveStatusTagColor.get(item.status);
      item.createdAt = momentUtc(item.createdAt)
      item.updatedAt = momentUtc(item.updatedAt)
    })
    dataSource.value = res.data.result;
  }
}

// onMounted(async () => {
//   await getList()
// })

// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await getList()
}, {
  immediate: true
})

</script>

<style scoped lang="less">
.content {
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

  .btns {
    //border-bottom: 1px solid #e9e9e9;
    //padding: 10px 16px;
    //background: #fff;
    //margin-bottom: 8px;
  }
}

.operation-a{
  text-align: center;
  display: inline-block;
  width: 80px;
}

</style>
