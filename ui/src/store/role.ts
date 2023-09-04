import {StoreModuleType} from "@/utils/store";
import {Action, Mutation} from "vuex";
import {setCache} from "@/utils/localCache";
import settings from "@/config/settings";
import {ResponseData} from "@/utils/request";
import {getAllSysRoles} from "@/services/role";
import {changeProject, getByUser} from "@/services/project";

export interface StateType {
    roles: any[]
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        saveSysRoles: Mutation<StateType>;
    };
    actions: {
        getAllRoles: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    roles: [],
}

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'SysRole',
    state: {
        ...initState
    },
    mutations: {
        saveSysRoles(state, payload) {
            state.roles = payload.result || [];
        },
    },
    actions: {
        async getAllRoles({ commit }) {
            try {
                const response: ResponseData = await getAllSysRoles();
                const { data } = response;
                commit('saveSysRoles', data || 0);

                return true;
            } catch (error) {
                return false;
            }
        },
    },
}

export default StoreModel;
