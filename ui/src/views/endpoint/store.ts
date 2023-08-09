import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';
import {
    Endpoint,
    QueryResult,
    QueryParams,
    QueryCaseTreeParams,
    filterFormState
} from './data.d';
import {
    query,
    get,
    save,
    remove,
    getYaml,
    updateStatus,
    getDocs,
    importEndpointData,
    upload,
    updateEndpointCaseName,
    removeEndpointCase,
    getEndpointCase,
    listEndpointCase,
    saveEndpointCase,
    saveEndpointCaseDebugData,
    batchUpdateField,
    tagList,
    updateTag,
    copyEndpoint,
    deleteEndpoint,
    expireEndpoint,
    getEndpointDetail,
    getEndpointList,
    saveEndpoint, copyEndpointCase,loadCaseTree,reBuildTree
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

import {genNodeMap, getNodeMap} from "@/services/tree";
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
    filterState: filterFormState;
    endpointDetail: any,
    endpointDetailYamlCode: any,
    serveServers: any[]; // serve list
    securityOpts: any[];
    interfaceMethodToObjMap: any;
    refsOptions: any;
    selectedMethodDetail: any;
    selectedCodeDetail: any;

    caseList: any[];
    caseDetail: any;
    tagList: any;
    caseTree:any;
    caseTreeMap:any;
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

        setEndpointCaseList: Mutation<StateType>;
        setEndpointCaseDetail: Mutation<StateType>;
        setEndpointTagList: Mutation<StateType>;

        setCaseTree: Mutation<StateType>;
        setCaseTreeMap:Mutation<StateType>;
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

        listCase: Action<StateType, StateType>;
        getCase: Action<StateType, StateType>;
        saveCase: Action<StateType, StateType>;
        copyCase: Action<StateType, StateType>;
        updateCaseName: Action<StateType, StateType>;
        removeCase: Action<StateType, StateType>;
        batchUpdateField: Action<StateType, StateType>;
        saveCaseDebugData: Action<StateType, StateType>;
        getEndpointList: Action<StateType, StateType>;
        getEndpointTagList: Action<StateType, StateType>;
        updateEndpointTag: Action<StateType, StateType>;
        getCaseTree: Action<StateType, StateType>;
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
        categoryId: null,
        tagNames:[]
    },
    endpointDetail: null,
    endpointDetailYamlCode: null,
    serveServers: [],
    securityOpts: [],
    interfaceMethodToObjMap: {},
    refsOptions: {},
    selectedMethodDetail: {},
    selectedCodeDetail: {},
    caseList: [],
    caseDetail: {},
    tagList:[],
    caseTree:[],
    caseTreeMap:[],
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
            state.filterState.categoryId = payload.categoryId || null ;
        },
        clearFilterState(state) {
            state.filterState.status = null;
            state.filterState.createUser = null;
            state.filterState.title =null ;
            state.filterState.categoryId =null ;
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
        },

        setEndpointCaseList(state, payload) {
            state.caseList = payload;
        },
        setEndpointCaseDetail(state, payload) {
            state.caseDetail = payload;
        },
        setEndpointTagList(state, payload) {
            state.tagList = payload;
        },
        setCaseTree(state,payload){
            state.caseTree = payload
        },
        setCaseTreeMap(state,payload){
            state.caseTreeMap = payload
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
        async removeCategoryNode({commit, dispatch, state}, payload: any) {
            try {
                await removeCategory(payload.id, payload.type);
                await dispatch('loadCategory');
                await dispatch('loadList', {projectId: payload.projectId});
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
        async loadList({commit, state, rootState}: any, {projectId, page, pageSize, ...opts}: any) {
            page = page || state.listResult.pagination.current;
            pageSize = pageSize || state.listResult.pagination.pageSize;
            const otherParams = {
                ...state.filterState, 
                serveId: rootState.ServeGlobal.currServe.id, 
                ...opts
            };

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
                // 删除接口后，需要重新拉取分类树
                await dispatch('loadCategory');
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
                await dispatch("getEndpointDetail", {id: res.data})
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
        async listCase({commit}, endpointId: number) {
            const resp = await listEndpointCase(endpointId);
            commit('setEndpointId', endpointId);

            if (resp.code === 0) {
                commit('setEndpointCaseList', resp.data);
                return true;
            } else {
                return false
            }
        },
        async getCase({commit}, id: number) {
            const resp: ResponseData = await getEndpointCase(id);

            if (resp.code === 0) {
                commit('setEndpointCaseDetail', resp.data);
                return true;
            } else {
                return false
            }
        },
        async saveCase({state, dispatch}, payload: any) {
            const jsn = await saveEndpointCase(payload)
            if (jsn.code === 0) {
                dispatch('listCase', state.endpointId);
                return true;
            } else {
                return false
            }
        },
        async copyCase({state, dispatch}, id: number) {
            const jsn = await copyEndpointCase(id)
            if (jsn.code === 0) {
                return jsn.data;
            } else {
                return null
            }
        },
        async updateCaseName({state, dispatch}, payload: any) {
            const jsn = await updateEndpointCaseName(payload)
            if (jsn.code === 0) {
                dispatch('listCase', state.endpointId);
                return true;
            } else {
                return false
            }
        },
        async removeCase({commit, dispatch, state}, record: any) {
            try {
                const jsn = await removeEndpointCase(record);
                if (jsn.code === 0) {
                    dispatch('listCase', state.endpointId);
                    return true;
                }
                return false;
            } catch (error) {
                return false;
            }
        },
        async batchUpdateField({commit, dispatch}, payload: any) {
            const res = await batchUpdateField(payload);
            if (res.code === 0) {
                await dispatch('loadList', {projectId: payload.projectId});
                await dispatch('loadCategory');
            } else {
                return null
            }
        },

        async saveCaseDebugData({ state, dispatch }, payload: any) {
            const jsn = await saveEndpointCaseDebugData(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },

        async getEndpointList({ commit }, payload: any) {
            const resp = await getEndpointList(payload)
            if (resp.code === 0) {
                commit('setEndpointCaseDetail', resp.data);
            } else {
                return false
            }
        },
        async getEndpointTagList({ commit } ) {
            const resp = await tagList()
            if (resp.code === 0) {

               const res =  resp.data.map((arrItem)=>{
                 return {label:arrItem.tagName,value:arrItem.tagName}
               })
               commit("setEndpointTagList", res);
               //console.log("setEndpointTagList",state.tagList);

            } else {
                return false
            }
        },
        async updateEndpointTag({ dispatch }, payload: any) {
            const jsn = await updateTag(payload)
           
            if (jsn.code === 0) {
                await dispatch("getEndpointTagList")
                return true;
            } else {
                return false
            }
        },
        async getCaseTree({ commit }, payload: QueryCaseTreeParams){
            try {
                const response: ResponseData = await loadCaseTree(payload);
             
                if (response.code != 0) return;
              
                const data = {id: 0,count:1, children: response.data} 
            
                const newData = reBuildTree(data,0)
        
                commit('setCaseTree', newData.children);
                const data1 = {id: 0, children: newData.children} 
                const mp = genNodeMap(data1)
                commit('setCaseTreeMap', mp);

                return true;
            } catch (error) {
                
                return false;
            }

        }


    },
};

export default StoreModel;
