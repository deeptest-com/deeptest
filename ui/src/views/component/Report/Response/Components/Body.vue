<template>
    <div class="editor-wrapper">
        <MonacoEditor 
            class="editor" 
            :value="data.content"
            :language="data.contentLang" 
            theme="vs" 
            :options="editorOptions" />
    </div>
</template>
<script setup lang="ts">
import { ref, defineProps, watch } from 'vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";

const props = defineProps({
    data: {
        type: Object,
        default: () => ({
            content: '',
            contentLang: ''
        })
    }
});

const editorOptions = ref(MonacoOptions);

watch(() => {
    return props.data;
}, val => {
    console.log(val);
}, {
    immediate: true,
    deep: true
})
</script>

<style scoped lang="less">
.editor-wrapper {
    width: 100%;
    height: calc(100vh - 120px);
    overflow-x: hidden;
    overflow-y: hidden;
    padding-top: 20px;

    &>div {
        height: 100%;
    }


    :deep(.monaco-editor) {
        height: 100% !important;
    }
}
</style>