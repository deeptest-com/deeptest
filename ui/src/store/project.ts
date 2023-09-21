import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import settings from '@/config/settings';
import {changeProject, checkProjectAndUser, getByUser} from '@/services/project';
import {setCache} from "@/utils/localCache";
import { message } from 'ant-design-vue';

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
    checkProjectAndUser: Action<StateType, StateType>;
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
      state.projects = payload.projects || [];
      state.currProject = payload.currProject;
      state.recentProjects = payload.recentProjects || [];
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

    async changeProject({ commit }, projectId) {
      try {
        await changeProject(projectId);

        const response: ResponseData = await getByUser(projectId);
        const { data } = response;
        commit('saveProjects', data || 0);

        return true;
      } catch (error) {
        return false;
      }
    },

    async checkProjectAndUser({ dispatch }, payload: { project_code: string }) {
      try {
        const result: any = await checkProjectAndUser(payload);
        if (result.code === 0) {
          dispatch('changeProject', result.data.id);
          return result.data;
        }
        if (result.code === 10600) {
          message.error('用户暂无无权限访问，请联系管理员');
          return false;
        }
        if (result.code === 10700) {
          message.error('访问项目异常');
          return false;
        }
        return false;
      } catch (error) {
        return false;
      }
    }
  },


}

export default StoreModel;
