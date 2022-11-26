import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {getProfile, queryMessage, queryProject, updateEmail, updateName, updatePassword} from "@/services/user";
import { removeToken } from "@/utils/localToken";

export interface CurrentUser {
  id: number;
  username: string;
  email: string;
  avatar: string;
  sysRoles: string[];
  projectRoles: any;
}

export interface StateType {
  currentUser: CurrentUser;
  message: number;
}

export interface ModuleType extends StoreModuleType<StateType> {
  state: StateType;
  mutations: {
    saveCurrentUser: Mutation<StateType>;
    saveMessage: Mutation<StateType>;
  };
  actions: {
    fetchCurrent: Action<StateType, StateType>;
    fetchMessage: Action<StateType, StateType>;
    logout: Action<StateType, StateType>;

    updateEmail: Action<StateType, StateType>;
    updateName: Action<StateType, StateType>;
    updatePassword: Action<StateType, StateType>;
  };
}

const initState: StateType = {
  currentUser: {
    id: 0,
    username: '',
    email: '',
    avatar: '',
    sysRoles: [],
    projectRoles: {},
  },
  message: 0,
}

const StoreModel: ModuleType = {
  namespaced: true,
  name: 'User',
  state: {
    ...initState
  },
  mutations: {
    saveCurrentUser(state, payload) {
      state.currentUser = {
        ...initState.currentUser,
        ...payload,
      }
    },
    saveMessage(state, payload) {
      state.message = payload;
    },
  },
  actions: {
    async fetchCurrent({ commit }) {
      try {
        const response: ResponseData = await getProfile();

        const { data } = response;
        commit('saveCurrentUser', data || {});
        return true;
      } catch (error) {
        return false;
      }
    },

    async fetchMessage({ commit }) {
      try {
        const response: ResponseData = await queryMessage();
        const { data } = response;
        commit('saveMessage', data || 0);
        return true;
      } catch (error) {
        return false;
      }
    },
    async logout({ commit }) {
      try {
        await removeToken();
        commit('saveCurrentUser', { ...initState.currentUser });
        return true;
      } catch (error) {
        return false;
      }
    },

    async updateEmail({ commit }, payload) {
      try {
        const json = await updateEmail(payload);
        if (json.code === 0) {
          commit('saveCurrentUser', json.data);
          return true;
        } else {
          return false
        }
      } catch (error) {
        return false;
      }
    },
    async updateName({ commit }, payload) {
      try {
        const json = await updateName(payload);
        if (json.code === 0) {
          commit('saveCurrentUser', json.data);
          return true;
        } else {
          return false
        }
      } catch (error) {
        return false;
      }
    },
    async updatePassword({ commit }, payload) {
      try {
        const json = await updatePassword(payload);
        if (json.code === 0) {
          commit('saveCurrentUser', json.data);
          return true;
        } else {
          return false
        }
        return true;
      } catch (error) {
        return false;
      }
    }
  }
}



export default StoreModel;
