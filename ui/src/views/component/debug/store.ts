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
    listShareVar,
    getSnippet,
    listPreConditions,
    listPostConditions,
    getScript,
    saveScript,
    removeScript,
    createPostConditions,
    createPreConditions,
    removePostConditions,
    movePostConditions,
    removePreConditions,
    movePreConditions, disablePreConditions, disablePostConditions, saveAsCase,
} from './service';
import {Checkpoint, DebugInfo, Extractor, Interface, Response, Script} from "./data";
import {UsedBy} from "@/utils/enum";
import {ResponseData} from "@/utils/request";
import {listEnvVarByServer} from "@/services/environment";

export interface StateType {
    debugInfo: DebugInfo
    debugData: any;

    requestData: any;
    responseData: Response;

    invocationsData: any[];

    preConditions: any[];
    postConditions: any[];

    extractorData: any;
    checkpointData: any;
    scriptData: any;
}
const initState: StateType = {
    debugInfo: {} as DebugInfo,
    debugData: {},

    requestData: {},
    responseData: {} as Response,

    invocationsData: [],

    preConditions: [],
    postConditions: [],

    extractorData: {} as Extractor,
    checkpointData: {} as Checkpoint,
    scriptData: {} as Script,
};

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setDebugInfo: Mutation<StateType>;
        setDebugData: Mutation<StateType>;

        setRequest: Mutation<StateType>;
        setResponse: Mutation<StateType>;

        setInvocations: Mutation<StateType>;
        setServerId: Mutation<StateType>;

        setPreConditions: Mutation<StateType>;
        setPostConditions: Mutation<StateType>;
        setExtractor: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;
        setScript: Mutation<StateType>;
        setScriptContent: Mutation<StateType>;

        setShareVars: Mutation<StateType>;
        setEnvVars: Mutation<StateType>;
        setGlobalVars: Mutation<StateType>;

        setUrl: Mutation<StateType>;
        setBaseUrl: Mutation<StateType>;
        setBody: Mutation<StateType>;
    };
    actions: {
        loadDataAndInvocations: Action<StateType, StateType>;
        resetDataAndInvocations: Action<StateType, StateType>;
        loadData: Action<StateType, StateType>;
        call: Action<StateType, StateType>;
        save: Action<StateType, StateType>;
        saveAsCase: Action<StateType, StateType>;

        listInvocation: Action<StateType, StateType>;
        getLastInvocationResp: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listPreCondition: Action<StateType, StateType>;
        createPreCondition: Action<StateType, StateType>;
        disablePreCondition: Action<StateType, StateType>;
        removePreCondition: Action<StateType, StateType>;
        movePreCondition: Action<StateType, StateType>;

        listPostCondition: Action<StateType, StateType>;
        createPostCondition: Action<StateType, StateType>;
        removePostCondition: Action<StateType, StateType>;
        disablePostCondition: Action<StateType, StateType>;
        movePostCondition: Action<StateType, StateType>;

        getExtractor: Action<StateType, StateType>;
        saveExtractor: Action<StateType, StateType>;
        createExtractorOrUpdateResult: Action<StateType, StateType>;
        removeExtractor: Action<StateType, StateType>;

        getCheckpoint: Action<StateType, StateType>;
        saveCheckpoint: Action<StateType, StateType>;
        removeCheckpoint: Action<StateType, StateType>;

        getScript: Action<StateType, StateType>;
        saveScript: Action<StateType, StateType>;
        removeScript: Action<StateType, StateType>;
        addSnippet: Action<StateType, StateType>;

        listShareVar: Action<StateType, StateType>;
        removeShareVar: Action<StateType, StateType>;
        clearShareVar: Action<StateType, StateType>;

        updateUrl: Action<StateType, StateType>;
        updateBaseUrl:Action<StateType, StateType>;
        updateBody: Action<StateType, StateType>;

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
        setRequest(state, payload) {
            state.requestData = payload;
        },
        setResponse(state, payload) {
            state.responseData = payload;
        },

        setInvocations(state, payload) {
            state.invocationsData = payload;
        },

        setPreConditions(state, payload) {
            state.preConditions = payload;
        },
        setPostConditions(state, payload) {
            state.postConditions = payload;
        },

        setExtractor(state, payload) {
            state.extractorData = payload;
        },
        setCheckpoint(state, payload) {
            state.checkpointData = payload;
        },
        setScript(state, payload) {
            state.scriptData = payload;
        },
        setScriptContent(state, content) {
            console.log('setScriptContent', content)
            state.scriptData.content = content;
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

        setUrl(state, payload) {
            state.debugData.url = payload;
        },
        setBaseUrl(state, payload) {
            state.debugData.baseUrl = payload;
        },
        setBody(state, payload) {
            state.debugData.body = payload;
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
            commit('setRequest', {});
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
        async saveAsCase({commit}, payload: any) {
            const resp = await  saveAsCase(payload)
            if (resp.code === 0) {
                // commit('setDebugData', resp.data);
                return true;
            } else {
                return false
            }
        },

        async call({commit, dispatch, state}, data: any) {
            commit('setRequest', {});
            commit('setResponse', {});
            const response = await call(data)

            if (response.code === 0) {
                commit('setRequest', response.data.req);
                commit('setResponse', response.data.resp);

                await dispatch('getLastInvocationResp')
                await dispatch('listInvocation')

                await dispatch('listShareVar');

                await dispatch('listPreCondition');
                await dispatch('listPostCondition');

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

            commit('setRequest', response.data.req);
            commit('setResponse', response.data.resp);
            return true;
        },
        async getInvocationAsInterface({commit}, id: number) {
            try {
                const resp = await getInvocationAsInterface(id);
                const {data} = resp;

                commit('setDebugData', data.debugData);
                commit('setRequest', data.req);
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

        // conditions
        async listPreCondition({commit, state}) {
            try {
                const resp = await listPreConditions(state.debugInfo.debugInterfaceId, state.debugData.endpointInterfaceId);
                const {data} = resp;
                commit('setPreConditions', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createPreCondition({commit, dispatch, state}, payload: any) {
            try {
                await createPreConditions(payload);
                dispatch('listPreCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async disablePreCondition({commit, dispatch, state}, id: number) {
            try {
                await disablePreConditions(id);
                dispatch('listPreCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removePreCondition({commit, dispatch, state}, id: number) {
            try {
                await removePreConditions(id);
                dispatch('listPreCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async movePreCondition({commit, dispatch, state}, payload: any) {
            try {
                await movePreConditions(payload);
                dispatch('listPreCondition');
                return true;
            } catch (error) {
                return false;
            }
        },

        async listPostCondition({commit, state}) {
            try {
                const resp = await listPostConditions(state.debugInfo.debugInterfaceId, state.debugData.endpointInterfaceId);
                const {data} = resp;
                commit('setPostConditions', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async createPostCondition({commit, dispatch, state}, payload: any) {
            try {
                await createPostConditions(payload);
                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async disablePostCondition({commit, dispatch, state}, id: number) {
            try {
                await disablePostConditions(id);
                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removePostCondition({commit, dispatch, state}, id: number) {
            try {
                await removePostConditions(id);
                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async movePostCondition({commit, dispatch, state}, payload: any) {
            try {
                await movePostConditions(payload);
                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },

        // extractor
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
                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },
        async createExtractorOrUpdateResult({commit, dispatch, state}, payload: any) {
            try {
                await createExtractorOrUpdateResult(payload);
                dispatch('listPostCondition');
                dispatch('listShareVar');
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeExtractor({commit, dispatch, state}, payload) {
            try {
                await removeExtractor(payload.id);

                dispatch('listPostCondition');
                return true;
            } catch (error) {
                return false;
            }
        },

        // checkpoint
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
                dispatch('listPostCondition', UsedBy.InterfaceDebug);
                return true
            } catch (error) {
                return false;
            }
        },
        async removeCheckpoint({commit, dispatch, state}, id: number) {
            try {
                await removeCheckpoint(id);

                dispatch('listPostCondition', UsedBy.InterfaceDebug);
                return true;
            } catch (error) {
                return false;
            }
        },

        // script
        async getScript({commit}, id: number) {
            try {
                const response = await getScript(id);
                const {data} = response;

                commit('setScript', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveScript({commit, dispatch, state}, payload: any) {
            try {
                await saveScript(payload);
                dispatch('listPostCondition', UsedBy.InterfaceDebug);
                return true
            } catch (error) {
                return false;
            }
        },
        async removeScript({commit, dispatch, state}, id: number) {
            try {
                await removeScript(id);

                dispatch('listPostCondition', UsedBy.InterfaceDebug);
                return true;
            } catch (error) {
                return false;
            }
        },
        async addSnippet({commit, dispatch, state}, name: string) {
            const json = await getSnippet(name)
            if (json.code === 0) {
                let script = state.scriptData.content ? state.scriptData.content: '' + '\n' +  json.data.script
                script = script.trim()
                commit('setScriptContent', script);
            }

            return true;
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

        async changeServer({commit, dispatch, state}, serverId: number) {
            console.log('changeServer in DebugStore', serverId)

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
