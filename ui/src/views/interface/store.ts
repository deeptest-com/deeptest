import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";

import {
    changeEnvironment,
    clearEnvironmentVar,
    copyEnvironment,
    create,
    get,
    invoke,
    listInvocation,
    getEnvironment,
    getInvocationAsInterface,
    listEnvironment,
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

    listExtractor,
    getExtractor,
    saveExtractor,
    removeExtractor,

    listCheckpoint,
    getCheckpoint,
    saveCheckpoint,
    removeCheckpoint,
} from './service';
import {Checkpoint, Extractor, Interface, Response} from "@/views/interface/data";
import {getNodeMap} from "@/services/tree";

export interface StateType {
    treeData: any[];
    treeDataMap: any,
    interfaceData: Interface;
    responseData: Response;

    invocationsData: any[];

    environmentsData: any[];
    environmentData: any;

    extractorsData: any[];
    extractorData: any;

    checkpointsData: any[];
    checkpointData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        setTreeDataMapItem: Mutation<StateType>;
        setTreeDataMapItemProp: Mutation<StateType>;

        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;

        setInvocations: Mutation<StateType>;

        setEnvironments: Mutation<StateType>;
        setEnvironment: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setExtractor: Mutation<StateType>;

        setCheckpoints: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;
    };
    actions: {
        invoke: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        saveTreeMapItem: Action<StateType, StateType>;
        saveTreeMapItemProp: Action<StateType, StateType>;

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

        listExtractor: Action<StateType, StateType>;
        getExtractor: Action<StateType, StateType>;
        saveExtractor: Action<StateType, StateType>;
        removeExtractor: Action<StateType, StateType>;

        listCheckpoint: Action<StateType, StateType>;
        getCheckpoint: Action<StateType, StateType>;
        saveCheckpoint: Action<StateType, StateType>;
        removeCheckpoint: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    treeData: [],
    treeDataMap: {},

    interfaceData: {} as Interface,
    responseData: {} as Response,

    invocationsData: [],

    environmentsData: [],
    environmentData: [],

    extractorsData: [],
    extractorData: {} as Extractor,

    checkpointsData: [],
    checkpointData: {} as Checkpoint,
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Interface',
    state: {
        ...initState
    },
    mutations: {
        setTreeData(state, payload) {
            payload.name = '所有接口'
            state.treeData = [payload];
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
        setEnvironment(state, payload) {
            state.environmentData = payload;
        },

        setExtractors(state, payload) {
            state.extractorsData = payload;
        },
        setExtractor(state, payload) {
            state.extractorData = payload;
        },

        setCheckpoints(state, payload) {
            state.checkpointsData = payload;
        },
        setCheckpoint(state, payload) {
            state.checkpointData = payload;
        },
    },
    actions: {
        async invoke({commit, dispatch, state}, payload: any) {
            invoke(payload).then((json) => {
                if (json.code === 0) {
                    commit('setResponse', json.data);

                    dispatch('listInvocation', payload.id);
                    dispatch('listExtractor', payload.id);
                    dispatch('listCheckpoint', payload.id);

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
        async saveTreeMapItem({commit}, payload: any) {
            commit('setTreeDataMapItem', payload);
        },
        async saveTreeMapItemProp({commit}, payload: any) {
            commit('setTreeDataMapItemProp', payload);
        },

        async loadInterface({commit, dispatch, state}) {
            const response = await load();
            if (response.code != 0) return;

            const {data} = response;
            commit('setTreeData', data || {});

            const mp = {}
            getNodeMap(data, mp)
            commit('setTreeDataMap', mp);

            return true;
        },
        async getInterface({commit}, payload: any) {
            if (payload.isDir) {
                commit('setInterface', {
                    bodyType: 'application/json'
                });
                commit('setResponse', {headers: [], contentLang: 'html', content: ''});
                return true;
            }

            try {
                const response = await get(payload.id);
                const {data} = response;

                commit('setInterface', data);
                commit('setResponse', {headers: [], contentLang: 'html', content: ''});
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

        // extractor
        async listExtractor({commit, dispatch, state}) {
            try {
                const resp = await listExtractor(state.interfaceData.id);
                const {data} = resp;
                commit('setExtractors', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getExtractor({commit}, id: number) {
            try {
                const response = await getExtractor(id);
                const {data} = response;

                commit('setExtractor', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveExtractor({commit, dispatch, state}, payload: any) {
            try {
                await saveExtractor(payload);
                dispatch('listExtractor');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeExtractor({commit, dispatch, state}, id: number) {
            try {
                await removeExtractor(id);

                dispatch('listExtractor', state.interfaceData.id);
                return true;
            } catch (error) {
                return false;
            }
        },

        // checkpoint
        async listCheckpoint({commit, state}) {
            try {
                const resp = await listCheckpoint(state.interfaceData.id);
                const {data} = resp;
                commit('setCheckpoints', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getCheckpoint({commit}, id: number) {
            try {
                const response = await getCheckpoint(id);
                const {data} = response;

                commit('setCheckpoint', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveCheckpoint({commit, dispatch, state}, payload: any) {
            try {
                await saveCheckpoint(payload);
                dispatch('listCheckpoint', state.interfaceData.id);
                return true
            } catch (error) {
                return false;
            }
        },
        async removeCheckpoint({commit, dispatch, state}, id: number) {
            try {
                await removeCheckpoint(id);

                dispatch('listCheckpoint', state.interfaceData.id);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
