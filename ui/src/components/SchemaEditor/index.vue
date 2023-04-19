<template>
  <div class="tab-content" :style="tabContentStyle">
    <div class="tab-header" v-if="!activeGenSchemaMode">
      <div>
        <a-button :type="activeTab === 'schema' ? 'link': 'text'" :size="'small'" @click="switchTab('schema')">Schema
        </a-button>
        <a-button :type="activeTab === 'examples' ? 'link': 'text'" :size="'small'" @click="switchTab('examples')">
          Examples
        </a-button>
      </div>
      <div>
        <a-button class="gen-btn"
                  size="small"
                  @click="genSchema"
                  type="text">
          <template #icon>
            <PlusOutlined/>
          </template>
          Generate from JSON
        </a-button>
      </div>
    </div>
    <div class="tab-body" v-if="!activeGenSchemaMode">
      <!-- ::::Schema Tab -->
      <div class="tab-body-schema" v-if="activeTab=== 'schema'">
        <SchemaEditor
            :value="content"
            :refsOptions="refsOptions || []"
            @change="handleContentChange"
            :contentStyle="tabContentStyle"/>
      </div>
      <!--::::示例Tab -->
      <div class="tab-body-examples" v-if="activeTab=== 'examples'">
        <div class="left">
          <a-button class="new-btn" size="small" type="text" @click="addExample">
            <template #icon>
              <PlusOutlined/>
            </template>
            New Example
          </a-button>
          <a-button v-for="(item,index) in examples"
                    :key="item.name + index"
                    class="new-btn tab-btn"
                    size="small"
                    :type="activeExampleIndex === index ? 'primary': 'text'"
                    @click="clickExampleItem(index)">
            {{ item.name }}
          </a-button>
        </div>
        <div class="right">
          <div v-if="!activeExample?.content"
               class="nodata-tip"
               title="Your operation has been executed">
            <InfoCircleOutlined class="tip-icon"/>
            <div class="tip-text">No Example. Click '+ New Example' to get started.</div>
          </div>
          <div class="activeExampleInfo" v-if="activeExample?.content">
            <div class="activeExampleInfo-header">
              <a-input
                  class="input exampleName-input"
                  @change="handleExampleNameChange"
                  :value="activeExample.name"
                  placeholder="Basic usage"/>
              <div class="btns">
                <a-button type="text" @click="deleteExample">
                  <template #icon>
                    <DeleteOutlined/>
                  </template>
                </a-button>
              </div>
            </div>
            <div class="activeExampleInfo-body">
              <div style="border: 1px solid #f0f0f0; padding: 8px 0;">
                <MonacoEditor
                    class="editor"
                    :value="activeExample?.content"
                    :language="'json'"
                    :height="200"
                    theme="vs"
                    :options="{...MonacoOptions}"
                    @change="handleExampleContentChange"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="genSchemaFromCode" v-if="activeGenSchemaMode">
      <div class="btns">
        <a-button @click="generate" type="primary" :disabled="hasSyntaxError">
          <template #icon>
            <EditOutlined/>
          </template>
          生成
        </a-button>
        <a-button @click="cancalGen">
          <template #icon>
            <CloseOutlined/>
          </template>
          取消
        </a-button>
      </div>
      <div class="info">
        <a-alert message="Paste or write a JSON example below, then click Generate above to build a schema."
                 type="info" show-icon/>
      </div>
      <div style="border: 1px solid #f0f0f0; padding: 8px 0;">
        <MonacoEditor
            class="editor"
            :value="exampleJsonStr"
            :language="'json'"
            :height="200"
            theme="vs"
            :options="{...MonacoOptions}"
            @change="handleJSONDemoChange"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {
  CloseOutlined,
  EditOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  PlusOutlined,
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";
import SchemaEditor from './schema';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';

const props = defineProps<{
  tabContentStyle?: object,
  contentStr?: string,
  exampleStr?: string,
  schemeVisibleKey?: string | number,
  refsOptions?: Array<any>,
}>();

const emit = defineEmits<{
  (e: 'generateFromJSON', jsonStr?: string): void,
  (e: 'generateExample', jsonStr?: string): void,
  (e: 'change', json?: object): void,
}>();

const content: any = ref({type: 'object', properties: {}});
const examples: any = ref([]);

const activeExample: any = ref(null);
const activeExampleIndex: any = ref(0);

function addExample() {
  emit('generateExample', examples.value);
}

function clickExampleItem(index: number) {
  activeExampleIndex.value = index;
  activeExample.value = examples.value[index];
}

const activeTab = ref('schema');

function switchTab(tab) {
  activeTab.value = tab;
}

const activeGenSchemaMode: any = ref(false);

function genSchema() {
  activeGenSchemaMode.value = true;
}

function deleteExample() {
  examples.value.splice(activeExampleIndex.value, 1);
  activeExampleIndex.value = activeExampleIndex.value - 1 === -1 ? 0 : activeExampleIndex.value - 1;
  activeExample.value = examples.value[activeExampleIndex.value] || null;
}

function handleExampleNameChange(e) {
  activeExample.value.name = e.target.value;
  examples.value[activeExampleIndex.value].name = e.target.value;
}

function handleExampleContentChange(val) {
  activeExample.value.content = val;
}

function handleContentChange(val) {
  emit('change', {
    examples: examples.value,
    content: val
  });
}

const exampleJsonStr = ref('');
const hasSyntaxError = ref(true);

function handleJSONDemoChange(val, event, syntaxError) {
  exampleJsonStr.value = val;
  hasSyntaxError.value = !syntaxError;
}

function cancalGen() {
  activeGenSchemaMode.value = false;
}

function generate() {
  activeGenSchemaMode.value = false;
  activeTab.value = 'schema';
  emit('generateFromJSON', exampleJsonStr.value);
}


watch(() => {
  return props?.contentStr
}, (newVal: any) => {
  try {
    const obj = JSON.parse(newVal);
    content.value = obj?.type ? obj : {
      type: 'object'
    };
  }catch (e){
    console.log('e',e);
  }
}, {
  immediate: true,
  deep: true
});

watch(() => {
  return props?.exampleStr
}, (newVal: any) => {
  if(!newVal){
    return;
  }
  try {
    const obj = JSON.parse(newVal);
    examples.value = obj || [];
  }catch (e){
    console.log('e',e);
  }
}, {
  immediate: true,
  deep: true
});

watch(() => {
  return props.contentStr
}, (newVal: any) => {
  if(!newVal){
    return;
  }
  try {
    content.value = JSON.parse(newVal)
  }catch (e){
    console.log('e',e);
  }
});

watch(() => {
  return examples.value
}, (newVal: any) => {
  console.log('examples',examples.value)
  emit('change', {
    examples: newVal,
    content: content.value
  });
}, {
  immediate: false,
  deep: true
});

watch(() => {
  return examples.value.length
}, (newVal) => {
  activeExample.value = newVal > 0 ? examples.value[newVal - 1] : null;
}, {
  immediate: true
})

</script>


<style lang="less" scoped>
@import "var.less";

.tab-content {
  border: 1px solid @border-color;
  border-radius: 3px;
  width: @content-width;
}

.tab-header {
  border-bottom: 1px solid @border-color;
  height: 36px;
  line-height: 36px;
  display: flex;
  justify-content: space-between;
}

.tab-body-schema {
  //margin-left: -8px;
}

.tab-body-examples, .tab-body-extensions {
  display: flex;
  min-height: 200px;

  .left {
    flex: 1;
    border-right: 1px solid @border-color;

    .new-btn {
      text-align: left;
      width: 100%;
      height: 36px;
      line-height: 36px;
    }
  }

  .right {
    flex: 3;
    //padding-left:16px;
  }
}

.activeExampleInfo-header {
  display: flex;
  justify-content: space-between;
  height: 48px;
  align-items: center;

  .input {
    width: 80%;
  }

  .btns {

  }
}

.nodata-tip {
  padding: 24px;
  //display: flex;
  //justify-content: center;
  //flex-direction: column;

  .tip-text {
    text-align: center;
    margin: 16px auto;
    display: block;
    text-align: center;
  }

  .tip-icon {
    display: block;
    text-align: center;

    svg {
      width: 1.5em;
      height: 1.5em;
    }
  }
}

.exampleName-input {
  border: none;
  height: 24px;
  //margin-left: 16px;

  &:hover {
    border: 1px solid #1aa391;
  }
}

.genSchemaFromCode {
  .btns {
    height: 36px;
    line-height: 36px;
    margin-top: 8px;

    button {
      margin-left: 8px;
    }
  }

  .info {
    margin: 8px;
  }
}

</style>

