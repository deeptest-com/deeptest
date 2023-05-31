<template>
  <div class="card-list p-2">
    <div class="p-2 bg-white">
      <List
        :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: grid, xxl: grid }"
        :data-source="searchValue != '' ? filterList : tableList"
        :pagination="paginationProp"
        :loading="loading"
      >
        <template #header> </template>
        <template #renderItem="{ item }">
          <ListItem>
            <Card class="card" @click="goProject(item)">
              <div class="card-title">
                <div class="title">
                  <img :src="getProjectLogo(item?.logo)" alt="" />

                  <span class="card-title-text" :title="item?.projectName">{{
                    item?.projectName.length > 13
                      ? item?.projectName.substring(0, 13) + "..."
                      : item?.projectName
                  }}</span>
                </div>

                <div class="action">
                  <a-dropdown>
                    <span class="ant-dropdown-link" @click.prevent>
                      <EllipsisOutlined key="ellipsis" />
                    </span>
                    <template #overlay>
                      <a-menu>
                        <a-menu-item
                          v-if="item.accessible === 0"
                          @click="handleJoin(item)"
                        >
                          <a href="javascript:;">申请加入</a>
                        </a-menu-item>
                        <a-menu-item @click="handleEdit(item)">
                          <a href="javascript:;">编辑</a>
                        </a-menu-item>
                        <!-- <a-menu-item>
                          <a href="javascript:;">启用/禁用</a>
                        </a-menu-item> -->
                        <a-menu-item>
                          <a
                            href="javascript:;"
                            @click.stop="handleDelete(item.projectId)"
                            >删除</a
                          >
                        </a-menu-item>
                      </a-menu>
                    </template>
                  </a-dropdown>
                </div>
              </div>
              <div class="card-desc" :title="item?.projectDescr">
                {{
                  item?.projectDescr.length > 38
                    ? item?.projectDescr.substring(0, 38) + "..."
                    : item?.projectDescr
                    ? item?.projectDescr
                    : "&nbsp;"
                }}
              </div>

              <div class="card-static">
                <div>
                  <span>测试场景数：{{ item.scenarioTotal }}个</span>
                  <span>接口数：{{ item.interfaceTotal }}个</span>
                </div>
                <div>
                  <span> 测试覆盖率：{{ item.coverage }}%</span>
                  <span> 执行次数：{{ item.execTotal }}次</span>
                </div>
                <div>
                  <span> 测试通过率：{{ item.passRate }}%</span>
                  <span>发现缺陷数：{{ item.bugTotal }}个</span>
                </div>
              </div>

              <template #actions> </template>
            </Card>
          </ListItem>
        </template>
      </List>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {
  computed,
  onMounted,
  ref,
  defineProps,
  defineEmits,
  watch,
  nextTick,
} from "vue";
import {
  EditOutlined,
  UserOutlined,
  EllipsisOutlined,
  RedoOutlined,
  TableOutlined,
} from "@ant-design/icons-vue";
import { List, Card, Image, Typography, Tooltip, Slider } from "ant-design-vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { StateType } from "../../store";
import { grid } from "./data";
import { getProjectLogo } from "@/components/CreateProjectModal";
// 组件接收参数
const props = defineProps({
  // 请求API的参数
  // params: propTypes.object.def({}),
  activeKey: {
    type: Number,
  },
  searchValue: {
    type: String,
  },
});
const router = useRouter();
const store = useStore<{ Home: StateType }>();
const ListItem = List.Item;
const CardMeta = Card.Meta;
const list = computed<any>(() => store.state.Home.queryResult.list);
const projectLoading = computed<any>(() => store.state.Home.loading);
const TypographyText = Typography.Text;
const tableList = ref<any>([]);
const filterList = ref<any>([]);
const loading = ref(true);
//分页相关
const page = ref(1);
const pageSize = ref(8);
const total = ref(0);
const paginationProp = ref({
  showSizeChanger: false,
  showQuickJumper: true,
  pageSize,
  current: page,
  total,
  showTotal: (total) => `总 ${total} 条`,
  onChange: pageChange,
  onShowSizeChange: pageSizeChange,
});
//暴露内部方法
const emit = defineEmits(["join", "edit", "delete"]);
//数据
const data = ref([]);
const height = computed(() => {
  return `h-${120 - grid.value * 6}`;
});

function sliderChange(n) {
  pageSize.value = n * 4;
}

// 监听关键词搜索
watch(
  () => {
    return props.searchValue;
  },
  async (newVal) => {
    console.log("watch props.searchValue", props.searchValue);
    if (!props.searchValue) {
      total.value = tableList.value.length;
      return;
    }
    const searchText = props.searchValue.toLowerCase();
    filterList.value = tableList.value.filter((item) => {
      // console.log(item)
      // 根据你的数据结构，修改下面的属性名
      return item.projectName.toLowerCase().includes(searchText);
    });
    total.value = filterList.value.length;
  },
  {
    immediate: true,
  }
);
// 监听项目数据变化
watch(
  () => {
    return list.value;
  },
  async (newVal) => {
    console.log("watch list.value", list.value);
    fetch(list.value.userProjectList);
  },
  {
    immediate: true,
  }
);
// 监听我的项目、所有项目切换
watch(
  () => {
    return props.activeKey;
  },
  async (newVal) => {
    if (newVal == 1) {
      fetch(list.value.userProjectList || []);
    } else {
      fetch(list.value.projectList || []);
    }
  },
  {
    immediate: true,
  }
);
// 监听项目loading变化
watch(
  () => {
    return projectLoading.value;
  },
  async (newVal) => {
    loading.value = projectLoading.value.loading;
  },
  {
    immediate: true,
  }
);
async function fetch(data) {
  tableList.value = data;

  if (tableList.value && tableList.value.length > 0) {
    total.value = tableList.value.length;
  }
}

function pageChange(p, pz) {
  page.value = p;
  pageSize.value = pz;
}
function pageSizeChange(_current, size) {
  pageSize.value = size;
}

async function handleJoin(item) {
  emit("join", item);
}
async function handleEdit(item) {
  emit("edit", item);
}
async function handleDelete(id) {
  emit("delete", id);
}
async function goProject(item: any) {
  if (item?.accessible === 0) {
    handleJoin(item);
    return false;
  }
  await store.dispatch("ProjectGlobal/changeProject", item?.projectId);
  // 更新左侧菜单以及按钮权限
  await store.dispatch("Global/getPermissionList");
  // 项目切换后，需要重新更新可选服务列表
  await store.dispatch("ServeGlobal/fetchServe");
  router.push(`/workbench/index`);
}
</script>


<style lang="less" scoped>
.card-list {
  :deep(.ant-list-header) {
    border: none;
    padding: 0;
  }
  .card {
    cursor: pointer;
    height: 220px;
    max-height: 220px;
    &-title {
      font-size: 18px;
      font-weight: 500;
      display: flex;
      align-items: center;
      justify-content: space-between;
      &-text {
        padding-left: 8px;
      }
    }
    &-desc {
      margin-top: 8px;
    }
    &-static {
      margin-top: 8px;
      div {
        display: flex;
        // justify-content: space-between;
        span {
          flex: 1;
        }
      }
    }
  }
}
.add-card {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
}
</style>
