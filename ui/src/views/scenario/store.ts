import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    remove,

    listInvocation, getInvocationAsInterface, removeInvocation,
    loadExecResult, getInterface, getLastInvocationResp,

    loadScenario,
    getNode,
    createNode,
    updateNode,
    removeNode,
    moveNode,
    addInterfaces, addProcessor,
    saveProcessorName, saveProcessor, saveInterface,
} from './service';

import {
    loadCategory,
    getCategory,
    createCategory,
    updateCategory,
    removeCategory,
    moveCategory,
    updateCategoryName} from "@/services/category";

// below use same apis with interface controller
import {
    invokeInterface,
    listExtractor,
    listCheckpoint,
    listValidExtractorVariableForInterface,
} from "@/views/interface1/service";

import {getNodeMap} from "@/services/tree";
import {Interface, Response} from "@/views/interface1/data";
import {UsedBy} from "@/utils/enum";

export interface StateType {
    scenarioId: number;

    listResult: QueryResult;
    detailResult: Scenario;
    queryParams: any;

    treeData: Scenario[];
    treeDataMap: any,
    nodeData: any;

    treeDataCategory: any[];
    treeDataMapCategory: any,
    nodeDataCategory: any;

    execResult: any;

    interfaceData: Interface;
    invocationsData: [],
    responseData: Response;
    extractorsData: any[];
    checkpointsData: any[];
    validExtractorVariablesData: any[];
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setScenarioId: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;

        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        setTreeDataMapItem: Mutation<StateType>;
        setTreeDataMapItemProp: Mutation<StateType>;
        setNode: Mutation<StateType>;

        setTreeDataCategory: Mutation<StateType>;
        setTreeDataMapCategory: Mutation<StateType>;
        setTreeDataMapItemCategory: Mutation<StateType>;
        setTreeDataMapItemPropCategory: Mutation<StateType>;
        setNodeCategory: Mutation<StateType>;

