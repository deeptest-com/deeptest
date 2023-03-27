<template>
    <a-modal
        v-model:visible="modalVisible"
        title="关联服务"
        @ok="handleOk">
        <a-form-item class="selecte-default select-service" :labelCol="{span: 6}" :wrapperCol="{span: 16}" label="请选择服务">
        <a-select
            v-model:value="selectedValue"
            :options="modalOptions"
            placeholder="请选择服务"
            style="width: 200px"/>
        </a-form-item>

    </a-modal>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';
import { any } from 'vue-types';

const props = defineProps({
    modalVisible: {
        type: Boolean,
        required: true
    },
    selectedValue: {
        type: String,
        required: true
    },
    modalOptions: {
        type: Array,
        required: true
    },
    selectedClass: {
        type: String,
        required: false
    },
    placeholder: {
        type: String,
        required: false
    },
    handleOk: {
        type: any,
        required: true
    }

});
</script>