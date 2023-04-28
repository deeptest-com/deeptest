<template>
  <div class="card-list p-2">
    <div class="p-2 bg-white">
      <List
        :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: 3, xxl: grid }"
        :data-source="tableList"
        :pagination="paginationProp"
        :loading="loading"
      >
        <template #header> </template>
        <template #renderItem="{ item }">
          <ListItem>
            <Card class="card" @click="goProject(item?.project_id)">
              <!-- <template #title> -->
              <div class="card-title">
                <div class="title">
                  <!-- <Avatar style="background-color: #1890ff">
                    <template #icon> -->
                      <img :src="getProjectLogo(item?.logo)" alt="" />
                    <!-- </template>
                  </Avatar> -->
                  <span class="card-title-text">{{
                    item?.project_chinese_name
                  }}</span>
                </div>

                <div class="action">
                  <a-dropdown>
                    <span class="ant-dropdown-link" @click.prevent>
                      <EllipsisOutlined key="ellipsis" />
                    </span>
                    <template #overlay>
                      <a-menu>
                        <a-menu-item @click="handleEdit(item)">
                          <a href="javascript:;">编辑</a>
                        </a-menu-item>
                        <!-- <a-menu-item>
                          <a href="javascript:;">启用/禁用</a>
                        </a-menu-item> -->
                        <a-menu-item>
                          <a href="javascript:;" @click="handleDelete(item.project_id)">删除</a>
                        </a-menu-item>
                      </a-menu>
                    </template>
                  </a-dropdown>
                </div>
              </div>

              <!-- </template> -->
              <div class="card-desc">
                {{
                  item.project_des.length > 35
                    ? item.project_des.substring(0, 35) + "..."
                    : item.project_des
                    ? item.project_des
                    : "&nbsp;"
                }}
              </div>

              <div class="card-static">
                <div>
                  测试场景数： {{ item.scenario_total }}个 接口数： {{}}个
                </div>
                <div>
                  测试覆盖率：{{ item.coverage }}% 执行次数：
                  {{ item.exec_total }}次
                </div>
                <div>
                  测试通过率： {{ item.pass_rate }}% 发现缺陷数：
                  {{ item.bug_total }}个
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
import {
  List,
  Card,
  Image,
  Typography,
  Tooltip,
  Slider,
  Avatar,

} from "ant-design-vue";
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
});
const router = useRouter();
const store = useStore<{ Home: StateType }>();
const ListItem = List.Item;
const CardMeta = Card.Meta;
const list = computed<any>(() => store.state.Home.queryResult.list);
const TypographyText = Typography.Text;
const tableList = ref([]);
const loading = ref(true);
//分页相关
const page = ref(1);
const pageSize = ref(6);
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
const emit = defineEmits(["edit", "delete"]);
//数据
const data = ref([]);
const height = computed(() => {
  return `h-${120 - grid.value * 6}`;
});

function sliderChange(n) {
  pageSize.value = n * 4;
}

// 监听项目数据变化
watch(
  () => {
    return list.value;
  },
  async (newVal) => {
    console.log("watch list.value", list.value);
    fetch(list.value.current_user_project_list);
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
      fetch(list.value.current_user_project_list);
    } else {
      fetch(list.value.all_project_list);
    }
  },
  {
    immediate: true,
  }
);

async function fetch(data) {
  tableList.value = data;
  if (tableList.value && tableList.value.length > 0) {
    total.value = tableList.value.length;
    loading.value = false;
  } else {
    setTimeout(() => {
      loading.value = false;
    }, 5000);
  }
}

function pageChange(p, pz) {
  page.value = p;
  pageSize.value = pz;
}
function pageSizeChange(_current, size) {
  pageSize.value = size;
}

async function handleEdit(item) {
  emit("edit", item);
}
async function handleDelete(id) {
  emit("delete",id)
  
}
function goProject(projectId: number) {
  store.dispatch("ProjectGlobal/changeProject", projectId);
  // store.dispatch("Environment/getEnvironment", { id: 0, projectId: projectId });

  // 项目切换后，需要重新更新可选服务列表
  store.dispatch("ServeGlobal/fetchServe");
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
