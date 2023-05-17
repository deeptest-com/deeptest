<template>
  <div class="project-members">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="成员">
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
                  :disabled="currentUser.projectRoles[projectId] !== 'admin'"
                  >移除</a-button
                >
              </template>
            </a-table>
          </div>
        </a-card></a-tab-pane
      >
      <a-tab-pane key="2" tab="审批" force-render>
        <a-table
          row-key="id"
          :columns="auditColumns"
          :data-source="auditLst.list"
          :loading="loading"
          :pagination="{
            ...auditLst.pagination,
            onChange: (page) => {
              getAudits(page);
            },
            onShowSizeChange: (page, size) => {
              pagination1.pageSize = size;
              getAudits(page);
            },
          }"
        >
          <template #status="{ text }">
            {{ text == 0 ? "待审批" : text == 1 ? "已同意" : "已拒绝" }}
          </template>
          <template #action="{ record }">
            <a-button type="link" @click="() => audit(record.id)"
              >审批</a-button
            >
          </template>
        </a-table></a-tab-pane
      >
    </a-tabs>
  </div>
  <a-modal
    v-model:visible="auditModal"
    title="审批"
    @cancel="auditModal = false"
  >
    <p>是否通过该用户的审批？</p>

    <template #footer>
      <a-button key="back" danger @click="handleAudit(2)">拒绝</a-button>
      <a-button
        key="submit"
        type="primary"
        :loading="loading"
        @click="handleAudit(1)"
        >同意</a-button
      >
    </template>
  </a-modal>
  <EditPage :record="data" :title="title" :visible="visible" @cancal="cancal" />
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, Ref, ref } from "vue";
import { PaginationConfig, QueryParams, Project, Member } from "../data.d";
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
  getAuditList,
  doAudit,
} from "@/views/project/service";
import { StateType as UserStateType } from "@/store/user";
import EditPage from "../edit/invite.vue";
import { SelectTypes } from "ant-design-vue/lib/select";

const router = useRouter();
const store = useStore<{ Project: StateType; User: UserStateType }>();
const currentUser = computed<any>(() => store.state.User.currentUser);

let pagination = computed<PaginationConfig>(
  () => store.state.Project.queryResult.pagination
);

let queryParams = reactive<QueryParams>({
  keywords: "",
  enabled: "1",
  page: pagination.value.current,
  pageSize: pagination.value.pageSize,
});
let activeKey = ref("1");
const members = ref({});
const auditLst = ref({});
const auditModal = ref(false);
const auditId = ref(0);
const data = reactive<Member>({
  userId: 0,
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
const auditColumns = [
  {
    title: "申请人",
    dataIndex: "applyUserId",
    // slots: { customRender: "username" },
  },
  {
    title: "申请角色",
    dataIndex: "projectRoleName",
    // slots: { customRender: "role" },
  },
  {
    title: "申请原因",
    dataIndex: "description",
    width: 260,
    // slots: { customRender: "email" },
  },
  {
    title: "申请日期",
    dataIndex: "createdAt",
    // slots: { customRender: "email" },
  },
  {
    title: "状态",
    dataIndex: "status",
    slots: { customRender: "status" },
  },
  {
    title: "操作",
    key: "action",
    width: 50,
    slots: { customRender: "action" },
  },
];

const invite = () => {
  title.value = "邀请用户";
  data.email = "";
  data.userId = 0;
  visible.value = true;
  getSelectUserList();
};

onMounted(() => {
  console.log("onMounted");
  getMembers(1);
  getRoles();
  getAudits(1);
});
const pagination1 = ref({
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
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
    id: projectId,
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
const getAudits = (page: number) => {
  loading.value = true;

  getAuditList({
    pageSize: pagination1.value.pageSize,
    page: page,
  })
    .then((json) => {
      console.log("审批列表", json);
      if (json.code === 0) {
        // auditLst.value = json.data.result;
        auditLst.value = {
          current: 1,
          list: json.data.result || [],
          pagination: {
            ...pagination1.value,
            current: page,
            total: json.data.total || 0,
          },
        };
      }
    })
    .finally(() => {
      loading.value = false;
    });
};
const audit = (id: number) => {
  console.log("remove");
  auditModal.value = true;
  auditId.value = id;
};
const handleAudit = async (type: number) => {
  await doAudit({ id: auditId.value, status: type }).then((json) => {
    if (json.code === 0) {
      getAudits(1);
      notification.success({
        key: NotificationKeyCommon,
        message: `审批成功`,
      });
    } else {
      notification.error({
        key: NotificationKeyCommon,
        message: `审批失败`,
      });
    }
  });

  auditModal.value = false;
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

const visible = ref(false);

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

const title = ref("邀请用户");
const cancal = () => {
  visible.value = false;
  getMembers(1);
};

const handleChangeRole = async (val: any, record: any) => {
  await changeRole({
    projectId: projectId,
    projectRoleId: val,
    userId: record.id,
  });
};
</script>

<style lang="less" scoped>
.project-members {
}
</style>
