<template>
  <div class="param-content-main">
    <div class="head">
      <a-row type="flex">
        <a-col class="title" flex="1">{{title}}</a-col>
        <a-col flex="80px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>复制</template>
            <CopyOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="content">
      {{content}}
    </div>
  </div>
</template>

<script setup lang="ts">
import {CopyOutlined} from '@ant-design/icons-vue';
import {useI18n} from "vue-i18n";
import {defineProps, PropType} from "vue";
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
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
    .title {
      font-weight: bolder;
    }
  }
  .content {
    padding: 10px 10px;
    word-break: break-all;
  }
}

</style>