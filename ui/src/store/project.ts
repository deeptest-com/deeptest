import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import settings from '@/config/settings';
import {getByUser} from '@/services/project';
import {setCache} from "@/utils/localCache";

export interface StateType {
  projects: any[]
  currProject: any
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    saveProjects: Mutation<StateType>;
  };
  actions: {
    fetchProject: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  projects: [],
  currProject: {},
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'ProjectData',
  state: {
    ...initState
  },
  mutations: {
    saveProjects(state, payload) {
      console.log('payload', payload)

      setCache(settings.currProjectId, payload.currProject.id);

      state.projects = payload.projects;
      state.currProject = payload.currProject;
    },
  },
  actions: {
    async fetchProject({ commit }, currProjectId) {
      try {
        const response: ResponseData = await getByUser(currProjectId);
        const { data } = response;
        commit('saveProjects', data || 0);

        return true;
      } catch (error) {
        return false;
      }
    },
  }
}

export default StoreModel;
  