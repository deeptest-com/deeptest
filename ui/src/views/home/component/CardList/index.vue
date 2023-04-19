<template>
  <div class="card-list p-2">
    <div class="p-2 bg-white">
      <List
        :grid="{ gutter: 5, xs: 1, sm: 2, md: 4, lg: 4, xl: 3, xxl: grid }"
        :data-source="list"
        :pagination="paginationProp"
      >
        <template #header> </template>
        <template #renderItem="{ item }">
          <ListItem>
            <Card>
              <template #title>
                <Avatar style="background-color: #1890ff">
                  <template #icon><UserOutlined /></template>
                </Avatar>
                {{item.project_chinese_name}}
              </template>
              <!-- <template #avatar>
                <a-avatar style="background-color: #1890ff">
                  <template #icon><UserOutlined /></template>
                </a-avatar>

              </template> -->
              <template #description>{{ item.project_chinese_name }}</template>
              <template #cover>
                <!-- <div :class="height">
                  <Image :src="item.imgs[0]" />
                </div> -->
              </template>
              <template #actions>
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
              </template>

              <CardMeta>
                <template #title>
                  <!-- <TypographyText :content="item.name" :ellipsis="{ tooltip: item.address }" /> -->
                </template>
                <template #avatar>
                  <!-- <Avatar :src="item.avatar" /> -->
                </template>
                <template #description>
                  <!-- {{ item.project_chinese_name }} -->
                </template>
              </CardMeta>
            </Card>
          </ListItem>
        </template>
      </List>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref, defineProps, defineEmits } from "vue";
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
// import { Dropdown } from '/@/components/Dropdown';
// import { propTypes } from '/@/utils/propTypes';
// import { Button } from '/@/components/Button';
import { useStore } from "vuex";
import { StateType } from "../../store";
import { useSlider, grid } from "./data";
const store = useStore<{ Home: StateType }>();
const ListItem = List.Item;
const CardMeta = Card.Meta;
const list = computed<any[]>(() => store.state.Home.queryResult.list);
const TypographyText = Typography.Text;
// 获取slider属性
const sliderProp = computed(() => useSlider(4));
// 组件接收参数
const props = defineProps({
  // 请求API的参数
  // params: propTypes.object.def({}),
  params: {
    type: Object,
  },
  //api
  // api: propTypes.func,
  api: {
    type: String,
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
function handleTabClick(e: number) {
  // queryParams.userId=e;
  //  getList(1);
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
</script>


<style lang="less" scoped>
.card-list {
  :deep(.ant-list-header) {
    border: none;
    padding: 0;
  }
}
</style>
