<template>
  <div class="response-meta">
    <a-row type="flex" :gutter="15">
      <a-col>
        <span>状态：{{ responseData.statusContent }}</span>
      </a-col>
      <a-col>
        <span>耗时：{{ responseData.time }} 毫秒</span>
      </a-col>
      <a-col>
        <span>大小：{{ responseData.contentLength }} 字节</span>
      </a-col>
    </a-row>
  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
const {t} = useI18n();

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const responseData = computed<any>(() => store.state.Debug.responseData);

</script>

<style lang="less" scoped>
.response-meta {
  padding: 6px;
}

</style>