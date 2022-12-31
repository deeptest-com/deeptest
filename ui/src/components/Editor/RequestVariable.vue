<template>
  <a-modal
      title="使用变量"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="requestVariableVisible"
      :footer="null"
      width="800px"
      height="600px"
  >
    <div>
      <a-row>
        <a-col flex="100px" class="dp-border">环境变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in environmentData.vars" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

      <br/>

      <a-row>
        <a-col flex="100px" class="dp-border">共享变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in validExtractorVariablesData" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue==='extractor_err'? t(item.rightValue+'_short') : item.value}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, defineEmits, onMounted, reactive, ref, Ref, computed, onUnmounted} from "vue";
import {message, Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getEnvironment, saveEnvironment} from "@/views/interface/service";
import {useStore} from "vuex";
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as EnvironmentStateType} from "@/store/environment";
import {Interface} from "@/views/interface/data";
import {StateType as ProjectStateType} from "@/store/project";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const { t } = useI18n();

const store = useStore<{ Interface: InterfaceStateType, ProjectGlobal: ProjectStateType, Environment: EnvironmentStateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
const environmentData = computed<any>(() => store.state.Environment.environmentData);
const validExtractorVariablesData = computed(() => store.state.Interface.validExtractorVariablesData);

const requestVariableVisible = ref(false)

onMounted(()=> {
  console.log('onMounted')
  bus.on(settings.eventVariableSelectionStatus, onVariableSelectionStatus);
})

onUnmounted(() => {
  console.log('onUnmounted')
  bus.off(settings.eventVariableSelectionStatus, onVariableSelectionStatus);
})

const src = ref('')
const onVariableSelectionStatus = (data) => {
  src.value = data.src
  requestVariableVisible.value = data.showVariableSelection
}

const select = async (item) => {
  console.log('select', item)
  bus.emit(settings.eventVariableSelectionResult, {src: src.value, item: item});
  requestVariableVisible.value = false
}

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>

<style lang="less">
.request-variable-main {

}
</style>