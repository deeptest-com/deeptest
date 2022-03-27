import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {
    load, get, remove, create, update, expandAllKeys, move,
} from './service';
import {ApiKey, BasicAuth, BearerToken, OAuth20, Param, Request, Response} from "@/views/interface/data";

export interface StateType {
    treeData: any[];
    requestData: Request;
    responseData: Response;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTree: Mutation<StateType>;
        setRequest: Mutation<StateType>;
        setResponse: Mutation<StateType>;
    };
    actions: {
        loadInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        createInterface: Action<StateType, StateType>;
        updateInterface: Action<StateType, StateType>;
        deleteInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;
        sendRequest: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    treeData: [],
    requestData: {} as Request,
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
        setRequest(state, data) {
            if (!data.params) data.params = [{} as Param]
            if (!data.basicAuth) data.basicAuth = {} as BasicAuth
            if (!data.bearerToken) data.bearerToken = {} as BearerToken
            if (!data.oAuth20) data.oAuth20 = {} as OAuth20
            if (!data.apiKey) data.apiKey = {transferMode: 'headers'} as ApiKey

            //debug
            data.url = 'http://127.0.0.1:8085/api/v1/exec/test'
            data.params = [{name: 'param1', value: 1} as Param,
                {} as Param]
            data.headers = [{name: 'token', value: 'uuid'} as Param,
                {} as Param]

            state.requestData = data;
        },
        setResponse(state, payload) {
            state.responseData = payload;
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
        async getInterface({ commit }, payload: any ) {
            if (payload.isDir) {
                commit('setRequest', {});
                return true;
            }

            try {
                const response: ResponseData = await get(payload.id);
                const { data } = response;

                commit('setRequest',data);
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
        async sendRequest({ commit }, payload: any ) {
            try {
                // await sendRequest(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
