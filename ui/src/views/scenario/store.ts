import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    remove,
    load,
    getNode,
    createNode,
    updateNode,
    removeNode,
    moveNode,
    addInterfaces,
    addProcessor,
    saveProcessorName, saveProcessor,

    loadExecResult, getInterface,
} from './service';
import {getNodeMap} from "@/services/tree";
import {Interface, Response} from "@/views/interface/data";
import {
    invokeInterface,
    listCheckpoint,
    listExtractor,
    listValidExtractorVariableForInterface,
    saveInterface
} from "@/views/interface/service";

export interface StateType {
    scenarioId: number;

    listResult: QueryResult;
    detailResult: Scenario;
    queryParams: any;

    treeData: Scenario[];
    treeDataMap: any,
    nodeData: any;

    execResult: any;

    interfaceData: Interface;
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

        setExecResult: Mutation<StateType>;

        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;
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

        saveProcessorName: Action<StateType, StateType>;
        saveProcessor: Action<StateType, StateType>;

        saveTreeMapItem: Action<StateType, StateType>;
        saveTreeMapItemProp: Action<StateType, StateType>;

        loadExecResult: Action<StateType, StateType>;
        updateExecResult: Action<StateType, StateType>;

        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        invokeInterface: Action<StateType, StateType>;

        listExtractor: Action<StateType, StateType>;
        listCheckpoint: Action<StateType, StateType>;
        listValidExtractorVariableForInterface: Action<StateType, StateType>;
    };
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
    execResult: {},

    interfaceData: {} as Interface,
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

        setExecResult(state, data) {
            state.execResult = data;
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },

        setInterface(state, data) {
            state.interfaceData = data;

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
        async listScenario({ commit, dispatch }, params: QueryParams ) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;

                commit('setList',{
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

        async getScenario({ commit }, id: number ) {
            if (id === 0) {
                commit('setDetail',{
                    ...initState.detailResult,
                })
                return
            }
            try {
                const response: ResponseData = await get(id);
                const { data } = response;
                commit('setDetail',{
                    ...initState.detailResult,
                    ...data,
                });
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

        async saveScenario({commit}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async removeScenario({ commit, dispatch, state }, payload: number ) {
            try {
                await remove(payload);
                await dispatch('listScenario', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },

        async loadScenario({commit}, scenarioId) {
            const response = await load(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setTreeData', data || {});
            commit('setScenarioId', scenarioId );

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

        async loadExecResult({commit, dispatch, state}, scenarioId) {
            const response = await loadExecResult(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setExecResult', data || {});
            commit('setScenarioId', scenarioId );

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
                data.headers.push({name:'', value:''})
                data.params.push({name:'', value:''})
                data.bodyFormData.push({name:'', value:'', type: 'text'})
                data.bodyFormUrlencoded.push({name:'', value:''})

                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },

        async saveInterface({commit}, payload: any) {
            const json = await  saveInterface(payload)
            if (json.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async invokeInterface({commit, dispatch, state}, data: any) {
            const response = await invokeInterface(data)
            // console.log('=invoke=', response.data)

            if (response.code === 0) {
                commit('setResponse', response.data);

                dispatch('listExtractor', data.usedBy);
                dispatch('listCheckpoint', data.usedBy);

                return true;
            } else {
                return false
            }
        },

        async listExtractor({commit, dispatch, state}, usedBy) {
            try {
                const resp = await listExtractor(state.interfaceData.id, usedBy);
                const {data} = resp;
                commit('setExtractors', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async listCheckpoint({commit, state}, usedBy) {
            try {
                const resp = await listCheckpoint(state.interfaceData.id, usedBy);
                const {data} = resp;
                commit('setCheckpoints', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async listValidExtractorVariableForInterface({commit, dispatch, state}, usedBy) {
            try {
                console.log('listValidExtractorVariableForInterface')
                const resp = await listValidExtractorVariableForInterface(state.interfaceData.id, usedBy);
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
