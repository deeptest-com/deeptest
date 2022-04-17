import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";

import {
    load, get, remove, create, update, move, invokeInterface, saveInterface, listRequest, removeRequest, loadHistory,
} from './service';
import {Interface, Response} from "@/views/interface/data";

export interface StateType {
    treeData: any[];
    interfaceData: Interface;
    responseData: Response;

    invocationsData: any[];
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTree: Mutation<StateType>;
        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;

        setInvocations: Mutation<StateType>;
    };
    actions: {
        invoke: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;

        loadInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        createInterface: Action<StateType, StateType>;
        updateInterface: Action<StateType, StateType>;
        deleteInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;

        listRequest: Action<StateType, StateType>;
        loadHistory: Action<StateType, StateType>;
        removeRequest: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    treeData: [],
    interfaceData: {} as Interface,
    responseData: {} as Response,

    invocationsData: [],
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Interface',
    state: {
        ...initState
    },
    mutations: {
        setTree(state, payload) {
            payload.name = '所有接口'
            state.treeData = [payload];
        },
        setInterface(state, data) {
            state.interfaceData = data;

        },
        setResponse(state, payload) {
            state.responseData = payload;
        },

        setInvocations(state, payload) {
            state.invocationsData = payload;
        },
    },
    actions: {
        async invoke({ commit }, payload: any ) {
            invokeInterface(payload).then((json) => {
                if (json.code === 0) {
                    commit('setResponse', json.data);
                    this.dispatch('Interface/listRequest', payload.id);
                    return true;
                } else {
                    return false
                }
            })
        },
        async saveInterface({ commit }, payload: any ) {
            saveInterface(payload).then((json) => {
                if (json.code === 0) {
                    return true;
                } else {
                    return false
                }
            })
        },

        async loadInterface({ commit }) {
            const response = await load();
            if (response.code != 0) return;

            const { data } = response;
            console.log('data', data)

            commit('setTree',data || {});
            return true;
        },
        async getInterface({ commit }, payload: any ) {
            if (payload.isDir) {
                commit('setInterface', {});
                return true;
            }

            try {
                const response = await get(payload.id);
                const { data } = response;

                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createInterface({ commit }, payload: any ) {
            try {
                const resp = await create(payload);

                await this.dispatch('Interface/loadInterface');
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async updateInterface({ commit }, payload: any ) {
            try {
                const { id, ...params } = payload;
                await update(id, { ...params });
                return true;
            } catch (error) {
                return false;
            }
        },
        async deleteInterface({ commit }, payload: number ) {
            try {
                await remove(payload);
                await this.dispatch('Interface/loadInterface');
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveInterface({ commit }, payload: any ) {
            try {
                await move(payload);
                await this.dispatch('Interface/loadInterface');
                return true;
            } catch (error) {
                return false;
            }
        },

        async listRequest({ commit }, interfaceId: number ) {
            try {
                const resp = await listRequest(interfaceId);
                const { data } = resp;
                commit('setInvocations', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async loadHistory({ commit }, requestId: number ) {
            try {
                const resp = await loadHistory(requestId);
                const { data } = resp;
                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeRequest({ commit }, data: any ) {
            try {
               await removeRequest(data.id);
                await this.dispatch('Interface/listRequest', data.interfaceId);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
