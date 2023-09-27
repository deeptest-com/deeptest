<template>
  <a-modal
      width="600px"
      :visible="!!nodeInfo"
      @ok="ok"
      @cancel="cancel"
      :title="(!nodeInfo.id ? '新建' : '修改') + (formState.type === 'interface' ? '接口' : '目录')">
    <a-form
        ref="tagFormRef"
        :rules="rules"
        :model="formState"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }">

      <a-form-item :label="(formState.type === 'interface' ? '接口' : '目录') + '名称'" name="name">
        <a-input placeholder="请输入名称" v-model:value="formState.title"/>
      </a-form-item>

      <!-- <a-form-item :label="(formState.type === 'interface' ? '接口' : '目录') + '备注'" name="desc">
        <a-input placeholder="请输入备注" v-model:value="formState.desc"/>
      </a-form-item> -->

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

const emit = defineEmits(['ok', 'cancel']);

const tagFormRef = ref();
const formState = ref({
  id: 0,
  title: '',
  desc: '',
  type: '',
  parentId: 0,
});

watch(props.nodeInfo, () => {
  console.log('watch props.nodeInfo', props?.nodeInfo?.type)
  formState.value = {
    id: props?.nodeInfo?.id,
    title: props?.nodeInfo?.title,
    desc: props?.nodeInfo?.desc,
    type: props?.nodeInfo?.type,
    parentId: props?.nodeInfo?.parentId,
  }
}, { immediate: true, deep: true })

function ok() {
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
  title: [
    {required: true, message: '请输入名称', trigger: 'blur'},
    {min: 1, max: 50, message: '最少1个字符，最长50个字符', trigger: 'blur'},
  ],
  desc: [{required: false}],
};


</script>

<style lang="less" scoped>
.modal-btns {
  display: flex;
  justify-content: flex-end;
}
</style>
