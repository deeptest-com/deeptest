<template>
  <a-menu
      theme="dark"
      mode="inline"
      :inline-collapsed="collapsed"
      :selectedKeys="selectedKeys"
      :openKeys="openKeys"
      @openChange="(key)=>{
          openChange(key);
        }"
  >
    <sider-menu-item
        v-for="item in newMenuData"
        :key="item.path"
        :routeItem="item"
        :topNavEnable="topNavEnable"
        :belongTopMenu="belongTopMenu"
    >
    </sider-menu-item>

  </a-menu>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, PropType, toRefs } from "vue";
import { useStore } from "vuex";
import SiderMenuItem from './SiderMenuItem.vue';
import { RoutesDataItem } from '@/utils/routes';
import { StateType as GlobalStateType } from "@/store/global";
import { RouteMenuType } from "@/types/permission";

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
    const newMenuData = computed<RoutesDataItem[]>(() => {
      if(!topNavEnable.value) {
        return menuData.value as RoutesDataItem[];
      }
      const MenuItems: RoutesDataItem[] = [];
      for (let index = 0, len = menuData.value.length; index < len; index += 1) {
        const element = menuData.value[index];
        if (element.children) {
          const childrenRoute = element.children.length > 1 ? element.children.slice(1) : [];
          const routeDataItem: RoutesDataItem = { ...element.children[0], children: [] };
          if (childrenRoute.length > 0) {
            childrenRoute.forEach((routeItem: RoutesDataItem) => {
              if (!routeItem.hidden && permissionRouteMenuMap.value[RouteMenuType[`${routeItem.meta?.code}`]]) {
                routeDataItem.children?.push(routeItem);
              }
            })
          }
          // 根据可访问权限路由表来做匹配可展示的路由menu
          if (permissionRouteMenuMap.value && permissionRouteMenuMap.value[RouteMenuType[`${routeDataItem.meta?.code}`]]) {
            MenuItems.push({ ...routeDataItem });
          }
        }
      }
      return MenuItems;
    })

    const openChange = (key: string): void => {
      props.onOpenChange && props.onOpenChange(key);
    }


    return {
      newMenuData,
      openChange
    }
  }
})
</script>