<template>
  <a-empty :image="simpleImage">
    <template #description>
    </template>
  </a-empty>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import { Empty } from 'ant-design-vue';
export default defineComponent({
  setup() {
    return {
      simpleImage: Empty.PRESENTED_IMAGE_SIMPLE,
    };
  },
});
</script>

<style lang="less" scoped>
  .ant-empty-normal {
    margin-bottom: 0;
    height: auto;
  }
</style>