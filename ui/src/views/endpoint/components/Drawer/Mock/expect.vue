<template>
  <div class="endpoint-mock-expect-main">
    1
  </div>
</template>

<script setup lang="ts">
import {ref, computed, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";
import {getUrls} from "@/utils/request";

const {t} = useI18n()

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);


</script>

<style lang="less" scoped>
.endpoint-mock-expect-main {
  height: 100%;
}
</style>

