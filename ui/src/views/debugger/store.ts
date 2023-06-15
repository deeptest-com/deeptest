import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { TestInterface, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    remove,
    move,
    clone,
} from './service';

export interface StateType {
    interfaceId: number;
    interfaceData: any;

    queryParams: any;

    treeData: any[] | null;
    treeDataMap: any,
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setInterfaceId: Mutation<StateType>;
        setInterfaceData: Mutation<StateType>;

        setQueryParams: Mutation<StateType>;

        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        changeTreeDataMapItem: Mutation<StateType>;
        changeTreeDataMapItemProp: Mutation<StateType>;
    };
    actions: {
        loadTree: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        removeInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;
        cloneInterface: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    interfaceId: 0,
    interfaceData: null,
    queryParams: {},

    treeData: [],
    treeDataMap: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'TestInterface',
    state: {
        ...initState
    },
    mutations: {
        setInterfaceId(state, id) {
            state.interfaceId = id;
        },
        setInterfaceData(state, payload) {
            state.interfaceData = payload;
        },

        setTreeData(state, data) {
            state.treeData = data
        },
        setTreeDataMap(state, payload) {
            state.treeDataMap = payload
        },
        changeTreeDataMapItem(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id] = payload
        },
        changeTreeDataMapItemProp(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id][payload.prop] = payload.value
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
    },
    actions: {
        async loadTree({ commit, state, dispatch }, params: any) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                commit('setQueryParams', params);
                commit('setTreeData', response.data);

                return true;
            } catch (error) {
                return false;
            }
        },
        async getInterface({ commit }, node: any) {
            if (node.type !== 'interface') {
                commit('setInterfaceData', null)
                return
            }

            try {
                const resp: ResponseData = await get(node.id);
                const { data } = resp;
                commit('setInterfaceData', {
                    ...data,
                });
                return true;
            } catch (error) {
                return false;
            }
        },

        async saveInterface({ state, dispatch }, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                dispatch('loadTree', state.queryParams);
                return true;
            } else {
                return false
            }
        },
        async removeInterface({ commit, dispatch, state }, payload: any) {
            try {
                const jsn = await remove(payload.id, payload.type);
                if (jsn.code === 0) {
                    dispatch('loadTree', state.queryParams);
                    return true;
                }
                return false;
            } catch (error) {
                return false;
            }
        },
        async moveInterface({commit, dispatch, state}, payload: any) {
            try {
                await move(payload);
                dispatch('loadTree', state.queryParams);
                return true;
            } catch (error) {
                return false;
            }
        },
        async cloneInterface({ dispatch, state }, payload: number) {
            try {
                const jsn = await clone(payload);
                if (jsn.code === 0) {
                    dispatch('listInterface', state.queryParams);
                    return true;
                }
                return false;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
