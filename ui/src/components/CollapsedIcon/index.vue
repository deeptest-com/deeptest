<template>
  <div class="expand-icon" @click="toggleCollapsed"
       :class="{'collapsed':collapsed}"
       :style="collapsed ? collapsedStyle  :  style">
    <MenuUnfoldOutlined class="icon" v-if="collapsed"/>
    <MenuFoldOutlined class="icon" v-if="!collapsed"/>
  </div>
</template>
<script lang="ts">

import {defineComponent, onMounted, PropType, ref} from "vue";

import {
  // CaretRightOutlined,
  // CaretLeftOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined
} from '@ant-design/icons-vue';

export default defineComponent({
  name: 'CollapsedIcon',
  props: {
    style: {
      type: Object,
    },
    collapsedStyle: {
      type: Object,
    },
    collapsed: {
      type: Boolean,
      default: false
    },
    toggleCollapsed: {
      type: Function as PropType<() => void>
    },
  },
  components: {
    // CaretRightOutlined,
    // CaretLeftOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined
  },

})

</script>

<style lang="less" scoped>
.expand-icon {
  position: absolute;
  height: 20px;
  width: 20px;
  opacity: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  left: 8px;
  top: 22px;
  z-index: 9;
  cursor: pointer;
  background-color: #FFF;
  &.collapsed {
    left: 8px;
  }
  .icon {
    color: #8A8A8A;
    font-size: 18px;
  }
}
</style>
