import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";

import {
  changeEnvironment, clearEnvironmentVar,
  copyEnvironment,
  getEnvironment, listEnvironment,
  removeEnvironment, removeEnvironmentVar,
  saveEnvironment, saveEnvironmentVar
} from "@/views/interface/service";

export interface StateType {
  environmentsData: [],
  environmentData: [],
  projectId: number,
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    setEnvironments: Mutation<StateType>;
    setEnvironment: Mutation<StateType>;
    setProjectId: Mutation<StateType>;
  };
  actions: {
    listEnvironment: Action<StateType, StateType>;
    getEnvironment: Action<StateType, StateType>;
    changeEnvironment: Action<StateType, StateType>;
    saveEnvironment: Action<StateType, StateType>;
    copyEnvironment: Action<StateType, StateType>;
    removeEnvironment: Action<StateType, StateType>;

    saveEnvironmentVar: Action<StateType, StateType>;
    removeEnvironmentVar: Action<StateType, StateType>;
    clearEnvironmentVar: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  environmentsData: [],
  environmentData: [],
  projectId: -1,
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'EnvironmentData',
  state: {
    ...initState
  },
  mutations: {
    setEnvironments(state, payload) {
      state.environmentsData = payload;
    },
    setEnvironment(state, payload) {
      state.environmentData = payload;
    },
    setProjectId(state, payload) {
      state.projectId = payload;
    },
  },
  actions: {
    async listEnvironment({commit}) {
      try {
        const resp = await listEnvironment();
        const {data} = resp;
        commit('setEnvironments', data);
        return true;
      } catch (error) {
        return false;
      }
    },

    async getEnvironment({commit}, payload: any) {
      try {
        const response = await getEnvironment(payload.id, payload.projectId);
        const {data} = response;

        commit('setEnvironment', data);
        commit('setProjectId', payload.projectId);

        return true;
      } catch (error) {
        return false;
      }
    },
    async saveEnvironment({commit, dispatch, state}, payload: any) {
      try {
        const resp = await saveEnvironment(payload);

        dispatch('listEnvironment');
        dispatch('getEnvironment', {id: 0, projectId: payload.projectId})
        return resp.data;
      } catch (error) {
        return false;
      }
    },
    async copyEnvironment({commit, dispatch, state}, payload: any) {
      try {
        const resp = await copyEnvironment(payload.id);

        dispatch('listEnvironment');
        dispatch('getEnvironment', {id: 0, projectId: payload.projectId})
        return resp.data;
      } catch (error) {
        return false;
      }
    },
    async removeEnvironment({commit, dispatch, state}, payload: any) {
      try {
        await removeEnvironment(payload.id);

        dispatch('listEnvironment', payload.id);
        dispatch('getEnvironment', {id: 0, projectId: payload.projectId})
        return true;
      } catch (error) {
        return false;
      }
    },
    async changeEnvironment({commit, dispatch, state}, payload: any) {
      await changeEnvironment(payload.id, payload.projectId);

      dispatch('listEnvironment');
      dispatch('getEnvironment', {id: 0, projectId: payload.projectId})
      return true
    },

    // environment var
    async saveEnvironmentVar({commit}, payload: any) {
      try {
        const resp = await saveEnvironmentVar(payload);
        const {data} = resp;
        commit('setEnvironment', data);
        return true;
      } catch (error) {
        return false;
      }
    },
    async removeEnvironmentVar({commit, dispatch, state}, id: number) {
      try {
        const resp = await removeEnvironmentVar(id);
        const {data} = resp;
        dispatch('getEnvironment', {id: 0, projectId: state.projectId})

        return true;
      } catch (error) {
        return false;
      }
    },
    async clearEnvironmentVar({commit, dispatch, state}, environmentId: number) {
      try {
        const resp = await clearEnvironmentVar(environmentId);
        const {data} = resp;
        commit('setEnvironment', data);

        return true;
      } catch (error) {
        return false;
      }
    },
  },

}

export default StoreModel;
