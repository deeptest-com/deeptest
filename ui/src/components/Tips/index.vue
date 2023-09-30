<template>
  <span class="tips-main">
    <a-tooltip overlayClassName="dp-tip-small dp-tip-white">
      <template #title>
        <div class="tips">
          <div>{{title}}</div>
          <div @click="openHelp(section)" class="dp-link-primary">
            了解更多 <ArrowRightOutlined/>
          </div>
        </div>
      </template>
      <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
    </a-tooltip>
  </span>
</template>

<script setup lang="ts">
import {defineProps} from "vue";
import { QuestionCircleOutlined, ArrowRightOutlined } from '@ant-design/icons-vue';
import {openHelp} from "@/utils/link";

defineProps<{
  title: string
  section: string
}>()

</script>

<style lang="less" scoped>
.tips-main {

}
</style>