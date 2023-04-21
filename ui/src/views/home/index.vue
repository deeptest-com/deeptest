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
          <!-- <a-input-search
          @change="onSearch"
          @search="onSearch"
          v-model:value="queryParams.keywords"
          placeholder="输入关键字搜索"
          style="width: 270px; margin-left: 16px"
        /> -->
          <a-radio-group v-model:value="showMode" button-style="solid">
            <a-radio-button value="card">卡片</a-radio-button>
            <a-radio-button value="list">列表</a-radio-button>
          </a-radio-group>
        </template>

        <div>
          <HomeList v-if="showMode == 'list'" params />
          <!-- <CardList
            v-else
            :params="{ params: 1 }"
            :api="'demoListApi/111'"
            @get-method="getMethod"
            @delete="handleDel"
          /> -->
            <CardList
            v-else
           
          />
        </div>
      </a-card>
    </div>
  </div>
</template>


<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import StatisticHeader from "@/components/StatisticHeader/index.vue";
import HomeList from "./component/HomeList/index.vue";
import CardList from "./component/CardList/index.vue";
import { useStore } from "vuex";
import { StateType } from "./store";
import { PaginationConfig, QueryParams } from "./data.d";
const store = useStore<{ Home: StateType }>();
const mode = computed<any[]>(() => store.state.Home.mode);
const activeKey = ref(1);
const showMode = ref("card");
const currentUser = computed<any>(() => store.state.User.currentUser);
let queryParams = reactive<QueryParams>({
  keywords: "",
  enabled: "1",
  userId: activeKey.value == 0 ? 0 : currentUser.value?.id,
  // page: pagination.value.current,
  // pageSize: pagination.value.pageSize,
});

onMounted(() => {
  getList(1);
});

const getList = async (current: number): Promise<void> => {
  // console.log('queryParams.keywords',queryParams.keywords)
  await store.dispatch("Home/queryProject", {
    keywords: queryParams.keywords,
    enabled: queryParams.enabled,
    userId: queryParams.userId,
    currProjectId:0,
    // pageSize: pagination.value.pageSize,
    // page: current,
  });
};
function handleTabClick(e: number) {
  queryParams.userId=e;
   getList(1);
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