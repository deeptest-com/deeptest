<template>
  <div class="con-box-title" :style="backgroundStyle"  @click="expandInfo">
    <span>{{ title }}</span>
    <span v-if="showArrow" class="icon">
        <DownOutlined v-if="expand"/>
        <RightOutlined v-if="!expand"/>
      </span>
  </div>
</template>
<script setup lang="ts">
import {defineProps, ref,defineEmits} from 'vue';
import {RightOutlined, DownOutlined} from '@ant-design/icons-vue'
const emit = defineEmits(['expand']);

defineProps<{
  title: string
  backgroundStyle?: string,
  showArrow?: boolean,
}>()


const expand = ref(true);

function expandInfo() {
  expand.value = !expand.value;
  emit('expand', expand.value);
}
</script>
<style scoped lang="less">
.con-box-title {
  padding-left: 12px;
  cursor: pointer;
  width: 100%;
  height: 36px;
  line-height: 36px;
  display: flex;
  align-items: center;
  text-align: left;
  font-size: 16px;
  color: rgba(0, 0, 0, 0.85);
  position: relative;
  font-weight: bold;
  display: flex;
  justify-content: space-between;

  &:before {
    content: '';
    position: absolute;
    width: 3px;
    height: 16px;
    background: #447DFD;
    top: 10px;
    left: 0;
  }
}
.icon {
  margin-right: 8px;
}
</style>
