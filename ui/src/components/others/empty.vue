<template>
  <a-empty :image="simpleImage">
    <template #description>
      {{desc}}
    </template>
  </a-empty>
</template>
<script lang="ts">
import { defineComponent, ref } from 'vue';
import { Empty } from 'ant-design-vue';
export default defineComponent({
  name: 'EmptyPage',
  props: {
    desc: {
      type: String,
      default: ''
    },
  },
  setup(props) {
    const desc = ref(props.desc)
    if (!desc.value) {
      desc.value = '暂无数据'
    }

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