        setExecResult: Mutation<StateType>;

        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;
        setInvocations: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setCheckpoints: Mutation<StateType>;
        setValidExtractorVariables: Mutation<StateType>;
    };
    actions: {
        listScenario: Action<StateType, StateType>;
        getScenario: Action<StateType, StateType>;
        removeScenario: Action<StateType, StateType>;

        loadScenario: Action<StateType, StateType>;
        saveScenario: Action<StateType, StateType>;
        getNode: Action<StateType, StateType>;

        addInterfaces: Action<StateType, StateType>;
        addProcessor: Action<StateType, StateType>;

        createNode: Action<StateType, StateType>;
        updateNode: Action<StateType, StateType>;
        removeNode: Action<StateType, StateType>;
        moveNode: Action<StateType, StateType>;
        saveTreeMapItem: Action<StateType, StateType>;
        saveTreeMapItemProp: Action<StateType, StateType>;

        saveProcessorName: Action<StateType, StateType>;
        saveProcessor: Action<StateType, StateType>;

        loadExecResult: Action<StateType, StateType>;
        updateExecResult: Action<StateType, StateType>;

        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        invokeInterface: Action<StateType, StateType>;
        getLastInvocationResp: Action<StateType, StateType>;

        listInvocation: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listExtractor: Action<StateType, StateType>;
        listCheckpoint: Action<StateType, StateType>;
        listValidExtractorVariableForInterface: Action<StateType, StateType>;

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
    scenarioId: 0,

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
    detailResult: {} as Scenario,
    queryParams: {},

    treeData: [],
    treeDataMap: {},
    nodeData: {},

    treeDataCategory: [],
    treeDataMapCategory: {},
    nodeDataCategory: {},

    execResult: {},

    interfaceData: {} as Interface,
    invocationsData: [],
    responseData: {} as Response,
    extractorsData: [],
    checkpointsData: [],
    validExtractorVariablesData: [],
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Scenario',
    state: {
        ...initState
    },
    mutations: {
        setScenarioId(state, id) {
            state.scenarioId = id;
        },

        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },

        setTreeData(state, data) {
            state.treeData = [data];
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
        setNode(state, data) {
            state.nodeData = data;
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

        setExecResult(state, data) {
            state.execResult = data;
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },

        setInterface(state, data) {
            state.interfaceData = data;
        },
        setInvocations(state, payload) {
            state.invocationsData = payload;
        },
        setResponse(state, payload) {
            state.responseData = payload;
        },
        setExtractors(state, payload) {
            state.extractorsData = payload;
        },
        setCheckpoints(state, payload) {
            state.checkpointsData = payload;
        },
        setValidExtractorVariables(state, payload) {
            state.validExtractorVariablesData = payload;
        },
    },
    actions: {
        async listScenario({commit, dispatch}, params: QueryParams) {
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
        async getScenario({commit}, id: number) {
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

        async saveScenario({commit}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async removeScenario({commit, dispatch, state}, payload: number) {
            try {
                await remove(payload);
                await dispatch('listScenario', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },

        async addInterfaces({commit, dispatch, state}, payload: any) {
            try {
                const resp = await addInterfaces(payload);

                await dispatch('loadScenario', state.scenarioId);
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async addProcessor({commit, dispatch, state}, payload: any) {
            try {
                const resp = await addProcessor(payload);

                await dispatch('loadScenario', state.scenarioId);
                return resp.data;
            } catch (error) {
                return false;
            }
        },

        // scenario tree
        async loadScenario({commit}, scenarioId) {
            const response = await loadScenario(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setTreeData', data || {});
            commit('setScenarioId', scenarioId);

            const mp = {}
            getNodeMap(data, mp)

            commit('setTreeDataMap', mp);

            return true;
        },
        async getNode({commit}, payload: any) {
            try {
                if (!payload) {
                    commit('setNode', {});
                    return true;
                }

                const response = await getNode(payload.id);
                const {data} = response;

                commit('setNode', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createNode({commit, dispatch, state}, payload: any) {
            try {
                const resp = await createNode(payload);

                await dispatch('loadScenario');
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async updateNode({commit}, payload: any) {
            try {
                const {id, ...params} = payload;
                await updateNode(id, {...params});
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeNode({commit, dispatch, state}, payload: number) {
            try {
                await removeNode(payload);
                await dispatch('loadScenario', state.scenarioId);
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveNode({commit, dispatch, state}, payload: any) {
            try {
                await moveNode(payload);
                await dispatch('loadScenario', state.scenarioId);
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveTreeMapItem({commit}, payload: any) {
            commit('setTreeDataMapItem', payload);
        },
        async saveTreeMapItemProp({commit}, payload: any) {
            commit('setTreeDataMapItemProp', payload);
        },
        async saveProcessor({commit, dispatch, state}, payload: any) {
            const jsn = await saveProcessor(payload)
            if (jsn.code === 0) {
                commit('setNode', jsn.data);
                await dispatch('loadScenario', state.scenarioId);
                return true;
            } else {
                return false
            }
        },
        async saveProcessorName({commit, dispatch, state}, payload: any) {
            const jsn = await saveProcessorName(payload)
            if (jsn.code === 0) {
                await dispatch('loadScenario', state.scenarioId);
                return true;
            } else {
                return false
            }
        },

        // category tree
        async loadCategory({commit}) {
            const response = await loadCategory("scenario");
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
            const jsn = await saveProcessor(payload)
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

        async loadExecResult({commit, dispatch, state}, scenarioId) {
            const response = await loadExecResult(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setExecResult', data || {});
            commit('setScenarioId', scenarioId);

            return true;
        },
        async updateExecResult({commit, dispatch, state}, payload) {
            commit('setExecResult', payload);
            commit('setScenarioId', payload.scenarioId);

            return true;
        },

        async getInterface({commit}, interfaceId: number) {
            try {
                const response = await getInterface(interfaceId);

                const {data} = response;
                data.headers.push({name: '', value: ''})
                data.params.push({name: '', value: ''})
                data.bodyFormData.push({name: '', value: '', type: 'text'})
                data.bodyFormUrlencoded.push({name: '', value: ''})

                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },

        async getLastInvocationResp({commit, dispatch, state}, id: number) {
            const response = await getLastInvocationResp(id);
            // console.log('=getLastInvocationResp=', response.data)

            const {data} = response;

            commit('setResponse', data);
            return true;
        },

        async saveInterface({commit}, payload: any) {
            const json = await saveInterface(payload)
            if (json.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async invokeInterface({commit, dispatch, state}, data: any) {
            const response = await invokeInterface(data)
            // console.log('=invoke in processor=', response.data)

            if (response.code === 0) {
                commit('setResponse', response.data);

                dispatch('listInvocation', state.interfaceData.id);
                dispatch('listValidExtractorVariableForInterface');

                dispatch('listExtractor');
                dispatch('listCheckpoint');

                return true;
            } else {
                return false
            }
        },

        // invocation
        async listInvocation({commit}, interfaceId: number) {
            try {
                const resp = await listInvocation(interfaceId);
                const {data} = resp;
                commit('setInvocations', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getInvocationAsInterface({commit}, id: number) {
            try {
                const resp = await getInvocationAsInterface(id);
                const {data} = resp;

                commit('setInterface', data.req);
                commit('setResponse', data.resp);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeInvocation({commit, dispatch, state}, data: any) {
            try {
                await removeInvocation(data.id);
                dispatch('listInvocation', data.interfaceId);
                return true;
            } catch (error) {
                return false;
            }
        },

        async listExtractor({commit, dispatch, state}) {
            try {
                const resp = await listExtractor(state.interfaceData.id, UsedBy.ScenarioDebug);
                const {data} = resp;
                commit('setExtractors', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async listCheckpoint({commit, state}) {
            try {
                const resp = await listCheckpoint(state.interfaceData.id, UsedBy.ScenarioDebug);
                const {data} = resp;
                commit('setCheckpoints', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async listValidExtractorVariableForInterface({commit, dispatch, state}) {
            try {
                console.log('listValidExtractorVariableForInterface')
                const resp = await listValidExtractorVariableForInterface(state.interfaceData.id, UsedBy.ScenarioDebug);
                const {data} = resp;
                commit('setValidExtractorVariables', data);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
