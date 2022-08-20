<template>
  <div class="scenario-exec-log-main">
    <div v-for="(item, index) in logs" :key="index" class="log">
      <div>
        <span v-if="!item.respContent">
          {{ item.name }}&nbsp; {{ joinArr(item.summary) }}
        </span>

        <a-collapse v-if="item.respContent" expand-icon-position="right">
          <a-collapse-panel key="index" :header="item.name + '&nbsp;' + joinArr(item.summary)">
            <div class="resp-content">{{ item.respContent }}</div>
          </a-collapse-panel>
        </a-collapse>
      </div>

      <Log v-if="item.logs" :logs="item.logs"></Log>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineProps} from "vue";
import {Collapse, CollapsePanel} from 'ant-design-vue';
import Log from "./Log.vue"

defineProps<{
  logs: []
}>()

const joinArr = (arr : string[]) => {
  if (!arr) return ''

  if (Array.isArray(arr)) {
    return arr.join(' ')
  } else {
    return arr + ''
  }
}

</script>

<style lang="less" scoped>
.scenario-exec-log-main {
  height: 100%;
  padding: 10px;
  .log {
    .resp-content {
      width: 100%;
      word-break:break-word;
      overflow: auto;
    }
  }
}

</style>