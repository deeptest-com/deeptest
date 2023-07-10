<template>
  <a-modal
      width="640px"
      :visible="visible"
      @ok="ok"
      @cancel="cancel"
      title="接口批量修改">
    <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 5 }"
        :wrapper-col="{ span: 17 }">
      <a-form-item label="修改字段" name="fieldName" style="width: 100%">
        <a-select v-model:value="formState.fieldName"
                  :options="FieldNameOpts"
                  @change="selectedFieldName"
                  placeholder="请选择"/>
      </a-form-item>
      <a-form-item label="修改为" name="value" style="width: 100%">
        <a-select v-if="formState.fieldName === 'status'"
                  v-model:value="formState.value"
                  :options="endpointStatusOpts"
                  @change="selectedValue" />
        <a-tree-select v-else-if="formState.fieldName === 'categoryId'"
            @change="selectedValue"
                       v-model:value="formState.value"
            show-search
            :multiple="false"
            :treeData="treeData"
            :treeDefaultExpandAll="true"
            :replaceFields="{ title: 'name',value:'id'}"
            :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
            placeholder="请选择所属分类"
            allow-clear/>
        <a-select v-else
                  :value="formState.value"/>
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
import {endpointStatusOpts} from '@/config/constant';
import {useStore} from "vuex";
import {NewEndpointFormState} from "@/views/Endpoint/data";

const store = useStore<{ Endpoint }>();
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);

const treeData: any = computed(() => {
  const data = treeDataCategory.value;
  return  data?.[0]?.children || [];
});

const FieldNameOpts = [

  {
    label: '状态',
    value: 'status',
  },
  {
    label: '所属分类',
    value: 'categoryId',
  },
]

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
})

const emit = defineEmits(['ok', 'cancel']);

const formRef = ref();

const formState = ref({
  fieldName: '',
  value: null as any,
  endpointIds: [],
});

function ok() {
  formRef.value
      .validate()
      .then(() => {
        emit('ok', formState);
        formRef.value.resetFields();
      })
      .catch((error: ValidateErrorEntity) => {
        console.log('error', error);
      });
}

function cancel() {
  emit('cancel', formState);
  formRef.value.resetFields();
}

function selectedValue(value) {
  formState.value.value = value;
}

function selectedFieldName() {
  formState.value.value = null;
}

const rules = {
  fieldName: [
    {required: true, message: '请选择字段', trigger: 'change'},
  ],
};

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
