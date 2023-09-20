<template>
  <div id="indexlayout-right-top" :class="{'topNavEnable': !topNavEnable, 'tabNavEnable': !tabNavEnable }">
    <div class="indexlayout-right-top-top">
      <div class="indexlayout-top-menu">
        <RightTopProject/>
        <RightTopServer v-if="showServerSelector" />
      </div>
      <div class="indexlayout-top-menu-right">
        <RightTopSettings/>
      </div>
    </div>
    <RightTopWebsocket/>
    <RightTopUpdate />
  </div>
</template>
<script lang="ts">
import {defineComponent, PropType, toRefs, watch, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import {BreadcrumbType, RoutesDataItem} from '@/utils/routes';
import RightTopProject from './RightTopProject.vue';
import RightTopServer from "./RightTopServer.vue";
import RightTopSettings from './RightTopSettings.vue';
import RightTopWebsocket from './RightTopWebsocket.vue';
import RightTopUpdate from './RightTopUpdate.vue';
import useTopMenuWidth from "../composables/useTopMenuWidth";

export default defineComponent({
  name: 'RightTop',
  components: {
    RightTopProject,
    RightTopSettings,
    RightTopWebsocket,
    RightTopUpdate,
    RightTopServer
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false
    },
    tabNavEnable: {
      type: Boolean,
      default: true
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
    breadCrumbs: {
      type: Array as PropType<BreadcrumbType[]>,
      default: () => {
        return [];
      }
    },
    menuData: {
      type: Array as PropType<RoutesDataItem[]>,
      default: () => {
        return [];
      }
    },
    routeItem: {
      type: Object as PropType<RoutesDataItem>,
      required: true
    }
  },
  setup(props) {
    const {t} = useI18n();
    const {topNavEnable} = toRefs(props);
    const showServerSelector = ref(false);
    const router = useRouter();

    const {topMenuCon, topMenuWidth} = useTopMenuWidth(topNavEnable);

    watch(() => {return router.currentRoute.value.path;}, (val: string) => {
      showServerSelector.value = val.includes('IM');
    }, {immediate: true})

    return {
      t,
      topMenuCon,
      topMenuWidth,
      showServerSelector
    }
  }
})
</script>
<style lang="less">
@import '../../../assets/css/global.less';

#indexlayout-right-top {
  width: 100%;
  height: (@headerHeight + @headerBreadcrumbHeight + @headerTabNavHeight);
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  z-index: 9;
  min-width: 1217px;

  .indexlayout-right-top-top {
    display: flex;
    width: 100%;
    height: @headerHeight;
    background-color: @menu-white-bg;
    color: #c0c4cc;
    .indexlayout-flexible {
      width: 18px;
      //height: 50px;
      //line-height: 50px;
      text-align: center;
      cursor: pointer;
      display: flex;
      align-items: center;
      margin-left: 8px;

    }

    .indexlayout-left-logo {
      width: 150px;
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

    .indexlayout-top-menu {
      height: @headerHeight;
      line-height: @headerHeight;
      flex: 1;
      display: flex;
      align-items: center;
      overflow: hidden;
      overflow-x: auto;
      width: 280px;
      padding-left: 16px;

      .indexlayout-top-menu-li {
        display: inline-block;
        padding: 0 15px;
        height: @headerHeight;
        text-decoration: none;
        color: #c0c4cc;
        font-size: 15px;
        //border-bottom: solid 3px transparent;
        &:hover,
        &.active {
          background-color: @menu-dark-bg;
          color: @menu-dark-highlight-color;
          //border-bottom-color: @primary-color;
        }
      }

      .breadcrumb {
        line-height: @headerHeight;
      }
    }

    .indexlayout-top-menu-right {
      width: 380px;
      height: @headerHeight;
      line-height: @headerHeight;
      .indexlayout-top-project {
        float: left;
        padding: 10px 10px;
      }
      .indexlayout-top-settings {
        //float: right;
        //padding: 15px 0 15px 16px;
        //width: 100px;
        //color: #c0c4cc;
        color: #FFFFFF;
        display: flex;
        align-items: center;
        justify-content: flex-end;
        .msgs {
          //float: left;
          width: 40px;
        }
        .indexlayout-top-usermenu {
          //float: left;
          color: #c0c4cc;
        }
      }
    }

    .scrollbar();

  }

  .indexlayout-right-top-bot {
    display: flex;
    width: 100%;
    height: @headerBreadcrumbHeight;
    background-color: @mainBgColor;

    .indexlayout-right-top-bot-home {
      width: @headerBreadcrumbHeight;
      height: @headerBreadcrumbHeight;
      line-height: @headerBreadcrumbHeight;
      text-align: center;
    }

    .breadcrumb {
      line-height: @headerBreadcrumbHeight;
      margin-left: 10px;
    }
  }

  &.tabNavEnable {
    height: (@headerHeight) // + @headerBreadcrumbHeight);
  }

  &.topNavEnable {
    height: (@headerHeight + @headerTabNavHeight);

    .indexlayout-right-top-top {
      background-color: #ffffff;
      color: @text-color;

      .indexlayout-flexible {
        &:hover {
          background-color: @mainBgColor;
          color: @heading-color;
        }
      }

      .indexlayout-top-menu-right {
        .indexlayout-top-message {
          color: @heading-color;
        }

        .indexlayout-top-usermenu {
          color: @heading-color;
        }
      }
    }

    &.tabNavEnable {
      height: (@headerHeight);
    }
  }
}
</style>
