import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Interface, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    remove,
    loadExecResult,
} from './service';

import {
    loadCategory,
    getCategory,
    createCategory,
    updateCategory,
    removeCategory,
    moveCategory,
    updateCategoryName} from "@/services/category";

import {getNodeMap} from "@/services/tree";

export interface StateType {
    interfaceId: number;

    listResult: QueryResult;
    detailResult: Interface;
    queryParams: any;

    execResult: any;

    treeDataCategory: any[];
    treeDataMapCategory: any,
    nodeDataCategory: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setInterfaceId: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;

        setExecResult: Mutation<StateType>;

        setTreeDataCategory: Mutation<StateType>;
        setTreeDataMapCategory: Mutation<StateType>;
        setTreeDataMapItemCategory: Mutation<StateType>;
        setTreeDataMapItemPropCategory: Mutation<StateType>;
        setNodeCategory: Mutation<StateType>;
    };
    actions: {
        listInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        removeInterface: Action<StateType, StateType>;

        loadCategory: Action<StateType, StateType>;
        getCategoryNode: Action<StateType, StateType>;
        createCategoryNode: Action<StateType, StateType>;
        updateCategoryNode: Action<StateType, StateType>;
        removeCategoryNode: Action<StateType, StateType>;
        moveCategoryNode: Action<StateType, StateType>;
        saveTreeMapItemCategory: Action<StateType, StateType>;
        saveTreeMapItemPropCategory: Action<StateType, StateType>;
        saveCategory: Action<StateType, StateType>;
        updateCategoryName: Action<StateType, StateType>;

        loadExecResult: Action<StateType, StateType>;
        updateExecResult: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    interfaceId: 0,

    listResult: {
        list: [],
        pagination: {
            total: 0,
            current: 1,
            pageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
        },
    },
    detailResult: {} as Interface,
    queryParams: {},
    execResult: {},

    treeDataCategory: [],
    treeDataMapCategory: {},
    nodeDataCategory: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Interface',
    state: {
        ...initState
    },
    mutations: {
        setInterfaceId(state, id) {
            state.interfaceId = id;
        },

        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },

        setExecResult(state, data) {
            state.execResult = data;
        },

        setTreeDataCategory(state, data) {
            state.treeDataCategory = [data];
        },
        setTreeDataMapCategory(state, payload) {
            state.treeDataMapCategory = payload
        },
        setTreeDataMapItemCategory(state, payload) {
            if (!state.treeDataMapCategory[payload.id]) return
            state.treeDataMapCategory[payload.id] = payload
        },
        setTreeDataMapItemPropCategory(state, payload) {
            if (!state.treeDataMapCategory[payload.id]) return
            state.treeDataMapCategory[payload.id][payload.prop] = payload.value
        },
        setNodeCategory(state, data) {
            state.nodeDataCategory = data;
        },

        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
    },
    actions: {
        async listInterface({commit, dispatch}, params: QueryParams) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;

                commit('setList', {
                    ...initState.listResult,
                    list: data.result || [],
                    pagination: {
                        ...initState.listResult.pagination,
                        current: params.page,
                        pageSize: params.pageSize,
                        total: data.total || 0,
                    },
                });
                commit('setQueryParams', params);

                return true;
            } catch (error) {
                return false;
            }
        },
        async getInterface({commit}, id: number) {
            if (id === 0) {
                commit('setDetail', {
                    ...initState.detailResult,
                })
                return
            }
            try {
                const response: ResponseData = await get(id);
                const {data} = response;
                commit('setDetail', {
                    ...initState.detailResult,
                    ...data,
                });
                return true;
            } catch (error) {
                return false;
            }
        },

        async saveInterface({commit}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async removeInterface({commit, dispatch, state}, payload: number) {
            try {
                await remove(payload);
                await dispatch('listInterface', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },

        async loadExecResult({commit, dispatch, state}, scenarioId) {
            const response = await loadExecResult(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setExecResult', data || {});
            commit('setInterfaceId', scenarioId);

            return true;
        },
        async updateExecResult({commit, dispatch, state}, payload) {
            commit('setExecResult', payload);
            commit('setInterfaceId', payload.scenarioId);

            return true;
        },

        // category tree
        async loadCategory({commit}) {
            const response = await loadCategory('interface');
            if (response.code != 0) return;

            const {data} = response;
            commit('setTreeDataCategory', data || {});

            const mp = {}
            getNodeMap(data, mp)

            commit('setTreeDataMapCategory', mp);

            return true;
        },
        async getCategoryNode({commit}, payload: any) {
            try {
                if (!payload) {
                    commit('setNodeCategory', {});
                    return true;
                }

                const response = await getCategory(payload.id);
                const {data} = response;

                commit('setNodeCategory', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createCategoryNode({commit, dispatch, state}, payload: any) {
            try {
                const resp = await createCategory(payload);

                await dispatch('loadCategory');
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async updateCategoryNode({commit}, payload: any) {
            try {
                const {id, ...params} = payload;
                await updateCategory(id, {...params});
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeCategoryNode({commit, dispatch, state}, payload: number) {
            try {
                await removeCategory(payload);
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveCategoryNode({commit, dispatch, state}, payload: any) {
            try {
                await moveCategory(payload);
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveTreeMapItemCategory({commit}, payload: any) {
            commit('setTreeDataMapItemCategory', payload);
        },
        async saveTreeMapItemPropCategory({commit}, payload: any) {
            commit('setTreeDataMapItemPropCategory', payload);
        },
        async saveCategory({commit, dispatch, state}, payload: any) {
            const jsn = await updateCategory(payload.id, payload)
            if (jsn.code === 0) {
                commit('setCategory', jsn.data);
                await dispatch('loadCategory');
                return true;
            } else {
                return false
            }
        },
        async updateCategoryName({commit, dispatch, state}, payload: any) {
            const jsn = await updateCategoryName(payload.id, payload.name)
            if (jsn.code === 0) {
                await dispatch('loadCategory');
                return true;
            } else {
                return false
            }
        },
    }
};

export default StoreModel;
