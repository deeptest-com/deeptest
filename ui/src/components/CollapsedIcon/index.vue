<template>
  <div class="expand-icon" @click="toggleCollapsed"
       :class="{'collapsed':collapsed}"
       :style="collapsed ? collapsedStyle  :  style">
    <CaretRightOutlined class="icon" v-if="collapsed"/>
    <CaretLeftOutlined class="icon" v-if="!collapsed"/>
  </div>
</template>
<script lang="ts">

import {defineComponent, onMounted, PropType, ref} from "vue";

import {
  CaretRightOutlined,
  CaretLeftOutlined
} from '@ant-design/icons-vue';

export default defineComponent({
  name: 'CollapsedIcon',
  props: {
    style: {
      type: Object,
      default: () => {
        return {};
      }
    },
    collapsedStyle: {
      type: Object,
      default: () => {
        return {};
      }
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
    CaretRightOutlined,
    CaretLeftOutlined,
  },

})

</script>

<style lang="less" scoped>
.expand-icon {
  position: fixed;
  width: 12px;
  height: 56px;
  top: 115px;
  opacity: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  left: 223px;
  z-index: 9;
  cursor: pointer;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 0px 8px 8px 0px;
  transition: left 0.2s;
  &.collapsed {
    left: 54px;
  }
  .icon {
    font-size: 12px;
    color: #8A8A8A;;
  }
}
</style>
