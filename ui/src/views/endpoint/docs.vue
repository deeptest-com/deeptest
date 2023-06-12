<template>
  <div class="container">
    <Docs :show-basic-info="true" :show-menu="true" :data="data"/>
  </div>
</template>
<script setup lang="ts">
import {
  computed,
  watch,
  ref
} from 'vue';

import Docs from '@/components/Docs/index.vue';

import {useStore} from "vuex";


const store = useStore<{ Endpoint, ProjectGlobal }>();

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const data = ref<any>(null);
watch(() => {
  return currProject.value.id;
}, async (newVal) => {
  if (newVal) {
    const docs = await store.dispatch('Endpoint/getDocs', {
      projectId: currProject.value.id,
    });
    data.value = docs;
  }
}, {
  immediate: true
})

</script>
<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
  min-height: calc(100vh - 106px);
  overflow: hidden;
}

</style>
