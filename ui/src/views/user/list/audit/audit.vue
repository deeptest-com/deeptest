<template>
<a-card :bordered="false">
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
          pagination.pageSize = size;
          getAudits(page);
        },
        showTotal: (total) => {
          return `共 ${total} 条数据`;
        },
      }"
    >
      <template #status="{ text }">
        {{ text == 0 ? "待审批" : text == 1 ? "已同意" : "已拒绝" }}
      </template>
      <template #action="{ record }">
        <a-button
          v-if="record.status == 0"
          type="link"
          style="padding: 0"
          @click="() => audit(record.id)"
          >审批</a-button
        >
      </template>
    </a-table>
  </a-card>
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
</template>
<script setup lang="ts">

import {  User } from "../../data.d";
import { computed, ref, watch,onMounted } from "vue";
import { useStore } from "vuex";
import { Modal, notification } from "ant-design-vue";
import { NotificationKeyCommon } from "@/utils/const";
import debounce from "lodash.debounce";
import { getAuditList, doAudit } from "@/views/project/service";


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

const getRoleName = (val:any)=>{
  let rolesList = roles()
  return rolesList[val.text]
}
const auditLst = ref({});
const auditModal = ref(false);
const auditId = ref(0);

const auditColumns = [
{
    title: "ID",
    dataIndex: "id",
  },
  {
    title: "申请人",
    dataIndex: "applyUserName",

  },
  {
    title: "申请加入项目",
    dataIndex: "projectName",

  },
  {
    title: "申请角色",
    dataIndex: "projectRoleName",
    customRender:getRoleName
  },
  {
    title: "申请原因",
    dataIndex: "description",
  },
  {
    title: "申请日期",
    dataIndex: "createdAt",
    width: 240,

  },
  {
    title: "状态",
    dataIndex: "status",
    width: 100,
    slots: { customRender: "status" },
  },
  {
    title: "操作",
    key: "action",
    width: 50,
    slots: { customRender: "action" },
  },
];

const pagination = ref({
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
});
const loading = ref<boolean>(true);

const getAudits = (page: number) => {
  loading.value = true;

  getAuditList({
    pageSize: pagination.value.pageSize,
    page: page,
    type: 0,
  })
    .then((json) => {
      console.log("审批列表", json);
      if (json.code === 0) {
        // auditLst.value = json.data.result;
        auditLst.value = {
          current: 1,
          list: json.data.result || [],
          pagination: {
            ...pagination.value,
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

const visible = ref(false);

onMounted(() => {
  getAudits(1);
});

watch(
  () => activeKey.value,
  (val) => {
    if (val) {
      if (val == "2" || val == "3") {
        getAudits(1);
      }
    }
  },
  {
    immediate: true,
  }
);

</script>
