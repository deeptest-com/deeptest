import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query,
    get,
    save,
    load,
    getNode,
    createNode,
    updateNode,
    removeNode,
    moveNode,
    addInterfaces,
    addProcessor,
    saveProcessorName, saveProcessor,

    loadExecResult,
} from './service';

export interface StateType {
    scenarioId: number;

    listResult: QueryResult;
    detailResult: Scenario;

    treeData: Scenario[];
    nodeData: any;

    execResult: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setScenarioId: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;

        setTree: Mutation<StateType>;
        setNode: Mutation<StateType>;

        setExecResult: Mutation<StateType>;
    };
    actions: {
        listScenario: Action<StateType, StateType>;
        getScenario: Action<StateType, StateType>;

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

        loadExecResult: Action<StateType, StateType>;
        updateExecResult: Action<StateType, StateType>;
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

    treeData: [],
    nodeData: {},
    execResult: {},
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

        setTree(state, data) {
            state.treeData = [data];
        },
        setNode(state, data) {
            state.nodeData = data;
        },

        setExecResult(state, data) {
            state.execResult = data;
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
        async saveScenario({commit}, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                return true;
            } else {
                return false
            }
        },

        async loadScenario({commit}, scenarioId) {
            const response = await load(scenarioId);
            if (response.code != 0) return;

            const {data} = response;
            commit('setTree', data || {});
            commit('setScenarioId', scenarioId );

            return true;
        },

        async getNode({commit}, payload: any) {
            try {
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
    }
};

export default StoreModel;
