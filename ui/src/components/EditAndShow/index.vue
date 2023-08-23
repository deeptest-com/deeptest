<template>
  <div class="editor show-on-hover"
       v-if="isEditing"
       v-on-click-outside="cancelEdit">
    <a-input  class="input"
             :placeholder="placeholder || '请输入内容'"
             :size="'small'"
              ref="inputRef"
             @keydown.enter="updateField"
             v-model:value="fieldValue" />

    <a-space :size="8">
      <CloseOutlined @click.stop="cancelEdit"/>
      <CheckOutlined
          @click.stop="updateField"
          class="editor-icon"
          :class="{disabled: !fieldValue}"/>&nbsp;
    </a-space>
  </div>

  <div :class="['editor','show-on-hover', customClass]" v-else>
    <span class="title" :title="fieldValue" @click.stop="handleClick">
      {{ fieldValue || emptyValue }}
    </span> &nbsp;&nbsp;

    <span class="edit-icon">
      <EditOutlined @click.stop="edit"/>
    </span>
  </div>

</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits,
  nextTick,
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
  },
  autoFocus: {
    required: false,
    type: Boolean,
    default: false,
  },
  canEmpty: {
    required: false,
    type: Boolean,
    default: false,
  },
  emptyValue: {
    required: true,
    type: String,
    default: '暂无',
  },
})
const emit = defineEmits(['update', 'edit']);

function updateField() {
  if (!props.canEmpty && !fieldValue.value) {
    message.warning('请请输入内容');
    return;
  }
  emit('update', fieldValue.value);
  isEditing.value = false;
}

function edit() {
  isEditing.value = true;
  nextTick(() => {
    inputRef?.value?.focus();
  })
}
function cancelEdit() {
  fieldValue.value = props.value;
  isEditing.value = false;
  if (props.canEmpty) {
    emit('update', fieldValue.value);
  }
}

function handleClick() {
  emit('edit');
}
const inputRef:any = ref(null);

watch(() => {return props.value}, (newVal) => {
  fieldValue.value = newVal
}, {immediate: true})

watch(() => {return props.autoFocus}, (newVal) => {
  if (newVal) {
    isEditing.value = true;

    nextTick(() => {
      inputRef?.value?.focus();
    })

  }
}, {immediate: true})

</script>

<style lang="less" scoped>
.editor {
  display: flex;
  align-items: center;
  overflow: hidden;
  //flex: 1;
  height: 24px;

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
