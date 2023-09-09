<template>
  <div class="editor show-on-hover" v-if="isEditing"    v-on-click-outside="cancelEdit">
      <a-tree-select
          :value="fieldValue"
          :multiple="false"
          :treeData="treeData"
          style="width: 200px"
          :size="'small'"
          :show-search="showSearch"
          @select="updateField"
          :treeDefaultExpandAll="true"
          :searchPlaceholder="'请输入'"
          :searchValue="searchValue"
          @search="search"
          @treeExpand="treeExpand"
          tree-node-filter-prop="name"
          :replaceFields="{ title: 'name',value:'id'}"
          :dropdown-style="{ maxHeight: '400px', overflow: 'auto' }"
          placeholder="请选择"/>
  </div>
  <div :class="['editor', 'show-on-hover']" v-else>
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
  },
  showSearch: {
    required: false,
    type: Boolean,
    default: false,
  },
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
  isOpen.value = open;
}

const searchValue = ref(props.label || '');
function search(value) {
    searchValue.value = value;
}

function handleClick() {
  emit('edit');
}

watch(() => {return props.value}, (newVal) => {
  fieldValue.value = newVal
}, {immediate: true})

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
