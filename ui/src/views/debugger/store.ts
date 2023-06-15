import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { TestInterface, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    remove,
    clone,
} from './service';

import {
    loadCategory,
    getCategory,
    createCategory,
    updateCategory,
    removeCategory,
    moveCategory,
    updateCategoryName
} from "@/services/category";

import { getNodeMap } from "@/services/tree";

export interface StateType {
    testInterfaceId: number;
    currInterface: any;

    listResult: QueryResult;
    detailResult: any;
    queryParams: any;

    treeData: any[] | null;
    treeDataMap: any,
    nodeDataCategory: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setInterfaceId: Mutation<StateType>;
        setCurrInterface: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;

        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        setTreeDataMapItem: Mutation<StateType>;
        setTreeDataMapItemProp: Mutation<StateType>;
        setNodeCategory: Mutation<StateType>;
    };
    actions: {
        loadTree: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        removeInterface: Action<StateType, StateType>;
        cloneInterface: Action<StateType, StateType>;

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
    }
}

const initState: StateType = {
    testInterfaceId: 0,
    currInterface: null,

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
    detailResult: {} as TestInterface,
    queryParams: {},

    treeData: [],
    treeDataMap: {},
    nodeDataCategory: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'TestInterface',
    state: {
        ...initState
    },
    mutations: {
        setInterfaceId(state, id) {
            state.testInterfaceId = id;
        },
        setCurrInterface(state, payload) {
            state.currInterface = payload;
        },
        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },

        setTreeData(state, data) {
            state.treeData = data
        },
        setTreeDataMap(state, payload) {
            state.treeDataMap = payload
        },
        setTreeDataMapItem(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id] = payload
        },
        setTreeDataMapItemProp(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id][payload.prop] = payload.value
        },
        setNodeCategory(state, data) {
            state.nodeDataCategory = data;
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
        async getInterface({ commit }, id: number) {
            if (id === 0) {
                commit('setDetail', {
                    ...initState.detailResult,
                })
                return
            }
            try {
                const response: ResponseData = await get(id);
                const { data } = response;
                commit('setDetail', {
                    ...initState.detailResult,
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
        // category tree
        async loadCategory({ commit }) {
            const response = await loadCategory('plan');
            if (response.code != 0) return;

            const { data } = response;
            commit('setTreeData', data || {});

            const mp = {}
            getNodeMap(data, mp)

            commit('setTreeDataMap', mp);

            return true;
        },
        async getCategoryNode({ commit }, payload: any) {
            try {
                if (!payload) {
                    commit('setNodeCategory', {});
                    return true;
                }

                const response = await getCategory(payload.id);
                const { data } = response;
                console.log(data);
                commit('setNodeCategory', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createCategoryNode({ commit, dispatch, state }, payload: any) {
            try {
                const resp = await createCategory(payload);

                await dispatch('loadCategory');
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async updateCategoryNode({ commit }, payload: any) {
            try {
                const { id, ...params } = payload;
                await updateCategory(id, { ...params });
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeCategoryNode({ commit, dispatch, state }, payload: number) {
            try {
                await removeCategory(payload);
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveCategoryNode({ commit, dispatch, state }, payload: any) {
            try {
                await moveCategory(payload);
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveTreeMapItemCategory({ commit }, payload: any) {
            commit('setTreeDataMapItem', payload);
        },
        async saveTreeMapItemPropCategory({ commit }, payload: any) {
            commit('setTreeDataMapItemProp', payload);
        },
        async saveCategory({ commit, dispatch, state }, payload: any) {
            const jsn = await updateCategory(payload.id, payload)
            if (jsn.code === 0) {
                commit('setCategory', jsn.data);
                await dispatch('loadCategory');
                return true;
            } else {
                return false
            }
        },
        async updateCategoryName({ commit, dispatch, state }, payload: any) {
            const jsn = await updateCategoryName(payload.id, payload.name)
            if (jsn.code === 0) {
                await dispatch('loadCategory');
                return true;
            } else {
                return false
            }
        }
    }
};

export default StoreModel;
