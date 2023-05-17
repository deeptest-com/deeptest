<template>
    <div class="headers">
        <a-table
            :data="data"
            :columns="columns" />
    </div>
</template>
<script setup lang="ts">
import { defineProps } from 'vue';

defineProps<{
    data: Array<any>
}>();

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
    },
    {
        title: 'Value',
        dataIndex: 'value',
    }
];
</script>
<style scoped lang="less">
.headers {
    padding-top: 20px;
}
</style>