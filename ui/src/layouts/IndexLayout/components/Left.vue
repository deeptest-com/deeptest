<template>
  <div id="indexlayout-left" :class="{'narrow': collapsed}">
    <div class="indexlayout-left-logo">
      <router-link to="/" class="logo-url">
        <div :class="{
          'logo-title':true,
          'logo-title-collapsed':collapsed,
          'leyan-logo':isLeyanEnv}"/>
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
    <div v-if="version" class="version">
      V{{ version }}
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent, onMounted, PropType, ref} from "vue";
import {RoutesDataItem} from '@/utils/routes';
import SiderMenu from './SiderMenu.vue';

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
  setup(props) {
    let isLeyanEnv = process.env.VUE_APP_DEPLOY_ENV === 'ly';
    return {
      isLeyanEnv
    };
  }

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
        background-image: url("../../../assets/images/logo.png");
        background-size: 105px 35px;
        background-repeat: no-repeat;
        background-position: center;
        &.leyan-logo{
          transform: scale(1.1) translateX(-16px);
          background-image: url("https://od-1310531898.cos.ap-beijing.myqcloud.com/202306291016448.svg");
        }
      }

      .logo-title-collapsed {
        background-image: url("../../../assets/images/logo-mini.png");
        background-size: 22px 22px;
        background-repeat: no-repeat;
        background-position: center;
        &.leyan-logo{
          background-size: 24px 24px;
          background-image: url("https://od-1310531898.cos.ap-beijing.myqcloud.com/202306291016780.svg");
          transform: scale(1) translateX(-2px);
        }
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

    :deep(.ant-menu-submenu.ant-menu-submenu-inline.ant-menu-submenu-open) {
      color: rgba(255, 255, 255, 0.4);
    }
    :deep(.ant-menu-submenu.ant-menu-submenu-inline.ant-menu-submenu-open .ant-menu-submenu-arrow) {
      opacity: 0.45;
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
