<template>
  <div class="content" :class="{'full-content':isDocsFullPage}"  v-if="data?.name && serviceList?.length">
    <DocsHeader v-if="showHeader"
                :data="data"
                :items="serviceList"
                @select="selectSugRes"
                @changeVersion="changeVersion"/>
    <a-divider style="margin:0" v-if="showHeader"/>
    <div class="doc-container">
      <div class="left" v-if="showMenu">
        <LeftTreeView :serviceList="serviceList" @select="selectMenu" :selectedKeys="selectedKeys"/>
      </div>
      <div class="right" :class="{'only-docs':!showMenu}">
        <EndpointDoc v-if="selectedItem" :info="selectedItem" :onlyShowDocs="onlyShowDocs"/>
      </div>
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
  computed, onMounted, watch,
} from 'vue';
import {useStore} from "vuex";

import LeftTreeView from "./components/LeftTreeView.vue";
import EndpointDoc from "./components/EndpointDoc.vue";
import DocsHeader from "./components/DocsHeader.vue";

const store = useStore<{ Endpoint, ProjectGlobal }>();
const props = defineProps(['showMenu', 'data', 'onlyShowDocs', 'showHeader']);
const emit = defineEmits(['changeVersion','switchToDefineTab']);

const isEndpointPage = window.location.href.includes('/endpoint/index');
const isDocsFullPage = window.location.href.includes('/docs/share') || window.location.href.includes('/docs/view');

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
  return [selectedItem.value?.id];
})

watch(() => {
  return serviceList.value
}, (newVal) => {
  if (newVal.length > 0) {
    selectedItem.value = newVal.find((item) => {
      return item.endpointInfo && item.serveInfo;
    })
    if (!selectedItem.value) {
      selectedItem.value = newVal[0];
    }
  }
}, {immediate: true});


function selectSugRes(item) {
  selectedItem.value = item
}


function selectMenu(item) {
  selectedItem.value = item
}


function changeVersion(docId) {
  // debugger;
  emit('changeVersion', docId);
}


</script>

<style lang="less" scoped>
.content {
  height: calc(100vh - 100px);
  position: relative;
}

.doc-container {
  display: flex;
  height: 100%;

  .left {
    width: 300px;
    height: 100%;
    overflow: hidden;
    //margin-left: 24px;
    //padding: 0 12px;
    //border-right: 1px solid #f0f0f0;
    overflow-y: scroll;
    position: relative;

    &:before {
      content: '';
      position: absolute;
      top: 0;
      right: 0;
      height: 100%;
      z-index: 99;
      background-color: #f0f0f0;
      width: 1px;
    }
  }

  .right {
    flex: 1;
    height: 100%;
    overflow: auto;
    padding: 12px 24px 96px 24px;
  }

  .only-docs {
    padding: 0;
  }
}
.full-content{
  height: calc(100vh - 56px);
}
</style>
