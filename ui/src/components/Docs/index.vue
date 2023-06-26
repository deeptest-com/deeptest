<!-- :::: 接口定义模块 -->
<template>
  <div class="content" v-if="data?.name">
<!--    <BasicDetail  :items="items" v-if="showBasicInfo"/>-->
    <DocsHeader v-if="showHeader" :data="items" :items="serviceList" @select="selectSugRes"/>
    <a-divider style="margin:0" v-if="showBasicInfo"/>
    <div class="doc-container">
      <div class="left" v-if="showMenu">
        <LeftTreeView :serviceList="serviceList" @select="selectMenu" :selectedKeys="selectedKeys"/>
      </div>
      <div class="right" :class="{'only-docs':!showMenu}">
        <EndpointDoc v-if="selectedItem" :info="selectedItem" :onlyShowDocs="onlyShowDocs"/>
      </div>
    </div>
  </div>
  <a-skeleton v-else/>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, onMounted, watch,
} from 'vue';
import {useStore} from "vuex";

import BasicDetail from "./components/BasicDetail.vue";
import LeftTreeView from "./components/LeftTreeView.vue";
import EndpointDoc from "./components/EndpointDoc.vue";
import DocsHeader from "./components/DocsHeader.vue";

const store = useStore<{ Endpoint, ProjectGlobal }>();
const props = defineProps(['showBasicInfo', 'showMenu', 'data', 'onlyShowDocs','showHeader']);
const emit = defineEmits([]);

const items = computed(() => {
  return [
    {
      label: '项目名称',
      value: props?.data?.name,
    },
    {
      label: '项目描述',
      value: props?.data?.desc,
    },
  ]
})

const serviceList = computed(() => {
  // 组装数据以兼容组件 LeftTreeMenu
  let items: any = [];

  props?.data?.serves.forEach((item: any) => {
    console.log(832,'item', item)
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
  console.log(832,'items', items)
  return items;
})

const selectedItem: any = ref(null);

const selectedKeys = computed(() => {
  if (!selectedItem.value?.id) {
    return [];
  }
  return [selectedItem.value?.id];
})

watch(() => {return serviceList.value}, (newVal) => {
  if (!selectedItem.value && newVal.length > 0) {
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

</script>

<style lang="less" scoped>
.content {
  //padding: 24px;
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
</style>
