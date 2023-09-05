<template>
  <!-- 权限按钮 -->
  <!-- 这里判断当前页面按钮是否在权限列表中,反之则提示用户 -->
  <a-tooltip :title="disabled ? disabledTooltip : null" color="#1677ff">
    <a-button
      class="permission-btn"
      :disabled="disabled"
      :type="type || 'default'"
      :loading="loading || false"
      :html-type="htmlType"
      :size="size"
      :danger="danger || false"
      @click="
        (e) => {
          handleClick(e);
        }
      "
    >
      <!-- 前置icon -->
      <slot name="before"></slot>
      {{ text }}
      <!-- 后置icon -->
      <slot name="after"></slot>
    </a-button>
  </a-tooltip>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, computed, ref, watch } from "vue";
import { useStore } from "vuex";
import { StateType as GlobalStateType } from "@/store/global";
import { PermissionButtonType } from "@/types/permission";

const props = defineProps<{
  code: String;
  text: String;
  disabled?: Boolean;
  type?: String;
  htmlType?: String;
  danger?: Boolean;
  size?: String;
  loading?: Boolean;
  dataCreateUser?: String;
  action?: String;
  tip?: String;
}>();

const emits = defineEmits(["handleAccess"]);
const disabledTooltip = computed(() => props.tip || "暂无权限，请联系管理员");
const store = useStore<{ Global: GlobalStateType; User; ProjectGlobal }>();
const permissionButtonMap = computed<any[]>(
  () => store.state.Global.permissionButtonMap
);
const disabled = computed(() => {
  console.log(props);
  if (props.disabled) {
    return true;
  }
  if (permissionButtonMap.value && Object.keys(permissionButtonMap.value).length > 0) {
    const permission = !hasPermission();
    return permission;
  }
  return false;
});
const currentUser = computed<any>(() => store.state.User.currentUser);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const handleClick = (e) => {
  e.preventDefault();
  emits("handleAccess");
};

const hasPermission = () => {
  // let hasPermissionButton = false
  // let filterUserName = ["undefined", "", currentUser.value.username]
  // if (currentUser.value.sysRoles.indexOf('admin') != -1) {
  //   return true
  // }
  // if (currentUser.value.projectRoles[currProject.value.id] == 'admin') {
  //   return true
  // }
  // if (permissionButtonMap.value[PermissionButtonType[`${props.code}`]] && props.action != 'delete') {
  //   return true
  // }
  // if (permissionButtonMap.value[PermissionButtonType[`${props.code}`]] && props.action == 'delete' && filterUserName.indexOf(props.dataCreateUser) > -1)  {
  //   return true
  // }
  if (!props.code) {
    return true;
  }
  const hasPermissionButton =
    permissionButtonMap.value[PermissionButtonType[`${props.code}`]];
  return hasPermissionButton;
};

</script>
<style scoped lang="less">
.permission-btn {
  
  &.envDetail-btn {
    margin-bottom: 16px;
  }

  &.action-new {
    margin-right: 8px;
  }

  &.save-btn {
    margin-right: 16px;
  }
}
</style>
