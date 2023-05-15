<template>
  <div class="scenario-edit-main">
    <a-modal :title="'新建测试场景'"
             :visible="visible"
             @ok="submitForm"
             @cancel="cancal"
             class="scenario-edit"
             width="600px">
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="名称" v-bind="validateInfos.name">
          <a-input v-model:value="modelRef.name"
                   placeholder="请输入场景名称"
                   @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>
        <a-form-item label="优先级">
          <a-select v-model:value="modelRef.priority" :options="priorityOptions" placeholder="请选择"/>
        </a-form-item>
        <a-form-item label="所属分类" v-bind="validateInfos.categoryId">
          <a-tree-select
              v-model:value="modelRef.categoryId"
              show-search
              :multiple="false"
              :treeData="treeData"
              style="width: 100%"
              :treeDefaultExpandAll="true"
              :replaceFields="{ title: 'name',value:'id'}"
              :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
              placeholder="请选择所属分类"
              allow-clear/>
        </a-form-item>
        <a-form-item label="测试类型" v-bind="validateInfos.type">
          <a-select v-model:value="modelRef.type" placeholder="请选择" :options="testTypeOptions"/>
        </a-form-item>
        <a-form-item label="描述" v-bind="validateInfos.desc">
          <a-textarea v-model:value="modelRef.desc"
                      @blur="validate('desc', { trigger: 'blur' }).catch(() => {})"/>
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
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form} from 'ant-design-vue';

const useForm = Form.useForm;
import {StateType} from "../store";
import {priorityOptions, testTypeOptions} from "@/config/constant"
import {getSelectedKey} from "@/utils/cache";

const router = useRouter();
const {t} = useI18n();

const props = defineProps({
  visible: {
    type: Boolean,
    required: true
  },
  onFinish: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: StateType, ProjectGlobal }>();

const treeDataCategory = computed<any>(() => store.state.Scenario.treeDataCategory);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return data?.[0]?.children || [];
});


const modelRef = ref({} as any)

watch(() => props.visible, async (val) => {
  if (val) {
    // 重新打开时，清楚表单数据
    modelRef.value = {
      name: '',
      priority: null,
      // 从缓存中 获取当前 默认选中的分类
      categoryId: await getSelectedKey('category-scenario', currProject.value.id),
      type: null,
      desc: null,
      projectId: currProject.value.id
    };
  }
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = async () => {
  validate().then(() => {
    // modelRef.value.categoryId = props.categoryId
    store.dispatch('Scenario/saveScenario', modelRef.value).then((res) => {
      console.log('res', res)
      if (res === true) {
        props.onFinish()
      }
    })
  })
      .catch(err => {
        console.log('error', err);
      });
};

const emit = defineEmits(['cancel']);

function cancal() {
  emit('cancel')
}

const labelCol = {span: 4}
const wrapperCol = {span: 18}

</script>

<style lang="less" scoped>
.scenario-edit-main {

}
</style>
