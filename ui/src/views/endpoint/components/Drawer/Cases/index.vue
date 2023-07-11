<template>
  <div class="endpoint-debug-cases-main">
    <CaseList
        v-if="show === 'list'"
        :onDesign="design" />

    <CaseDesign
        v-if="show === 'design'"
        :onBack="back" />
  </div>
</template>

<script setup lang="ts">
import {provide, ref, computed} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {UsedBy} from "@/utils/enum";
import CaseList from "./list.vue";
import CaseDesign from "./design.vue";

const {t} = useI18n()

const store = useStore<{ Endpoint }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);

const show = ref('list')

const design = (record) => {
  console.log('design', record)
  show.value = 'design'

  store.commit('Endpoint/setEndpointCaseDetail', record);
}

const back = () => {
  console.log('back')
  show.value = 'list'
}

</script>

<style lang="less" scoped>
.endpoint-debug-cases-main {
  position: relative;
  height: 100%;
  .toolbar {
    position: absolute;
    top: -42px;
    right: 0;
    height: 50px;
    width: 100px;
  }
  .content {
    height: 100%;
  }
}
</style>

