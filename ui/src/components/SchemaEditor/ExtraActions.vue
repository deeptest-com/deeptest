<script lang="ts" setup>
import {
  DeleteOutlined,
  InfoCircleOutlined,
  ReadOutlined,
  InfoCircleTwoTone
} from '@ant-design/icons-vue';
import {computed, defineProps, defineEmits, ref, watch} from "vue";

const props = defineProps<{
  isFirst?: boolean | undefined,
  isLast?: boolean | undefined,
  isRoot?: boolean | undefined,
  isRefChildNode?: boolean | undefined,
  value?: any
}>();

const emit = defineEmits<{
  (e: 'setRequire', id?: number): void,
  (e: 'addDesc', value?: string): void,
  (e: 'del', value?: string): void
}>()

const visible = ref(false);

const disableSetRequire = computed(() => {
  if (props.isRoot) {
    return true;
  }
  if (props.isRefChildNode) {
    return true;
  }
  return false;
});
const disableAddDesc = computed(() => {
  return false;
});
const disableDel = computed(() => {
  if (props.isRoot) {
    return true;
  }
  if (props.isRefChildNode) {
    return true;
  }
  return false;
});

// 该字段是否必填
const isRequired = computed(() => {
  const extraViewInfo = props.value?.extraViewInfo;
  return !!(extraViewInfo && extraViewInfo?.parent?.required?.includes(extraViewInfo?.keyName));
});

const description = ref('');
watch(() => {return props.value}, (newVal) => {
  if (newVal) {
    description.value = newVal.description;
  }
}, {immediate: true})

watch(() => {return visible.value}, (newVal) => {
  if (!newVal && description.value) {
    emit('addDesc', description.value);
  }
})

</script>

<template>
  <a-tooltip placement="topLeft" :title="disableSetRequire ? null :  '是否必填？'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableSetRequire" type="text" @click="emit('setRequire')">
      <template #icon>
        <InfoCircleOutlined  v-if="!isRequired" />
        <InfoCircleTwoTone   v-if="isRequired" />
      </template>
    </a-button>
  </a-tooltip>

  <a-popover :title="null"
             trigger="click"
             v-model:visible="visible"
             placement="left"
             :overlayClassName="'container'">
    <template #content>
      <div class="content">
        <a-input v-model:value="description"
                 :disabled="isRefChildNode"
                 placeholder="请输入描述信息"/>
      </div>
    </template>
    <a-tooltip placement="topLeft" :title="disableAddDesc ? null :  '描述'" arrow-point-at-center>
      <a-button :size="'small'" :disabled="disableAddDesc" type="text" @click="visible = true">
        <template #icon>
          <ReadOutlined/>
        </template>
      </a-button>
    </a-tooltip>
  </a-popover>
  <a-tooltip placement="topLeft" :title="disableDel ? null :  '删除'" arrow-point-at-center>
    <a-button :size="'small'" :disabled="disableDel" type="text" @click="emit('del')">
      <template #icon>
        <DeleteOutlined/>
      </template>
    </a-button>
  </a-tooltip>
</template>


