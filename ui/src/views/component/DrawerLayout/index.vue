<template>
  <a-drawer class="dp-drawer-container"
            :width="1200"
            :placement="'right'"
            :closable="true"
            :visible="visible"
            :headerStyle="{padding: '12px 16px',height: '48px',zIndex: 99,width: '100%'}"
            :bodyStyle="{padding: '0px',height: 'calc(100vh - 48px)' ,'overflow':'hidden'}"
            :wrapClassName="'abc-1'"
            @close="onCloseDrawer">
    <!-- 头部信息  -->
    <a-spin tip="Loading..." :spinning="spinning" style="z-index: 2000;">
    <template #title>
      <div  class="dp-drawer-header">
        <slot name="header"/>
      </div>

    </template>
    <div class="dp-drawer-content" ref="contentRef">
      <!-- 基本信息区域 -->
      <div class="dp-drawer-content-basic-info">
        <slot name="basicInfo"/>
      </div>
      <!-- Tab 切换区域头部信息 -->
      <div class="dp-drawer-content-tabs-header">
        <slot name="tabHeader"/>
      </div>
      <!-- Tab 切换区域内容区域 -->
      <div class="dp-drawer-content-tabs-content">
        <slot name="tabContent"/>
      </div>
    </div>
    </a-spin>
  </a-drawer>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  nextTick,
  defineEmits,
  computed, watch,
  onMounted,
} from 'vue';

import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {message} from "ant-design-vue";
import {StateType as ServeStateType} from "@/store/serve";



const store = useStore<{ Global}>();

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  // 每次变化时，触发吸顶操作
  stickyKey: {
    type: Number,
    required: true,
  }
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

function onCloseDrawer() {
  emit('close');
}

const contentRef: any = ref(null)

const spinning = computed( ()=>store.state.Global.spinning )


watch(() => {
  return props.stickyKey;
}, (newVal) => {
  if (newVal && contentRef?.value) {
    contentRef?.value?.scrollTo(0, 78);
  }
})


</script>

<style lang="less" scoped>
.dp-drawer-container {
  width: 100%;
  height: 100vh;

  .dp-drawer-content {
    height: calc(100vh - 48px);
    overflow-y: scroll;
    overflow-x: hidden;
  }

  .dp-drawer-content-basic-info {
    padding: 16px 24px 0 24px;
  }

  .dp-drawer-content-tabs-header {
    //position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    height: 48px;
    border-bottom: 1px solid #f0f0f0;
    margin: 0 16px;
    z-index: 9999;
    background-color: #ffffff;

    :deep(.tab-header-items) {
      width: 80%;
      display: flex;
      align-items: center;
    }

    :deep(.tab-header-btns) {
      width: 20%;
      display: flex;
      justify-content: flex-end;
    }

    :deep(.tab-header-items .tab-header-item) {
      color: #000000d9;
      position: relative;
      margin: 0 32px 0 0;
      padding: 12px 16px;
      text-decoration: none;
      cursor: pointer;
    }

    :deep(.tab-header-items .tab-header-item:hover) {
      color: #40a9ff;
    }

    :deep(.tab-header-items .tab-header-item.active) {
      color: #1890ff;
    }

    :deep(.tab-header-items  .tab-header-item.active:after) {
      content: "";
      position: absolute;
      left: 0;
      bottom: 0;
      height: 2px;
      background-color: #1890ff;
      width: 100%;
    }
  }

  .dp-drawer-content-tabs-content {
    padding: 0 16px;
    :deep(.tab-pane) {
      position: relative;
      //margin-top: 16px;
      min-height: calc(100vh - 96px);
    }
  }
  .dp-drawer-header{
    :deep(.header-text) {
      display: flex;
      max-width: 80%;
    }
    :deep(.header-text .serialNumber) {
      margin-right: 6px;
    }
  }
}

</style>
