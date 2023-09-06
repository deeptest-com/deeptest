<template>
  <div class="endpoint-mock-expect-main">
    <div class="toolbar">
      <a-button @click="createExpect" type="primary" class="btn-create">新建期望</a-button>
    </div>
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

const createExpect = () => {
  console.log('createExpect')
}

</script>

<style lang="less" scoped>
.endpoint-mock-expect-main {
  height: 100%;
  position: relative;

  .toolbar {
    position: absolute;
    top: -80px;
    right: 0px;

    text-align: right;

    .btn-create {
      margin-left: 30px;
    }
  }
}
</style>

