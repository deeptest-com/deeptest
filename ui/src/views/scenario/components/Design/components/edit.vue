<template>
  <a-modal
      width="600px"
      :visible="!!nodeInfo"
      @ok="ok"
      @cancel="cancel"
      :title="(!nodeInfo.id ? '新建' : '修改') + (formState.entityType === 'processor_interface_default' ? '接口' : '目录')">
    <a-form
        ref="tagFormRef"
        :rules="rules"
        :model="formState"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }">

      <a-form-item :label="(formState.entityType === 'processor_interface_default' ? '接口' : '目录') + '名称'" name="name">
        <a-input placeholder="请输入名称" v-model:value="formState.name" @change="removeSpace"/>
      </a-form-item>

      <a-form-item :label="(formState.entityType === 'processor_interface_default' ? '接口' : '目录') + '备注'" name="comments">
        <a-input placeholder="请输入备注" v-model:value="formState.comments"/>
      </a-form-item>

    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, defineProps, defineEmits, watch} from 'vue';

const props = defineProps({
  nodeInfo: {
    required: true,
    type: Object,
  }
})

const removeSpace = () => {
  console.log(formState.value.name)
  formState.value.name = formState.value.name.trim()
}

const emit = defineEmits(['ok', 'cancel']);

const tagFormRef = ref();
const formState = ref({
  id: 0,
  name: '',
  comments: '',
  entityType: '',
  entityCategory: '',
  parentId: 0,
} as any);

watch(props.nodeInfo, () => {
  console.log('watch props.nodeInfo', props?.nodeInfo?.type)
  formState.value = {
    id: props?.nodeInfo?.id,
    name: props?.nodeInfo?.name,
    comments: props?.nodeInfo?.comments,
    entityCategory: props?.nodeInfo?.entityCategory,
    entityType: props?.nodeInfo?.entityType,
    processorInterfaceSrc: props?.nodeInfo?.processorInterfaceSrc,

    targetProcessorCategory: props?.nodeInfo?.targetProcessorCategory,
    targetProcessorType: props?.nodeInfo?.targetProcessorType,
    targetProcessorId: props?.nodeInfo?.targetProcessorId,
    mode: props?.nodeInfo?.mode,
  }
}, { immediate: true, deep: true })

function ok() {
  console.log('edit ok', formState.value.name)
  tagFormRef.value
      .validate()
      .then(() => {
        emit('ok', {
          ...formState.value
        });
        tagFormRef.value.resetFields();
      })
      .catch((error) => {
        console.log('error', error);
      });
}

function cancel() {
  emit('cancel');
  tagFormRef.value.resetFields();
}

const rules = {
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
    {required: true, max: 30, message: '最少1个字符，最长30个字符', trigger: 'blur'},
  ],
  comments: [{required: false}],
};

</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
