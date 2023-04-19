<!-- ::::请求体设置 -->
<template>
  <!-- 增加请求体 -->
  <a-row class="form-item-request-item">
    <a-col :span="3" class="form-label">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">增加请求体</span>
    </a-col>
    <a-col :span="18">
      <a-select
          v-if="selectedMethodDetail.requestBody"
          :value="selectedMethodDetail.requestBody.mediaType || null"
          @change="handleChangeMediaType"
          placeholder="请选择请求格式"
          style="width: 300px"
          :options="mediaTypesOpts.filter(item => !item.disabled)"
      ></a-select>
    </a-col>
  </a-row>
  <!-- 增加请求体 - 描述  -->
  <a-row class="form-item-request-item" v-if="collapse">
    <a-col :span="3" class="form-label"></a-col>
    <a-col :span="20">
      <a-input @change="handleChangeDesc"
               placeholder="描述信息"
               :value="selectedMethodDetail.requestBody.description"/>
    </a-col>
  </a-row>
  <!-- 增加请求体 - scheme定义 -->
  <a-row class="form-item-request-item" v-if="collapse">
    <a-col :span="3" class="form-label"></a-col>
    <a-col :span="21">
      <SchemaEditor
          @generateFromJSON="generateFromJSON"
          @generateExample="handleGenerateExample"
          @change="handleChange"
          :refsOptions="refsOptions"
          :contentStr="JSON.stringify(activeReqBodySchema?.content)"
          :exampleStr="JSON.stringify(activeReqBodySchema?.examples)"
          :tab-content-style="{width:'100%'}"/>
    </a-col>
  </a-row>
</template>
<script lang="ts" setup>
import {computed, defineEmits, defineProps, onMounted, ref, watch,} from 'vue';
import {useStore} from "vuex";
import {mediaTypesOpts,} from '@/config/constant';
import {Endpoint} from "@/views/endpoint/data";
import {DownOutlined, RightOutlined} from '@ant-design/icons-vue';
import SchemaEditor from '@/components/SchemaEditor/index.vue';

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User,ServeGlobal }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const selectedMethodDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const props = defineProps({});
const emit = defineEmits([]);
// 是否折叠,默认展开
const collapse = ref(true);
const activeReqBodySchema: any = ref({
  content: null,
  examples: [],
});
watch(() => {
  return selectedMethodDetail?.value?.requestBody?.schemaItem?.content
}, (newVal, oldValue) => {
  activeReqBodySchema.value.content = JSON.parse(newVal || 'null')
}, {immediate: true});

watch(() => {
  return selectedMethodDetail?.value?.requestBody?.examples
}, (newVal, oldValue) => {
  activeReqBodySchema.value.examples = JSON.parse(newVal || '[]')
}, {immediate: true});

async function generateFromJSON(JSONStr: string) {
  activeReqBodySchema.value.content = await store.dispatch('Endpoint/example2schema', {data: JSONStr});
}

async function handleGenerateExample(examples: any) {
  const res = await store.dispatch('Endpoint/schema2example', {
    data: JSON.stringify(activeReqBodySchema.value.content)
  });
  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res),
  };
  activeReqBodySchema.value.examples.push(example);
}

function handleChangeMediaType(mediaType: string) {
  selectedMethodDetail.value.requestBody.mediaType = mediaType;
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function handleChangeDesc(e: any) {
  selectedMethodDetail.value.requestBody.description = e.target.value;
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
}

function handleChange(json: any) {
  const {content, examples} = json;
  if (selectedMethodDetail?.value?.requestBody) {
    selectedMethodDetail.value.requestBody.schemaItem.content = JSON.stringify(content);
    selectedMethodDetail.value.requestBody.examples = JSON.stringify(examples);
    selectedMethodDetail.value.requestBody.schemaItem.type = content.type;
  }
  store.commit('Endpoint/setSelectedMethodDetail', {
    ...selectedMethodDetail.value
  })
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
  position: relative;
  left: -18px;
}

.form-label-first {
  font-weight: bold;
  position: relative;
  left: -18px;
}

.label-name {
  display: inline-block;
  margin-left: 4px;
  margin-top: 4px;
}

.form-item-request-item {
  margin-top: 16px;
}
</style>
