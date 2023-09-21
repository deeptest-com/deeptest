<!-- :::: 接口定义模块 -->
<template>
  <div class="content">
    <div class="mode-btns">
      <a-button  :type="showMode === 'form' ? 'default' : 'default'" @click="switchMode('form')">
        <template #icon>
          <BarsOutlined/>
        </template>图形
      </a-button>
      <a-button   :type="showMode === 'code' ? 'default' : 'default'" @click="switchMode('code')">
        <template #icon>
          <CodeOutlined/>
        </template>YAML
      </a-button>
    </div>

    <EndpointForm v-if="showMode === 'form'"/>

    <div class="endpoint-code" v-if="showMode === 'code'">
      <MonacoEditor
          class="editor"
          :value="endpointDetailYamlCode"
          :language="'yaml'"
          :height="600"
          theme="vs"
          :options="{...MonacoOptions}"
          @change="handleYamlCodeChange"
          :timestamp="timestamp" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';
import {useStore} from "vuex";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {CodeOutlined, BarsOutlined} from '@ant-design/icons-vue';
import {Endpoint} from "@/views/endpoint/data";
import {MonacoOptions} from '@/utils/const';
const store = useStore<{ Endpoint, ProjectGlobal }>();
const endpointDetail = computed<Endpoint[]>(() => store.state.Endpoint.endpointDetail);
const endpointDetailYamlCode = computed<any>(() => store.state.Endpoint.endpointDetailYamlCode);
import EndpointForm from './Form/index.vue'

const timestamp = ref('')
watch(endpointDetailYamlCode, (newVal) => {
  timestamp.value = Date.now() + ''
}, {immediate: true, deep: true})

const props = defineProps({});
const emit = defineEmits(['switchMode']);
const showMode = ref('form');

async function switchMode(val) {
  showMode.value = val;
  // 需求去请求YAML格式
  if (val === 'code') {
    await store.dispatch('Endpoint/getYamlCode', endpointDetail.value);
  }
  emit('switchMode', val);

}

function handleYamlCodeChange(code) {
  console.log(code);
  // store.commit("Endpoint/setYamlCode", code);
}

</script>

<style lang="less" scoped>
.content {
  //padding: 16px;
  //height: 100%;
  //min-height: calc(100vh - 200px);
  position: relative;
  //margin-top: 24px;
  //padding-top: 16px;
  .mode-btns {
    position: absolute;
    right: 0;
    top:16px;
    z-index: 99;
  }
}
</style>
