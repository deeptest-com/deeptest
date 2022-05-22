import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    saveScenario, load, getNode, createNode, updateNode, removeNode, moveNode,
} from './service';

export interface StateType {
    treeData: Scenario[];
    scenarioData: Scenario;
    nodeData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setScenario: Mutation<StateType>;
        setNode: Mutation<StateType>;
    };
    actions: {
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
        setScenario(state, data) {
            state.scenarioData = data;
        },
        setNode(state, data) {
            state.nodeData = data;
        },
    },
    actions: {
        async loadScenario({commit}) {
            const response = await load();
            if (response.code != 0) return;

            const {data} = response;
            commit('setScenario', data || {});
            return true;
        },
        async saveScenario({commit}, payload: any) {
            saveScenario(payload).then((json) => {
                if (json.code === 0) {
                    return true;
                } else {
                    return false
                }
            })
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
