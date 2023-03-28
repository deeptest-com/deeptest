<template>
    <a-form layout="inline" :model="formState">
        <a-form-item v-for="(schemaItem) in formSchemaList" :key=schemaItem.type>
            <span v-if="schemaItem.type === 'tooltip'" style="cursor: pointer;font-weight: bold">{{ schemaItem.text }}
                <a-tooltip placement="topLeft" arrow-point-at-center
                    :title="schemaItem.title">
                    <QuestionCircleOutlined class="icon"
                        style="position: relative;top:-8px; font-size: 12px;transform: scale(0.9)" />
                </a-tooltip>
            </span>
            <a-input v-if="schemaItem.type === 'input'" style="width: 150px" v-model:value="formState[schemaItem.stateName!]"
                :placeholder="schemaItem.placeholder" />
            <a-select v-if="schemaItem.type === 'select'" v-model:value="formState[schemaItem.stateName!]" mode="tags"
                style="width: 150px" :placeholder="schemaItem.placeholder" :options="schemaItem.options"
                @change="schemaItem.action">
            </a-select>
            <a-button v-if="schemaItem.type === 'button'" class="editable-add-btn" @click="schemaItem.action($event, formState)" type="primary" style="margin-bottom: 8px">{{
                schemaItem.text }}
            </a-button>
        </a-form-item>
    </a-form>
    <a-input-search v-if="needSearch" v-model:value="keyword" :placeholder="searchPlaceHolder" style="width: 200px" @search="handleSearch"/>
</template>
<script setup lang="ts">
import { ref, defineProps, reactive, defineEmits } from 'vue';
import { Schema } from '../../data';
const props = defineProps<{
    formSchemaList: Schema[]
    needSearch?: boolean
    handleOk?: any;
    searchPlaceHolder?: string;
}>()

const emits = defineEmits(['handleSearch']);

const keyword = ref('');

const setFormState = () => {
    const formObj = {};
    props.formSchemaList.forEach((e: Schema) => {
        if (e.stateName) {
            formObj[e.stateName] = e.valueType === 'string' ? '' : e.valueType === 'array' ? [] : '';
        }
    })
    return formObj;
}

const formState = reactive(setFormState());

const handleSearch = (e: string) => {
    emits('handleSearch', e);
}
</script>