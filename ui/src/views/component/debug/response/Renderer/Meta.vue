<template>
  <div class="response-meta">
    <div class="item" :class="[responseData.statusCode===200? 'dp-color-pass': 'dp-color-fail']">
      状态：{{ responseData.statusContent }}
    </div>

    <div class="item">耗时：{{ responseData.time }} 毫秒</div>
    <div class="item">大小：{{ responseData.contentLength }} 字节</div>
  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {useI18n} from "vue-i18n";
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const responseData = computed<any>(() => store.state.Debug.responseData);

</script>

<style lang="less" scoped>
.response-meta {
  padding: 0 6px;

  .item {
    padding: 2px 0;
  }
}

</style>