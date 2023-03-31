import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';
import {
    Interface,
    QueryResult,
    QueryParams,
    filterFormState
} from './data.d';
import {
    query,
    get,
    save,
    remove,
    loadExecResult, getYaml, updateStatus
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
import {
    copyInterface,
    deleteInterface,
    expireInterface,
    getInterfaceDetail,
    getInterfaceList,
    saveInterface,
} from './service';

import {getNodeMap} from "@/services/tree";
import {momentUtc} from "@/utils/datetime";
import {getEnvList, getSecurityList, serverList} from "@/views/projectSetting/service";

export interface StateType {
    interfaceId: number;
    listResult: QueryResult;
    detailResult: Interface;
    queryParams: any;
    execResult: any;
    treeDataCategory: any[];
    treeDataMapCategory: any,
    nodeDataCategory: any;
    filterState: filterFormState,
    interfaceDetail: any,
    interfaceDetailYamlCode: any,
    serveServers: any[], // serve list
    securityOpts: any[]
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
        setFilterState: Mutation<StateType>;
        setInterfaceDetail: Mutation<StateType>;
        setInterfaceDetailByIndex: Mutation<StateType>;
        setServerList: Mutation<StateType>;
        setSecurityOpts: Mutation<StateType>;
        setYamlCode: Mutation<StateType>;
        setStatus: Mutation<StateType>;
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
        loadList: Action<StateType, StateType>;
        createApi: Action<StateType, StateType>;
        disabled: Action<StateType, StateType>;
        del: Action<StateType, StateType>;
        copy: Action<StateType, StateType>;
        getInterfaceDetail: Action<StateType, StateType>;
        updateInterfaceDetail: Action<StateType, StateType>;
        getServerList: Action<StateType, StateType>;
        getSecurityList: Action<StateType, StateType>;
        getYamlCode: Action<StateType, StateType>;
        updateStatus: Action<StateType, StateType>;
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
    filterState: {
        "status": null,
        "createUser": null,
        "title": null,
    },
    interfaceDetail: null,
    interfaceDetailYamlCode: null,
    serveServers: [],
    securityOpts: [],
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
        setFilterState(state, payload) {
            state.filterState = payload;
        },
        setInterfaceDetail(state, payload) {
            state.interfaceDetail = payload;
        },
        setInterfaceDetailByIndex(state, payload) {
            if (payload.codeIndex === -1 || payload.codeIndex) {
                payload.codeIndex = state.interfaceDetail.interfaces[payload.methodIndex]['responseBodies'].length;
            }
            state.interfaceDetail.interfaces[payload.methodIndex]['responseBodies'][payload.codeIndex] = payload.value;
        },
        setServerList(state, payload) {
            state.serveServers = payload;
        },
        setSecurityOpts(state, payload) {
            state.securityOpts = payload;
        },
        setYamlCode(state, payload) {
            state.interfaceDetailYamlCode = payload;
        },
        setStatus(state, payload) {
            state.listResult.list.forEach((item) => {
                if (item.id === payload.id) {
                    item.status = payload.status;
                }
            });
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
        async loadList({commit, dispatch, state}, {currProjectId, page, pageSize, opts}: any) {
            page = page || state.listResult.pagination.current;
            pageSize = pageSize || state.listResult.pagination.pageSize;
            const otherParams = {...state.filterState, ...opts};

            const res = await getInterfaceList({
                "projectId": currProjectId,
                "page": page,
                "pageSize": pageSize,
                ...otherParams,
            });
            if (res.code === 0) {
                const {result, total} = res.data;
                result.forEach((item, index) => {
                    item.index = index + 1;
                    item.key = `${index + 1}`;
                    item.updatedAt = momentUtc(item.updatedAt);
                })
                commit('setList', {
                    list: result || [],
                    pagination: {
                        ...initState.listResult.pagination,
                        "current": page,
                        "pageSize": pageSize,
                        total: total || 0,
                    },
                });
                commit('setFilterState', {
                    ...otherParams
                });
                return true;
            } else {
                return false
            }
        },
        async createApi({commit, dispatch, state}, params: any) {
            const res = await saveInterface({
                ...params
            });
            if (res.code === 0) {
                await dispatch('loadList', {currProjectId: params.projectId});
            } else {
                return false
            }
        },
        async disabled({commit, dispatch, state}, payload: any) {
            const res = await expireInterface(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        async del({commit, dispatch, state}, payload: any) {
            const res = await deleteInterface(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        async copy({commit, dispatch, state}, payload: any) {
            const res = await copyInterface(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        // 用于新建接口时选择接口分类
        async getInterfaceDetail({commit}, payload: any) {
            const res = await getInterfaceDetail(payload.id);
            res.data.createdAt = momentUtc(res.data.createdAt);
            res.data.updatedAt = momentUtc(res.data.updatedAt);
            if (res.code === 0) {
                commit('setInterfaceDetail', res.data || null);
            } else {
                return false
            }
        },
        // 用于新建接口时选择接口分类
        async updateInterfaceDetail({commit, dispatch}, payload: any) {
            const res = await saveInterface({
                ...payload
            });
            if (res.code === 0) {
                await dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        // 获取项目的服务
        async getServerList({commit}, payload: any) {
            const res = await serverList({
                serveId: payload.id
            });
            if (res.code === 0) {
                res.data.forEach((item: any) => {
                    item.label = item.url;
                    item.value = item.url;
                })
                commit('setServerList', res.data || null);
            } else {
                return false
            }
        },
        async getSecurityList({commit}, payload: any) {
            const res = await getSecurityList({
                serveId: payload.id,
                "page": 1,
                "pageSize": 100
            });
            if (res.code === 0) {
                res.data.result.forEach((item: any) => {
                    item.label = item.name;
                    item.value = item.name;
                })
                commit('setSecurityOpts', res.data.result || []);
            } else {
                return false
            }
        },
        async getYamlCode({commit}, payload: any) {
            const res = await getYaml(payload);
            if (res.code === 0) {
                commit('setYamlCode', res.data);
            } else {
                return false
            }
        },
        async updateStatus({commit}, payload: any) {
            const res = await updateStatus(payload);
            if (res.code === 0) {
                commit('setStatus', payload);
            } else {
                return false
            }
        },

    }
};

export default StoreModel;
