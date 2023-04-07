<template>
  <!-- 响应定义  -->
  <a-row class="form-item-response">
    <a-col :span="3" class="form-label">
      <RightOutlined v-if="!collapse" @click="collapse = !collapse"/>
      <DownOutlined v-if="collapse" @click="collapse = !collapse"/>
      <span class="label-name">选择响应代码</span>
    </a-col>
    <a-col :span="21">
      <a-radio-group v-model:value="selectedCode" button-style="solid">
        <a-radio-button
            :class="{'has-defined': hasDefinedCode(code.value)}"
            :key="code.value" v-for="code in repCodeOpts"
            :value="code.value">
          {{ code.label }}
        </a-radio-button>
      </a-radio-group>
      <div class="form-item-response" v-if="collapse">
        <div v-if="selectedCodeDetail">
          <!-- Description -->
          <a-row class="form-item-response-item">
            <a-col :span="4" class="form-label">
              Description
            </a-col>
            <a-col :span="18">
              <a-input v-model:value="selectedCodeDetail.desc"/>
            </a-col>
          </a-row>
          <!-- 增加响应头 -->
          <a-row class="form-item-response-item">
            <a-col :span="4" class="form-label">
              增加响应头
            </a-col>
            <a-col :span="18">
              <div class="params-defined-btns">
                <a-button type="primary" @click="addResponseHeader">
                  <template #icon>
                    <PlusOutlined/>
                  </template>
                  {{ `添加` }}
                </a-button>
              </div>
            </a-col>
          </a-row>
          <!-- 响应头展示-->
          <a-row class="form-item-response-item">
            <a-col :span="4"></a-col>
            <a-col :span="20">
              <div class="params-defined">
                <div class="params-defined-content">
                  <div class="params-defined-item" v-if="selectedCodeDetail?.headers?.length">
                    <div class="header-defined header-defined-items">
                      <div v-for="(item,index) in selectedCodeDetail.headers" :key="item.id">
                        <Field
                            :fieldData="{...item,index:index}"
                            :showRequire="false"
                            @del="deleteResHeader(index)"
                            @change="(val) => {
                                        handleResHeaderChange(val);
                                      }"/>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </a-col>
          </a-row>
          <!-- 增加响应体体 -->
          <a-row class="form-item-response-item">
            <a-col :span="4" class="form-label">
              增加响应体
            </a-col>
            <a-col :span="18">
              <a-select
                  placeholder="请选择格式"
                  v-model:value="selectedCodeDetail.mediaType"
                  style="width: 300px"
                  :options="mediaTypesOpts"
              ></a-select>
            </a-col>
          </a-row>
          <!-- 增加响应体 - 描述  -->
          <a-row class="form-item-response-item">
            <a-col :span="4" class="form-label"></a-col>
            <a-col :span="18">
              <a-input placeholder="请输入描述" v-model:value="selectedCodeDetail.description"/>
            </a-col>
          </a-row>
          <!-- 增加响应体 - scheme定义 -->
          <a-row class="form-item-response-item">
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
        </div>
        <div v-if="!selectedCodeDetail">
          <a-button type="primary" @click="addCodeResponse">
            <template #icon>
              <PlusOutlined/>
            </template>
            {{ `Add Response` }}
          </a-button>
        </div>
      </div>
    </a-col>
  </a-row>
</template>
<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  watch,
  computed,
} from 'vue';
import {useStore} from "vuex";
import {
  mediaTypesOpts,
  repCodeOpts,
  defaultHeaderParams,
  defaultCodeResponse,
} from '@/config/constant';
import {PlusOutlined, DownOutlined, RightOutlined} from '@ant-design/icons-vue';
import Field from './Field.vue'
import {Endpoint} from "@/views/endpoint/data";
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {cloneByJSON} from "@/utils/object";

const store = useStore<{ Endpoint, Debug, ProjectGlobal, User }>();
const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const selectedMethodDetail = computed<any>(() => store.state.Endpoint.selectedMethodDetail);
const currentUser: any = computed<Endpoint>(() => store.state.User.currentUser);
// 是否折叠,默认展开
const collapse = ref(true);
const props = defineProps({});
const emit = defineEmits([]);
const selectedCode = ref('200');
// 当前选中的请求方法的响应体详情
const selectedCodeDetail: any = ref(null);

// 是否定义了请求方法的响应体
function hasDefinedCode(code: string) {
  return selectedMethodDetail?.value?.responseBodies?.some((item) => {
    return item.code === code;
  })
}

watch(() => {
  return selectedCode.value
}, (newVal, oldVal) => {
  selectedCodeDetail.value = selectedMethodDetail?.value?.responseBodies?.find((item) => {
    return item.code === newVal;
  })
}, {immediate: true});

const selectedMethodIndex: any = computed(() => {
  return endpointDetail?.value?.interfaces?.findIndex((item) => {
    return item.method === selectedMethodDetail.value.method;
  })
});



const selectedCodeIndex: any = computed(() => {
  return selectedMethodDetail?.value?.responseBodies?.findIndex((item) => {
    return item.code === selectedCode?.value;
  })
});

function addResponseHeader() {
  selectedCodeDetail.value.headers.push(cloneByJSON(defaultHeaderParams));
}

function addCodeResponse() {
  const item = {
    ...cloneByJSON(defaultCodeResponse),
    "code": selectedCode.value,
    "endpointId": selectedMethodDetail.value.id,
  }
  store.commit('Endpoint/setEndpointDetailByIndex', {
    methodIndex: selectedMethodIndex.value,
    codeIndex: selectedCodeIndex.value,
    value: item
  })
  selectedCodeDetail.value = item;
}

function deleteResHeader(index) {
  selectedCodeDetail.value.headers.splice(index, 1);
}

function handleResHeaderChange(data) {
  selectedCodeDetail.value.headers[data.index] = {...data};
}

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


async function generateFromJSON(type: string, JSONStr: string) {
  const res = await store.dispatch('Endpoint/example2schema',
      {data: JSONStr}
  );
  activeResBodySchema.value.content = res;
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

function handleChange( json: any) {
  const {content, examples} = json;
  // activeResBodySchema.value.content = content;
  // activeResBodySchema.value.type = content.type;
  // activeResBodySchema.value.examples = examples;
  if (selectedCodeDetail?.value) {
    selectedCodeDetail.value.schemaItem.content = JSON.stringify(content);
    selectedCodeDetail.value.examples = JSON.stringify(examples);
    selectedCodeDetail.value.schemaItem.type = content.type;
    // store.commit('Endpoint/setInterfaceDetailByIndex', {
    //   methodIndex: selectedMethodIndex.value,
    //   codeIndex: selectedCodeIndex.value,
    //   value: selectedCodeDetail.value
    // })
  }
}

</script>
<style lang="less" scoped>
.content {
  padding-top: 16px;
}

.form-item {
  margin-bottom: 16px;
  align-items: baseline;
}

.path-param-list {
  margin-top: 16px;
}

.form-label {
  font-weight: bold;
}

.form-item-request {
  margin-top: 16px;

  .form-item-request-item {
    margin-top: 16px;
    align-items: center;
  }

  .form-item-response {
    margin-top: 16px;

    .form-item-response-item {
      margin-top: 16px;
      align-items: center;
    }
  }
}

.params-defined-item-header {
  font-weight: bold;
  margin-bottom: 8px;
  margin-top: 8px;
}

.has-defined {
  color: #1890ff;
  //font-weight: bold;
}

.label-name {
  margin-left: 4px;
}
</style>
