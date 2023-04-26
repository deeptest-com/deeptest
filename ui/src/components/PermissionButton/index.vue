<template>
    <!-- 权限按钮 -->
    <!-- 这里判断当前页面按钮是否在权限列表中,反之则提示用户 -->
    <a-button 
        :type="type || 'default'" 
        :loading="loading || false" 
        :html-type="htmlType" 
        :size="size" 
        :danger="danger || false" 
        @click="handleClick">
            <!-- 前置icon -->
            <slot name="before"></slot>
            {{ text }}
            <!-- 后置icon -->
            <slot name="after"></slot>
    </a-button>
</template>
<script setup lang="ts">
import { message } from 'ant-design-vue';
import { defineProps, defineEmits, computed } from 'vue';
import { useStore } from 'vuex';
import { StateType as GlobalStateType } from "@/store/global";
import { PermissionButtonType } from '@/types/permission';

const props = defineProps<{
    code: String
    text: String
    type?: String
    htmlType?: String
    danger?: Boolean
    size?: String
    loading?: Boolean
}>();

const emits = defineEmits(['handleAccess']);
const store = useStore<{ Global: GlobalStateType }>();
const permissionButtonMap = computed<any[]>(() => store.state.Global.permissionButtonMap);


const handleClick = () => {
    const hasPermission = permissionButtonMap.value[PermissionButtonType[`${props.code}`]];
    console.log('%c ~~~~ hasPermission ~~~~', 'color: red;font-size: 24px', hasPermission);
    if (!hasPermission) {
        message.error('暂无权限,请联系管理员');
        return;
    }
    emits('handleAccess');
}

// 判断权限方法
</script>