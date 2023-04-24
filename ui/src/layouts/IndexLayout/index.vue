<template>
  <div id="indexlayout">
    <left
        :collapsed="collapsed"
        :topNavEnable="topNavEnable"
        :belongTopMenu="belongTopMenu"
        :selectedKeys="selectedKeys"
        :openKeys="leftOpenKeys"
        :menuData="permissionMenuData"
        :version="version"
        :onOpenChange="onOpenChange"
    >
    </left>

    <div
        id="indexlayout-right"
        :class="{'fiexd-header': headFixed}"
    >
      <right-top
          :collapsed="collapsed"
          :tabNavEnable="tabNavEnable"
          :topNavEnable="topNavEnable"
          :belongTopMenu="belongTopMenu"
          :breadCrumbs="breadCrumbs"
          :toggleCollapsed="toggleCollapsed"
          :menuData="permissionMenuData"
          :routeItem="routeItem"
      >
      </right-top>
      <div class="indexlayout-right-main">
        <permission :roles="routeItem.roles">
          <router-view></router-view>
        </permission>
        <right-footer></right-footer>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import {computed, defineComponent, nextTick, onMounted, ref, watch} from "vue";
import {useStore} from 'vuex';
import {useRoute} from 'vue-router';
import {StateType as GlobalStateType} from '@/store/global';
import {StateType as UserStateType} from "@/store/user";
import {
  BreadcrumbType,
  formatRoutePathTheParents,
  getBreadcrumbRoutes,
  getPermissionMenuData,
  getRouteBelongTopMenu,
  getRouteItem,
  getSelectLeftMenuPath,
  RoutesDataItem,
  vueRoutes
} from '@/utils/routes';
import {mergeUnique as ArrayMergeUnique} from '@/utils/array';
import useTitle from '@/composables/useTitle';
import IndexLayoutRoutes from './routes';
import Permission from '@/components/Permission/index.vue';
import Left from '@/layouts/IndexLayout/components/Left.vue';
import RightTop from '@/layouts/IndexLayout/components/RightTop.vue';
import RightFooter from '@/layouts/IndexLayout/components/RightFooter.vue';
import {useI18n} from "vue-i18n";

export default defineComponent({
  name: 'IndexLayout',
  components: {
    Permission,
    Left,
    RightTop,
    RightFooter,
  },
  setup() {
    const store = useStore<{ Global: GlobalStateType; User: UserStateType; }>();
    const route = useRoute();

    const version = ref('')

    onMounted(() => {
      console.log('onMounted')

      const isElectron = ref(!!window.require)
      if (isElectron.value) {
        const remote = window.require('@electron/remote')
        version.value = remote.getGlobal('sharedObj').version
      }
    })

    // 所有菜单路由
    const menuData: RoutesDataItem[] = vueRoutes(IndexLayoutRoutes);

    // 当前路由 item
    const routeItem = computed<RoutesDataItem>(() => getRouteItem(route.path, menuData));

    // 有权限的菜单
    const permissionMenuData = computed<RoutesDataItem[]>(() => getPermissionMenuData(store.state.User.currentUser.sysRoles, menuData));

    // 当前路由的顶部菜单path
    const belongTopMenu = computed<string>(() => getRouteBelongTopMenu(routeItem.value))

    // 当前路由的父路由path[]
    const routeParentPaths = computed<string[]>(() => formatRoutePathTheParents(routeItem.value.path));

    // 收缩左侧
    const collapsed = computed<boolean>(() => store.state.Global.collapsed);
    const toggleCollapsed = (): void => {
      store.commit('Global/changeLayoutCollapsed', !collapsed.value);
    }

    // 右侧顶部tabNav是否开启
    const tabNavEnable = computed<boolean>(() => store.state.Global.tabNavEnable);

    // 右侧顶部导航是否开启
    const topNavEnable = computed<boolean>(() => store.state.Global.topNavEnable);

    // 右侧顶部是否固定
    const headFixed = computed<boolean>(() => store.state.Global.headFixed);

    // 左侧选择菜单key
    const selectedKeys = computed<string[]>(() => {
      const selectedKeys = getSelectLeftMenuPath(routeItem.value);
      return selectedKeys
    });

    // 左侧展开菜单keys
    const leftOpenKeys = ref<string[]>(routeParentPaths.value);
    watch([routeParentPaths, collapsed], () => {
      if (routeParentPaths.value.length && !collapsed.value) {
        leftOpenKeys.value = ArrayMergeUnique<string>(leftOpenKeys.value, routeParentPaths.value);
      } else {
        nextTick(() => {
          leftOpenKeys.value = [];
        })
      }
    })

    // 左侧菜单展开收起
    const onOpenChange = (key: any) => {
      leftOpenKeys.value = key;
    }

    // 面包屑导航
    const breadCrumbs = computed<BreadcrumbType[]>(
        () => getBreadcrumbRoutes(routeItem.value, routeParentPaths.value, menuData));

    // 设置title
    useTitle(routeItem);

    return {
      collapsed,
      toggleCollapsed,
      tabNavEnable,
      topNavEnable,
      belongTopMenu,
      headFixed,
      selectedKeys,
      leftOpenKeys,
      breadCrumbs,
      permissionMenuData,
      version,
      onOpenChange,
      routeItem
    }
  }
})
</script>
<style lang="less">
@import '../../assets/css/variables.less';

#indexlayout {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

#indexlayout-right {
  position: relative;
  flex: 1;
  overflow: auto;
  background-color: @mainBgColor;

  &.fiexd-header {
    display: flex;
    flex-direction: column;

    .indexlayout-right-main {
      flex: 1;
      overflow: auto;
      min-width: 1217px;
    }
  }
}

.indexlayout-main-conent {
  margin: 24px;
  position: relative;
}
</style>
