<template>
  <div class="container" :class="{'full-container':isDocsSharePage || isDocsViewPage}">
    <a-spin tip="加载中..." :spinning="loading">
      <Docs :show-menu="true"
            :showHeader="true"
            :data="data"/>
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

const store = useStore<{ Docs, ProjectGlobal }>();
import {useRouter} from "vue-router";

const loading = ref(false);
const router = useRouter();
const query: any = router.currentRoute.value.query;
const path: any = router.currentRoute.value.path;
// 是否分享页面
const isDocsSharePage = path.includes('/share');
const isDocsViewPage = path.includes('/view');

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
const currDocId = computed<any>(() => store.state.Docs.currDocId);

const shareId: any = computed(() => {
  if (query.code) {
    console.log('query.code', query.code)
    return query.code;
  } else {
    return '';
  }
});

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
});


watch(() => {
  return endpointIds.value;
}, async (newVal) => {
  if (isDocsSharePage) {
    return;
  }
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
  if (isDocsSharePage) {
    return;
  }
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


// 监控文档版本的变化
watch(() => {
  return currDocId.value
}, async (newVal) => {
  if (isDocsSharePage || isDocsViewPage) {
    return;
  }
  if (newVal || newVal === 0) {
    loading.value = true;
    data.value = await store.dispatch('Docs/getDocs', {
      documentId: newVal,
      projectId: currProject.value.id,
    })
    loading.value = false;
  }
})

// 获取版本列表
onMounted(async () => {
  if (isDocsSharePage || isDocsViewPage) {
    return;
  }
  await store.dispatch('Docs/getVersionList', {
    needLatest: true,
  });
})

watch(() => {return shareId.value}, async (newVal) => {
  if (newVal) {
    loading.value = true;
    data.value = await store.dispatch('Docs/getShareContent', {
      code: newVal,
    });
    loading.value = false;
  }
},{
  immediate: true
})

</script>
<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
  //min-height: calc(100vh - 106px);
  overflow: hidden;
}
.full-container{
  margin: 0;
  //height: 100vh;
}

</style>
