<template>
  <HomeLayout>
    <div class="home">
      <StatisticHeader :params="cardData" :type="0"/>
      <div style="margin-top: 16px">
        <a-card :bordered="false">
          <template #title>
            <a-tabs v-model:activeKey="activeKey">
              <a-tab-pane :key="1" tab="我的项目"/>
              <a-tab-pane :key="0" tab="所有项目"/>
            </a-tabs
            >
          </template>
          <template #extra>
            <a-input-search
                v-model:value="keywordsMap[activeKey]"
                style="width: 200px; margin-right: 20px"
                placeholder="请输入项目名称搜索"/>
            <a-button
                type="primary"
                style="margin-right: 20px"
                @click="handleOpenAdd">新建项目
            </a-button>
            <a-radio-group v-model:value="showMode" button-style="solid">
              <a-radio-button value="card">卡片</a-radio-button>
              <a-radio-button value="list">列表</a-radio-button>
            </a-radio-group>
          </template>
          <div>
            <HomeList
                v-if="showMode === 'list'"
                :activeKey="activeKey"
                :searchValue="keywordsMap[activeKey]"
                @join="handleJoin"
                :isLoading="isLoadingList"
                @edit="handleOpenEdit"
                @delete="handleDelete"
                @exit="handleExit"/>
            <CardList v-if="showMode === 'card'"
                      :activeKey="activeKey"
                      :searchValue="keywordsMap[activeKey]"
                      :isLoading="isLoadingList"
                      @join="handleJoin"
                      @edit="handleOpenEdit"
                      @delete="handleDelete"
                      @exit="handleExit"/>
          </div>
        </a-card>
      </div>

      <!-- 创建项目弹窗 -->
      <CreateProjectModal
          :visible="createProjectModalVisible"
          :formState="formState"
          @update:visible="createProjectModalVisible = false"
          @handleSuccess="handleCreateSuccess"
      />
      <!-- 申请项目权限弹窗 -->
      <ApplyProPermissionsModal
          :visible="applyProPermissionsModalVisible"
          :item="applyItem"
          @update:visible="applyProPermissionsModalVisible = false"
          @handleSuccess="handleSuccess"
      />
    </div>
  </HomeLayout>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref, watch} from "vue";
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import {Modal, notification} from "ant-design-vue";

import StatisticHeader from "@/components/StatisticHeader/index.vue";
import CreateProjectModal from "@/components/CreateProjectModal/index.vue";
import ApplyProPermissionsModal from "@/components/ApplyProPermissions/index.vue";
import HomeList from "./component/HomeList/index.vue";
import CardList from "./component/CardList/index.vue";
import HomeLayout from "@/layouts/HomeLayout.vue";
import {StateType} from "./store";
import {removeMember} from "@/views/project/service";

import {NotificationKeyCommon} from "@/utils/const";
import {CurrentUser, StateType as UserStateType} from "@/store/user";
import {notifyError, notifySuccess} from "@/utils/notify";

// 获取当前登录用户信息
const router = useRouter();
const store = useStore<{ Home: StateType, User: UserStateType }>();
const currentUser = computed<CurrentUser>(() => store.state.User.currentUser);
const cardData = computed<any>(() => store.state.Home.cardData);
const activeKey = ref(1);
const keywordsMap = ref({
  1: null, // 我的项目
  0: null, // 所有项目
});
const showMode = ref("card");
const createProjectModalVisible = ref(false);
const applyProPermissionsModalVisible = ref(false);
const applyItem = ref({});
const isLoadingList = ref(true);

let formState = ref({
  id: 0,
  logo: "",
  name: "",
  shortName: "",
  adminId: "",
  includeExample: false,
  desc: "",
});

onMounted(async () => {
  if (router.currentRoute.value?.query?.type == 'all') {
    activeKey.value = 0
  }
  await getHearderData();
  await getList(1);
  await store.dispatch("User/fetchCurrent");
});

const getHearderData = async (): Promise<void> => {
  await store.dispatch("Home/queryCard", {projectId: 0});
  await store.dispatch("Home/queryPie", {projectId: 0});
};

// 获取全部项目数据
const getList = async (current: number): Promise<void> => {
  await store.dispatch("Home/queryProject", {});
  isLoadingList.value = false;
};

// 创建项目成功的回调
const handleCreateSuccess = () => {
  createProjectModalVisible.value = false;
  getList(1);
};

// 申请加入项目成功的回调
const handleSuccess = async () => {
  applyProPermissionsModalVisible.value = false;
};



function handleJoin(item) {
  Modal.confirm({
    title: "提示",
    content: "您还没有该项目的访问权限，是否申请更多角色权限？",
    okText: "申请权限",
    cancelText: "取消",
    onOk: async () => {
      applyProPermissionsModalVisible.value = true;
      applyItem.value = item;
    },
  });
}

function handleOpenAdd() {
  createProjectModalVisible.value = true;
  formState.value.id = 0;
}

function handleOpenEdit(item) {
  formState.value.id = item.projectId;
  formState.value.name = item.projectName;
  formState.value.logo = item.logo;
  formState.value.shortName = item.projectShortName;
  formState.value.adminId = item.adminId;
  formState.value.includeExample = item.includeExample;
  formState.value.desc = item.projectDescr;
  createProjectModalVisible.value = true;
}

async function handleDelete(id) {
  Modal.confirm({
    title: "删除项目",
    content: "确定删除指定的项目？",
    okText: "确认",
    cancelText: "取消",
    onOk: async () => {
      store.dispatch("Project/removeProject", id).then((res) => {
        console.log("res", res);
        if (res === true) {
          notifySuccess(`删除成功`);
          getList(1);
        } else {
          notifyError(`删除失败`);
        }
      });
    },
  });
}

const handleExit = (item) => {
  console.log("exit");

  Modal.confirm({
    title: "退出项目",
    content: "确定要推出项目" + item.projectName + "？",
    okText: "确认",
    cancelText: "取消",
    onOk: async () => {
      removeMember(currentUser.value.id, item.projectId).then((json) => {
        if (json.code === 0) {
          notifySuccess(`退出成功`);
          getList(1);
        } else {
          notifyError(`退出失败`);
        }
      });
    },
  });
};

</script>

<style lang="less" scoped>
.home {
  padding: 16px;

  :deep(.ant-card-head .ant-tabs-bar) {
    border-bottom: none;
  }
}
</style>
