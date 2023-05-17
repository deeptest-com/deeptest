<template>
  <div class="author-basic-main author-content">
    <div class="params">
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">键值</span>
        </a-col>

        <a-col flex="1">
          <a-input v-model:value="debugData.apiKey.key" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="160px">
          <span class="label">取值</span>
        </a-col>

        <a-col flex="1">
          <a-input v-model:value="debugData.apiKey.value" class="dp-bg-input-transparent" />
        </a-col>
      </a-row>
      <a-row class="param">
        <a-col flex="80px">传递方式</a-col>
        <a-col flex="1">
          <a-select
              v-model:value="debugData.apiKey.transferMode"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
            <a-select-option value="headers">请求头</a-select-option>
            <a-select-option value="queryParams">查询参数</a-select-option>
          </a-select>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ArrowRightOutlined} from '@ant-design/icons-vue';
import {UsedBy} from "@/utils/enum";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {Param} from "@/views/component/debug/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

</script>

<style lang="less" scoped>
.author-basic-main {
}

</style>