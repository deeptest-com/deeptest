<template>
  <div class="param-content-main">
    <div class="head">
      <ConBoxTitle :title="title" />
    </div>

    <div class="content">
      {{content}}
    </div>
  </div>
</template>

<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {defineProps, PropType} from "vue";
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
const {t} = useI18n();

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  content: {
    type: String,
    required: true
  },
})
</script>

<style lang="less" scoped>
.param-content-main {
  margin-bottom: 10px;
  .head {

  }
  .content {
    padding: 10px 10px;
    word-break: break-all;
  }
}

</style>