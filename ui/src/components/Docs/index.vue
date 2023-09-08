<template>
  <div v-if="data?.name && serviceList?.length"
       :class="{'content':isDocsPage,'full-content':isDocsFullPage,'drawer-content':isEndpointPage}">
    <DocsHeader v-if="showHeader"
                :data="data"
                :items="serviceList"
                @select="selectMenu"
                @changeVersion="changeVersion"/>
    <a-divider style="margin:0" v-if="showHeader"/>
    <div class="doc-container">
      <ContentPane :containerStyle="{padding:0,margin:0,height:'100%',width:'100%'}">
        <template #left>
          <div class="left" v-if="showMenu">
            <LeftTreeView :serviceList="serviceList" @select="selectMenu" :selectedKeys="selectedKeys"/>
          </div>
        </template>
        <template #right>
          <div class="right" :class="{'only-docs':!showMenu}">
            <EndpointDoc v-if="selectedItem" :info="selectedItem" :onlyShowDocs="onlyShowDocs"/>
          </div>
        </template>
      </ContentPane>
    </div>
  </div>
  <a-skeleton v-if="!data?.name"/>
  <!--    没有定义接口文档，则展示空信息   -->
  <div v-if="data?.name && serviceList?.length ===0" style="margin-top: 48px;">
    <a-empty
        image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
        :image-style="{height: '60px',}">
      <template #description>
                <span>
                  您还未定义接口，请先定义接口
                </span>
      </template>
      <a-button v-if="isEndpointPage" type="primary" @click="emit('switchToDefineTab')">接口定义</a-button>
    </a-empty>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed,
  watch,
} from 'vue';
import {useStore} from "vuex";

import LeftTreeView from "./components/LeftTreeView.vue";
import EndpointDoc from "./components/EndpointDoc.vue";
import DocsHeader from "./components/DocsHeader.vue";
import ContentPane from '@/views/component/ContentPane/index.vue';

const store = useStore<{ Docs }>();
const props = defineProps(['showMenu', 'data', 'onlyShowDocs', 'showHeader']);
const emit = defineEmits(['changeVersion', 'switchToDefineTab']);
const currDocId = computed<any>(() => store.state.Docs.currDocId);
const isEndpointPage = window.location.href.includes('/endpoint/index');
const isSharePage = window.location.href.includes('/docs/share');
const isViewPage = window.location.href.includes('/docs/view');
const isDocsPage = window.location.href.includes('/docs/index');

const isDocsFullPage = isSharePage || isViewPage;

const serviceList = computed(() => {
  // 组装数据以兼容组件 LeftTreeMenu
  let items: any = [];
  props?.data?.serves?.forEach((item: any) => {
    // 只显示文档，不展示服务信息
    if (!props.onlyShowDocs) {
      items.push(item);
    }
    item?.endpoints?.forEach((endpoint: any) => {
      endpoint?.interfaces?.forEach((interfaceItem: any) => {
        items.push({
          ...interfaceItem,
          endpointInfo: endpoint,
          serveInfo: item,
          serveId: item.id,
        });
      })
    })
  })
  return items;
})

const selectedItem: any = ref(null);

const selectedKeys = computed(() => {
  if (!selectedItem.value?.id) {
    return [];
  }
  return [`${selectedItem.value?.method || 'serve'}-${selectedItem.value.id}`];
})

watch(() => {
  return serviceList.value
}, async (newVal) => {
  if (newVal.length > 0) {
    selectedItem.value = newVal[0];
  }
}, {immediate: true});


async function selectMenu(item) {
  selectedItem.value = item;

  // 如果是接口定义页面，则不请求文档详情,会一次性请求所有接口的文档详情
  if (isEndpointPage) {
    return;
  }

  // 仅选择文档时，才请求文档详情
  if (item.endpointInfo && item.serveInfo) {
    let res: any = null;
    if (isSharePage) {
      res = await store.dispatch('Docs/getShareDocsDetail', {
        currProjectId: item.serveInfo.projectId,
        interfaceId: item.id,
        documentId: props?.data?.documentId || 0,
        endpointId: item.endpointInfo.id,
      });
    } else {
      res = await store.dispatch('Docs/getDocsDetail', {
        currProjectId: item.serveInfo.projectId,
        interfaceId: item.id,
        documentId: currDocId?.value || 0,
        endpointId: item.endpointInfo.id,
      });
    }
    if (res?.interface) {
      selectedItem.value = {
        ...selectedItem.value,
        ...res.interface,
        mock: res?.mock || null
      };
    }
  }
}

function changeVersion(docId) {
  emit('changeVersion', docId);
}


</script>

<style lang="less" scoped>

.right {
  flex: 1;
  height: 100%;
  overflow: auto;
  padding: 12px 24px 96px 24px;
}

// 文档页面
.content {
  height: calc(100vh - 100px);

  .doc-container {
    display: flex;
    height: calc(100vh - 156px);

    .left {
      margin-left: 12px;
    }

    .only-docs {
      padding: 0;
    }
  }
}

// 文档分享页和查看页
.full-content {
  min-height: 100vh !important;

  .doc-container {
    display: flex;
    top: 0;
    height: calc(100vh - 56px);
  }
}

// 抽屉里的文档区域
.drawer-content {
  height: calc(100vh - 96px) !important;

  .doc-container {
    position: relative;

    .left {
      margin-left: 0 !important;
      height: calc(100vh - 96px);
    }
  }
}
</style>
