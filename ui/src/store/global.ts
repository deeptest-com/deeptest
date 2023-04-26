
import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { TabNavItem } from '@/utils/routes';
import settings from '@/config/settings';
import router from '@/config/routes';
import { getPermissionMenuList } from '@/services/project';

export interface StateType {
  // 左侧展开收起
  collapsed: boolean;
  // 顶部菜单开启
  topNavEnable: boolean;
  // 头部固定开启
  headFixed: boolean;
  // tab菜单开启
  tabNavEnable: boolean;
  // 头部tab导航列表
  headTabNavList: TabNavItem[];

  permissionMenuMap: any[];
  permissionButtonMap: any[];
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    changeLayoutCollapsed: Mutation<StateType>;
    setTopNavEnable: Mutation<StateType>;
    setHeadFixed: Mutation<StateType>;
    setTabNavEnable: Mutation<StateType>;
    setHeadTabNavList: Mutation<StateType>;
    setPermissionMenuAndBtn: Mutation<StateType>;
  };
  actions: {
    getPermissionList: Action<StateType, StateType>;
  };
}

const homeRoute = router.resolve(settings.homeRouteItem.path);

const initState: StateType = {
  collapsed: false,
  topNavEnable: settings.topNavEnable,
  headFixed: settings.headFixed,
  tabNavEnable: settings.tabNavEnable,
  headTabNavList: [
    {
      route: homeRoute,
      menu: settings.homeRouteItem
    }
  ],
  permissionMenuMap: [],
  permissionButtonMap: []
};

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'Global',
  state: {
    ...initState
  },
  mutations: {
    changeLayoutCollapsed(state, payload) {
      state.collapsed = payload;
    },
    setTopNavEnable(state, payload) {
      state.topNavEnable = payload;
    },
    setHeadFixed(state, payload) {
      state.headFixed = payload;
    },
    setTabNavEnable(state, payload) {
      state.tabNavEnable = payload;
    },
    setHeadTabNavList(state, payload) {
      state.headTabNavList = payload;
    },
    setPermissionMenuAndBtn(state, payload) {
      const { permissionButtonMap, permissionMenuMap } = payload;
      state.permissionButtonMap = permissionButtonMap;
      state.permissionMenuMap = permissionMenuMap
    }
  },
  actions: {
    async getPermissionList({ commit }) {
      const result = await getPermissionMenuList();
      if (result.code === 0) {
        const menuData = {};
        const buttonData = {};
        result.data.result.forEach(e => {
          if (e.type === 'menu') {
            menuData[e.code] = e;
          } else if (e.type === 'button') {
            buttonData[e.code] = e;
          }
        })
        console.log('~permissionMenu --', menuData);
        console.log('~permissionButton --', buttonData);
        commit('setPermissionMenuAndBtn', { permissionButtonMap: buttonData, permissionMenuMap: menuData });
      }
    }
  }
}



export default StoreModel;
