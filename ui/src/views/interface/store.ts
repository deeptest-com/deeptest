import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";

import {
    create,
    get,
    invoke,
    listInvocation,
    getInvocationAsInterface,
    load,
    move,
    remove,
    removeInvocation,
    saveInterface,
    update,

    listExtractor,
    listValidExtractorVariable,
    getExtractor,
    saveExtractor,
    saveExtractorOrUpdateResult,
    removeExtractor,
    removeShareVar, clearShareVar,

    listCheckpoint,
    getCheckpoint,
    saveCheckpoint,
    removeCheckpoint, getLastInvocationResp, loadExecData, submitInvokeResult,
} from './service';
import {Checkpoint, Extractor, Interface, Response} from "@/views/interface/data";
import {getNodeMap} from "@/services/tree";

export interface StateType {
    treeData: any[];
    treeDataMap: any,
    interfaceData: Interface;
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
        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        setTreeDataMapItem: Mutation<StateType>;
        setTreeDataMapItemProp: Mutation<StateType>;

        setInterface: Mutation<StateType>;
        setResponse: Mutation<StateType>;

        setInvocations: Mutation<StateType>;

        setExtractors: Mutation<StateType>;
        setExtractor: Mutation<StateType>;
        setValidExtractorVariables: Mutation<StateType>;

        setCheckpoints: Mutation<StateType>;
        setCheckpoint: Mutation<StateType>;
    };
    actions: {
        invoke: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        saveTreeMapItemProp: Action<StateType, StateType>;

        loadInterface: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        getLastInvocationResp: Action<StateType, StateType>;
        createInterface: Action<StateType, StateType>;
        updateInterface: Action<StateType, StateType>;
        deleteInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;

        listInvocation: Action<StateType, StateType>;
        getInvocationAsInterface: Action<StateType, StateType>;
        removeInvocation: Action<StateType, StateType>;

        listExtractor: Action<StateType, StateType>;
        getExtractor: Action<StateType, StateType>;
        saveExtractor: Action<StateType, StateType>;
        saveExtractorOrUpdateResult: Action<StateType, StateType>;
        removeExtractor: Action<StateType, StateType>;
        removeShareVar: Action<StateType, StateType>;
        clearShareVar: Action<StateType, StateType>;
        listValidExtractorVariable: Action<StateType, StateType>;

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

    extractorsData: [],
    extractorData: {} as Extractor,
    validExtractorVariablesData: [],

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
    },
    actions: {
        async invoke({commit, dispatch, state}, request: any) {
            const updatedRequest = await loadExecData(request)

            const response = await invoke(updatedRequest.data)

            const submitResult = await submitInvokeResult({
                request: updatedRequest.data,
                response: response.data,
            })

            if (submitResult.code === 0) {
                commit('setResponse', response.data);

                dispatch('listInvocation', request.id);
                dispatch('listExtractor', request.id);
                dispatch('listCheckpoint', request.id);
                dispatch('listValidExtractorVariable', state.interfaceData.id);

                return true;
            } else {
                return false
            }
        },
        async saveInterface({commit}, payload: any) {
            const json = await  saveInterface(payload)
            if (json.code === 0) {
                return true;
            } else {
                return false
            }
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
                    bodyType: 'application/json',
                    headers: [{name:'', value:''}],
                    params: [{name:'', value:''}],
                    bodyFormData: [{name:'', value:'', type: 'text'}],
                    bodyFormUrlencoded: [{name:'', value:''}],
                });
                commit('setResponse', {headers: [], contentLang: 'html', content: ''});
                return true;
            }

            try {
                const response = await get(payload.id);
                const {data} = response;
                data.headers.push({name:'', value:''})
                data.params.push({name:'', value:''})
                data.bodyFormData.push({name:'', value:'', type: 'text'})
                data.bodyFormUrlencoded.push({name:'', value:''})

                commit('setInterface', data);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getLastInvocationResp({commit, dispatch, state}, id: number) {
            const response = await getLastInvocationResp(id);
            const {data} = response;

            commit('setResponse', data);
            return true;
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
        async saveExtractorOrUpdateResult({commit, dispatch, state}, payload: any) {
            try {
                await saveExtractorOrUpdateResult(payload);
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

        // extractor variable
        async removeShareVar({commit, dispatch, state}, id: any) {
            try {
                const resp = await removeShareVar(id);
                const {data} = resp;
                dispatch('listValidExtractorVariable', state.interfaceData.id);

                return true;
            } catch (error) {
                return false;
            }
        },
        async clearShareVar({commit, dispatch, state}, interfaceId: any) {
            try {
                const resp = await clearShareVar(interfaceId);
                const {data} = resp;
                dispatch('listValidExtractorVariable', state.interfaceData.id);

                return true;
            } catch (error) {
                return false;
            }
        },

        async listValidExtractorVariable({commit, dispatch, state}) {
            try {
                const resp = await listValidExtractorVariable(state.interfaceData.id);
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
