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
import {listEnvVarByServer} from "@/services/environment";

export interface StateType {
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
        setDebugInfo: Mutation<StateType>;
        setDebugData: Mutation<StateType>;
        setResponse: Mutation<StateType>;
        setInvocations: Mutation<StateType>;
        setServerId: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setExtractor: Mutation<StateType>;
        setShareVars: Mutation<StateType>;
        setEnvVars: Mutation<StateType>;
        setGlobalVars: Mutation<StateType>;

        setCheckpoints: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;

        setUrl: Mutation<StateType>;
        setBaseUrl: Mutation<StateType>;
        setBody: Mutation<StateType>;
        setPreRequestScript: Mutation<StateType>;
    };
    actions: {
        loadDataAndInvocations: Action<StateType, StateType>;
        resetDataAndInvocations: Action<StateType, StateType>;
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
        updateBaseUrl:Action<StateType, StateType>;
        updateBody: Action<StateType, StateType>;
        addSnippet: Action<StateType, StateType>;

        changeServer: Action<StateType, StateType>;
    };
}

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Debug',
    state: {
        ...initState
    },
    mutations: {
        setDebugInfo(state, payload) {
            state.debugInfo = payload;
        },
        setServerId(state, id) {
            state.debugData.serverId = id;
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
        setEnvVars(state, payload) {
            state.debugData.envVars = payload;
        },
        setGlobalVars(state, payload) {
            state.debugData.globalVars = payload;
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
        setBaseUrl(state, payload) {
            state.debugData.baseUrl = payload;
        },
        setBody(state, payload) {
            state.debugData.body = payload;
        },
        setPreRequestScript(state, payload) {
            console.log('setPreRequestScript', payload)
            state.debugData.preRequestScript = payload;
        },
    },
    actions: {
        // debug
        async loadDataAndInvocations({commit, dispatch, state}, data) {
            try {
                await dispatch('loadData', data)

                dispatch('getLastInvocationResp', {
                    debugInterfaceId: state.debugInfo.debugInterfaceId,
                    endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                    diagnoseInterfaceId: state.debugInfo.diagnoseInterfaceId,
                    caseInterfaceId: state.debugInfo.caseInterfaceId,
                })
                dispatch('listInvocation', {
                    debugInterfaceId: state.debugInfo.debugInterfaceId,
                    endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                    diagnoseInterfaceId: state.debugInfo.diagnoseInterfaceId,
                    caseInterfaceId: state.debugInfo.caseInterfaceId,
                })

                return true;
            } catch (error) {
                return false;
            }
        },
        async resetDataAndInvocations({commit, dispatch, state}) {
            commit('setDebugInfo', {});
            commit('setDebugData', {});
            commit('setResponse', {});
            commit('setInvocations', []);
        },

        async loadData({commit, state, dispatch}, data) {
            try {
                const resp: ResponseData = await loadData(data);
                if (resp.code != 0) return false;

                await commit('setDebugInfo', {
                    debugInterfaceId: resp.data.debugInterfaceId,
                    endpointInterfaceId: data.endpointInterfaceId,
                    scenarioProcessorId  : data.scenarioProcessorId,
                    diagnoseInterfaceId  : data.diagnoseInterfaceId,
                    caseInterfaceId: data.caseInterfaceId,
                    usedBy:          data.usedBy,
                } as DebugInfo);
                console.log('set debugInfo', state.debugInfo)

                await commit('setDebugData', resp.data);

                return true;
            } catch (error) {
                return false;
            }
        },
        async save({commit}, payload: any) {
            const resp = await  save(payload)
            if (resp.code === 0) {
                commit('setDebugData', resp.data);

                return true;
            } else {
                return false
            }
        },

        async call({commit, dispatch, state}, data: any) {

            // 发送时请求时，先清空response
            // 比如，手动清空 response.content中的内容后，再次点击发送，还是现实空的response.content

            commit('setResponse', {});
            const response = await call(data)

            if (response.code === 0) {
                commit('setResponse', response.data);

                await dispatch('getLastInvocationResp')
                await dispatch('listInvocation')

                await dispatch('listShareVar');

                await dispatch('listExtractor');
                await dispatch('listCheckpoint');

                return true;
            } else {
                return false
            }
        },

        // invocation
        async listInvocation({commit, state}) {
            try {
                const resp = await listInvocation({
                    debugInterfaceId: state.debugInfo.debugInterfaceId,
                    endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                    diagnoseInterfaceId: state.debugInfo.diagnoseInterfaceId,
                    caseInterfaceId: state.debugInfo.caseInterfaceId,
                } as DebugInfo);
                const {data} = resp;
                commit('setInvocations', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getLastInvocationResp({commit, dispatch, state}) {
            const response = await getLastInvocationResp( {
                debugInterfaceId: state.debugInfo.debugInterfaceId,
                endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                diagnoseInterfaceId: state.debugInfo.diagnoseInterfaceId,
                caseInterfaceId: state.debugInfo.caseInterfaceId,
            } as DebugInfo);

            const {data} = response;
            console.log('getLastInvocationResp', data)

            commit('setResponse', data);
            return true;
        },
        async getInvocationAsInterface({commit}, id: number) {
            try {
                const resp = await getInvocationAsInterface(id);
                const {data} = resp;

                commit('setDebugData', data.debugData);
                commit('setResponse', data.resp);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeInvocation({commit, dispatch, state}, id: number) {
            try {
                await removeInvocation(id);
                dispatch('listInvocation', {
                    endpointInterfaceId: state.debugInfo.endpointInterfaceId,
                });
                return true;
            } catch (error) {
                return false;
            }
        },

        // extractor
        async listExtractor({commit, dispatch, state}) {
            try {
                const resp = await listExtractor(state.debugInfo.debugInterfaceId, state.debugInfo.endpointInterfaceId);
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

        // checkpoint
        async listCheckpoint({commit, state}) {
            try {
                const resp = await listCheckpoint(state.debugInfo.debugInterfaceId, state.debugData.endpointInterfaceId);
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
                console.log('debugInfo', state.debugInfo)

                await clearShareVar(state.debugInfo);
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
        async updateUrl({commit, dispatch, state}, body: string) {
            commit('setUrl', body);
            return true;
        },
        async updateBaseUrl({commit, dispatch, state}, body: string) {
            commit('setBaseUrl', body);
            return true;
        },
        async updateBody({commit, dispatch, state}, body: string) {
            commit('setBody', body);
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

        async changeServer({commit, dispatch, state}, serverId: number) {
            const json = await listEnvVarByServer(serverId)
            if (json.code === 0) {
                commit('setServerId', serverId);
                commit('setEnvVars', json.data);
            }

            return true;
        },
    }
};

export default StoreModel;
