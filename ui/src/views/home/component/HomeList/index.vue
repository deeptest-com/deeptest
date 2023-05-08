<template>
  <div class="home-list">
     <a-table
      row-key="id"
      :columns="columns"
      :data-source="tableList"
      :loading="loading"
    
    >
      <template #name="{ text, record }">
        <div class="project-name" :title="text" @click="goProject(record.projectId)">
          {{ text.length>16? text.substring(0,16)+'...':text}}
        </div>
      </template>

      <template #action="{ record }">
        <a-dropdown>
          <MoreOutlined />
          <template #overlay>
            <a-menu>
              <a-menu-item key="1">
                <a-button style="width: 80px" @click="handleEdit(record)" type="link" size="small"
                  >编辑</a-button
                >
              </a-menu-item>
              <!-- <a-menu-item key="2">
                <a-button style="width: 80px" type="link" size="small"
                  >禁用/启用</a-button
                >
              </a-menu-item> -->
              <a-menu-item key="3">
                <a-button
                  style="width: 80px"
                  type="link"
                  size="small"
                  @click="handleDelete(record.projectId)"
                  >删除</a-button
                >
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>
    </a-table>
  </div>
</template>


<script setup lang="ts">
import { computed, onMounted, reactive, ref, defineProps,defineEmits, watch } from "vue";
import debounce from "lodash.debounce";
import { PaginationConfig, QueryParams } from "../../data.d";
import { SelectTypes } from "ant-design-vue/es/select";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import { StateType } from "../../store";
import { StateType as UserStateType } from "@/store/user";
import { StateType as ProjectStateType } from "@/store/project";
import { MoreOutlined } from "@ant-design/icons-vue";
const router = useRouter();
const store = useStore<{
  ProjectGlobal: ProjectStateType;
  Home: StateType;
  User: UserStateType;
}>();
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currentUser = computed<any>(() => store.state.User.currentUser);
const list = computed<any>(() => store.state.Home.queryResult.list);
const projectLoading = computed<any>(() => store.state.Home.loading);
const loading = ref<boolean>(false);
const showMode = ref("list");
const activeKey = ref(1);
const tableList = ref([]);
// 组件接收参数
const props = defineProps({
  // 请求API的参数
  // params: propTypes.object.def({}),
  activeKey: {
    type: Number,
  },
});

const total = ref(0);
let queryParams = reactive<QueryParams>({
  keywords: "",

});
//暴露内部方法
const emit = defineEmits(["edit", "delete"]);
const columns = [
  // {
  //   title: '序号',
  //   dataIndex: 'index',
  //   width: 80,
  //   customRender: ({
  //                    text,
  //                    index
  //                  }: { text: any; index: number }) => (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  // },
   {
    title: "项目名称",
    dataIndex: "projectName",
    slots: { customRender: "name" },
      width: 200,
       ellipsis: true,
  },
  {
    title: "英文缩写",
    dataIndex: "projectShortName",
     ellipsis: true,
     width: 150,
  },
 

  {
    title: "管理员",
    dataIndex: "adminName",
      ellipsis: true,
     width: 150,
  },

  {
    title: "测试场景数",
    dataIndex: "scenarioTotal",
  },
  {
    title: "测试覆盖率",
    dataIndex: "coverage",
  },
  {
    title: "执行次数",
    dataIndex: "execTotal",
  },
  {
    title: "测试通过率",
    dataIndex: "passRate",
  },
  {
    title: "发现缺陷",
    dataIndex: "bugTotal",
  },
  {
    title: "创建时间",
    dataIndex: "createdAt",
      ellipsis: true,
     width: 200,
  },
  {
    title: '操作',
    key: 'action',
    width: 60,
    slots: {customRender: 'action'},
  },
];
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
      fetch(list.value.userProjectList);
    } else {
      fetch(list.value.projectList);
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
const getList = async (current: number): Promise<void> => {
 
  console.log("queryParams.keywords", queryParams.keywords);
  await store.dispatch("Home/queryProject", {
    keywords: queryParams.keywords,

  });

};
function handleTabClick(e: number) {
 
  getList(1);
}
async function goProject(projectId: number) {
  await store.dispatch("ProjectGlobal/changeProject", projectId);
  // 更新左侧菜单以及按钮权限
  await store.dispatch('Global/getPermissionList');
  // 项目切换后，需要重新更新可选服务列表
  await store.dispatch("ServeGlobal/fetchServe");
  router.push(`/workbench/index`);
}
function changeMode(e) {
  console.log("changemode", e.target.value);
  store.dispatch("Home/changemode", {
    mode: showMode.value,
  });
}
const onSearch = debounce(() => {
  getList(1);
}, 500);
const members = (id: number) => {
  console.log("members");
  router.push(`/project/members/${id}`);
};

const visible = ref(false);
const currentProjectId = ref(0);

const handleOk = (e: MouseEvent) => {
  console.log(e);
  visible.value = false;
};

const closeModal = () => {
  visible.value = false;
};
async function handleEdit(item) {
  emit("edit", item);
}
async function handleDelete(id) {
  emit("delete", id);
}

</script>

<style lang="less" scoped>
.home-list {
  // padding:0 16px;

  .project-name {
    color: #447dfd;
    cursor: pointer;
  }
}
</style>