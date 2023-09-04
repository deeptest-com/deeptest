<template>
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <a-radio-group v-model:value="language" size="small" v-if="data.contentLang=='json'">
            <a-radio-button :value="data.contentLang">Pretty</a-radio-button>
            <a-radio-button value="raw">Raw</a-radio-button>
          </a-radio-group>
        <span style="margin-left:5px;">{{data.contentLang.toUpperCase()}}</span>
        </a-col>
        <a-col flex="100px" class="dp-right">
        </a-col>
      </a-row>
    </div>
    <div class="editor-wrapper">
        <MonacoEditor 
            v-if="!isImage(data.contentType)"
            class="editor" 
            :value="content"
            :language="language" 
            theme="vs"
            :key = "language"
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

const language =  ref(props.data.contentLang)
const content = ref(props.data.content)

const editorOptions = ref(MonacoOptions);

const isImage = (type) => {
  return type && type.indexOf("image") > -1;
};

watch (()=>{return language.value} ,(val)=>{
  if (val == 'raw') {
    content.value = props.data.content.replace(/\r\n/g,'').replace(/\n/g,'').replace(/\s+/g,'')
  }
  console.log(content.value)

}, {immediate: true}) 

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
    padding-top: 5px;

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