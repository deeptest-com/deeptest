import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";

import {
    clearShareVar,
    createExtractorOrUpdateResult,
    get,
    getCheckpoint,
    getExtractor,
    getInvocationAsInterface,
    getLastInvocationResp,

    removeCheckpoint,
    removeExtractor,
    removeInvocation,
    removeShareVar,
    saveCheckpoint,
    saveExtractor,

    listInvocation,
    invokeInterface,
    listExtractor,
    listCheckpoint,
    listValidExtractorVariableForInterface, getSnippet, loadData,
} from './service';
import {Checkpoint, Extractor, Interface, Response} from "./data";
import {UsedBy} from "@/utils/enum";
import {ResponseData} from "@/utils/request";

export interface StateType {
    currEndpointId: number;
    currInterface: any;
    debugData: any;

    responseData: Response;

    invocationsData: any[];

    extractorsData: any[];
    extractorData: any;
    validExtractorVariablesData: any[];

    checkpointsData: any[];
    checkpointData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEndpointId: Mutation<StateType>;
        setInterface: Mutation<StateType>;
        setDebugData: Mutation<StateType>;
        setResponse: Mutation<StateType>;
        setInvocations: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setExtractor: Mutation<StateType>;
        setValidExtractorVariables: Mutation<StateType>;

        setCheckpoints: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;

