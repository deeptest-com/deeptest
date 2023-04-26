<template>
  <div class="card-list p-2">
    <div class="p-2 bg-white">
      <List
        :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: 3, xxl: grid }"
        :data-source="tableList.concat(addCard)"
        :pagination="paginationProp"
      >
        <template #header> </template>
        <template #renderItem="{ item }">
          <ListItem>
            <div v-if="item.type && item.type == 'add'">
              <Card class="card add-card" @click="addProject(0)">创建项目+</Card>
            </div>
            <div v-else>
              <Card class="card" @click="goProject(item.id)">
                <!-- <template #title> -->
                <div class="card-title">
                  <Avatar style="background-color: #1890ff">
                    <template #icon><UserOutlined /></template>
                  </Avatar>
                  <span class="card-title-text">{{
                    item.project_chinese_name
                  }}</span>
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

                <!-- <template #actions> -->

                <!--              <SettingOutlined key="setting" />-->
                <!-- <EditOutlined key="edit" /> -->
                <!-- <Dropdown
                  :trigger="['hover']"
                  :dropMenuList="[
                    {
                      text: '删除',
                      event: '1',
                      popConfirm: {
                        title: '是否确认删除',
                        confirm: handleDelete.bind(null, item.id),
                      },
                    },
                  ]"
                  popconfirm
                >
                  <EllipsisOutlined key="ellipsis" />
                </Dropdown> -->
                <!-- </template> -->
              </Card>
            </div>
          </ListItem>
        </template>
      </List>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref, defineProps, defineEmits, watch } from "vue";
import {
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
// import { Dropdown } from '/@/components/Dropdown';
// import { propTypes } from '/@/utils/propTypes';
// import { Button } from '/@/components/Button';
import { useStore } from "vuex";
import { StateType } from "../../store";
import { grid } from "./data";
const router = useRouter();
const store = useStore<{ Home: StateType }>();
const ListItem = List.Item;
const CardMeta = Card.Meta;
const list = computed<any>(() => store.state.Home.queryResult.list);
const TypographyText = Typography.Text;
const tableList = ref([]);
const addCard = ref([{ type: "add" }]);

// 组件接收参数
const props = defineProps({
  // 请求API的参数
  // params: propTypes.object.def({}),
  activeKey: {
    type: Number,
  },
});
//暴露内部方法
const emit = defineEmits(["getMethod", "delete"]);
//数据
const data = ref([]);
// 切换每行个数
// cover图片自适应高度
//修改pageSize并重新请求数据

const height = computed(() => {
  return `h-${120 - grid.value * 6}`;
});

function sliderChange(n) {
  pageSize.value = n * 4;
  fetch();
}

// 自动请求并暴露内部方法
onMounted(() => {
  // fetch();
  // emit('getMethod', fetch);
});
watch(
  () => {
    return props.activeKey;
  },
  (newVal) => {
    if (newVal == 1) {
      setTimeout(() => {
        tableList.value = list.value.current_user_project_list;
      }, 500);
    } else {
      setTimeout(() => {
        // tableList.value =[]
        tableList.value = list.value.all_project_list;
      }, 500);
    }
  },
  {
    immediate: true,
  }
);
function handleTabClick(e: number) {
  // queryParams.userId=e;
  //  getList(1);
}
function addProject(id:number){
 


}
async function fetch(p = {}) {
  const { api, params } = props;
  if (api && typeof api === "function") {
    const res = await api({
      ...params,
      page: page.value,
      pageSize: pageSize.value,
      ...p,
    });
    data.value = res.items;
    total.value = res.total;
  }
}
//分页相关
const page = ref(1);
const pageSize = ref(36);
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

function pageChange(p, pz) {
  page.value = p;
  pageSize.value = pz;
  fetch();
}
function pageSizeChange(_current, size) {
  pageSize.value = size;
  fetch();
}

async function handleDelete(id) {
  emit("delete", id);
}
function goProject(id: number) {
  router.push(`/workbench/index/${id}`);
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
.add-card{
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
}
</style>
