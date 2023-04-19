<template>
  <div class="list">

    <a-card :bordered="false">
      <template #title>
        测试场景贡献排名
        </template>
      <template #extra>
       <a-radio-group v-model:value="timeframe" button-style="solid">
        <a-radio-button value="a">当月</a-radio-button>
        <a-radio-button value="b">所有</a-radio-button>
       
      </a-radio-group>
      </template>

      <div>
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
        <div class="project-name" @click="goProject(record.id)">
          {{text}}
        </div>
      </template>
       
    

        </a-table>
      </div>
    </a-card>
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
const store = useStore<{ ProjectGlobal: ProjectStateType, workbench: StateType, User: UserStateType }>();
const projects = computed<any>(() => store.state.ProjectGlobal.projects);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currentUser = computed<any>(() => store.state.User.currentUser);
const list = computed<any[]>(() => store.state.workbench.queryResult.list);
const projectId = +router.currentRoute.value.params.id
const timeframe = ref<string>('a');
let pagination = computed<PaginationConfig>(() => store.state.workbench.queryResult.pagination);
let queryParams = reactive<QueryParams>({
 projectId:projectId,cycle:0,
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
    title: '排名',
    dataIndex: 'project_name',
  },
  {
    title: '用户名',
    dataIndex: 'project_chinese_name',
    slots: {customRender: 'name'},
  },

  {
    title: '较上周',
    dataIndex: 'admin_user',
  },

  {
    title: '用例数',
    dataIndex: 'scenario_total',
  },
  {
    title: '场景数',
    dataIndex: 'coverage',
  },
  {
    title: '最近更新日期',
    dataIndex: 'exec_total',
  },
 
];

onMounted(() => {
  console.log('onMounted',currentUser.value.id)
  getList(1);
})

const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;
  await store.dispatch('workbench/queryRanking', {
  
    projectId: queryParams.projectId,
    cycle: queryParams.cycle,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
}

function goProject(id:number){
  // TODO
  //  router.push(`/project/index/${id}`)

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
      store.dispatch('workbench/removeProject', id).then((res) => {
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
.list {
  // padding:0 16px;
  :deep(.ant-card-head .ant-tabs-bar){
    border-bottom:none ;
  }
  .project-name{
    color: #447DFD;
    cursor: pointer;
  }
}
</style>