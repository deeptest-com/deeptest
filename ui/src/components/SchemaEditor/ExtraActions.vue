<script lang="ts" setup>
import {
  DeleteOutlined,
  InfoCircleOutlined,
  ReadOutlined
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits} from "vue";


const props = defineProps<{
  isFirst?: boolean | undefined,
  isLast?: boolean | undefined,
  isRoot?: boolean | undefined
}>();

const emit = defineEmits<{
  (e: 'setRequire', id?: number): void,
  (e: 'addDesc', value?: string): void,
  (e: 'del', value?: string): void
}>()

const disableSetRequire = computed(() => {
  return props.isRoot;
});
const disableAddDesc = computed(() => {
  return false
});
const disableDel = computed(() => {
  return props.isRoot;
});

</script>

<template>
  <a-tooltip placement="topLeft" :title="disableSetRequire ? null :  '是否必填？'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableSetRequire" type="text" @click="emit('setRequire')">
      <template #icon>
        <InfoCircleOutlined/>
      </template>
    </a-button>
  </a-tooltip>

  <a-tooltip placement="topLeft" :title="disableAddDesc ? null :  '添加描述'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableAddDesc" type="text" @click="emit('addDesc')">
      <template #icon>
        <ReadOutlined/>
      </template>
    </a-button>
  </a-tooltip>

  <a-tooltip placement="topLeft" :title="disableDel ? null :  '删除'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableDel" type="text" @click="emit('del')">
      <template #icon>
        <DeleteOutlined/>
      </template>
    </a-button>
  </a-tooltip>
</template>


