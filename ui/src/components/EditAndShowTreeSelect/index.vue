<template>
  <div class="editor" v-if="isEditing"    v-on-click-outside="cancelEdit">
      <a-tree-select
          :value="fieldValue"
          :multiple="false"
          :treeData="treeData"
          style="width: 200px"
          :size="'small'"
          @select="updateField"
          :treeDefaultExpandAll="true"
          @treeExpand="treeExpand"
          :replaceFields="{ title: 'name',value:'id'}"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          placeholder="请选择所属分类"
          allow-clear/>
  </div>
  <div :class="['editor', customClass]" v-else>
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
const isEditing:any = ref(false);
const fieldValue:any = ref('');
const props = defineProps({
  value: {
    required: true,
    type:  Number,
  },
  label: {
    required: true,
    type: String || Number,
  },
  treeData: {
    required: true,
    type: Object,
  },
  customClass: {
    required: false,
    type: String,
  }
})
const emit = defineEmits(['update', 'edit']);

function updateField(value) {
  fieldValue.value = value;
  emit('update', fieldValue.value || null);
  isEditing.value = false;
}


function cancelEdit() {
  if(isEditing.value){
    fieldValue.value = props.value;
    isEditing.value = false;
  }
}

const isOpen = ref(false);

function treeExpand(open) {
  debugger;
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
