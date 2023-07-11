<template>
  <div class="multipane-resizer" :class="{'collapsed':collapsed}">
    <slot></slot>
  </div>
</template>

<script>


export default {
  name: 'multipane-resizer',
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },

  },

};
</script>
<style lang="less" scoped>
.multipane-resizer {
  position: relative;
  z-index: 2;
  width: 1px;
  display: flex;
  align-items: center;
  &:after {
    width: 1px;
    content: '';
    height: 100%;
    position: absolute;
    background-color: #f0f0f0;
    left:0;
  }

  &:hover:after {
    cursor: col-resize;
    width: 2px;
  }

  &.collapsed {
    &:after {
      display: none;
    }
  }

}


</style>
