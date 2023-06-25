<template>
  <a-modal
      width="600px"
      :visible="visible"
      @ok="ok"
      @cancel="cancel"
      :title="mode === 'new' ? '新建分类' : '修改分类'">
    <a-form
        ref="tagFormRef"
        :rules="rules"
        :model="formState"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 14 }">
      <a-form-item label="分类名称" name="name">
        <a-input placeholder="请输入分类名称" v-model:value="formState.name"/>
      </a-form-item>
      <a-form-item label="备注信息" name="desc">
        <a-input placeholder="请输入备注信息" v-model:value="formState.desc"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref, defineProps, defineEmits, watch} from 'vue';

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  nodeInfo: {
    required: true,
  },
  mode: {
    required: true,
    type: String,
  }
})

const emit = defineEmits(['ok', 'cancel']);

const tagFormRef = ref();
const formState = ref({
  name: '',
  desc: '',
});


watch(() => {return props.visible}, (newVal) => {
  if(newVal){
    formState.value.name = props?.nodeInfo?.name || '';
    formState.value.desc = props?.nodeInfo?.desc || '';
    if (props.mode === 'new') {
      formState.value.name = '';
      formState.value.desc = '';
    }
  }
})

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
  tagFormRef.value.resetFields();
  emit('cancel');
}

const rules = {
  name: [
    {required: true, message: '请输入分类名称', trigger: 'blur'},
    {min: 1, max: 50, message: '最少 1 个字符，最长 50 个字符', trigger: 'blur'},
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
