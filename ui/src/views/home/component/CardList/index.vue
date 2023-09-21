<template>
  <div class="card-list p-2">
    <div class="p-2 bg-white">
      <List :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: 4, xxl: 4 }"
            :data-source="filterList"
            :loading="isLoading"
            :pagination="{
                pageSize: 8,
                current: current,
                total: total,
                onChange: (page) => {
                  handlePageChange(page);
                },
                showTotal: (total) => {
                  return `共 ${total} 条数据`;
                },
            }">
        <template #header></template>
        <template #renderItem="{ item }">
          <ListItem>
            <Card class="card" @click="goProject(item)">
              <div class="card-title">
                <div class="title">
                  <img :src="getProjectLogo(item?.logo)" alt=""/>
                  <span class="card-title-text" :title="item?.projectName">{{
                      item?.projectName.length > 13
                          ? item?.projectName.substring(0, 13) + "..."
                          : item?.projectName
                    }}</span>
                </div>

                <div class="action">
                  <a-dropdown>
                    <span class="ant-dropdown-link" @click.prevent.stop>
                      <EllipsisOutlined key="ellipsis"/>
                    </span>
                    <template #overlay>
                      <a-menu>
                        <a-menu-item
                            v-if="item.accessible === 0"
                            @click="handleJoin(item)"
                        >
                          <a href="javascript:;">申请加入</a>
                        </a-menu-item>
                        <a-menu-item
                            v-if="item.accessible === 1"
                            @click="handleEdit(item)">
                          <a href="javascript:;">编辑</a>
                        </a-menu-item>
                        <!-- <a-menu-item>
                          <a href="javascript:;">启用/禁用</a>
                        </a-menu-item> -->
                        <a-menu-item v-if="item.accessible === 1">
                          <a
                              href="javascript:;"
                              @click.stop="handleDelete(item.projectId)"
                          >删除</a
                          >
                        </a-menu-item>
                        <a-menu-item v-if="item.accessible === 1">
                          <a
                              href="javascript:;"
                              @click.stop="handleExit(item)"
                          >退出项目</a
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

              <template #actions></template>
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
  ref,
  defineProps,
  defineEmits,
  watch,
} from "vue";
import {
  EllipsisOutlined,
} from "@ant-design/icons-vue";
import {List, Card, Image, Typography, Tooltip, Slider} from "ant-design-vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {StateType} from "../../store";
import {getProjectLogo} from "@/components/CreateProjectModal";

// 组件接收参数
const props = defineProps({
  // 请求API的参数
  activeKey: {
    type: Number,
  },
  searchValue: {
    type: String,
  },
  isLoading: {
    type: Boolean,
    default: false,
  },
});
const router = useRouter();
const store = useStore<{ Home: StateType }>();
const ListItem = List.Item;
const list = computed<any>(() => store.state.Home.queryResult.list);

const filterList = computed(() => {
  const items = props?.activeKey === 0 ? list?.value?.projectList || [] : list?.value?.userProjectList || [];
  if(!items?.length) return [];
  return items.filter((item) => {
    const projectName = (item.projectName || '').toLowerCase();
    const keyword = (props?.searchValue || '').toLowerCase();
    const projectShortName = (item.projectShortName || '').toLowerCase();
    return projectName.includes(keyword) || projectShortName.includes(keyword);
  })
})

const loading = ref(true);
// 分页相关
const current = ref(1);
const total = computed(() => filterList.value.length);

watch(() => props?.searchValue, (val) => {
  current.value = 1;
})

//暴露内部方法
const emit = defineEmits(["join", "edit", "delete", "exit"]);

function handlePageChange(page) {
  current.value = page;
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

async function handleExit(item) {
  emit("exit", item);
}

async function goProject(item: any) {
  if (item?.accessible === 0) {
    handleJoin(item);
    return false;
  }
  await store.dispatch("ProjectGlobal/changeProject", item?.projectId);
  // 更新左侧菜单以及按钮权限
  await store.dispatch("Global/getPermissionList", { projectId: item.id });
  // 项目切换后，需要重新更新可选服务列表
  await store.dispatch("ServeGlobal/fetchServe");
  router.push(`/${item.projectShortName}/workspace`);
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
