<script lang="ts" setup>
import {
  ArrowUpOutlined,
  ArrowDownOutlined,
  CopyOutlined,
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits} from "vue";


const props = defineProps<{
  isFirst: boolean,
  isLast: boolean,
  isRoot: boolean
}>();

const emit = defineEmits<{
  (e: 'moveUp', id?: number): void,
  (e: 'moveDown', value?: string): void,
  (e: 'copy', value?: string): void
}>()

const disableMoveUp = computed(() => {
  if(props.isRoot){
    return true;
  }
  if(props.isFirst){
    return true;
  }
  return false;
});
const disableMoveDown = computed(() => {
  if(props.isRoot){
    return true;
  }
  if(props.isLast){
    return true;
  }
  return false;
});
const disableCopy = computed(() => {
  return props.isRoot;
});

</script>

<template>
  <a-tooltip placement="topLeft" :title="disableMoveUp ? null :  '向上移动'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableMoveUp" type="text" @click="emit('moveUp')">
      <template #icon>
        <ArrowUpOutlined/>
      </template>
    </a-button>
  </a-tooltip>
  <a-tooltip placement="topLeft" :title="disableMoveDown ? null :  '向下移动'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableMoveDown" type="text" @click="emit('moveDown')">
      <template #icon>
        <ArrowDownOutlined/>
      </template>
    </a-button>
  </a-tooltip>
  <a-tooltip placement="topLeft" :title="disableCopy ? null :  '复制'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableCopy" type="text" @click="emit('copy')">
      <template #icon>
        <CopyOutlined/>
      </template>
    </a-button>
  </a-tooltip>
</template>


