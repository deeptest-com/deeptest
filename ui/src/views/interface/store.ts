import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";

import {
    load, get, remove, create, update, expandAllKeys, move,
} from './service';
import {Interface, Response} from "@/views/interface/data";

export interface StateType {
    treeData: any[];
    interfaceData: Interface;
    responseData: Response;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTree: Mutation<StateType>;
        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;
    };
    actions: {
        loadInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        createInterface: Action<StateType, StateType>;
        updateInterface: Action<StateType, StateType>;
        deleteInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;
        sendInterface: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    treeData: [],
    interfaceData: {} as Interface,
    responseData: {} as Response,
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
    },
    actions: {
        async loadInterface({ commit }) {
            try {
                const response = await load();
                if (response.code != 0) return;

                const { data } = response;
                console.log('data', data)

                commit('setTree',data || {});
                return true;
            } catch (error) {
                return false;
            }
        },
        async getInterface({ commit }, payload: any ) {
            if (payload.isDir) {
                commit('setInterface', {});
                return true;
            }

            try {
                const response = await get(payload.id);
                const { data } = response;

                commit('setInterface',data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createInterface({ commit }, payload: any ) {
            try {
                const resp = await create(payload);
                console.log('resp', resp.data)

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
        async sendInterface({ commit }, payload: any ) {
            try {
                // await sendInterface(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
