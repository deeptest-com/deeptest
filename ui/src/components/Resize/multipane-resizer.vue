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
.dp-multipane-con.layout-v .multipane-resizer {
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
    background-color: #1890ff;
  }

  &.collapsed {
    &:after {
      display: none;
    }
  }

}

.dp-multipane-con.layout-h .multipane-resizer {
  position: relative;
  z-index: 2;
  height: 2px;
  width: 100%;
  display: flex;
  align-items: center;
  &:after {
    height: 2px;
    content: '';
    width: 100%;
    position: absolute;
    background-color: #f0f0f0;
    left:0;
  }

  &:hover:after {
    cursor: row-resize;
    height: 2px;
    width: 100%;
    background-color: #1890ff;
  }

  &.collapsed {
    &:after {
      display: none;
    }
  }

}


</style>
