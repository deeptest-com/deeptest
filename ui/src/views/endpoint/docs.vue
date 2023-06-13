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
import {useRouter} from "vue-router";


const router = useRouter();
const query: any = router.currentRoute.value.query;
const endpointIds: any = computed(() => {
  if (query.endpointIds) {
    return query.endpointIds.split(',').map((item: any) => {
      return +item
    })
  } else {
    return [];
  }
});
const serveIds: any = computed(() => {
  if (query.serveIds) {
    return query.serveIds.split(',').map((item: any) => {
      return +item;
    })
  } else {
    return [];
  }
});
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const data = ref<any>(null);

watch(() => {
  return currProject.value.id;
}, async (newVal) => {
  if (endpointIds.value.length > 0 || serveIds.value.length > 0) {
    return;
  }
  if (newVal) {
    data.value = await store.dispatch('Endpoint/getDocs', {
      projectId: currProject.value.id,
    });
  }
}, {
  immediate: true
})

watch(() => {
  return endpointIds.value;
}, async (newVal) => {
  if (newVal && newVal.length > 0) {
    data.value = await store.dispatch('Endpoint/getDocs', {
      endpointIds: newVal,
    })
  }
}, {
  immediate: true
})

watch(() => {
  return serveIds.value;
}, async (newVal) => {
  if (newVal && newVal.length > 0) {
    data.value = await store.dispatch('Endpoint/getDocs', {
      serveIds: newVal,
    })
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
