<template>
  <a-menu
      theme="dark"
      mode="inline"
      :inline-collapsed="collapsed"
      :selectedKeys="selectedKeys"
      :openKeys="openKeys"
      @openChange="(key)=>{openChange(key);}">

    <sider-menu-item
        v-for="item in newMenuData"
        :key="item.path"
        :routeItem="item"
        :selectedKeys="selectedKeys"
        :topNavEnable="topNavEnable"
        :belongTopMenu="belongTopMenu">
    </sider-menu-item>
  </a-menu>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, PropType, toRefs, ref, watch } from "vue";
import { useStore } from "vuex";
import SiderMenuItem from './SiderMenuItem.vue';
import { RoutesDataItem } from '@/utils/routes';
import { StateType as GlobalStateType } from "@/store/global";
import { RouteMenuType } from "@/types/permission";
import {isLeyan} from "@/utils/comm";

export default defineComponent({
  name: 'SiderMenu',
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
    }
  },
  components: {
    SiderMenuItem
  },
  setup(props) {
    const store = useStore<{ Global: GlobalStateType }>();
    // 后端获取的权限路由可访问列表
    const permissionRouteMenuMap = computed(() => store.state.Global.permissionMenuMap);
    const { menuData, topNavEnable }  = toRefs(props);
    const newMenuData = ref<RoutesDataItem[]>([]);
    const isLeyanEnv = isLeyan();

    const getNewMenuData = () => {
      if(!topNavEnable.value) {
        return menuData.value as RoutesDataItem[];
      }
      const MenuItems: RoutesDataItem[] = [];
      /**
         * 项目设置位置调整到左侧底部  设置按钮
         */
      const sourceMenuData = menuData.value.filter(e => !e.path.includes('project-setting'));
      for (let index = 0, len = sourceMenuData.length; index < len; index += 1) {
        const element = sourceMenuData[index];

        if (element.children) {
          const routeDataItem: RoutesDataItem = { ...element.children[0], children: [] };

          /**
           * 左侧菜单无需展示二级菜单了。
           * 二级菜单将在 一级菜单页面以tab的形式切换
           */
          // const childrenRoute = element.children.length > 1 ? element.children.slice(1) : [];
          // if (childrenRoute.length > 0) {
          //   childrenRoute.forEach((routeItem: RoutesDataItem) => {
          //     if (!routeItem.hidden && permissionRouteMenuMap.value && permissionRouteMenuMap.value[RouteMenuType[`${routeItem.meta?.code}`]]) {
          //       routeDataItem.children?.push(routeItem);
          //     }
          //   })
          // }

          // 根据可访问权限路由表来做匹配可展示的路由menu
          if (permissionRouteMenuMap.value && permissionRouteMenuMap.value[RouteMenuType[`${routeDataItem.meta?.code}`]]) {
            MenuItems.push({ ...routeDataItem });
          }
        }
      }
      newMenuData.value = MenuItems;
      console.log('getNewMenuData---,', newMenuData);
    }

    const openChange = (key: string): void => {
      props.onOpenChange && props.onOpenChange(key);
    }

    watch(() => {return permissionRouteMenuMap.value;}, () => {
      getNewMenuData();
    }, {
      immediate: true
    })

    return {
      newMenuData,
      openChange
    }
  }
})
</script>