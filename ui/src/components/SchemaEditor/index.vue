<script lang="ts" setup>
import {
  ArrowUpOutlined,
  ArrowDownOutlined,
  CopyOutlined,
  CloseOutlined,
  EditOutlined,
  DeleteOutlined,
  InfoCircleOutlined,
  PlusOutlined,
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";
import SchemaEditor from './schema';
import {message} from "ant-design-vue";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';

const props = defineProps<{
  value: object,
}>();

const emit = defineEmits<{
  (e: 'generateFromJSON', jsonStr?: string): void,
  (e: 'contentChange', jsonStr?: string): void,
  (e: 'exampleChange', jsonStr?: string): void,
  (e: 'schemaTypeChange', type?: string): void
}>();

const content: any = ref(null);
const examples: any = ref([]);

function addExample() {
  const data = {
    name: `Example ${examples.value.length + 1}`,
    content: `{
      "a":1
    }`
  };
  examples.value.push(data);
  activeExample.value = examples.value[examples.value.length - 1];
}


const activeExample: any = ref(null);

function clickExampleItem(item: any) {
  activeExample.value = item;
}

const activeTab = ref('schema');

function switchTab(tab) {
  activeTab.value = tab;
}

function copyExample() {
  message.info('复制成功');
}

const activeGenSchemaMode: any = ref(false);

function genSchema() {
  activeGenSchemaMode.value = true;
}

function deleteExample(value) {
  let index = examples.value.findIndex((item) => {
    return item.name === value.name;
  })
  examples.value.splice(index, 1);
  if (examples.value.length === 1) {
    activeExample.value = examples.value[0];
  }
  if (examples.value.length === 0) {
    activeExample.value = null;
  }
}

function handleExampleContentChange() {
  console.log('编辑器改变')
}

const exampleJsonStr = ref('');

function handleJSONDemoChange(val) {
  exampleJsonStr.value = val;
}


function cancalGen() {
  activeGenSchemaMode.value = false;
}

function generate() {
  activeGenSchemaMode.value = false;
  emit('generateFromJSON', exampleJsonStr.value);
}

const key = ref(0)


watch(() => {
  return props.value
}, (newVal: any) => {
  key.value++;
  content.value = newVal?.content || null;
  examples.value = newVal?.examples || [];
}, {
  immediate: true
});

watch(() => {
  return content.value
}, (newVal: any) => {
  // let newObj = JSON.parse(JSON.stringify(newVal));
  emit('contentChange', JSON.stringify(newVal));
  emit('schemaTypeChange', newVal.type);
}, {
  immediate: false,
  deep: true
});


watch(() => {
  return examples.value
}, (newVal: any) => {
  emit('exampleChange', JSON.stringify(newVal));
}, {
  immediate: false,
  deep: true
});


</script>

<template>
  <div class="tab-content">
    <div class="tab-header" v-if="!activeGenSchemaMode">
      <div>
        <a-button :type="activeTab === 'schema' ? 'link': 'text'" :size="'small'" @click="switchTab('schema')">Schema
        </a-button>
        <a-button :type="activeTab === 'examples' ? 'link': 'text'" :size="'small'" @click="switchTab('examples')">
          Examples
        </a-button>
        <a-button :type="activeTab === 'extensions' ? 'link': 'text'" :size="'small'" @click="switchTab('extensions')">
          Extensions
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
        <SchemaEditor :value="content"/>
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

          <a-button v-for="item in examples"
                    :key="item.name"
                    class="new-btn tab-btn"
                    size="small"
                    :type="activeExample?.name === item.name ? 'primary': 'text'"
                    @click="clickExampleItem(item)">
            {{ item.name }}
          </a-button>

        </div>
        <div class="right">
          <div v-if="!activeExample?.content" class="nodata-tip" title="Your operation has been executed">
            <InfoCircleOutlined class="tip-icon"/>
            <div class="tip-text">No Example. Click '+ New Example' to get started.</div>
          </div>
          <div class="activeExampleInfo" v-if="activeExample?.content">
            <div class="activeExampleInfo-header">
              <a-input class="input exampleName-input"
                       v-model:value="activeExample.name"
                       placeholder="Basic usage"/>
              <div class="btns">
                <a-button type="text" @click="copyExample(activeExample)">
                  <template #icon>
                    <CopyOutlined/>
                  </template>
                </a-button>
                <a-button type="text" @click="deleteExample(activeExample)">
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
      <!--::::扩展 -->
      <div class="tab-body-extensions" v-if="activeTab=== 'extensions'">
        <div class="left">
          <a-button class="new-btn" size="small" type="text">
            <template #icon>
              <PlusOutlined/>
            </template>
            New Extension
          </a-button>
        </div>
        <div class="right">

        </div>
      </div>
    </div>
    <div class="genSchemaFromCode" v-if="activeGenSchemaMode">
      <div class="btns">
        <a-button @click="generate">
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

<style lang="less" scoped>
@import "var.less";

.tab-content {
  //overflow: hidden;
  border: 1px solid @border-color;
  //overflow: hidden;
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

