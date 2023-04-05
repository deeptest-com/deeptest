<template>
  <div class="bearer-token-main author-content">
    <div class="params">
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">Token</span>
        </a-col>

        <a-col flex="1">
          <a-input v-model:value="debugData.bearerToken.token" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
    </div>
    <div class="tips">
      <div class="dp-light">授权头将会在你发送请求时自动生成。</div>
      <div class="dp-link-primary">了解更多 <ArrowRightOutlined /></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ArrowRightOutlined} from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

</script>

<style lang="less" scoped>
.bearer-token-main {
}

</style>