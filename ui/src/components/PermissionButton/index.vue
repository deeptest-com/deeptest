<template>
  <!-- 权限按钮 -->
  <!-- 这里判断当前页面按钮是否在权限列表中,反之则提示用户 -->
  <template v-if="disabled">
    <a-tooltip :title="disabledTooltip" color="#1677ff">
      <div class="permission-btn">
        <slot name="before"></slot>
        {{ text }}
        <!-- 后置icon -->
        <slot name="after"></slot>
      </div>
    </a-tooltip>
  </template>
  <template v-else>
    <a-button
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
  </template>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, computed, ref, watch } from "vue";
import { useStore } from "vuex";
import { StateType as GlobalStateType } from "@/store/global";
import { PermissionButtonType } from "@/types/permission";

const props = defineProps<{
  code: String;
  text: String;
  type?: String;
  htmlType?: String;
  danger?: Boolean;
  size?: String;
  loading?: Boolean;
  dataCreateUser?: String;
  action?: String;
}>();

const emits = defineEmits(["handleAccess"]);
const disabled = ref<Boolean>(false);
const disabledTooltip = "暂无权限，请联系管理员";
const store = useStore<{ Global: GlobalStateType; User; ProjectGlobal }>();
const permissionButtonMap = computed<any[]>(
  () => store.state.Global.permissionButtonMap
);
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

// 判断权限方法
watch(
  () => {
    return permissionButtonMap.value;
  },
  (val: any) => {
    if (val && Object.keys(val).length > 0) {
      disabled.value = !hasPermission();
    }
  },
  { immediate: true }
);
</script>
<style scoped lang="less">
.permission-btn {
  line-height: 1.75;
  height: 32px;
  padding: 4px 15px;
  font-size: 14px;
  border-radius: 2px;
  display: inline-block;
  font-weight: 400;
  white-space: nowrap;
  text-align: center;
  color: rgba(0, 0, 0, 0.25);
  background: #f5f5f5;
  border-color: #d9d9d9;
  text-shadow: none;
  box-shadow: none;
  cursor: pointer;

  &.envDetail-btn {
    margin: 16px 0;
  }

  &.action-new {
    margin-right: 8px;
  }

  &.save-btn {
    margin-right: 16px;
  }
}
</style>
