import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import settings from '@/config/settings';
import {changeProject, getByUser} from '@/services/project';
import {setCache} from "@/utils/localCache";

export interface StateType {
  projects: any[]
  currProject: any,
  recentProjects: any[],
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    saveProjects: Mutation<StateType>;
  };
  actions: {
    fetchProject: Action<StateType, StateType>;
    changeProject: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  projects: [],
  recentProjects:[],
  currProject: {},
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'ProjectGlobal',
  state: {
    ...initState
  },
  mutations: {
    saveProjects(state, payload) {
      setCache(settings.currProjectId, payload.currProject.id);
      state.projects = payload.projects;
      state.currProject = payload.currProject;
      state.recentProjects = payload.recentProjects;
    },
  },
  actions: {
    async fetchProject({ commit }, currProjectId) {
      try {
        const response: ResponseData = await getByUser(currProjectId);
        const { data } = response;
        data.recentProjects = response?.data?.projects || [];
        commit('saveProjects', data || 0);

        return true;
      } catch (error) {
        return false;
      }
    },

    async changeProject({ commit }, projectId) {
      try {
        await changeProject(projectId);

        const response: ResponseData = await getByUser(projectId);
        const { data } = response;
        data.recentProjects = response?.data?.projects || [];
        console.log('data', data);
        commit('saveProjects', data || 0);

        return true;
      } catch (error) {
        return false;
      }
    },
  },


}

export default StoreModel;
