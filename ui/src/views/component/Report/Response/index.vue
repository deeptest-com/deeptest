<template>
    <a-drawer 
        :visible="responseDrawerVisible" 
        :closable="true" :width="1000" 
        :bodyStyle="{ padding: '24px' }"
        @close="onClose">
        <template #title>
           <span>接口运行结果</span>
        </template>
        <div class="drawer-content">
            <a-tabs v-model:activeKey="activeKey">
                <a-tab-pane key="Body" tab="Body">
                    <BodyInfo />
                </a-tab-pane>
                <a-tab-pane key="Cookie" tab="Cookie" force-render>
                    <CookieInfo />
                </a-tab-pane>
                <a-tab-pane key="Header" tab="Header">
                    <HeaderInfo />
                </a-tab-pane>
                <a-tab-pane key="Console" tab="控制台">
                    <ConsoleInfo />
                </a-tab-pane>
                <a-tab-pane key="Request" tab="实际请求">
                    <RequestInfo />
                </a-tab-pane>
            </a-tabs>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref } from 'vue';
import { BodyInfo, ConsoleInfo, CookieInfo, RequestInfo, HeaderInfo } from './Components/index';

defineProps({
    responseDrawerVisible: {
        type: Boolean
    }
});

const emits = defineEmits(['onClose']);

const activeKey = ref('Body');

function onClose() {
    emits('onClose');
}
</script>