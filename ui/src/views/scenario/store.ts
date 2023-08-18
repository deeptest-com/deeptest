import {Mutation, Action} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {ResponseData} from '@/utils/request';
import {Scenario, QueryResult, QueryParams} from './data.d';
import {
    query,
    get,
    save,
    remove,
    loadScenario,
    getNode,
    createNode,
    updateNode,
    removeNode,
    disableNodeOrNot,
    moveNode,
    addInterfacesFromDefine, addInterfacesFromDiagnose, addProcessor,
    saveProcessor, loadExecResult,
    getScenariosReports,
    getScenariosReportsDetail,
    addPlans,
    getPlans,
    removePlans, updatePriority, updateStatus, genReport, saveDebugData, syncDebugData, saveProcessorInfo,importCurl,addInterfacesFromCase
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

import {getNodeMap} from "@/services/tree";
import {getSnippet} from "@/views/component/debug/service";

export interface StateType {
    scenarioId: number;
    scenarioProcessorIdForDebug: number;
    // endpointInterfaceIdForDebug: number;

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
    reportsDetail: any;

    interfaceData: any;
    invocationsData: [],
    responseData: any;
    extractorsData: any[];
    checkpointsData: any[];
    scenariosReports: any[];
    linkedPlans: any[];
    notLinkedPlans: any[];
    linkedPlansPagination: any;

    scenarioCount: number,
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setScenarioId: Mutation<StateType>;
        setScenarioProcessorIdForDebug: Mutation<StateType>;
        // setEndpointInterfaceIdForDebug: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
        setReportsDetail: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;

        // tree of scenario nodes
        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        setTreeDataMapItem: Mutation<StateType>;
        setTreeDataMapItemProp: Mutation<StateType>;
        setNode: Mutation<StateType>;
        setCodeContent: Mutation<StateType>;

        // tree of scenario categories
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
        setScenariosReports: Mutation<StateType>;
        setLinkedPlans: Mutation<StateType>;
        setNotLinkedPlans: Mutation<StateType>;
        setLinkedPlansPagination: Mutation<StateType>;

        increaseScenarioCount: Mutation<StateType>;
    };
    actions: {
        setScenarioProcessorIdForDebug: Action<StateType, StateType>;
        // setEndpointInterfaceIdForDebug: Action<StateType, StateType>;
        listScenario: Action<StateType, StateType>;
        getScenario: Action<StateType, StateType>;
        removeScenario: Action<StateType, StateType>;

        loadScenario: Action<StateType, StateType>;
        saveScenario: Action<StateType, StateType>;
        getNode: Action<StateType, StateType>;
        updateCategoryId: Action<StateType, StateType>;

        addInterfacesFromDefine: Action<StateType, StateType>;
        addInterfacesFromDiagnose: Action<StateType, StateType>;
        addProcessor: Action<StateType, StateType>;

        createNode: Action<StateType, StateType>;
        updateNode: Action<StateType, StateType>;
        removeNode: Action<StateType, StateType>;
        disableNodeOrNot: Action<StateType, StateType>;
        moveNode: Action<StateType, StateType>;
        saveTreeMapItem: Action<StateType, StateType>;
        saveTreeMapItemProp: Action<StateType, StateType>;

        saveProcessorInfo: Action<StateType, StateType>;
        saveProcessor: Action<StateType, StateType>;
        addSnippet: Action<StateType, StateType>;

        loadExecResult: Action<StateType, StateType>;
        updateExecResult: Action<StateType, StateType>;
        getExecResultList: Action<StateType, StateType>;
        addPlans: Action<StateType, StateType>;
        removePlans: Action<StateType, StateType>;
        getPlans: Action<StateType, StateType>;
        updatePriority: Action<StateType, StateType>;
        updateStatus: Action<StateType, StateType>;
        getScenariosReportsDetail: Action<StateType, StateType>;
        genReport: Action<StateType, StateType>;
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

        saveDebugData: Action<StateType, StateType>;
        syncDebugData: Action<StateType, StateType>;

        importCurl: Action<StateType, StateType>;
        addInterfacesFromCase: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    scenarioId: 0,
    scenarioProcessorIdForDebug: 0,
    // endpointInterfaceIdForDebug: 0,

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
    reportsDetail: {},
    interfaceData: {},
    invocationsData: [],
    responseData: {},
    extractorsData: [],
    checkpointsData: [],
    scenariosReports: [],
    linkedPlans: [],
    notLinkedPlans: [],
    linkedPlansPagination: {
        current: 1,
        total: 0,
        pageSize: 10,
        showSizeChanger: false,
    },

    scenarioCount: 0,
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
        setScenarioProcessorIdForDebug(state, id) {
            state.scenarioProcessorIdForDebug = id;
        },
        // setEndpointInterfaceIdForDebug(state, id) {
        //     state.endpointInterfaceIdForDebug = id;
        // },

        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },
        setReportsDetail(state, payload) {
            state.reportsDetail = payload;
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
            console.log('=== setNode', data)
            state.nodeData = data;
        },
        setCodeContent(state, content) {
            console.log('setCodeContent', content)
            state.nodeData.content = content;
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
        setScenariosReports(state, payload) {
            state.scenariosReports = payload;
        },
        setLinkedPlans(state, payload) {
            state.linkedPlans = payload;
        },
        setNotLinkedPlans(state, payload) {
            state.notLinkedPlans = payload;
        },
        setLinkedPlansPagination(state, payload) {
            state.linkedPlansPagination = payload;
        },

        increaseScenarioCount(state) {
            state.scenarioCount += 1;
        },
    },
    actions: {
        async setScenarioProcessorIdForDebug({commit, dispatch, state}, id) {
            commit('setScenarioProcessorIdForDebug', id);
            return true;
        },
        // async setEndpointInterfaceIdForDebug({commit, dispatch, state}, id) {
        //     commit('setEndpointInterfaceIdForDebug', id);
        //     return true;
        // },
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
            commit('increaseScenarioCount')

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


        async saveScenario({commit ,dispatch, state}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },
        async updateCategoryId({commit, dispatch, state}, payload) {
            const res = await save(payload);
            if (res.code === 0) {
                commit('setDetail', {
                    ...state.detailResult,
                    categoryId: payload.categoryId
                });
                return res;
            }
            return false;
        },
        async removeScenario({commit, dispatch, state}, payload: number) {
            try {
                await remove(payload);
                await dispatch('listScenario', state.queryParams)
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },

        async addInterfacesFromDefine({commit, dispatch, state}, payload: any) {
            try {
                const resp = await addInterfacesFromDefine(payload);

                await dispatch('loadScenario', state.scenarioId);
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async addInterfacesFromDiagnose({commit, dispatch, state}, payload: any) {
            try {
                const resp = await addInterfacesFromDiagnose(payload);

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
        async disableNodeOrNot({commit, dispatch, state}, payload: number) {
            try {
                await disableNodeOrNot(payload);
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
        async saveProcessorInfo({commit, dispatch, state}, payload: any) {
            const jsn = await saveProcessorInfo(payload)
            if (jsn.code === 0) {
                commit('setTreeDataMapItemProp', {
                    id: payload.id,
                    prop: 'name',
                    value: payload.name});
                commit('setTreeDataMapItemProp', {
                    id: payload.id,
                    prop: 'comments',
                    value: payload.comments});
                return true;
            } else {
                return false
            }
        },
        async saveProcessor({commit, dispatch, state}, payload: any) {
            const jsn = await saveProcessor(payload)
            if (jsn.code === 0) {
                commit('setNode', jsn.data);
                commit('setTreeDataMapItemProp', {
                    id: jsn.data.processorID,
                    prop: 'name',
                    value: jsn.data.name});

                return true;
            } else {
                return false
            }
        },
        async addSnippet({commit, dispatch, state}, name: string) {
            const json = await getSnippet(name)
            if (json.code === 0) {
                let script = (state.nodeData.content ? state.nodeData.content: '') + '\n' +  json.data.script
                script = script.trim()
                commit('setCodeContent', script);
            }

            return true;
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
                if (payload) {
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
                return resp?.data;
            } catch (error) {
                return false;
            }
        },

        async updateCategoryNode({commit, dispatch}, payload: any) {
            try {
                const {id, ...params} = payload;
                await updateCategory(id, {...payload});
                await dispatch('loadCategory');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeCategoryNode({commit, dispatch, state}, payload: any) {
            try {
                await removeCategory(payload.id, payload.type);
                await dispatch('listScenario', state.queryParams)
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
        async getExecResultList({commit, dispatch, state}, payload) {
            const res = await getScenariosReports(payload.data || {});
            if (res.code === 0) {
                commit('setScenariosReports', res?.data?.result || []);
            }
            return true;
        },
        async addPlans({commit, dispatch, state}, payload) {
            const res = await addPlans(payload);
            if (res.code === 0) {
                return res;
            }
            return false;
        },
        async removePlans({commit, dispatch, state}, payload) {
            const res = await removePlans(payload);
            if (res.code === 0) {
                return res;
            }
            return false;
        },
        async getPlans({commit, dispatch, state}, payload) {
            const res = await getPlans(payload);
            if (res.code === 0) {
                if (payload.data.ref) {
                    commit('setLinkedPlans', res?.data?.result || []);
                } else {
                    commit('setNotLinkedPlans', res?.data?.result || []);
                }
                commit('setLinkedPlansPagination', {
                    ...state.linkedPlansPagination,
                    current: res.data.page || 1,
                    total: res.data.total || 0,
                })
            }
            return false;
        },
        async updatePriority({commit, dispatch, state}, payload) {
            const res = await updatePriority(payload);
            if (res.code === 0) {
                commit('setDetail', {
                    ...state.detailResult,
                    priority: payload.priority
                });
                return res;
            }
            return false;
        },
        async updateStatus({commit, dispatch, state}, payload) {
            const res = await updateStatus(payload);
            if (res.code === 0) {
                commit('setDetail', {
                    ...state.detailResult,
                    status: payload.status
                });
                return res;
            }
            return false;
        },
        async getScenariosReportsDetail({commit, dispatch, state}, payload) {
            const res = await getScenariosReportsDetail(payload);
            if (res.code === 0) {
                commit('setReportsDetail', {
                    ...res.data
                });
                return res;
            }
            return false;
        },
        async genReport({commit, dispatch, state}, payload) {
            const res = await genReport(payload);
            if (res.code === 0) {
                return res;
            }
            return false;
        },

        async saveDebugData({commit}, payload: any) {
            const resp = await  saveDebugData(payload)
            return resp.code === 0;
        },
        async syncDebugData({commit, state, dispatch}) {
            const resp = await  syncDebugData(state.scenarioProcessorIdForDebug)
            dispatch('loadScenario', state.scenarioId);
            commit('setScenarioProcessorIdForDebug', resp.data.id)
            return resp.code === 0;
        },
        async importCurl({state,dispatch}, payload) {
            const res = await importCurl(payload);
            await dispatch('loadScenario', state.scenarioId);
            if (res.code === 0) {
                return res.data;
            }
            return false;
        },
        async addInterfacesFromCase({commit, dispatch, state}, payload: any) {
            try {
                const resp = await addInterfacesFromCase(payload);

                await dispatch('loadScenario', state.scenarioId);
                return resp.data;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;


