<template>
  <div class="param-grid-main">
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

    <div class="items">
      <a-row v-for="(item, idx) in list" :key="idx" type="flex" class="item">
        <a-col flex="1">
          <a-input v-model:value="item.name" class="dp-bg-input-transparent" />
        </a-col>
        <a-col flex="1">
          <a-input v-model:value="item.value" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>

      <div v-if="list?.length==0" class="empty">
        无
      </div>
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
  list: {
    type: Array,
    required: true
  },
})

</script>

<style lang="less" scoped>
.param-grid-main {
  margin-bottom: 10px;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
    .title {
      font-weight: bolder;
    }
  }
  .items {
    .item {
      padding: 10px 3px;
      border-bottom: 1px solid #d9d9d9;
    }
  }
  .empty {
    padding: 10px 10px;
  }
}

</style>