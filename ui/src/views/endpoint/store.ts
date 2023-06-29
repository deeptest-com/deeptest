import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';
import {
    Endpoint,
    QueryResult,
    QueryParams,
    filterFormState
} from './data.d';
import {
    query,
    get,
    save,
    remove,
    getYaml, updateStatus, getDocs,
    importEndpointData,
    upload
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
    copyEndpoint,
    deleteEndpoint,
    expireEndpoint,
    getEndpointDetail,
    getEndpointList,
    saveEndpoint,
} from './service';

import {getNodeMap} from "@/services/tree";
import {momentUtc} from "@/utils/datetime";
import {
    example2schema,
    schema2example,
    getEnvList,
    getSecurityList,
    serverList,
    getSchemaList, getSchemaDetail
} from "@/views/project-settings/service";

export interface StateType {
    endpointId: number;
    listResult: QueryResult;
    detailResult: Endpoint;
    queryParams: any;
    execResult: any;
    treeDataCategory: any[];
    treeDataMapCategory: any,
    nodeDataCategory: any;
    filterState: filterFormState,
    endpointDetail: any,
    endpointDetailYamlCode: any,
    serveServers: any[], // serve list
    securityOpts: any[]
    interfaceMethodToObjMap: any;
    refsOptions: any;
    selectedMethodDetail: any,
    selectedCodeDetail: any,
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEndpointId: Mutation<StateType>;
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
        clearFilterState: Mutation<StateType>;
        setEndpointDetail: Mutation<StateType>;
        setServerList: Mutation<StateType>;
        setSecurityOpts: Mutation<StateType>;
        setYamlCode: Mutation<StateType>;
        setStatus: Mutation<StateType>;

