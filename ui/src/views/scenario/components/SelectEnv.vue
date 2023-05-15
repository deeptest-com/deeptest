<template>
  <div class="scenario-exec-main">
    <a-modal :title="'选择执行环境'"
             :visible="visible"
             @ok="ok"
             @cancel="cancal"
             class="scenario-edit"
             width="600px">
      <a-form :model="formState" :label-col="{span: 4}"  :wrapper-col="{span: 14}" ref="formRef">
        <a-form-item label="选择执行环境">
          <a-select v-model:value="formState.env"  :options="envOptions" placeholder="请选择环境"/>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import {
  defineComponent,
  computed,
  ref,
  reactive,
  ComputedRef,
  defineProps,
  PropType,
  defineEmits,
  onMounted, watch
} from "vue";
import {useStore} from "vuex";

import {StateType} from "../store";

const props = defineProps({
  visible: {
    type: Boolean,
    required: true
  },
  scenarioInfo:{
    required:true
  }
})
const emit = defineEmits(['cancel','ok']);

const rulesRef = reactive({
  env: [
    {required: true, message: '请选择环境', trigger: 'blur'},
  ],

});
const rules = {
  env: [
    {required: true, message: '请选择环境', trigger: 'change'},
  ],
};

const store = useStore<{ Scenario: StateType, ProjectGlobal,Environment }>();
const treeDataCategory = computed<any>(() => store.state.Scenario.treeDataCategory);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const environmentsData = computed<any>(() => store.state.Environment.environmentsData);
console.log('environmentsData',environmentsData.value);
const envOptions = [
  {
    label: '测试环境',
    value: 'test'
  },
  {
    label: '预发布环境',
    value: 'pre'
  },
  {
    label: '生产环境',
    value: 'prod'
  }
]
const formState = ref({
  env: null
});
watch(() => props.visible, async (val) => {
  if (val) {
    // 重新打开时，清楚表单数据
    formState.value = {
      env: null,
    };
  }
});


const formRef = ref();

function ok() {
  formRef.value
      .validate()
      .then(() => {
        console.log('formState', formState.value);
        emit('ok', formState.value, props.scenarioInfo);
      })
      .catch((error: any) => {
        console.log('error', error);
      });
}


function cancal() {
  emit('cancel')
}


</script>

<style lang="less" scoped>
.scenario-exec-main {

}
</style>
