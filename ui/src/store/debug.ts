import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';

import {loadData} from "@/services/debug";

export interface StateType {
    endpointId: number;
    interfaceId: number;
    debugData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEndpointId: Mutation<StateType>;
        setInterfaceId: Mutation<StateType>;
        setDebugData: Mutation<StateType>;
    };
    actions: {
        loadDebugData: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    endpointId: 0,
    interfaceId: 0,
    debugData: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Debug',
    state: {
        ...initState
    },
    mutations: {
        setEndpointId(state, id) {
            state.interfaceId = id;
        },
        setInterfaceId(state, id) {
            state.interfaceId = id;
        },
        setDebugData(state, payload) {
            state.debugData = payload;
        }
    },
    actions: {
        async loadDebugData({commit, dispatch}, data) {
            try {
                const resp: ResponseData = await loadData(data);
                if (resp.code != 0) return false;

                commit('setDebugData', resp.data);

                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
