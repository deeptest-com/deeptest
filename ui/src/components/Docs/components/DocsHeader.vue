<!--
 项目基本信息
-->
<template>
  <div class="docs-header">
    <div class="logo">
      <!--      <img height="36" src="../../../assets/images/logo.png" alt="">-->
      项目信息
    </div>
    <div class="search">
      <span class="left-divider"/>
      <SearchOutlined class="icon"/>
      <a-input class="search-input"  placeholder="输入关键词搜索文档..."/>
      <span class="search-shortcut">⌘ K</span>
    </div>
    <div class="space"/>
    <div class="action">
      <a-tooltip placement="bottom">
        <template #title>点击分享文档</template>
        <ShareAltOutlined class="action-item"/>
      </a-tooltip>

    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  ref,
  defineProps,
  defineEmits,
  computed, watch,
} from 'vue';

import {DownOutlined, RightOutlined, SearchOutlined,ShareAltOutlined} from '@ant-design/icons-vue';


const props = defineProps({
  items: {
    required: true,
    type: Object,
  },
})


const emit = defineEmits(['ok', 'close', 'refreshList']);

const expand = ref(true);

function switchExpand() {
  expand.value = !expand.value;
}


</script>
<style lang="less" scoped>
.docs-header {
  display: flex;
  justify-content: space-between;
  height: 56px;
  align-items: center;

  .logo {
    width: 294px;
    padding-left: 24px;
  }

  .search {
    width: 280px;
    display: flex;
    align-items: center;

    .search-input, .search-input:active, .search-input:hover, .search-input:focus {
      width: 180px;
      border: none !important;
      outline: none;
      box-shadow: none;

    }
    .left-divider{
      border-left: 1px solid rgba(0, 0, 0, 0.06) !important;
      width: 1px;
      padding-right: 12px;
      height: 18px;
    }
    .icon{
      opacity: 0.6;
    }
    .search-shortcut{
      opacity: 0.9;
      color: #ced4d9;
      background-color: rgba(150, 150, 150, 0.06);
      border-color: rgba(100, 100, 100, 0.2);
      border-radius: 4px;
      width: 46px;
      height: 22px;
      line-height: 22px;
      //display: inline-block;
      padding: 0px 8px;
      font-size: 12px;
      text-align: center;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    //:deep(.ant-input-affix-wrapper-focused) {
    //  border: none !important;
    //  outline: none;
    //}
  }

  .space {
    flex: 1;
  }

  .action {
    width: 200px;
    margin-right: 24px;
    display: flex;
    justify-content: flex-end;
    .action-item{
      cursor: pointer;

    }
  }
}


</style>
