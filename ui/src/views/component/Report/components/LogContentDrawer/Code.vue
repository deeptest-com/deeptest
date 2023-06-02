<template>
    <div class="editor-wrapper">
      <div style="border: 1px solid #f0f0f0; padding: 8px 0;">
        <MonacoEditor
            class="editor"
            :value="contentStr"
            :language="'json'"
            :height="600"
            theme="vs"
            :options="{...MonacoOptions}"
            @change="() => {

            }"
        />
      </div>
    </div>
</template>
<script setup lang="ts">
import {ref, defineProps, watch, computed} from 'vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import { MonacoOptions } from "@/utils/const";

const props = defineProps({
    content: {
        type: Object,
    }
});

const editorOptions = ref(MonacoOptions);

const contentStr = computed(() => {
    return JSON.stringify(props.content || null);
});


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
