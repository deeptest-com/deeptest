import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {loadSpec} from './service';

export interface StateType {
    mode: string,
    specData: any;
    pathData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setMode: Mutation<StateType>;
        setSpecData: Mutation<StateType>;
        setPathData: Mutation<StateType>;
    };
    actions: {
        setMode: Action<StateType, StateType>;
        loadSpec: Action<StateType, StateType>;
        setPath: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    mode: 'desc',
    specData: {info: {}},
    pathData: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Spec',
    state: {
        ...initState
    },
    mutations: {
        setMode(state, payload) {
            state.mode = payload;
        },
        setSpecData(state, payload) {
            state.specData = payload;
        },
        setPathData(state, payload) {
            state.pathData = payload;
        }
    },
    actions: {
        async loadSpec({commit, dispatch, state}, payload: any) {
            loadSpec(payload).then((json) => {
                if (json.code === 0) {
                    commit('setSpecData', json.data);
                    return true;
                } else {
                    return false
                }
            })
        },
        async setMode({commit, dispatch, state}, payload: any) {
            commit('setMode', payload);
        },
        async setPath({commit, dispatch, state}, payload: any) {
            commit('setPathData', payload);
            commit('setMode', 'path');
        },
    }
};

export default StoreModel;
