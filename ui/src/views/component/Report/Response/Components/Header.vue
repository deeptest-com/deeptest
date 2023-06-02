<template>
    <div class="headers">
        <a-table :dataSource="data" :columns="columns"/>
    </div>
</template>
<script setup lang="ts">
import { defineProps ,watch} from 'vue';

const props = defineProps<{
    data: Array<any>
}>();

watch(() => {
    return props.data;
}, () => {
    console.log('6666666',props.data)
}, {
    immediate: true
})
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
