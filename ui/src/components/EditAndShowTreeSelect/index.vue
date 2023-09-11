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
          v-model:searchValue="searchValue"
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

function updateField(value,node) {
  fieldValue.value = value;
  // searchValue.value = node.title;
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

const searchValue = ref('');
// function change(value,label) {
//     searchValue.value = label;
// }

function handleClick() {
  emit('edit');
}

watch(() => {return props.value}, (newVal) => {
  fieldValue.value = newVal
}, {immediate: true})

watch(() => {return props.label}, (newVal) => {
  searchValue.value = newVal
}, {immediate: true})

watch(() => {
  return isEditing.value
},(newVal) => {
  if(newVal){
    searchValue.value = props.label;
  }
})

/**
 * TODO 未来优化，antdv 2.0版本，下拉菜单展开时，点击输入框，会收起下拉菜单，导致无法输入内容，
 * 
 * 该组件，在不能判断下拉菜单是否展开，antdv 3.0可以支持,
 * 所以会存在，输入框输入内容时，但是下拉菜单没有展开的情况，
 * */

</script>

<style lang="less" scoped>
.editor {
  display: flex;
  align-items: center;

  :deep(.ant-select-selection-item){
    display: none;
  }

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
