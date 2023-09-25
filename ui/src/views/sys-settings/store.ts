import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import {notifyError, notifySuccess} from "@/utils/notify";
import {listJsLib} from './service';

export interface StateType {
    jsLibs: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setJsLibs: Mutation<StateType>,
    };
    actions: {
        listJsLib: Action<StateType, StateType>,
    }
}

const initState: StateType = {
    jsLibs: [],
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'SysSetting',
    state: {
        ...initState
    },
    mutations: {
        setJsLibs(state, payload) {
            state.jsLibs = payload;
        },
    },
    actions: {
        async listJsLib({ commit }) {
           const res = await listJsLib()
            if (res.code === 0) {
                commit('setJsLibs', res.data);
                return true;
            } else {
                return false;
            }
        },
    }
};

export default StoreModel;
