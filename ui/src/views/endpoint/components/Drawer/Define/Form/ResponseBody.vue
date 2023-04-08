<!-- 响应体定义 -->
<template>
  <!-- 增加响应体体 -->
  <a-row class="form-item-response-item">
    <a-col :span="4" class="form-label form-label-first">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">增加响应体</span>
    </a-col>
    <a-col :span="18">
      <a-select
          placeholder="请选择格式"
          :value="selectedCodeDetail.mediaType"
          @change="handleResBodyMediaTypeChange"
          style="width: 300px"
          :options="mediaTypesOpts"
      ></a-select>
    </a-col>
  </a-row>
  <!-- 增加响应体 - 描述  -->
  <a-row class="form-item-response-item" v-if="collapse">
    <a-col :span="4" class="form-label"></a-col>
    <a-col :span="18">
      <a-input placeholder="请输入描述" @change="handleResDescriptionChange" :value="selectedCodeDetail.description"/>
    </a-col>
  </a-row>
  <!-- 增加响应体 - scheme定义 -->
  <a-row class="form-item-response-item" v-if="collapse">
    <a-col :span="4" class="form-label"></a-col>
    <a-col :span="20">
      <SchemaEditor
          @generateFromJSON="generateFromJSON"
          @change="handleChange"
          @generateExample="handleGenerateExample"
          :tab-content-style="{width:'600px'}"
          :value="activeResBodySchema"/>
    </a-col>
  </a-row>
</template>
<script lang="ts" setup>
import {computed, defineEmits, defineProps, ref, watch,} from 'vue';
import {useStore} from "vuex";
import {mediaTypesOpts,} from '@/config/constant';
import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';
import {Endpoint} from "@/views/endpoint/data";
import SchemaEditor from '@/components/SchemaEditor/index.vue';

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const selectedCodeDetail = computed<any>(() => store.state.Endpoint.selectedCodeDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
// 是否折叠,默认展开
const collapse = ref(true);
const props = defineProps({});
const emit = defineEmits([]);
const activeResBodySchema: any = ref({
  content: null,
  examples: [],
});

watch(() => {
  return selectedCodeDetail?.value?.schemaItem?.content
}, (newVal, oldValue) => {
  activeResBodySchema.value.content = JSON.parse(newVal || 'null')
}, {immediate: true});

watch(() => {
  return selectedCodeDetail?.value?.examples
}, (newVal, oldValue) => {
  activeResBodySchema.value.examples = JSON.parse(newVal || '[]')
}, {immediate: true});

function handleResBodyMediaTypeChange(e) {
  selectedCodeDetail.value.mediaType = e;
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail?.value);
}

function handleResDescriptionChange(e) {
  selectedCodeDetail.value.description = e.target.value;
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail.value);
}

async function generateFromJSON(JSONStr: string) {
  activeResBodySchema.value.content = await store.dispatch('Endpoint/example2schema', {data: JSONStr});
}

async function handleGenerateExample(examples: any) {
  const res = await store.dispatch('Endpoint/schema2example',
      {data: JSON.stringify(activeResBodySchema.value.content)}
  );
  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res),
  };
  activeResBodySchema.value.examples.push(example);
}

function handleChange(json: any) {
  const {content, examples} = json;
  if (selectedCodeDetail?.value) {
    selectedCodeDetail.value.schemaItem.content = JSON.stringify(content);
    selectedCodeDetail.value.examples = JSON.stringify(examples);
    selectedCodeDetail.value.schemaItem.type = content.type;
  }
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail?.value);
}

</script>
<style lang="less" scoped>
.form-label {
  font-weight: bold;
}

.form-label-first {
  font-weight: bold;
  position: relative;
  left: -18px;
}

.form-item-response-item {
  margin-top: 16px;
  align-items: center;
}

.label-name {
  margin-left: 4px;
}
</style>
