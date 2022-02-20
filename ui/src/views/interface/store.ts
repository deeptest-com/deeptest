import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {
    load, get, remove, create, update, expandAllKeys, move,
} from './service';

export interface StateType {
    treeResult: any[];
    expandedKeys: number[];
    detailResult: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTree: Mutation<StateType>;
        setItem: Mutation<StateType>;
    };
    actions: {
        loadInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        createInterface: Action<StateType, StateType>;
        updateInterface: Action<StateType, StateType>;
        deleteInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    treeResult: [],
    expandedKeys: [],
    detailResult: {},
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
            state.treeResult = [payload];
        },
        setItem(state, payload) {
            state.detailResult = payload;
        },
    },
    actions: {
        async loadInterface({ commit }) {
            try {
                const response: ResponseData = await load();
                if (response.code != 0) return;

                const { data } = response;
                console.log('data', data)

                commit('setTree',data || {});
                return true;
            } catch (error) {
                return false;
            }
        },
        async getInterface({ commit }, payload: number ) {
            try {
                const response: ResponseData = await get(payload);
                const { data } = response;

                commit('setItem',{
                    ...initState.detailResult,
                    ...data,
                });
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
    }
};

export default StoreModel;
