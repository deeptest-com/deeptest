import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";

import {
    clearShareVar,
    createExtractorOrUpdateResult,
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
    loadData,
    call,
    save,
    listInvocation,
    listExtractor,
    listCheckpoint,
    listShareVar, getSnippet,
} from './service';
import {Checkpoint, DebugInfo, Extractor, Interface, Response} from "./data";
import {UsedBy} from "@/utils/enum";
import {ResponseData} from "@/utils/request";

export interface StateType {
    defineEndpoint: any,
    defineInterface: any,
    debugInfo: DebugInfo
    debugData: any;

    responseData: Response;

    invocationsData: any[];

    extractorsData: any[];
    extractorData: any;

    checkpointsData: any[];
    checkpointData: any;
}
const initState: StateType = {
    defineEndpoint: {},
    defineInterface: {},
    debugInfo: {} as DebugInfo,
    debugData: {},
    responseData: {} as Response,

    invocationsData: [],

    extractorsData: [],
    extractorData: {} as Extractor,

    checkpointsData: [],
    checkpointData: {} as Checkpoint,
};

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setDefineEndpoint: Mutation<StateType>;
        setDefineInterface: Mutation<StateType>;
        setDebugInfo: Mutation<StateType>;

        setDebugData: Mutation<StateType>;
        setResponse: Mutation<StateType>;
        setInvocations: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setExtractor: Mutation<StateType>;
        setShareVars: Mutation<StateType>;

        setCheckpoints: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;

        setUrl: Mutation<StateType>;
        setBody: Mutation<StateType>;
        setParam: Mutation<StateType>;
        setHeader: Mutation<StateType>;
        setPreRequestScript: Mutation<StateType>;
    };
    actions: {
        setDefineEndpoint: Action<StateType, StateType>;
        setDefineInterface: Action<StateType, StateType>;

        loadData: Action<StateType, StateType>;
        call: Action<StateType, StateType>;
        save: Action<StateType, StateType>;

        listInvocation: Action<StateType, StateType>;
        getLastInvocationResp: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listExtractor: Action<StateType, StateType>;
        getExtractor: Action<StateType, StateType>;
        saveExtractor: Action<StateType, StateType>;
        createExtractorOrUpdateResult: Action<StateType, StateType>;
        removeExtractor: Action<StateType, StateType>;

        listShareVar: Action<StateType, StateType>;
        removeShareVar: Action<StateType, StateType>;
        clearShareVar: Action<StateType, StateType>;

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

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Debug',
    state: {
        ...initState
    },
    mutations: {
        setDefineEndpoint (state, payload) {
            state.defineEndpoint = payload;
        },
        setDefineInterface(state, payload) {
            state.defineInterface = payload;
        },
        setDebugInfo(state, payload) {
            state.debugInfo = payload;
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

        setShareVars(state, payload) {
            state.debugData.shareVars = payload;
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
        async setDefineEndpoint({commit, dispatch}, id) {
            commit('setDefineEndpoint', id);
        },
        async setDefineInterface({commit, dispatch}, id) {
            commit('setDefineInterface', id);
        },

        // debug
        async loadData({commit, dispatch}, data) {
            try {
                const resp: ResponseData = await loadData(data);
                if (resp.code != 0) return false;

                commit('setDebugData', resp.data);

                commit('setDebugInfo', {
                    endpointInterfaceId: resp.data.endpointInterfaceId,
                    debugInterfaceId: resp.data.debugInterfaceId,
                    usedBy:          resp.data.usedBy,
                    processorId  : resp.data.processorId,
                } as DebugInfo);

                return true;
            } catch (error) {
                return false;
            }
        },
        async call({commit, dispatch, state}, data: any) {
            const response = await call(data)

            if (response.code === 0) {
                commit('setResponse', response.data);

                dispatch('listInvocation', {
                    endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                    debugInterfaceId: state.debugInfo.debugInterfaceId,
                });
                dispatch('listShareVar');

                dispatch('listExtractor');
                dispatch('listCheckpoint');

                return true;
            } else {
                return false
            }
        },
        async save({commit}, payload: any) {
            const json = await  save(payload)
            if (json.code === 0) {
                return true;
            } else {
                return false
            }
        },

        async listInvocation({commit}, info: any) {
            try {
                const resp = await listInvocation(info);
                const {data} = resp;
                commit('setInvocations', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getLastInvocationResp({commit, dispatch, state}, payload: any) {
            const response = await getLastInvocationResp(payload);

            const {data} = response;

            commit('setResponse', data);
            return true;
        },
        async getInvocationAsInterface({commit}, id: number) {
            try {
                const resp = await getInvocationAsInterface(id);
                const {data} = resp;

                commit('setDebugData', data.req);
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
                dispatch('listShareVar');
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

        // shared variable
        async listShareVar({commit, dispatch, state}) {
            try {
                const resp = await listShareVar(state.debugInfo, UsedBy.InterfaceDebug);
                const {data} = resp;
                commit('setShareVars', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async clearShareVar({commit, dispatch, state}, payload: any) {
            try {
                let id = 0
                if (state.debugInfo.usedBy === UsedBy.InterfaceDebug)
                    id = state.debugInfo.endpointInterfaceId ?
                        state.debugInfo.endpointInterfaceId : state.debugInfo.debugInterfaceId
                else if (state.debugInfo.usedBy === UsedBy.ScenarioDebug)
                    id = state.debugInfo.processorId
                else
                    return false

                await clearShareVar(id, state.debugInfo.usedBy);
                dispatch('listShareVar');

                return true;
            } catch (error) {
                return false;
            }
        },
        async removeShareVar({commit, dispatch, state}, payload: any) {
            try {
                const resp = await removeShareVar(payload.id);
                const {data} = resp;
                dispatch('listShareVar');

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
