<template>
  <a-drawer class="dp-drawer-container"
            :width="1200"
            :placement="'right'"
            :closable="true"
            :visible="visible"
            :headerStyle="{padding: '12px 16px',position: 'fixed',zIndex: 9,width: '100%'}"
            :bodyStyle="{padding: '0px',height: '100vh'}"
            @close="onCloseDrawer">
    <!-- 头部信息  -->
    <template #title>
      <slot name="title"/>
    </template>
    <div class="dp-drawer-content" ref="tabsRef">
      <!-- 基本信息区域 -->
      <div class="dp-drawer-content-basic-info">
        <slot name="basicInfo"/>
      </div>
      <!-- Tab 切换区域 -->
      <div class="dp-drawer-content-tabs-sticky" >
        <slot name="Tabs"/>
      </div>
    </div>

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


const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

function onCloseDrawer() {
  emit('close');
}

const tabsRef: any = ref(null);

onMounted(() => {

  setTimeout(() => {
    debugger;
    tabsRef.value?.addEventListener('scroll', (e) => {
      console.log('scroll', e);
    })
  }, 2000);
})

// watch(() => {
//   return props.visible;
// }, (newVal) => {
//   if (newVal) {
//     setTimeout(() => {
//       debugger;
//       tabsRef.value?.addEventListener('scroll', (e) => {
//         console.log('scroll', e);
//       })
//     }, 1000);
//
//   }
// })  // 监听抽屉的显示隐藏

</script>

<style lang="less" scoped>
.dp-drawer-container {
  width: 100%;
  height: 100vh;

  .dp-drawer-content {
    height: 100vh;
    margin-top: 48px;
  }

  .dp-drawer-content-basic-info {
    padding: 16px 24px;
  }

  .dp-drawer-content-tabs-sticky {

  }

}

.drawer-btns {
  background: #ffffff;
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  position: absolute;
  bottom: 0;
  //right: 0;
  width: 100%;
  padding-right: 24px;
  height: 56px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-right: 16px;
  z-index: 99;
}

.header-text {
  display: flex;
  max-width: 80%;

  .serialNumber {
    margin-right: 6px;
  }
}
</style>
