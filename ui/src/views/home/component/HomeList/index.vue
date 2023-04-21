<template>
  <div class="home-list">

  
        <a-table
            row-key="id"
            :columns="columns"
            :data-source="list"
            :loading="loading"
            :pagination="{
                ...pagination,
                onChange: (page) => {
                    getList(page);
                },
                onShowSizeChange: (page, size) => {
                    pagination.pageSize = size
                    getList(page);
                },
            }"
        >
          <template #name="{ text, record }">
        <div class="project-name" @click="goProject(record.project_id)">
          {{text}}
        </div>
      </template>
       
     <template #action="{record}">
            <a-dropdown>
              <MoreOutlined/>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">
                    <a-button style="width: 80px" type="link" size="small" >编辑</a-button>
                  </a-menu-item>
                   <a-menu-item key="2">
                    <a-button style="width: 80px" type="link" size="small">禁用/启用</a-button>
                  </a-menu-item>
                  <a-menu-item key="3">
                    <a-button style="width: 80px" type="link" size="small" @click="del(record)">删除</a-button>
                  </a-menu-item>
                 
                </a-menu>
              </template>
            </a-dropdown>
          </template>

        </a-table>
    
  </div>
</template>


<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import debounce from "lodash.debounce";
import { PaginationConfig, QueryParams } from "../../data.d";
import {SelectTypes} from 'ant-design-vue/es/select';
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import {Modal, notification} from "ant-design-vue";
import {NotificationKeyCommon} from "@/utils/const";
import {StateType} from "../../store";
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {MoreOutlined} from '@ant-design/icons-vue';
const router = useRouter();
const store = useStore<{ ProjectGlobal: ProjectStateType, Home: StateType, User: UserStateType }>();
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currentUser = computed<any>(() => store.state.User.currentUser);
const list = computed<any[]>(() => store.state.Home.queryResult.list);

const showMode=ref('list')
const activeKey = ref(1);
let pagination = computed<PaginationConfig>(() => store.state.Home.queryResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',userId:activeKey.value==0?0:currentUser.value?.id,
  page: pagination.value.current, pageSize: pagination.value.pageSize
});

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
    title: '英文缩写',
    dataIndex: 'project_name',
  },
  {
    title: '项目名称',
    dataIndex: 'project_chinese_name',
    slots: {customRender: 'name'},
  },

  {
    title: '管理员',
    dataIndex: 'admin_user',
  },

  {
    title: '测试场景数',
    dataIndex: 'scenario_total',
  },
  {
    title: '测试覆盖率',
    dataIndex: 'coverage',
  },
  {
    title: '执行次数',
    dataIndex: 'exec_total',
  },
   {
    title: '测试通过率',
    dataIndex: 'pass_rate',
  },
  {
    title: '发现缺陷',
    dataIndex: 'bug_total',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
   
  },
  // {
  //   title: '操作',
  //   key: 'action',
  //   // width: 260,
  //   slots: {customRender: 'action'},
  // },
];

onMounted(() => {
  console.log('onMounted',currentUser.value.id)
  // getList(1);
})

// const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  // loading.value = true;
console.log('queryParams.keywords',queryParams.keywords)
  await store.dispatch('Home/queryProject', {
    keywords: queryParams.keywords,
    enabled: queryParams.enabled,
    userId: queryParams.userId,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  // loading.value = false;
  // await store.dispatch('Home/getUserList');

  // store.dispatch("ProjectGlobal/fetchProject");
}
function handleTabClick(e: number) {
  queryParams.userId=e;
   getList(1);
}
function goProject(projectId:number){
  store.dispatch('ProjectGlobal/changeProject', projectId);
  store.dispatch('Environment/getEnvironment', {id: 0, projectId: projectId});

  // 项目切换后，需要重新更新可选服务列表
  store.dispatch("ServeGlobal/fetchServe");
  router.push(`/workbench/index`)

}
function changeMode(e){
  console.log('changemode',e.target.value)
    store.dispatch('Home/changemode', {
    mode: showMode.value
  });
}
const onSearch = debounce(() => {
  getList(1)
}, 500);
const members = (id: number) => {
  console.log('members')
  router.push(`/project/members/${id}`)
}

const visible = ref(false)
const currentProjectId = ref(0)
const edit = (id: number) => {
  console.log('edit')
  //router.push(`/project/edit/${id}`)
  currentProjectId.value = id
  console.log("currentProjectId", currentProjectId.value)
  visible.value = true
}
const handleOk = (e: MouseEvent) => {
  console.log(e);
  visible.value = false;
};

const closeModal = () => {
  visible.value = false;
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除项目',
    content: '确定删除指定的项目？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Home/removeProject', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          notification.success({
            key: NotificationKeyCommon,
            message: `删除成功`,
          });
        } else {
          notification.error({
            key: NotificationKeyCommon,
            message: `删除失败`,
          });
        }
      })
    }
  });
}
</script>

<style lang="less" scoped>
.home-list {
  // padding:0 16px;
 
  .project-name{
    color: #447DFD;
    cursor: pointer;
  }
}
</style>