<template>
  <div class="editor show-on-hover" v-if="isEditing" v-on-click-outside="cancelEdit">
    <a-input
             class="input"
             :placeholder="placeholder || '请输入内容'"
             :size="'small'"
             v-model:value="fieldValue"/>
    <a-space :size="8">
      <CloseOutlined @click.stop="cancelEdit"/>
      <CheckOutlined
          @click.stop="updateField"
          class="editor-icon"
          :class="{disabled: !fieldValue}"/>&nbsp;
    </a-space>
  </div>
  <div :class="['editor','show-on-hover', customClass]" v-else>
    <span class="title" :title="fieldValue" @click.stop="handleClick">{{ fieldValue || '暂无' }}</span> &nbsp;&nbsp;
    <span class="edit-icon"><EditOutlined @click.stop="isEditing = true"/></span>
  </div>
</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits,
  ref, watch,
} from 'vue';
import { message } from 'ant-design-vue';
import {
  EditOutlined,
  CheckOutlined,
  CloseOutlined
} from '@ant-design/icons-vue';
import { vOnClickOutside } from '@vueuse/components';
const isEditing = ref(false);
const fieldValue = ref('');
const editor = ref(null);
const props = defineProps({
  value: {
    required: true,
    type: String,
  },
  placeholder: {
    required: true,
    type: String,
  },
  customClass: {
    required: false,
    type: String,
  }
})
const emit = defineEmits(['update', 'edit']);

function updateField() {
  if (!fieldValue.value) {
    message.warning('请请输入内容');
    return;
  }
  emit('update', fieldValue.value);
  isEditing.value = false;
}

function cancelEdit() {
  fieldValue.value = props.value;
  isEditing.value = false;
}

function handleClick() {
  emit('edit');
}

watch(() => {
  return props.value
}, (newVal) => {
  fieldValue.value = newVal
}, {
  immediate: true
})




</script>

<style lang="less" scoped>
.editor {
  display: flex;
  align-items: center;
  overflow: hidden;
  flex: 1;

  &.custom-serve {
    color: #447DFD;
  }

  &.custom-endpoint {
    color: #447DFD;
  }

  &.show-on-hover {
    .edit-icon {
      display: none;
    }
    &:hover {
      .edit-icon {
        display: inline-block;
        color: #8A8A8A;
      }
    }
  }
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

  .title {
    cursor: pointer;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
}


</style>
