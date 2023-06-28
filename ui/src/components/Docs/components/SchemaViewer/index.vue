<template>
  <div class="body-schema">
    <SchemaEditor :value="content" :components="components"/>
    <div class="examples" v-if="examples?.length">
      <h4>示例</h4>
      <a-tabs :size="'small'">
        <a-tab-pane v-for="(example,index) in examples"
                    :key="index"
                    :tab="example.name">
          <MonacoEditor
              class="editor"
              :value="example.content"
              :language="'json'"
              :height="100"
              theme="vs"
              :options="{...MonacoOptions}"
          />
        </a-tab-pane>
      </a-tabs>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {computed, defineProps, defineEmits, ref, watch} from "vue";
import SchemaEditor from './schema';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';

const props = defineProps(['contentStr', 'examplesStr','components']);
const examples: any = ref([]);
const content: any = ref<any>(null);

watch(() => {
  return props.contentStr
}, (newVal) => {
  try {
    if (newVal) {
      content.value = JSON.parse(newVal);
    }
  }catch (e){
    console.log(e)
  }
}, {immediate: true})

watch(() => {
  return props.examplesStr
}, (newVal) => {
  try {
    if (newVal) {
      examples.value = JSON.parse(newVal);
    }
  }catch (e){
    console.log(e)
  }

}, {
  immediate: true
})

</script>


<style lang="less" scoped>
@import "var.less";

.body-schema {

  .examples{
    margin-top: 16px;
  }
}

</style>

