<template>
  <a-modal
      width="600px"
      :visible="visible"
      @ok="ok"
      @cancel="cancal"
      title="新建接口">
    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }">
      <a-form-item label="接口名称" name="title">
        <a-input placeholder="请输入接口名称" v-model:value="formState.title"/>
      </a-form-item>
      <a-form-item label="所属分类" name="parentId">
        <a-tree-select
            v-model:value="formState.parentId"
            show-search
            :treeData="interFaceCategoryOpt"
            style="width: 100%"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            placeholder="请选择所属分类"
            allow-clear
            tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="描述" name="description">
        <a-textarea
            v-model:value="formState.path"
            placeholder="清输入接口描述信息"
            :auto-size="{ minRows: 2, maxRows: 5 }"
        />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <span class="">
           注：接口请求方法可以通过详情页添加
        </span>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts" setup>
import {ValidateErrorEntity} from 'ant-design-vue/es/form/interface';
import {
  reactive,
  ref,
  toRaw,
  UnwrapRef,
  defineProps,
  defineEmits,
  computed
} from 'vue';
import {useStore} from "vuex";
import {NewInterfaceFormState} from "@/views/interfaceV2/data";

const store = useStore<{ InterfaceV2, ProjectGlobal, Project }>();
let interFaceCategoryOpt = computed<any>(() => store.state.InterfaceV2.interFaceCategoryOpt);

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  }
})

const emit = defineEmits(['ok', 'cancal']);


const formRef = ref();

function ok() {
  formRef.value
      .validate()
      .then(() => {
        emit('ok', formState);
        formRef.value.resetFields();
      })
      .catch((error: ValidateErrorEntity<NewInterfaceFormState>) => {
        console.log('error', error);
      });
}

function cancal() {
  emit('cancal', formState);
  formRef.value.resetFields();
}

const formState: UnwrapRef<NewInterfaceFormState> = reactive({
  title: '',
  parentId: null,
  description: '',
});

const rules = {
  title: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 1, max: 50, message: '最少 1 个字符，最长 100 个字符', trigger: 'blur'},
  ],
  parentId: [{required: false, message: '请选择', trigger: 'change'}],
  description: [{required: false, message: '请输入描述', trigger: 'change'}],
};


</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
