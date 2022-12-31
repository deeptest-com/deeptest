<template>
  <a-modal
      title="使用变量"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="true"
      :onCancel="onCancel"
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
          <span @click="editVar(item)" class="dp-link-primary">选择</span>
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
          <span @click="editVar(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, defineEmits, onMounted, reactive, ref, Ref, computed} from "vue";
import {message, Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getEnvironment, saveEnvironment} from "@/views/interface/service";
import {useStore} from "vuex";
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as EnvironmentStateType} from "@/store/environment";
import {Interface} from "@/views/interface/data";
import {StateType as ProjectStateType} from "@/store/project";


const props = defineProps({
  interfaceId:{
    type: Number,
    required: true
  },

  onCancel:{
    type: Function,
    required: true
  },
  onFinish:{
    type: Function,
    required: true
  },
});

const { t } = useI18n();

const store = useStore<{ Interface: InterfaceStateType, ProjectGlobal: ProjectStateType, Environment: EnvironmentStateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
const environmentData = computed<any>(() => store.state.Environment.environmentData);
const validExtractorVariablesData = computed(() => store.state.Interface.validExtractorVariablesData);

const variRef = ref<any>({})

const onSubmit = async () => {
  console.log('onSubmit', variRef.value)
  props.onFinish(variRef.value);
}

onMounted(()=> {
  console.log('onMounted')
})

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>

<style lang="less">
.request-variable-main {

}
</style>