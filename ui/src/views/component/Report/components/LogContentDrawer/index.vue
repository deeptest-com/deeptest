<template>
    <a-drawer
        :visible="visible"
        :closable="true" :width="1000"
        :bodyStyle="{ padding: '24px', height: '100%' }"
        @close="onClose">
        <template #title>
           <span>执行详细数据</span>
        </template>
        <div class="drawer-content">
           <Code :content="content"/>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, watch, computed } from 'vue';
import Code from './Code.vue';
const props = defineProps({
    visible: {
        type: Boolean
    },
    data: {
        type: Object
    }
});

const emits = defineEmits(['onClose']);
const content:any = ref(null);

function onClose() {
    emits('onClose');
}

watch(() => {
    return props.visible;
}, () => {
    content.value = props.data;
}, {
    immediate: true,
    deep: true
})
</script>

<style scoped lang="less">
.drawer-content {
    height: calc(100% - 46px);
}
</style>
