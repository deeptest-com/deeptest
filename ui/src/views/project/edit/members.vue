<template>
  <div class="project-members">
    <a-card :bordered="false">
      <template #title>
        <a-button type="primary" @click="() => invite()">邀请</a-button>
      </template>
      <template #extra>
        <a-input-search
          @change="onSearch"
          @search="onSearch"
          v-model:value="queryParams.keywords"
          placeholder="输入关键字搜索"
          style="width: 270px; margin-left: 16px"
        />
      </template>

      <div>
        <a-table
          row-key="id"
          :columns="columns"
          :data-source="members.list"
          :loading="loading"
          :pagination="{
            ...pagination,
            onChange: (page) => {
              getMembers(page);
            },
            onShowSizeChange: (page, size) => {
              pagination.pageSize = size;
              getMembers(page);
            },
          }"
        >
          <template #username="{ text }">
            {{ text }}
          </template>

          <template #email="{ text }">
            {{ text }}
          </template>

          <template #role="{ record }">
            <div class="customTitleColRender">
              <a-select
                :value="record.roleId"
                style="width: 100px"
                :size="'small'"
                placeholder="请选中角色"
                @change="
                  (val) => {
                    handleChangeRole(val, record);
                  }
                "
              >
                <a-select-option
                  v-for="(option, key) in roles"
                  :key="key"
                  :value="option.id"
                  >{{ option.label }}</a-select-option
                >
              </a-select>
            </div>
          </template>

          <template #action="{ record }">
            <a-button
              type="link"
              @click="() => remove(record.id)"
              :disabled="currentUser.projectRoles[projectId] !== 'admin' && currentUser.sysRoles.indexOf('admin') === -1"
              >移除</a-button
            >
          </template>
        </a-table>
      </div>
    </a-card>
  </div>
  <EditPage :visible="inviteVisible" @ok="ok"  @cancel="inviteVisible = false"/>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref,
watch } from "vue";
import { PaginationConfig, Project, Member } from "../data.d";
import { useStore } from "vuex";

import { StateType } from "../store";
import debounce from "lodash.debounce";
import { useRouter } from "vue-router";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import {
  queryMembers,
  removeMember,
  changeRole,
} from "../service";
import { StateType as UserStateType } from "@/store/user";
import EditPage from "../edit/invite.vue";
import { SelectTypes } from "ant-design-vue/lib/select";
import {QueryParams} from "@/types/data";
import {inviteUser} from "@/views/user/info/service";

const router = useRouter();
const store = useStore<{ Project: StateType; User: UserStateType,ProjectGlobal }>();
const currentUser = computed<any>(() => store.state.User.currentUser);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
let pagination = computed<PaginationConfig>(
  () => store.state.Project.queryResult.pagination
);

let queryParams = reactive<QueryParams>({
  keywords: "",
  enabled: "1",
  page: pagination.value.current,
  pageSize: pagination.value.pageSize,
});

const members = ref({});

const data = reactive<Member>({
  userId: "",
  email: "",
  roleName: "",
  username: "",
});

const projectId = Number(window.localStorage.getItem("currentProjectId"));
const columns = [
  {
    title: "序号",
    dataIndex: "index",
    width: 80,
    customRender: ({ text, index }: { text: any; index: number }) =>
      (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: "用户名",
    dataIndex: "username",
    slots: { customRender: "username" },
  },
  {
    title: "角色",
    dataIndex: "role",
    slots: { customRender: "role" },
  },
  {
    title: "邮箱",
    dataIndex: "email",
    slots: { customRender: "email" },
  },
  {
    title: "操作",
    key: "action",
    width: 260,
    slots: { customRender: "action" },
  },
];


onMounted(() => {
  console.log("onMounted");
  getMembers(1);
  getRoles()
  store.dispatch("User/fetchUserProjectRole");
});

const initState: StateType = {
  queryResult: {
    list: [],
    pagination: {
      total: 0,
      current: 1,
      pageSize: 10,
      showSizeChanger: true,
      showQuickJumper: true,
    },
  },
  detailResult: {} as Project,
  queryParams: {},
};

const loading = ref<boolean>(true);
const getMembers = (page: number) => {
  loading.value = true;

  queryMembers({
   // id: projectId,
    keywords: queryParams.keywords,
    pageSize: pagination.value.pageSize,
    page: page,
  })
    .then((json) => {
      if (json.code === 0) {
        members.value = {
          ...initState.queryResult,
          list: json.data.result || [],
          pagination: {
            ...initState.queryResult.pagination,
            current: page,
            pageSize: pagination.value.pageSize,
            total: json.data.total || 0,
          },
        };
      }
    })
    .finally(() => {
      loading.value = false;
    });
};

const remove = (userId: number) => {
  console.log("remove");

  Modal.confirm({
    title: "移除成员",
    content: "确定移除指定的项目成员？",
    okText: "确认",
    cancelText: "取消",
    onOk: async () => {
      removeMember(userId, projectId).then((json) => {
        if (json.code === 0) {
          getMembers(queryParams.page);

          notification.success({
            key: NotificationKeyCommon,
            message: `移除成功`,
          });
        } else {
          notification.error({
            key: NotificationKeyCommon,
            message: `移除失败`,
          });
        }
      });
    },
  });
};

const onSearch = debounce(() => {
  getMembers(1);
}, 500);

const inviteVisible = ref(false);

const roles = computed<SelectTypes["options"]>(() => store.state.Project.roles);

//角色列表
const getRoles = () => {
  store.dispatch("Project/getRoles");
  return;
};

const getSelectUserList = () => {
  store.dispatch("Project/getNotExistedUserList", projectId);
  return;
};

const invite = () => {
  inviteVisible.value = true;
  console.log( inviteVisible.value)
  getSelectUserList();
};


const ok= async (modelRef:any,callback:any)=>{
  inviteVisible.value = false;
   await inviteUser(modelRef, projectId).then((json) => {
      if (json.code === 0) {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存成功`,
        });
      } else {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存失败`,
        });
      }
      // close()
    })
  callback()
  getMembers(1);
}

const handleChangeRole = async (val: any, record: any) => {
  await changeRole({
    projectId: projectId,
    projectRoleId: val,
    userId: record.id,
  });
};


// 实时监听项目/服务 ID，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value.id;
}, async (newVal) => {
  if (newVal) {
    getMembers(1);
  }
}, {
  immediate: true
})


</script>

<style lang="less" scoped>
.project-members {
}
</style>
