<template>
  <div class="home">
    <StatisticHeader />
    <div style="margin-top: 16px">
      <a-card :bordered="false">
        <template #title>
          <a-tabs v-model:activeKey="activeKey" @change="handleTabClick">
            <a-tab-pane :key="1" tab="我的项目"> </a-tab-pane>
            <a-tab-pane :key="0" tab="所有项目"> </a-tab-pane> </a-tabs
        ></template>
        <template #extra>
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
          <HomeList v-if="showMode == 'list'" :activeKey="activeKey" />

          <CardList
            v-else
            @edit="handleOpenEdit"
            @delete="handleDelete"
            :activeKey="activeKey"
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
const activeKey = ref(1);
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
// let queryParams = reactive<QueryParams>({
// keywords: "",

// });

onMounted(() => {
  getList(1);
});

const getList = async (current: number): Promise<void> => {
  await store.dispatch("Home/queryProject", {
    // keywords: queryParams.keywords,
  });
};

// 创建项目成功的回调
const handleCreateSuccess = () => {
  createProjectModalVisible.value = false;
  // todo: 重新获取列表
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
  formState.value.id = item.project_id;
  formState.value.name = item.project_chinese_name;
  formState.value.logo = item.logo;
  formState.value.shortName = item.project_name;
  formState.value.adminId = item.admin_id;
  formState.value.includeExample = item.include_example;
  formState.value.desc = item.project_des;
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
