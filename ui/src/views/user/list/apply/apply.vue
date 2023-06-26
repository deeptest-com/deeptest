
<template>
<a-card :bordered="false">
    <a-table
      row-key="id"
      :columns="applyColumns"
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
      }"
    >
      <template #status="{ text }">
        {{ text == 0 ? "待审批" : text == 1 ? "已同意" : "已拒绝" }}
      </template>
    </a-table>
  </a-card>
</template>

<script setup lang="ts">

import {  ref, watch,onMounted } from "vue";
import { useStore } from "vuex";

import { getAuditList, doAudit } from "@/views/project/service";

const store = useStore();

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
const applyColumns = [
{
    title: "ID",
    dataIndex: "id",
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
    type: activeKey.value == "2" ? 0 : 1,
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

onMounted(() => {
  getAudits(1);
});


</script>