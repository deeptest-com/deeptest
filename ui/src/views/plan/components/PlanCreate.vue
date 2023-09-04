<template>
  <a-modal
    title="新建测试计划"
    :visible="createDrawerVisible"
    class="scenario-edit"
    :closable="true"
    @cancel="onCancel"
    @ok="save"
    width="600px"
  >
    <div class="plan-create-main">
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol" :ref="formRef">
        <a-form-item label="计划名称" v-bind="validateInfos.name">
          <a-input
            v-model:value="modelRef.name"
            placeholder="请输入计划名称"
            @blur="validate('name', { trigger: 'blur' }).catch(() => {})"
          />
        </a-form-item>
        <a-form-item
          label="负责人"
          has-feedback
          :rules="[{ required: true, message: '请选择负责人' }]"
        >
          <a-select
            v-model:value="modelRef.adminId"
            placeholder="请选择(默认当前用户)"
            :options="members"
          />
        </a-form-item>
        <a-form-item label="所属分类">
          <a-tree-select
            @change="selectedCategory"
            :value="modelRef.categoryId"
            show-search
            :multiple="false"
            :treeData="treeData"
            style="width: 100%"
            :treeDefaultExpandAll="true"
            :replaceFields="{ title: 'name', value: 'id' }"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            placeholder="请选择所属分类"
            allow-clear
          />
        </a-form-item>
        <a-form-item label="测试阶段">
          <a-select
            v-model:value="modelRef.testStage"
            placeholder="请选择"
            :options="testStageArr"
          />
        </a-form-item>
        <a-form-item label="描述">
          <a-textarea
            v-model:value="modelRef.desc"
            placeholder="请输入描述"
            :auto-size="{ minRows: 2, maxRows: 5 }"
          />
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, reactive, computed, watch } from "vue";
import { useStore } from "vuex";
import { Form } from "ant-design-vue";
import { StateType as PlanStateType } from "../store";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as UserStateType } from "@/store/user";
import {
  getExpandedKeys,
  setExpandedKeys,
  setSelectedKey,
  getSelectedKey,
} from "@/utils/cache";

const props = defineProps<{
  createDrawerVisible: Boolean;
  selectedCategoryId?: String;
}>();

const useForm = Form.useForm;
const emits = defineEmits(["onCancel", "getList"]);
const store = useStore<{
  Plan: PlanStateType;
  ProjectGlobal: ProjectStateType;
  User: UserStateType;
}>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const members = computed<any>(() => store.state.Plan.members);
const currentUser = computed<any>(() => store.state.User.currentUser);
const treeDataCategory = computed<any>(() => store.state.Plan.treeDataCategory);
const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return data?.[0]?.children || [];
});
const labelCol = { span: 5 };
const wrapperCol = { span: 17 };
const modelRef = reactive({
  name: "",
  categoryId: null,
  adminId: (currentUser.value && currentUser.value.id) || null,
  testStage: null,
  desc: "",
} as any);
const formRef = ref();
const testStageArr = [
  {
    label: "单元测试",
    value: "unit_test",
  },
  {
    label: "集成测试",
    value: "integration_test",
  },
  {
    label: "系统测试",
    value: "system_test",
  },
  {
    label: "验收测试",
    value: "acceptance_test",
  },
];
const rulesRef = reactive({
  name: [{ required: true, message: "请输入名称", trigger: "blur" }],
  adminId: [{ required: true, message: "请选择负责人" }],
});

const { validate, validateInfos } = useForm(modelRef, rulesRef);

function onCancel() {
  emits("onCancel");
}

function selectedCategory(value) {
  modelRef.categoryId = value;
}

async function save() {
  validate().then(() => {
    store.dispatch("Plan/savePlan", modelRef).then((res) => {
      if (res === true) {
        emits("onCancel");
        emits("getList");
      }
    });
  });
}

function onFinish() {
  console.log("完成创建");
}

watch(
  () => {
    return currProject.value;
  },
  (val) => {
    if (val.id) {
      store.dispatch("Plan/loadMembers", {
        id: val.id,
        keywords: "",
        pageSize: 100,
        page: 1,
      });
    }
  },
  {
    immediate: true,
  }
);

// 新建计划时，如果有选中的分类，则默认选中分类
watch(
  () => {
    return props.createDrawerVisible;
  },
  (newVal) => {
    if (newVal) {
      getSelectedKey("category-plan", currProject.value.id).then(
        async (keys) => {
          if (keys) {
            modelRef.categoryId = keys;
          } else {
            modelRef.categoryId = -1;
          }
        }
      );
      Object.assign(modelRef, {
        name: "",
        testStage: null,
        desc: "",
      });
    }
  }
);
</script>
