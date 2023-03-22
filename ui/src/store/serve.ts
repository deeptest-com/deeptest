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
      console.log('payload', payload)

      setCache(settings.currServeId, payload.currServe.id);

      state.serves = payload.projects;
      state.currServe = payload.currServe;
    },
  },
  actions: {
    async fetchServe({ commit }) {
      try {
        const response: ResponseData = await listServe();
        const { data } = response;
        commit('saveServes', data || 0);

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
