<template>
  <div class="home">
    <StatisticHeader :params="cardData" :type="0"/>
    <div style="margin-top: 16px">
      <a-card :bordered="false">
        <template #title>
          <a-tabs v-model:activeKey="activeKey" @change="handleTabClick">
            <a-tab-pane :key="1" tab="我的项目"> </a-tab-pane>
            <a-tab-pane :key="0" tab="所有项目"> </a-tab-pane> </a-tabs
        ></template>
        <template #extra>
          <a-input-search
            v-model:value="keywords"
            style="width: 200px; margin-right: 20px"
            placeholder="请输入项目名称搜索"
            @search="onSearch"
          />
          <a-button
            type="primary"
            style="margin-right: 20px"
            @click="handleOpenAdd"
            >新建项目</a-button
          >
          <a-radio-group v-model:value="showMode" button-style="solid">
            <a-radio-button value="card">卡片</a-radio-button>
            <a-radio-button value="list">列表</a-radio-button>
          </a-radio-group>
        </template>
        <div>
          <HomeList
            v-if="showMode == 'list'"
            :activeKey="activeKey"
            :searchValue="searchValue"
            @edit="handleOpenEdit"
            @delete="handleDelete"
          />

          <CardList
            v-else
            :activeKey="activeKey"
            :searchValue="searchValue"
            @edit="handleOpenEdit"
            @delete="handleDelete"
          />
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
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import { Modal, notification } from "ant-design-vue";
import StatisticHeader from "@/components/StatisticHeader/index.vue";
import CreateProjectModal from "@/components/CreateProjectModal/index.vue";
import HomeList from "./component/HomeList/index.vue";
import CardList from "./component/CardList/index.vue";
import { useStore } from "vuex";
import { StateType } from "./store";
import { PaginationConfig, QueryParams } from "./data.d";
import EditPage from "@/views/project/edit/edit.vue";
import { useRouter } from "vue-router";
import { setCache } from "@/utils/localCache";
import settings from "@/config/settings";
const store = useStore<{ Home: StateType }>();
const cardData = computed<any>(() => store.state.Home.cardData);
const activeKey = ref(1);
const keywords = ref<string>('');
const searchValue = ref("");
const showMode = ref("card");
const createProjectModalVisible = ref(false);
let formState = ref({
  id: 0,
  logo: "",
  name: "",
  shortName: "",
  adminId: "",
  includeExample: false,
  desc: "",
});

onMounted(() => {
  getHearderData()
  getList(1);
});
const onSearch = () => {
  searchValue.value = keywords.value;
};
const getHearderData = async (): Promise<void> => {
  await store.dispatch("Home/queryCard", {projectId:0});
  await store.dispatch("Home/queryPie", {projectId:0});
};
const getList = async (current: number): Promise<void> => {
  await store.dispatch("Home/queryProject", {});
};
// 创建项目成功的回调
const handleCreateSuccess = () => {
  createProjectModalVisible.value = false;
  getList(1);
};

function handleTabClick(e: number) {
  console.log("activeKey", activeKey);
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
          notification.success({
            // key: NotificationKeyCommon,
            message: `删除成功`,
          });
          getList(1);
        } else {
          notification.error({
            // key: NotificationKeyCommon,
            message: `删除失败`,
          });
        }
      });
    },
  });
}
</script>

<style lang="less" scoped>
.home {
  padding: 16px;
  :deep(.ant-card-head .ant-tabs-bar) {
    border-bottom: none;
  }
}
</style>