        setInterfaceMethodToObjMap: Mutation<StateType>;
        clearInterfaceMethodToObjMap: Mutation<StateType>;
        setRefsOptions: Mutation<StateType>;
        setSelectedMethodDetail: Mutation<StateType>;
        setSelectedCodeDetail: Mutation<StateType>;
    };
    actions: {
        listEndpoint: Action<StateType, StateType>;
        getEndpoint: Action<StateType, StateType>;
        saveEndpoint: Action<StateType, StateType>;
        removeEndpoint: Action<StateType, StateType>;
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
        updateExecResult: Action<StateType, StateType>;
        loadList: Action<StateType, StateType>;
        createApi: Action<StateType, StateType>;
        disabled: Action<StateType, StateType>;
        del: Action<StateType, StateType>;
        copy: Action<StateType, StateType>;
        getEndpointDetail: Action<StateType, StateType>;
        updateEndpointDetail: Action<StateType, StateType>;
        getServerList: Action<StateType, StateType>;
        getSecurityList: Action<StateType, StateType>;
        getYamlCode: Action<StateType, StateType>;
        updateStatus: Action<StateType, StateType>;
        example2schema: Action<StateType, StateType>;
        schema2example: Action<StateType, StateType>;
        getRefsOptions: Action<StateType, StateType>;
        getAllRefs: Action<StateType, StateType>;
        getRefDetail: Action<StateType, StateType>;
        getDocs: Action<StateType, StateType>;
        upload: Action<StateType, StateType>;
        importEndpointData: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    endpointId: 0,
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
    detailResult: {} as Endpoint,
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
    endpointDetail: null,
    endpointDetailYamlCode: null,
    serveServers: [],
    securityOpts: [],
    interfaceMethodToObjMap: {},
    refsOptions: {},
    selectedMethodDetail: {},
    selectedCodeDetail: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Endpoint',
    state: {
        ...initState
    },
    mutations: {
        setEndpointId(state, id) {
            state.endpointId = id;
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
            state.filterState.status = payload.status || null;
            state.filterState.createUser = payload.createUser || null;
            state.filterState.title = payload.title || null ;
        },
        clearFilterState(state) {
            state.filterState.status = null;
            state.filterState.createUser = null;
            state.filterState.title =null ;
        },
        setEndpointDetail(state, payload) {
            state.endpointDetail = payload;
        },
        setServerList(state, payload) {
            state.serveServers = payload;
        },
        setSecurityOpts(state, payload) {
            state.securityOpts = payload;
        },
        setYamlCode(state, payload) {
            state.endpointDetailYamlCode = payload;
        },
        setStatus(state, payload) {
            state.listResult.list.forEach((item) => {
                if (item.id === payload.id) {
                    item.status = payload.status;
                }
            });
        },
        setInterfaceMethodToObjMap(state, payload) {
            state.interfaceMethodToObjMap[payload.method] = payload.value;
        },
        clearInterfaceMethodToObjMap(state, payload) {
            state.interfaceMethodToObjMap = {};
        },
        setRefsOptions(state, payload) {
            state.refsOptions[payload.type] = payload.options;
        },
        setSelectedMethodDetail(state, payload) {
            state.selectedMethodDetail = payload;
            // 同步到接口详情
            const interfaces: any = [];
            state.endpointDetail.interfaces.forEach((item) => {
                if (item.method === payload.method) {
                    interfaces.push(payload);
                } else {
                    interfaces.push(item);
                }
            })
            state.endpointDetail.interfaces = [...interfaces];
        },
        setSelectedCodeDetail(state, payload) {
            state.selectedCodeDetail = payload;
            const methodIndex = state.endpointDetail?.interfaces?.findIndex((item) => item.method === state.selectedMethodDetail.method);
            const codeIndex = state.selectedMethodDetail?.responseBodies?.findIndex((item) => item.code === payload?.code);
            // 修改
            if (methodIndex !== -1 && codeIndex !== -1) {
                state.endpointDetail.interfaces[methodIndex]['responseBodies'][codeIndex] = {...payload};
            }
            // 新增
            if (methodIndex !== -1 && codeIndex === -1 && payload?.code) {
                state.endpointDetail.interfaces[methodIndex]['responseBodies'].push({...payload});
            }
        }
    },
    actions: {
        async listEndpoint({commit, dispatch, state}, params: QueryParams) {
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
        async getEndpoint({commit, state}, id: number) {
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
        async saveEndpoint({commit}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async removeEndpoint({commit, dispatch, state}, payload: number) {
            try {
                await remove(payload);
                await dispatch('listEndpoint', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },

        async updateExecResult({commit, dispatch, state}, payload) {
            commit('setExecResult', payload);
            commit('setEndpointId', payload.scenarioId);

            return true;
        },
        // category tree
        async loadCategory({commit}) {
            const response = await loadCategory('endpoint');
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
                const res = await createCategory(payload);
                await dispatch('loadCategory');
                return res;
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
            const res = await updateCategory(payload.id, payload);
            if (res.code === 0) {
                // commit('setCategory', res.data);
                await dispatch('loadCategory');
                return res;
            } else {
                return false
            }
        },
        async updateCategoryName({commit, dispatch, state}, payload: any) {
            const res = await updateCategoryName(payload.id, payload.name)
            if (res.code === 0) {
                await dispatch('loadCategory');
                return res;
            } else {
                return false
            }
        },
        async loadList({commit, dispatch, state}, {projectId, page, pageSize, opts}: any) {
            page = page || state.listResult.pagination.current;
            pageSize = pageSize || state.listResult.pagination.pageSize;
            const otherParams = {...state.filterState, ...opts};

            const res = await getEndpointList({
                "projectId": projectId,
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
            const res = await saveEndpoint({
                ...params
            });
            if (res.code === 0) {
                await dispatch('loadList', {projectId: params.projectId});
            } else {
                return false
            }
        },
        async disabled({commit, dispatch, state}, payload: any) {
            const res = await expireEndpoint(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {projectId: payload.projectId});
            } else {
                return false
            }
        },
        async del({commit, dispatch, state}, payload: any) {
            const res = await deleteEndpoint(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {projectId: payload.projectId});
                return true
            } else {
                return false
            }
        },
        async copy({commit, dispatch, state}, payload: any) {
            const res = await copyEndpoint(payload.id);
            if (res.code === 0) {
                await dispatch('loadList', {projectId: payload.projectId});
            } else {
                return false
            }
        },
        // 用于新建接口时选择接口分类
        async getEndpointDetail({commit, state}, payload: any) {
            // 请求数据之前先清空数据
            // await commit('setEndpointDetail',  {});
            await commit('clearInterfaceMethodToObjMap', {});

            const res = await getEndpointDetail(payload.id);
            res.data.createdAt = momentUtc(res.data.createdAt);
            res.data.updatedAt = momentUtc(res.data.updatedAt);

            if (res.code === 0) {
                await commit('setEndpointDetail', res.data || null);
                state.endpointDetail?.interfaces?.forEach((item) => {
                    commit('setInterfaceMethodToObjMap', {
                        method: item.method,
                        value: item,
                    });
                })


            } else {
                return false
            }
        },
        // 用于新建接口时选择接口分类
        async updateEndpointDetail({commit, dispatch}, payload: any) {
            const res = await saveEndpoint({
                ...payload
            });
            if (res.code === 0) {
                await dispatch('loadList', {projectId: payload.projectId});
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
                    item.label = item.description;
                    item.value = item.id;
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
        async example2schema({commit}, payload: any) {
            const res = await example2schema(payload);
            if (res.code === 0) {
                return res.data;
            } else {
                return null
            }
        },
        async schema2example({commit}, payload: any) {
            const res = await schema2example(payload);
            if (res.code === 0) {
                return res.data;
            } else {
                return null
            }
        },
        async getRefsOptions({commit}, payload: any) {
            const res = await getSchemaList({
                ...payload,
                "page": 1,
                "pageSize": 100
            });
            if (res.code === 0) {
                res.data.result.forEach((item: any) => {
                    item.label = item.ref;
                    item.value = item.ref;
                })
                commit('setRefsOptions', {
                    type: payload.type,
                    options: [...res.data.result]
                });
            } else {
                return null
            }
        },
        // 获取可选组件信息
        async getAllRefs({commit}, payload: any) {
            const res = await getSchemaList({
                ...payload,
                "page": 1,
                "pageSize": 100
            });
            if (res.code === 0) {
                res.data.result.forEach((item: any) => {
                    item.label = item.ref;
                    item.value = item.ref;

                })
                return res.data.result;
            } else {
                return null;
            }
        },
        // 获取可选组件信息
        async getRefDetail({commit}, payload: any) {
            const res = await getSchemaDetail({
                ...payload,
            });
            if (res.code === 0) {
                return res.data;
            } else {
                return null;
            }
        },

        // 获取可选组件信息
        async getDocs({commit}, payload: any) {
            const res = await getDocs({
                ...payload,
            });
            if (res.code === 0) {
                return res.data;
            } else {
                return null;
            }
        },

        // 获取可选组件信息
        async upload({commit}, payload: any) {
            let result = null;
            try {
                const res = await upload({
                    ...payload,
                });
                if (res.code === 0) {
                    result = res.data;
                } else {
                    result = null;
                }
            } catch (e) {
                result = null;
                console.log(e)
            }
            return result;
        },

        // 获取可选组件信息
        async importEndpointData({commit}, payload: any) {
            let result = null;
            try {
                const res = await importEndpointData({
                    ...payload,
                });
                if (res.code === 0) {
                    result = res;
                } else {
                    result = null;
                }
            } catch (e) {
                console.log(e)
            }
            return result;
        },
    }
};

export default StoreModel;
