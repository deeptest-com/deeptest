import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";

import {
  listDatapool, getDatapool,
  saveDatapool, removeDatapool,
} from "@/services/datapool";

export interface StateType {
  datapoolsData: [],
  datapoolData: [],
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    setDatapools: Mutation<StateType>;
    setDatapool: Mutation<StateType>;
  };
  actions: {
    listDatapool: Action<StateType, StateType>;
    getDatapool: Action<StateType, StateType>;
    saveDatapool: Action<StateType, StateType>;
    removeDatapool: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  datapoolsData: [],
  datapoolData: [],
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'Datapool',
  state: {
    ...initState
  },
  mutations: {
    setDatapools(state, payload) {
      state.datapoolsData = payload;
    },
    setDatapool(state, payload) {
      state.datapoolData = payload;
    },
  },
  actions: {
    async listDatapool({commit}) {
      try {
        const resp = await listDatapool();
        const {data} = resp;
        commit('setDatapools', data);
        return true;
      } catch (error) {
        return false;
      }
    },

    async getDatapool({commit}, payload: any) {
      try {
        const response = await getDatapool(payload.id, payload.projectId);
        const {data} = response;

        commit('setDatapool', data);
        commit('setProjectId', payload.projectId);

        return true;
      } catch (error) {
        return false;
      }
    },

    async saveDatapool({commit, dispatch, state}, payload: any) {
      try {
        const resp = await saveDatapool(payload);

        dispatch('listDatapool');
        dispatch('getDatapool', {id: 0, projectId: payload.projectId})
        return resp.data;
      } catch (error) {
        return false;
      }
    },

    async removeDatapool({commit, dispatch, state}, payload: any) {
      try {
        await removeDatapool(payload.id);

        dispatch('listDatapool', payload.id);
        dispatch('getDatapool', {id: 0, projectId: payload.projectId})
        return true;
      } catch (error) {
        return false;
      }
    },

  },
}

export default StoreModel;
