<!-- 结果antdv 的form 以及传入formConfig动态生成form表单，并暴露出部分方法，方便父组件使用 -->
<template>
    <a-form :model="formState" layout="inline">
        <template v-for="(item, index) in formConfig" :key="index">
            <a-form-item :name="item.name" :label="item.label" v-bind="validateInfos[item.modelName!]">
                <span v-if="item.type === 'tooltip'" style="cursor: pointer;font-weight: bold">{{ item.text }}
                    <a-tooltip placement="topLeft" arrow-point-at-center :title="item.title">
                        <QuestionCircleOutlined class="icon"
                            style="position: relative;top:-8px; font-size: 12px;transform: scale(0.9)" />
                    </a-tooltip>
                </span>
                <a-input v-if="item.type === 'input'" :style="item.attrs || 'width: 150px'"
                    v-model:value="formState[item.modelName]" :placeholder="item.placeholder" />
                <a-select v-if="item.type === 'select' && item.mode" v-model:value="formState[item.modelName]" :mode="item.mode"
                    :style="item.attrs || 'width: 200px'" :placeholder="item.placeholder" :options="item.options" allowClear>
                </a-select>
                <a-select v-if="item.type === 'select' && !item.mode" v-model:value="formState[item.modelName]"
                    :style="item.attrs || 'width: 200px'" :placeholder="item.placeholder" :options="item.options" allowClear>
                </a-select>
                <a-button v-if="item.type === 'button'" class="editable-add-btn" type="primary" html-type="submit"
                    style="margin-bottom: 8px" @click="onSubmit">
                    {{ item.text }}
                </a-button>
            </a-form-item>
        </template>
    </a-form>

    <a-input-search v-if="showSearch" v-model:value="keyword" :placeholder="searchPlaceholder" style="width: 200px"
        @search="handleSearch" />
</template>
<script setup lang="ts">
import { reactive, defineProps, toRaw, ref, defineEmits } from 'vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';
import { Form } from 'ant-design-vue';

const useForm = Form.useForm;

const props = defineProps<{
    formConfig: FormItem[]
    rules: any
    showSearch?: boolean
    searchPlaceholder?: string
}>();

const emits = defineEmits(['handleOk', 'handleSearch']);

interface FormItem {
    type: string; // 表单类型 input | selectMenuItem | button | or any more
    modelName?: string; // 表单model key值
    options?: any[]; // 表单类型为select时传入
    label?: string; // 是否需要label
    name?: string; // 是否需要name
    text?: string; // 表单文案  按钮或者 checkbox 等需要固定文案
    mode?: string; // 表单类型 select时 传入mode
    valueType?: string; // 表单值类型
    placeholder?: string; // input/ checkbox/ selectMenuItem 需要placeholder提示语
    attrs?: any; // 表单是否有自定义样式属性
    title?: string;
}

const getFormState = () => {
    const formObj = {};
    props.formConfig.forEach((formItem: FormItem) => {
        if (formItem.modelName) {
            formObj[formItem.modelName] = (formItem.type === 'select' && formItem.valueType === 'string') ? null : (formItem.type === 'select' && formItem.valueType === 'array') ? [] : '';
        }
    })
    return formObj;
}


const formState = reactive<any>(getFormState());
const rulesRef = reactive<any>(props.rules);
const keyword = ref<string>('');


const { validate, validateInfos ,resetFields} = useForm(formState, rulesRef);

const onSubmit = () => {
    validate()
        .then(() => {
            console.log(toRaw(formState));
            emits('handleOk', toRaw(formState));
            resetFields();
        })
        .catch(err => {
            console.log('捕捉表单错误信息：', err);
            // const { errorFields = [] } = err;
            // { name, errors: ['服务名称不能为空'] } = errorFileds[0]
        })
}

const handleSearch = (e: string) => {
    emits('handleSearch', e);
}

</script>
<style scoped lang="less">
:deep(.ant-form-inline .ant-form-item-with-help) {
    margin-bottom: 0 !important ;
    display: none;
}

:deep(.ant-form-item-explain.ant-form-item-explain-success) {
    display: none;
}

:deep(.ant-row.ant-form-item.ant-form-item-has-error.ant-form-item-with-help) {
    margin-bottom: 0 !important;
}

:deep(.ant-row.ant-form-item.ant-form-item-with-help.ant-form-item-has-success) {
    margin-bottom: 0 !important;
}
</style>
