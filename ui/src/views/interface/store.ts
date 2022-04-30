import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";

import {
    changeEnvironment,
    clearEnvironmentVar,
    copyEnvironment,
    create,
    get,
    getEnvironment,
    getInvocationAsInterface,
    invokeInterface,
    listEnvironment,
    listInvocation,
    load,
    move,
    remove,
    removeEnvironment,
    removeEnvironmentVar,
    removeInvocation,
    saveEnvironment,
    saveEnvironmentVar,
    saveInterface,
    update,
} from './service';
import {Interface, Response} from "@/views/interface/data";

export interface StateType {
    treeData: any[];
    interfaceData: Interface;
    responseData: Response;

    invocationsData: any[];

    environmentsData: any[];
    environmentData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTree: Mutation<StateType>;
        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;

        setInvocations: Mutation<StateType>;

        setEnvironments: Mutation<StateType>;
        setEnvironment: Mutation<StateType>;
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

        listInvocation: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listEnvironment: Action<StateType, StateType>;
        getEnvironment: Action<StateType, StateType>;
        changeEnvironment: Action<StateType, StateType>;
        saveEnvironment: Action<StateType, StateType>;
        copyEnvironment: Action<StateType, StateType>;
        removeEnvironment: Action<StateType, StateType>;

        saveEnvironmentVar: Action<StateType, StateType>;
        removeEnvironmentVar: Action<StateType, StateType>;
        clearEnvironmentVar: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    treeData: [],
    interfaceData: {} as Interface,
    responseData: {} as Response,

    invocationsData: [],

    environmentsData: [],
    environmentData: [],
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

        setEnvironments(state, payload) {
            state.environmentsData = payload;
        },
        setEnvironment(state, data) {
            state.environmentData = data;
        },
    },
    actions: {
        async invoke({commit, dispatch, state}, payload: any) {
            invokeInterface(payload).then((json) => {
                if (json.code === 0) {
                    commit('setResponse', json.data);
                    dispatch('listInvocation', payload.id);
                    return true;
                } else {
                    return false
                }
            })
        },
        async saveInterface({commit}, payload: any) {
            saveInterface(payload).then((json) => {
                if (json.code === 0) {
                    return true;
                } else {
                    return false
                }
            })
        },

        async loadInterface({commit}) {
            const response = await load();
            if (response.code != 0) return;

            const {data} = response;
            console.log('data', data)

            commit('setTree', data || {});
            return true;
        },
        async getInterface({commit}, payload: any) {
            if (payload.isDir) {
                commit('setInterface', {});
                return true;
            }

            try {
                const response = await get(payload.id);
                const {data} = response;

                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createInterface({commit, dispatch, state}, payload: any) {
            try {
                const resp = await create(payload);

                await dispatch('loadInterface');
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async updateInterface({commit}, payload: any) {
            try {
                const {id, ...params} = payload;
                await update(id, {...params});
                return true;
            } catch (error) {
                return false;
            }
        },
        async deleteInterface({commit, dispatch, state}, payload: number) {
            try {
                await remove(payload);
                await dispatch('loadInterface');
                return true;
            } catch (error) {
                return false;
            }
        },
        async moveInterface({commit, dispatch, state}, payload: any) {
            try {
                await move(payload);
                await dispatch('loadInterface');
                return true;
            } catch (error) {
                return false;
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
                commit('setInterface', data);
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

        // environment
        async listEnvironment({commit}) {
            try {
                const resp = await listEnvironment();
                const {data} = resp;
                commit('setEnvironments', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getEnvironment({commit}, payload: any) {
            try {
                const response = await getEnvironment(payload.id, payload.interfaceId);
                const {data} = response;

                commit('setEnvironment', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveEnvironment({commit, dispatch, state}, payload: any) {
            try {
                const resp = await saveEnvironment(payload);

                dispatch('listEnvironment');
                dispatch('getEnvironment', {id: 0, interfaceId: state.interfaceData.id})
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async copyEnvironment({commit, dispatch, state}, id: number) {
            try {
                const resp = await copyEnvironment(id);

                dispatch('listEnvironment');
                dispatch('getEnvironment', {id: 0, interfaceId: state.interfaceData.id})
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async removeEnvironment({commit, dispatch, state}, id: number) {
            try {
                await removeEnvironment(id);

                dispatch('listEnvironment', state.interfaceData.id);
                dispatch('getEnvironment', {id: 0, interfaceId: state.interfaceData.id})
                return true;
            } catch (error) {
                return false;
            }
        },
        async changeEnvironment({commit, dispatch, state}, id: Number) {
            await changeEnvironment(id, state.interfaceData.id);

            dispatch('listEnvironment');
            dispatch('getEnvironment', {id: 0, interfaceId: state.interfaceData.id})
            return true
        },

        // environment var
        async saveEnvironmentVar({commit}, payload: any) {
            try {
                const resp = await saveEnvironmentVar(payload);
                const {data} = resp;
                commit('setEnvironment', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeEnvironmentVar({commit}, id: number) {
            try {
                const resp = await removeEnvironmentVar(id);
                const {data} = resp;
                commit('setEnvironment', data);

                return true;
            } catch (error) {
                return false;
            }
        },
        async clearEnvironmentVar({commit, dispatch, state}) {
            try {
                const resp = await clearEnvironmentVar(state.environmentData.id);
                const {data} = resp;
                commit('setEnvironment', data);

                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
