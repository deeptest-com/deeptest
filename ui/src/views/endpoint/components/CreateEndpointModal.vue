<template>
  <a-modal
      width="640px"
      :visible="visible"
      @ok="ok"
      @cancel="cancel"
      title="新建接口">
    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 17 }">
      <a-form-item label="接口名称" name="title">
        <a-input placeholder="请输入接口名称" v-model:value="formState.title"/>
      </a-form-item>
      <a-form-item label="所属分类" name="categoryId">
        <a-tree-select
            @change="selectedCategory"
            :value="formState.categoryId"
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
      <a-form-item label="描述" name="description">
        <a-textarea
            v-model:value="formState.description"
            placeholder="清输入接口描述信息"
            :auto-size="{ minRows: 2, maxRows: 5 }"
        />
      </a-form-item>
      <a-form-item label="curl导入" name="curl">
        <a-textarea
            v-model:value="formState.curl"
            placeholder="请输入cURL (bash) 命令"
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
  UnwrapRef,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';
import {useStore} from "vuex";
import {NewEndpointFormState} from "@/views/Endpoint/data";

const store = useStore<{ Endpoint }>();
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);

const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return  data?.[0]?.children || [];
});

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  selectedCategoryId: {
    required: true,
  }
})

const emit = defineEmits(['ok', 'cancel']);

const formRef = ref();

function ok() {
  formRef.value
      .validate()
      .then(() => {
        emit('ok', formState);
        formRef.value.resetFields();
      })
      .catch((error: ValidateErrorEntity<NewEndpointFormState>) => {
        console.log('error', error);
      });
}

function cancel() {
  emit('cancel', formState);
  formRef.value.resetFields();
}

function selectedCategory(value) {
  formState.categoryId = value;
}

const formState: UnwrapRef<NewEndpointFormState> = reactive({
  title: '',
  categoryId: null,
  description: '',
  curl:'',
});

watch(() => {
  return props.visible
}, (newVal) => {
  if (newVal) {
    formState.categoryId = props.selectedCategoryId || -1;
  }
}, {
  immediate: true
})

const rules = {
  title: [
    {required: true, message: '请输入接口名称', trigger: 'blur'},
    {min: 1, max: 50, message: '最少 1 个字符，最长 100 个字符', trigger: 'blur'},
  ],
  categoryId: [{required: false}],
  description: [{required: false}],
};


</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
