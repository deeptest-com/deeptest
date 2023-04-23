<template>
  <div id="indexlayout-left" :class="{'narrow': collapsed}">
    <div class="indexlayout-left-logo">
      <router-link to="/" class="logo-url">
        <div :class="{'logo-title':true,'logo-title-collapsed':collapsed}"/>
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
      <CollapsedIcon @click="toggleCollapsed" :collapsed="collapsed"/>
    </div>

    <div v-if="version" class="version">
      V{{ version }}
    </div>

    <!--    <div class="indexlayout-left-menu-bottom">
          <SettingsMenu></SettingsMenu>
        </div>-->

  </div>
</template>

<script lang="ts">

import {defineComponent, onMounted, PropType, ref} from "vue";
import {RoutesDataItem} from '@/utils/routes';
import SiderMenu from './SiderMenu.vue';
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"
import SettingsMenu from './SettingsMenu.vue'
import {useI18n} from "vue-i18n";
import {
  CaretRightOutlined,
  CaretLeftOutlined
} from '@ant-design/icons-vue';

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
    toggleCollapsed: {
      type: Function as PropType<() => void>
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
    CollapsedIcon
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
  // padding: 10px;

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
        //display: inline-block;
        //margin: 0;
        //font-size: 16px;
        //font-family: Roboto, sans-serif;
        //color: #c0c4cc;
        height: 64px;
        line-height: 64px;
        font-family: 'Helvetica', sans-serif;
        font-weight: 400;
        font-style: normal;
        font-size: 18px;
        color: #FFFFFF;
        margin-bottom: 0px;
        background-image: url("../../../assets/images/logo-deep.png");
        background-size: 105px 35px;
        background-repeat: no-repeat;
        background-position: center;
      }

      .logo-title-collapsed {
        background-image: url("../../../assets/images/logo-deepmini.png");
        background-size: 22px 22px;
        background-repeat: no-repeat;
        background-position: center;
      }
    }

    img {
      vertical-align: middle;
    }
  }

  .indexlayout-left-menu {
    flex: 1;
    position: relative;
    overflow: hidden auto;
    // height: calc(100vh);
    .left-scrollbar {
      width: 100%;
      height: 100%;
    }
    .expand-icon{
      position: fixed;
      width: 12px;
      height: 56px;
      top:115px;
      opacity: 1;
      display: flex;
      align-items: center;
      justify-content: center;
      left: 223px;
      z-index: 9;
      cursor: pointer;
      background: rgba(255, 255, 255, 0.9);
      border-radius: 0px 8px 8px 0px;
      //transition: opacity 5s;
      &.collapsed{
        left: 54px;
      }
      .icon{
        font-size: 12px;
        color: #8A8A8A;;
      }
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

  :deep(.ant-menu-item) {
    margin-top: 0;
    margin-bottom: 8px;
  }

  :deep(.ant-menu-item.ant-menu-item-selected) {
    background-color: #2E3762 !important;
    border-radius: 4px;
  }

  :deep(.ant-menu-item.ant-menu-item-selected .svg-icon),
  :deep(.ant-menu-item.ant-menu-item-selected .left-menu-title) {
    color: #fff;
  }

  :deep(.ant-menu-item .svg-icon) {
    color: rgba(255, 255, 255, 0.4);
  }

  :deep(.ant-menu-dark .ant-menu-inline.ant-menu-sub) {
    background-color: @menu-dark-bg;
  }


  :deep(.ant-menu-submenu.ant-menu-submenu-open.ant-menu-submenu-selected .ant-menu-submenu-title .left-menu-title) {
      color: rgba(255, 255, 255, 0.4);
  }

  :deep(.indexlayout-left-menu .ant-menu) {
    padding: 10px;
  }
}
</style>
