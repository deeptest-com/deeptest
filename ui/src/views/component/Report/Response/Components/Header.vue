<template>
    <div class="headers">
        <a-table :dataSource="data" :columns="columns"/>
    </div>
</template>
<script setup lang="ts">
import {computed, defineProps} from 'vue';

const props = defineProps<{
    data: Array<any>
}>();

// 过滤掉没有 name=''  的header , 否则会导致antv表格组件渲染出错
const list = computed(() => {
    return props?.data.filter(item => {
        return item.name;
    })
});

const columns = [
    {
        title: 'name',
        dataIndex: 'name',
    },
    {
        title: 'value',
        dataIndex: 'value',
    }
];
</script>
<style scoped lang="less">
.headers {
    padding-top: 20px;
}
</style>
