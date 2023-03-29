<!-- :::: 接口定义模块 -->
<template>
  <div class="content">
    <div class="mode-btns">
      <a-button :type="showMode === 'form' ? 'primary' : 'default'" @click="switchMode('form')">
        <template #icon>
          <BarsOutlined/>
        </template>
        图形
      </a-button>
      <a-button :type="showMode === 'code' ? 'primary' : 'default'" @click="switchMode('code')">
        <template #icon>
          <CodeOutlined/>
        </template>
        YAML
      </a-button>
    </div>
    <InterfaceForm v-if="showMode === 'form'"/>
    <div class="interface-code" v-if="showMode === 'code'">
      <MonacoEditor
          class="editor"
          :value="interfaceDetailYamlCode"
          :language="'yaml'"
          :height="600"
          theme="vs"
          :options="{...MonacoOptions}"
          @change="handleYamlCodeChange"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed,
} from 'vue';
import {useStore} from "vuex";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {CodeOutlined, BarsOutlined} from '@ant-design/icons-vue';
import {Interface} from "@/views/interface/data";
import {MonacoOptions} from '@/utils/const';

const store = useStore<{ Interface, ProjectGlobal }>();
const interfaceDetail = computed<Interface[]>(() => store.state.Interface.interfaceDetail);
const interfaceDetailYamlCode = computed<Interface[]>(() => store.state.Interface.interfaceDetailYamlCode);
import InterfaceForm from './InterfaceForm.vue'

const props = defineProps({});
const emit = defineEmits(['ok', 'close', 'refreshList']);
const showMode = ref('form');

async function switchMode(val) {
  showMode.value = val;
  // 需求去请求YAML格式
  if (val === 'code') {
    await store.dispatch('Interface/getYamlCode', interfaceDetail.value);
  }
}

function handleYamlCodeChange(code) {
  store.commit("Interface/setYamlCode", code);
}

</script>

<style lang="less" scoped>
.content {
  position: relative;
  min-height: 50vh;

  .mode-btns {
    position: absolute;
    right: 16px;
    z-index: 999;
  }
}
</style>
