<template>
  <div class="editor" v-if="isEditing" >
    <a-select
        v-model:value="fieldValue"
        style="width: 100px;margin-right: 8px;"
        :size="'small'"
        @change="updateField"
        v-on-click-outside="cancelEdit"
        @dropdownVisibleChange="dropdownVisibleChange"
        placeholder="请选择"
        :options="options">
    </a-select>
  </div>
  <div :class="['editor','show-on-hover', customClass]" v-else>
    <span class="title" @click.stop="handleClick">{{ label }}</span> &nbsp;&nbsp;
    <span class="edit-icon"><EditOutlined @click.stop="isEditing = true"/></span>
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
} from '@ant-design/icons-vue';
import { vOnClickOutside } from '@vueuse/components';
const isEditing = ref(false);
const fieldValue:any = ref('');
const props = defineProps({
  value: {
    required: true,
  },
  label: {
    required: true,
  },
  options: {
    required: true,
    type: Object,
  },
  customClass: {
    required: false,
    type: String,
  }
})
const emit = defineEmits(['update', 'edit']);

function updateField() {
  // if (!fieldValue.value) {
  //   return;
  // }
  emit('update', fieldValue.value);
  isEditing.value = false;
}

function cancelEdit() {
  if(isOpen.value){
    return;
  }
  fieldValue.value = props.value;
  isEditing.value = false;
}

const isOpen = ref(false);

function dropdownVisibleChange(open) {
  isOpen.value = open;
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
