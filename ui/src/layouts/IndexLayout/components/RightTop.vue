<template>
    <div id="indexlayout-right-top" :class="{'topNavEnable': !topNavEnable, 'tabNavEnable': !tabNavEnable }">
        <div class="indexlayout-right-top-top">
          <div class="indexlayout-flexible"
               @click="() => {
                if(toggleCollapsed) {
                  toggleCollapsed();
                }
              }"
          >
            <MenuUnfoldOutlined v-if="collapsed" />
            <MenuFoldOutlined v-else />
          </div>

            <div class="indexlayout-top-menu">
            </div>
            <div class="indexlayout-top-menu-right">
                <right-top-project />
                <RightTopSettings></RightTopSettings>
            </div>
        </div>

    </div>
</template>
<script lang="ts">
import { defineComponent, onMounted, PropType, Ref, toRefs } from "vue";
import { useI18n } from "vue-i18n";
import { BreadcrumbType, RoutesDataItem } from '@/utils/routes';
import { MenuFoldOutlined, MenuUnfoldOutlined, EnvironmentOutlined } from '@ant-design/icons-vue';
import RightTopProject from './RightTopProject.vue';
import RightTopSettings from './RightTopSettings.vue';
import useTopMenuWidth from "../composables/useTopMenuWidth";

interface RightTopSetupData {
  t: (key: string | number) => string;
  topMenuCon: Ref;
  topMenuWidth: Ref;
}

export default defineComponent({
    name: 'RightTop',
    components: {
      MenuFoldOutlined, MenuUnfoldOutlined,
      RightTopProject,
      RightTopSettings,
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
      belongTopMenu: {
        type: String,
        default: ''
      },
      toggleCollapsed: {
        type: Function as PropType<() => void>
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
    setup(props): RightTopSetupData {
      const { t } = useI18n();
      const { topNavEnable } = toRefs(props);

      const { topMenuCon, topMenuWidth } = useTopMenuWidth(topNavEnable);

      return {
        t,
        topMenuCon,
        topMenuWidth
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
  .indexlayout-right-top-top {
    display: flex;
    width: 100%;
    height: @headerHeight;
    background-color: @menu-dark-bg;
    color: #c0c4cc;

    .indexlayout-flexible {
      width: 16px;
      height: 50px;
      line-height: 50px;
      text-align: center;
      cursor: pointer;
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
      /* display: flex; */
      overflow: hidden;
      overflow-x: auto;
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
      width: 238px;

      .indexlayout-top-project {
        float: left;
        padding: 10px 0;
      }

      .indexlayout-top-settings {
        float: left;
        padding: 15px 0 15px 20px;
        color: #c0c4cc;

        .indexlayout-top-usermenu {
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