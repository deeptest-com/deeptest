<template>
    <a-form layout="inline" :model="formState">
        <a-form-item v-for="(schemaItem) in formSchemaList" :key="schemaItem.type" :name="schemaItem.stateName" :rules="[{ required: schemaItem.required, message: schemaItem.message }]">
            <span v-if="schemaItem.type === 'tooltip'" style="cursor: pointer;font-weight: bold">{{ schemaItem.text }}
                <a-tooltip placement="topLeft" arrow-point-at-center
                    :title="schemaItem.title">
                    <QuestionCircleOutlined class="icon"
                        style="position: relative;top:-8px; font-size: 12px;transform: scale(0.9)" />
                </a-tooltip>
            </span>
            <a-input v-if="schemaItem.type === 'input'" style="width: 150px" v-model:value="formState[schemaItem.stateName!]"
                :placeholder="schemaItem.placeholder" />
            <a-select v-if="schemaItem.type === 'select'" v-model:value="formState[schemaItem.stateName!]" :mode="schemaItem.mode"
                style="width: 150px" :placeholder="schemaItem.placeholder" :options="options"
                @focus="onFocus($event, schemaItem)"
                @change="schemaItem.action">
            </a-select>
            <a-button v-if="schemaItem.type === 'button'" class="editable-add-btn" @click="schemaItem.action($event, formState)" type="primary" html-type="submit" style="margin-bottom: 8px">{{
                schemaItem.text }}
            </a-button>
        </a-form-item>
    </a-form>
    <a-input-search v-if="needSearch" v-model:value="keyword" :placeholder="searchPlaceHolder" style="width: 200px" @search="handleSearch"/>
</template>
<script setup lang="ts">
import { ref, defineProps, reactive, defineEmits, computed } from 'vue';
import { useStore } from 'vuex';
import {StateType as ProjectSettingStateType} from '../../store';
import { Schema } from '../../data';
const props = defineProps<{
    formSchemaList: Schema[]
    needSearch?: boolean
    handleOk?: any;
    searchPlaceHolder?: string;
}>()

const emits = defineEmits(['handleSearch']);
const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const userListOptions = computed<any>(() => store.state.ProjectSetting.userListOptions);

const keyword = ref('');
const options = reactive<any>([]);

const setFormState = () => {
    const formObj = {};
    props.formSchemaList.forEach((e: Schema) => {
        if (e.stateName) {
            formObj[e.stateName] = e.valueType === 'string' ? '' : [];
        }
        if (e.options) {
            options.value = e.options;
        }
    })
    return formObj;
}

const formState = reactive(setFormState());

const handleSearch = (e: string) => {
    emits('handleSearch', e);
}

const onFocus = async (e: any, schemaItem: Schema) => {
    if (schemaItem.focusType) {
        await store.dispatch('ProjectSetting/getUserOptionsList')
        options.value = userListOptions.value;
    }
}
</script>