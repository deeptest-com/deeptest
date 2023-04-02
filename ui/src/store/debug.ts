import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';

import {loadData} from "@/services/debug";

export interface StateType {
    currEndpointId: number;
    currInterface: any;
    debugData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEndpointId: Mutation<StateType>;
        setInterface: Mutation<StateType>;
        setDebugData: Mutation<StateType>;
    };
    actions: {
        loadDebugData: Action<StateType, StateType>;
        setEndpointId: Action<StateType, StateType>;
        setInterface: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    currEndpointId: 0,
    currInterface: {},
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
            state.currEndpointId = id;
        },
        setInterface(state, payload) {
            state.currInterface = payload;
        },
        setDebugData(state, payload) {
            state.debugData = payload;
        },
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
        async setEndpointId({commit, dispatch}, id) {
            commit('setEndpointId', id);
        },
        async setInterface({commit, dispatch}, id) {
            commit('setInterface', id);
        },
    }
};

export default StoreModel;
