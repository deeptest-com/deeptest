<template>
  <div id="indexlayout-left" :class="{'narrow': collapsed}">
    <div class="indexlayout-left-logo">
      <router-link to="/" class="logo-url">
        <h3 v-if="collapsed" class="logo-title">DT</h3>
        <h3 v-else class="logo-title">DeepTest</h3>
      </router-link>
    </div>
    <div class="indexlayout-left-menu">
      <sider-menu
          :collapsed="collapsed"
          :topNavEnable="topNavEnable"
          :belongTopMenu="belongTopMenu"
          :selectedKeys="selectedKeys"
          :openKeys="openKeys"
          :onOpenChange="onOpenChange"
          :menuData="menuData">
      </sider-menu>
    </div>

    <div class="version">
      V{{version}}
    </div>

<!--    <div class="indexlayout-left-menu-bottom">
      <SettingsMenu></SettingsMenu>
    </div>-->

  </div>
</template>

<script lang="ts">

import {defineComponent, onMounted, PropType, ref} from "vue";
import { RoutesDataItem } from '@/utils/routes';
import SiderMenu from './SiderMenu.vue';
import SettingsMenu from './SettingsMenu.vue'
import {useI18n} from "vue-i18n";

export default defineComponent({
  name: 'Left',
  props: {
    collapsed: {
      type: Boolean,
      default: false
    },
    topNavEnable: {
      type: Boolean,
      default: true
    },
    belongTopMenu: {
      type: String,
      default: ''
    },
    selectedKeys: {
      type: Array as PropType<string[]>,
      default: () => {
        return [];
      }
    },
    openKeys: {
      type: Array as PropType<string[]>,
      default: () => {
        return [];
      }
    },
    onOpenChange: {
      type: Function as PropType<(key: any) => void>
    },
    menuData: {
      type: Array as PropType<RoutesDataItem[]>,
      default: () => {
        return [];
      }
    },
    version: {
      type: String
    },
  },
  components: {
    SiderMenu,
  },

})

</script>

<style lang="less" scoped>
@import '../../../assets/css/global.less';
#indexlayout-left {
  display: flex;
  flex-direction: column;
  width: @leftSideBarWidth;
  height: 100vh;
  background-color: @menu-dark-bg;
  color: #c0c4cc;
  transition-duration: 0.1s;
  .indexlayout-left-logo {
    width: 100%;
    height: @headerHeight;
    line-height: @headerHeight;
    text-align: center;
    vertical-align: middle;
    /* background-color: $subMenuBg; */
    .logo-url {
      display: inline-block;
      width: 100%;
      height: 100%;
      overflow: hidden;
      .logo-title {
        display: inline-block;
        margin: 0;
        font-size: 16px;
        font-family: Roboto, sans-serif;
        color: #c0c4cc;
      }
    }
    img {
      vertical-align: middle;
    }
  }

  .indexlayout-left-menu {
    flex: 1;
    overflow: hidden auto;
    // height: calc(100vh);

    .left-scrollbar {
      width: 100%;
      height: 100%;
    }
  }
  .indexlayout-left-menu-bottom {
    position: absolute;
    bottom: 0;
  }

  &.narrow {
    width: @menu-collapsed-width;
  }

  .scrollbar();

  .version {
    position: absolute;
    bottom: 0;
    width: @leftSideBarWidth;
    text-align: center;
  }
}
</style>