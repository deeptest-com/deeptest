<template>
  <div class="container">
    <a-spin tip="加载中..." :spinning="loading">
      <Docs :show-basic-info="true" :show-menu="true" :showHeader="true" :data="data"/>
    </a-spin>
  </div>
</template>
<script setup lang="ts">
import {
  computed,
  watch,
  ref,
  onMounted
} from 'vue';

import Docs from '@/components/Docs/index.vue';

import {useStore} from "vuex";

const store = useStore<{ Endpoint, ProjectGlobal }>();
import {useRouter} from "vue-router";

const loading = ref(false);
const router = useRouter();
const query: any = router.currentRoute.value.query;

const docVersions: any = ref([
  {
    value: '',
    label: 'latest',
  }
]);

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
    loading.value = true;
    data.value = await store.dispatch('Docs/getDocs', {
      projectId: currProject.value.id,
    });
    loading.value = false;
  }
}, {
  immediate: true
})

watch(() => {
  return endpointIds.value;
}, async (newVal) => {
  if (newVal && newVal.length > 0) {
    loading.value = true;
    data.value = await store.dispatch('Docs/getDocs', {
      endpointIds: newVal,
    })
    loading.value = false;
  }
}, {
  immediate: true
})

watch(() => {
  return serveIds.value;
}, async (newVal) => {
  if (newVal && newVal.length > 0) {
    loading.value = true;
    data.value = await store.dispatch('Docs/getDocs', {
      serveIds: newVal,
    })
    loading.value = false;
  }
}, {
  immediate: true
})

onMounted(async () => {
  // if (query.version) {
  //   docVersions.value = await store.dispatch('Docs/getDocVersions', {
  //     projectId: currProject.value.id,
  //   });
  // }
  docVersions.value = await store.dispatch('Docs/getVersionList', {
    projectId: currProject.value.id,
  });
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
