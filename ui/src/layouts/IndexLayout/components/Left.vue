<template>
  <div id="indexlayout-left">
    <div class="indexlayout-left-logo">
      <router-link to="/" class="logo-url">
        <div :class="{
          'logo-title':true,
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
      <div v-if="hasSettingPermission" :class="['settings', isActive ? 'settings-active' : '' ]" @click="handleRedirect">
        <div class="settings-menu">
          <Icon :type="isActive ? 'settings-active' : 'settings'" />
          <span class="left-menu-title">项目设置</span>
        </div>
      </div>
    </div>
    <div v-if="version" class="version">
      V{{ version }}
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent, onMounted, PropType, ref, computed} from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";

import { StateType as GlobalStateType } from "@/store/global";
import { StateType as ProjectGlobalStateType } from "@/store/project";
import {RoutesDataItem} from '@/utils/routes';
import SiderMenu from './SiderMenu.vue';
import Icon from "./Icon.vue";

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
    Icon,
  },
  setup(props) {
    let isLeyanEnv = process.env.VUE_APP_DEPLOY_ENV === 'ly';
    const router = useRouter();
    const store = useStore<{ Global: GlobalStateType, ProjectGlobal: ProjectGlobalStateType }>();
    const permissionRouteMenuMap = computed(() => store.state.Global.permissionMenuMap);
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
    const isActive = computed(() => {
      return router.currentRoute.value.path.includes('project-setting');
    });

    const hasSettingPermission = computed(() => {
      if (permissionRouteMenuMap.value && permissionRouteMenuMap.value['project-setting']) {
        return true;
      }
      return false;
    });

    const handleRedirect = () => {
      router.push(`/${currProject.value.shortName}/project-setting/enviroment/var`);
    };

    return {
      isLeyanEnv,
      isActive,
      handleRedirect,
      hasSettingPermission,
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
        margin: 0px 6px;
        height: 64px;
        line-height: 64px;
        font-family: 'Helvetica', sans-serif;
        font-weight: 400;
        font-style: normal;
        font-size: 18px;
        color: #FFFFFF;
        background-image: url("../../../assets/images/logo.png");
        background-size: 100% 50%;
        background-repeat: no-repeat;
        background-position: center;
        &.leyan-logo{
          width: 100%;
          height: 40px;
          background-image: url("../../../assets/images/leyan-api-logo.png");
          background-size: 100% 100%;
          transform: unset;
        }
      }
    }

  }

  .indexlayout-left-menu {
    flex: 1;
    position: relative;
    overflow: hidden auto;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    .settings {
      width: 100%;
      padding: 6px 4px;
      border-top: 1px solid rgba(255,255,255,.1);
      box-sizing: border-box;

      &.settings-active {
        color: white;

        .settings-menu {
          background-color: #10131E !important;
          border-radius: 8px;
        }
      }

      .settings-menu {
        width: 64px;
        height: 64px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        cursor: pointer;

        .svg-icon {
          width: 24px;
          height: 24px;
        }

        .left-menu-title {
          line-height: 20px;
          font-size: 12px;
          margin-top: 4px;
          margin-left: 0 !important;
        }
      }
    }
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
    width: 64px;
    height: 64px;
    padding: 0 !important;

    .ant-menu-title-content {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      width: 100%;
      height: 100%;

      .svg-icon {
        width: 24px;
        height: 24px;
      }

      .left-menu-title {
        line-height: 20px;
        font-size: 12px;
        margin-top: 4px;
        margin-left: 0 !important;
      }
    }
  }

  :deep(.ant-menu-item.ant-menu-item-selected) {
    background-color: #10131E !important;
    border-radius: 8px;
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
    padding: 0 4px;
  }



}
</style>
