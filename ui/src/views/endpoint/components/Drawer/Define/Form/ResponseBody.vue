<!-- 响应体定义 -->
<template>
  <!-- 增加响应体体 -->
  <a-row class="form-item-response-item">
    <a-col :span="3" class="form-label form-label-first">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">增加响应体</span>
    </a-col>
    <a-col :span="21"/>
  </a-row>

  <a-row class="form-item-response-item" v-if="collapse">
    <a-col :span="1"/>
    <a-col :span="21">
      <a-select
          placeholder="请选择响应格式"
          :value="selectedCodeDetail.mediaType || null"
          @change="handleResBodyMediaTypeChange"
          style="width: 400px"
          :options="mediaTypesOpts.filter(item => !item.disabled)"
      ></a-select>
    </a-col>
  </a-row>

  <!-- 增加响应体 - scheme定义 -->
  <a-row class="form-item-response-item form-item-response-item-con" v-if="collapse">
    <a-col :span="1" class="form-label"></a-col>
    <a-col :span="23">
      <SchemaEditor
          @generateFromJSON="generateFromJSON"
          @changeContent="changeContent"
          @changeExamples="changeExamples"
          :serveId="currServe.id"
          :refsOptions="refsOptions"
          :contentStr="contentStr"
          :exampleStr="exampleStr"
          @generateExample="handleGenerateExample"
          :tab-content-style="{width:'720px'}"
          :value="activeResBodySchema"/>
    </a-col>
  </a-row>
</template>
<script lang="ts" setup>
import {computed, defineEmits, defineProps, onMounted, ref, watch,} from 'vue';
import {useStore} from "vuex";
import {mediaTypesOpts,} from '@/config/constant';
import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';
import {Endpoint} from "@/views/endpoint/data";
import SchemaEditor from '@/components/SchemaEditor/index.vue';

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User, ServeGlobal }>();
const selectedCodeDetail = computed<any>(() => store.state.Endpoint.selectedCodeDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);

// 是否折叠,默认展开
const collapse = ref(true);
const props = defineProps({});
const emit = defineEmits([]);
const activeResBodySchema: any = ref({
  content: null,
  examples: [],
});
const contentStr = ref('');
const exampleStr = ref('');

watch(() => {
  return selectedCodeDetail?.value?.schemaItem?.content
}, (newVal, oldValue) => {
  const str = newVal || 'null';
  activeResBodySchema.value.content = JSON.parse(str)
  contentStr.value = str;
}, {immediate: true});

watch(() => {
  return selectedCodeDetail?.value?.examples
}, (newVal, oldValue) => {
  activeResBodySchema.value.examples = JSON.parse(newVal || '[]')
  exampleStr.value = JSON.stringify(activeResBodySchema.value.examples);
}, {immediate: true});

function handleResBodyMediaTypeChange(e) {
  selectedCodeDetail.value.mediaType = e;
  store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail?.value);
}

async function generateFromJSON(JSONStr: string) {
  activeResBodySchema.value.content = await store.dispatch('Endpoint/example2schema', {data: JSONStr});
  contentStr.value = JSON.stringify(activeResBodySchema.value.content);
}

async function handleGenerateExample(examples: any) {
  const content = contentStr.value;
  const res = await store.dispatch('Endpoint/schema2example',
      {data: content, serveId: currServe.value.id,}
  );
  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res),
  };
  if(!activeResBodySchema.value?.examples) {
    activeResBodySchema.value.examples = [];
  }
  activeResBodySchema.value.examples.push(example);
  exampleStr.value = JSON.stringify(activeResBodySchema.value.examples);
}

function changeContent(content: any) {
  if (selectedCodeDetail?.value) {
    if (content?.type) {
      selectedCodeDetail.value.schemaItem.content = JSON.stringify(content);
      contentStr.value = JSON.stringify(content);
      selectedCodeDetail.value.schemaItem.type = content.type;
      store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail?.value);
    }
  }
}

function changeExamples(examples: any) {
  if (selectedCodeDetail?.value) {
    if (examples) {
      selectedCodeDetail.value.examples = JSON.stringify(examples);
      exampleStr.value = JSON.stringify(examples);
      store.commit('Endpoint/setSelectedCodeDetail', selectedCodeDetail?.value);
    }
  }
}

const refsOptions = ref([]);
onMounted(async () => {
  refsOptions.value = await store.dispatch('Endpoint/getAllRefs', {
    "serveId": currServe.value.id,
  });
})

</script>
<style lang="less" scoped>
.form-label {
  font-weight: bold;
}

.form-label-first {
  font-weight: bold;
  position: relative;
  left: -18px;
  margin-bottom: 16px;
}

.form-item-response-item {
  margin-top: 16px;
  align-items: center;
}

.form-item-response-item-con {
  position: relative;
  margin-bottom: 24px;

  &:before {
    content: "";
    position: absolute;
    left: -12px;
    top: -72px;
    width: 2px;
    background: #E5E5E5;
    min-height: 80vh;
    height: calc(100% + 50px);
  }
}

.label-name {
  margin-left: 4px;
}
</style>
