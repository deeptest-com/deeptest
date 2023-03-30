<template>
  <div class="editor" v-if="isEditing">
    <a-input class="input" placeholder="请输入接口名称" v-model:value="fieldValue"/>
    <a-space :size="8">
      <CloseOutlined @click.stop="cancelEdit"/>
      <CheckOutlined
          @click.stop="updateField"
          :class="{disabled: !fieldValue}"/>&nbsp;
    </a-space>
  </div>
  <div class="editor" v-else>
    <span>{{ fieldValue }}</span> &nbsp;
    <EditOutlined @click.stop="isEditing = true"/>
  </div>
</template>
<script lang="ts" setup>
import {
  defineProps,
  defineEmits,
  ref, watch,
} from 'vue';
import {
  EditOutlined,
  CheckOutlined,
  CloseOutlined
} from '@ant-design/icons-vue';
const isEditing = ref(false);
const fieldValue = ref('');
const props = defineProps({
  value: {
    required: true,
    type: String,
  }
})
const emit = defineEmits(['update']);

function updateField() {
  if (!fieldValue.value) {
    return;
  }
  emit('update', fieldValue.value);
  isEditing.value = false;
}

function cancelEdit() {
  fieldValue.value = props.value;
  isEditing.value = false;
}

watch(() => {
  return props.value
}, (newVal) => {
  console.log(newVal);
  fieldValue.value = newVal
}, {
  immediate: true
})

</script>

<style lang="less" scoped>
.editor {
  display: flex;
  align-items: center;
  .input {
    margin-right: 8px;
  }

  .btns {
    flex: 1;
    line-height: 30px;
    .disabled {
      color: #00000040;
    }
  }
}

</style>
