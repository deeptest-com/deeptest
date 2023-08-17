<template>
    <div class="editor-wrapper">
        <MonacoEditor 
            v-if="!isImage(data.contentType)"
            class="editor" 
            :value="data.content"
            :language="data.contentLang" 
            theme="vs" 
            :options="editorOptions" />
        <img v-else class="image" :src="'data:' + data.contentType + ';base64,' + data.content" />    
    </div>
</template>
<script setup lang="ts">
import { ref, defineProps, watch } from 'vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import { MonacoOptions } from "@/utils/const";

const props = defineProps({
    data: {
        type: Object,
        default: () => ({
            content: '',
            contentLang: '',
            contentType: '',
        })
    }
});

const editorOptions = ref(MonacoOptions);

const isImage = (type) => {
  return type && type.indexOf("image") > -1;
};

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

    .image {
      max-width: 100%;
      width: auto;
    }


    :deep(.monaco-editor) {
        height: 100% !important;
    }
}
</style>