        setUrl: Mutation<StateType>;
        setBody: Mutation<StateType>;
        setParam: Mutation<StateType>;
        setHeader: Mutation<StateType>;
        setPreRequestScript: Mutation<StateType>;
    };
    actions: {
        loadDebugData: Action<StateType, StateType>;
        setEndpointId: Action<StateType, StateType>;
        setInterface: Action<StateType, StateType>;

        invokeInterface: Action<StateType, StateType>;

        getLastInvocationResp: Action<StateType, StateType>;

        listInvocation: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listExtractor: Action<StateType, StateType>;
        getExtractor: Action<StateType, StateType>;
        saveExtractor: Action<StateType, StateType>;
        createExtractorOrUpdateResult: Action<StateType, StateType>;
        removeExtractor: Action<StateType, StateType>;
        removeShareVar: Action<StateType, StateType>;
        clearShareVar: Action<StateType, StateType>;
        listValidExtractorVariableForInterface: Action<StateType, StateType>;

        listCheckpoint: Action<StateType, StateType>;
        getCheckpoint: Action<StateType, StateType>;
        saveCheckpoint: Action<StateType, StateType>;
        removeCheckpoint: Action<StateType, StateType>;

        updateUrl: Action<StateType, StateType>;
        updateBody: Action<StateType, StateType>;
        updateParam: Action<StateType, StateType>;
        updateHeader: Action<StateType, StateType>;
        addSnippet: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    currEndpointId: 0,
    currInterface: {},
    debugData: {},
    responseData: {} as Response,

    invocationsData: [],

    extractorsData: [],
    extractorData: {} as Extractor,
    validExtractorVariablesData: [],

    checkpointsData: [],
    checkpointData: {} as Checkpoint,
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Debug',
    state: {
        ...initState
    },
    mutations: {
        setEndpointId(state, id) {
            state.currEndpointId = id;
        },
        setInterface(state, payload) {
            state.currInterface = payload;
        },
        setDebugData(state, payload) {
            state.debugData = payload;
        },
        setResponse(state, payload) {
            state.responseData = payload;
        },

        setInvocations(state, payload) {
            state.invocationsData = payload;
        },

        setExtractors(state, payload) {
            state.extractorsData = payload;
        },

        setExtractor(state, payload) {
            state.extractorData = payload;
        },

        setValidExtractorVariables(state, payload) {
            state.validExtractorVariablesData = payload;
        },

        setCheckpoints(state, payload) {
            state.checkpointsData = payload;
        },
        setCheckpoint(state, payload) {
            state.checkpointData = payload;
        },

        setUrl(state, payload) {
            state.debugData.url = payload;
        },
        setBody(state, payload) {
            state.debugData.body = payload;
        },
        setParam(state, payload) {
            state.debugData.params[payload.index].value = payload.value;
        },
        setHeader(state, payload) {
            console.log('setParam', payload)
            state.debugData.headers[payload.index].value = payload.value;
        },
        setPreRequestScript(state, payload) {
            console.log('setPreRequestScript', payload)
            state.debugData.preRequestScript = payload;
        },
    },
    actions: {
        async loadDebugData({commit, dispatch}, data) {
            try {
                const resp: ResponseData = await loadData(data);
                if (resp.code != 0) return false;

                commit('setDebugData', resp.data);

                return true;
            } catch (error) {
                return false;
            }
        },
        async setEndpointId({commit, dispatch}, id) {
            commit('setEndpointId', id);
        },
        async setInterface({commit, dispatch}, payload) {
            commit('setInterface', payload);
        },

        async invokeInterface({commit, dispatch, state}, data: any) {
            const response = await invokeInterface(data)
            console.log('=invoke in interface=', response.data)

            if (response.code === 0) {
                commit('setResponse', response.data);

                dispatch('listInvocation', state.currInterface.id);
                dispatch('listValidExtractorVariableForInterface');

                dispatch('listExtractor');
                dispatch('listCheckpoint');

                return true;
            } else {
                return false
            }
        },

        async getLastInvocationResp({commit, dispatch, state}, id: number) {
            const response = await getLastInvocationResp(id);
            // console.log('=getLastInvocationResp=', response.data)

            const {data} = response;

            commit('setResponse', data);
            return true;
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

                commit('setInterface', data.req);
                commit('setResponse', data.resp);
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

        // extractor
        async listExtractor({commit, dispatch, state}) {
            try {
                const resp = await listExtractor(state.debugData.id, UsedBy.InterfaceDebug);
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
                dispatch('listExtractor', UsedBy.InterfaceDebug);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createExtractorOrUpdateResult({commit, dispatch, state}, payload: any) {
            try {
                await createExtractorOrUpdateResult(payload);
                dispatch('listExtractor');
                dispatch('listValidExtractorVariableForInterface');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeExtractor({commit, dispatch, state}, payload) {
            try {
                await removeExtractor(payload.id);

                dispatch('listExtractor');
                return true;
            } catch (error) {
                return false;
            }
        },

        // extractor variable
        async clearShareVar({commit, dispatch, state}, payload: any) {
            try {
                const resp = await clearShareVar(state.debugData.id);
                const {data} = resp;
                dispatch('listValidExtractorVariableForInterface');

                return true;
            } catch (error) {
                return false;
            }
        },
        async removeShareVar({commit, dispatch, state}, payload: any) {
            try {
                const resp = await removeShareVar(payload.id);
                const {data} = resp;
                dispatch('listValidExtractorVariableForInterface');

                return true;
            } catch (error) {
                return false;
            }
        },

        async listValidExtractorVariableForInterface({commit, dispatch, state}) {
            try {
                console.log('listValidExtractorVariableForInterface')
                const resp = await listValidExtractorVariableForInterface(state.debugData.id, UsedBy.InterfaceDebug);
                const {data} = resp;
                commit('setValidExtractorVariables', data);
                return true;
            } catch (error) {
                return false;
            }
        },

        // checkpoint
        async listCheckpoint({commit, state}) {
            try {
                const resp = await listCheckpoint(state.debugData.id, UsedBy.InterfaceDebug);
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
                dispatch('listCheckpoint', UsedBy.InterfaceDebug);
                return true
            } catch (error) {
                return false;
            }
        },
        async removeCheckpoint({commit, dispatch, state}, id: number) {
            try {
                await removeCheckpoint(id);

                dispatch('listCheckpoint', UsedBy.InterfaceDebug);
                return true;
            } catch (error) {
                return false;
            }
        },

        async updateUrl({commit, dispatch, state}, url: string) {
            commit('setUrl', url);
            return true;
        },
        async updateBody({commit, dispatch, state}, body: string) {
            commit('setBody', body);
            return true;
        },
        async updateParam({commit, dispatch, state}, data: any) {
            commit('setParam', data);
            return true;
        },
        async updateHeader({commit, dispatch, state}, data: any) {
            commit('setHeader', data);
            return true;
        },
        async addSnippet({commit, dispatch, state}, name: string) {
            const json = await getSnippet(name)
            if (json.code === 0) {
                let script = state.debugData.preRequestScript + '\n' +  json.data.script
                script = script.trim()
                commit('setPreRequestScript', script);
            }

            return true;
        },
    }
};

export default StoreModel;
