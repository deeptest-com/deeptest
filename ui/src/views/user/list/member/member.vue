<template>
    <a-card :bordered="false">
    <template #title>
      <a-button type="primary" @click="() => edit(0)">新建用户</a-button>
    </template>
    <template #extra>
      <a-input-search
        @change="onSearch"
        @search="onSearch"
        v-model:value="queryParams.username"
        placeholder="输入用户名搜索"
        style="width: 270px; margin-left: 16px"
      />
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
            pagination.pageSize = size;
            getList(page);
          },
          showTotal: (total) => {
            return `共 ${total} 条数据`;
          },
        }"
      >
        <template #role_ids="{ record }">
          <div class="customTitleColRender">
            <a-select
                :value="record.role_ids"
                mode="multiple"
                style="width: 100px"
                :size="'small'"
                placeholder="请选中角色"
                @change="
                  (val) => {
                   updateSysRole(val, record);
                  }
                "
            >
              <a-select-option
                  v-for="(option, key) in sysRoles"
                  :key="key"
                  :value="option.id"
              >{{ option.displayName }}</a-select-option
              >
            </a-select>
          </div>
        </template>
        <template #action="{ record }">
          <a-tooltip>
            <template v-if="!isAdmin" #title>暂无权限，请联系管理员</template>
            <a-button type="link" :disabled="!isAdmin" @click="() => edit(record.id)">
              <span>编辑</span>
            </a-button>
          </a-tooltip>
          <a-tooltip>
            <template v-if="!isAdmin" #title>暂无权限，请联系管理员</template>
            <a-button type="link" :disabled="!isAdmin" @click="() => remove(record.id)">
              <span>删除</span>
            </a-button>
          </a-tooltip>
        </template>
      </a-table>
    </div>
  </a-card>
  <a-modal
    v-model:visible="visible"
    @ok="handleOk"
    width="700px"
    :footer="null"
  >
    <EditPage
      :currentUserId="currentUserId"
      :getList="getList"
      :closeModal="closeModal"
      :sysRoles="sysRoles"
    />
  </a-modal>
</template>

<script setup lang="ts">
import { PaginationConfig, QueryParams, User } from "../../data.d";
import { computed, onMounted, reactive, ref, watch } from "vue";
import { useStore } from "vuex";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import debounce from "lodash.debounce";
import { useRouter } from "vue-router";
import EditPage from "../../edit/edit.vue";
import { defineProps } from 'vue'
import {SelectTypes} from "ant-design-vue/lib/select";
import {notifyError, notifySuccess} from "@/utils/notify";

const props = defineProps(['isAdmin'])

const router = useRouter();
const store = useStore();

const list = computed<User[]>(() => store.state.UserInternal.queryResult.list);
const roles = ()=>{
  let rolesList = {}
  store.state.Project.roles.forEach((item:any)=>{
    rolesList[item.name] = item.displayName
  })
  return rolesList
}
let activeKey = ref("1");
let pagination = computed<PaginationConfig>(
  () => store.state.UserInternal.queryResult.pagination
);
let queryParams = reactive<QueryParams>({
  username: "",
  page: pagination.value.current,
  pageSize: pagination.value.pageSize,
});

const getRoleName = (val:any)=>{
  let rolesList = roles()
  return rolesList[val.text]
}
const sysRoles = computed<SelectTypes["options"]>(()=>store.state.SysRole.roles);
const columns = [
  {
    title: "序号",
    dataIndex: "index",
    align: "center",
    customRender: ({ text, index }: { text: any; index: number }) =>
      (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: "用户名",
    dataIndex: "username",
    align: "center",
  },
  {
    title: "姓名",
    dataIndex: "name",
    align: "center",
  },
  {
    title: "角色",
    dataIndex: "role_ids",
    slots: { customRender: "role_ids" },
  },
  {
    title: "邮箱",
    dataIndex: "email",
    align: "center",
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    slots: { customRender: "action" },
  },
];

onMounted(() => {
  getList(1);
  getRoles();
  getSysRoles()
  // getAudits(1);
});

const pagination1 = ref({
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
});
const loading = ref<boolean>(true);
const getList = async (current: number): Promise<void> => {
  loading.value = true;

  await store.dispatch("UserInternal/queryUser", {
    username: queryParams.username,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false;
};

const visible = ref(false);
const currentUserId = ref(0);
const edit = (id: number) => {
  currentUserId.value = id;
  visible.value = true;
};


const onSearch = debounce(() => {
  getList(1);
}, 500);

//角色列表
const getRoles = () => {
  store.dispatch("Project/getRoles");
  return;
};

const getSysRoles = () => {
  store.dispatch("SysRole/getAllRoles")
}

const closeModal = () => {
  visible.value = false;
  getList(1);
};

const remove = (id: number) => {
  console.log("remove");

  Modal.confirm({
    title: "删除用户",
    content: "确定删除指定的用户？",
    okText: "确认",
    cancelText: "取消",
    onOk: async () => {
      store.dispatch("UserInternal/removeUser", id).then((res) => {
        console.log("res", res);
        if (res === true) {
          notifySuccess(`删除成功`);
        } else {
          notifyError(`删除失败`);
        }
      });
    },
  });
};

const updateSysRole = async (val: any, record: any) => {
  await store.dispatch("UserInternal/updateSysRole",{userId:record.id, roleIds:val})
};
</script>
