<template>
  <div class="user-main-list">
    <a-tabs>
      <a-tab-pane key="1" v-if="isAdmin" tab="成员">
        <Member :isAdmin="isAdmin"/>
    
      </a-tab-pane>
      <a-tab-pane key="2" tab="我的审批" force-render>
        <Audit/>
        </a-tab-pane
      >
      <a-tab-pane key="3" tab="我的申请" force-render>
        <Apply/>
        </a-tab-pane
      >
    </a-tabs>
  </div>
</template>
<script setup lang="ts">

import Member from "./member/member.vue";
import Audit from "./audit/audit.vue";
import Apply from "./apply/apply.vue";
import {onMounted, ref} from "vue";
import {useStore} from "vuex";
import {StateType as UserStateType} from "@/store/user";

const store = useStore<{ User: UserStateType }>();
const isAdmin = ref<boolean>(false)
const getUserSysRole = () => {
  isAdmin.value = store.state.User.currentUser.sysRoles.indexOf('admin') != -1
}
onMounted(() => {
  getUserSysRole();
  console.log(isAdmin)
});
</script>
<style lang="less" scoped>
.user-main-list {
  background: #fff;
}
</style>
