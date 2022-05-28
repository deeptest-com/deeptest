import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query, get, save, load, getNode, createNode, updateNode, removeNode, moveNode,
} from './service';

export interface StateType {
    listResult: QueryResult;
    detailResult: Scenario;

    treeData: Scenario[];
    scenarioData: Scenario;
    nodeData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;

        setScenario: Mutation<StateType>;
        setNode: Mutation<StateType>;
    };
    actions: {
        listScenario: Action<StateType, StateType>;
        getScenario: Action<StateType, StateType>;

        loadScenario: Action<StateType, StateType>;
        saveScenario: Action<StateType, StateType>;
        getNode: Action<StateType, StateType>;
        createNode: Action<StateType, StateType>;
        updateNode: Action<StateType, StateType>;
        removeNode: Action<StateType, StateType>;
        moveNode: Action<StateType, StateType>;
    };
}
const initState: StateType = {
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
    scenarioData: {} as Scenario,
    nodeData: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Scenario',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },
        
        setScenario(state, data) {
            state.scenarioData = data;
        },
        setNode(state, data) {
            state.nodeData = data;
        },
    },
    actions: {
        async listScenario({ commit }, params: QueryParams ) {
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

        async loadScenario({commit}) {
            const response = await load();
            if (response.code != 0) return;

            const {data} = response;
            commit('setScenario', data || {});
            return true;
        },

        async getNode({commit}, payload: any) {
            if (payload.isDir) {
                commit('setNode', {});
                return true;
            }

            try {
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
                await dispatch('loadScenario');
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveNode({commit, dispatch, state}, payload: any) {
            try {
                await moveNode(payload);
                await dispatch('loadScenario');
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
