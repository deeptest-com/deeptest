import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import settings from '@/config/settings';
import {changeServe, listServe} from '@/services/serve';
import {setCache} from "@/utils/localCache";

export interface StateType {
  serves: any[]
  currServe: any
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    saveServes: Mutation<StateType>;
  };
  actions: {
    fetchServe: Action<StateType, StateType>;
    changeServe: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  serves: [],
  currServe: {},
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'ServeGlobal',
  state: {
    ...initState
  },
  mutations: {
    saveServes(state, payload) {
      console.log('~~~~~~~ getServeData From Server ~~~~~', payload)
      // 这里暂时取消indexDB对当前选中服务的存储.避免切换项目时查询列表受到该参数影响
      // setCache(settings.currServeId, currServe.id || 0);
      state.serves = payload.serves;
      state.currServe = payload.currServe;
    },
  },
  actions: {
    async fetchServe({ commit }) {
      try {
        const response: ResponseData = await listServe();
        const { data } = response;
        let currServe = {};
        if (data && data.serves && data.serves.length > 0) {
          currServe = data.serves[0];
        }
        const payload = { currServe, serves: (data && data.serves) || [] };
        commit('saveServes', payload);

        return true;
      } catch (error) {
        return false;
      }
    },

    async changeServe({ commit }, serveId) {
      try {
        await changeServe(serveId);

        const response: ResponseData = await listServe();
        const { data } = response;
        commit('saveServes', data || 0);

        return true;
      } catch (error) {
        return false;
      }
    },
  },


}

export default StoreModel;
