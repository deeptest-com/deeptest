import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import {
    copyEnv,
    copyServe,
    copySchema,
    deleteEnv,
    deleteServe,
    deleteSchema,
    deleteSecurity,
    disableServe,
    example2schema,
    getEnvironmentsParamList,
    getEnvList,
    getGlobalVarsList,
    getServeList,
    getUserList,
    getSchemaList,
    saveEnv,
    saveEnvironmentsParam,
    saveGlobalVars,
    saveServe,
    saveSchema,
    schema2example,
    getSecurityList
} from './service';
import { message } from 'ant-design-vue';
import { BasicSchemaParams, ParamsChangeState, SaveSchemaReqParams, SchemaListReqParams, VarsChangeState ,SecurityListReqParams} from './data';
import { serveStatus, serveStatusTagColor } from '@/config/constant';
import { momentUtc } from '@/utils/datetime';

export interface StateType {
    envList: any;
    serviceOptions: any;
    globalVarsData: any;
    globalParamsData: any;
    userListOptions: any;
    schemaList: any;
    securityList: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEnvsList: Mutation<StateType>,
        setGlobalVarsList: Mutation<StateType>,
        setGlobalParamsList: Mutation<StateType>,
        setServersList: Mutation<StateType>,
        setUserList: Mutation<StateType>,
        setSchemaList: Mutation<StateType>,
        setSecurityList: Mutation<StateType>,
        
    };
    actions: {
        getEnvsList: Action<StateType, StateType>,
        getServersList: Action<StateType, StateType>,
        addEnvData: Action<StateType, StateType>,
        deleteEnvData: Action<StateType, StateType>,
        copyEnvData: Action<StateType, StateType>,
        getEnvironmentsParamList: Action<StateType, StateType>,
        getGlobalVarsList: Action<StateType, StateType>,
        saveEnvironmentsParam: Action<StateType, StateType>,
        saveGlobalVars: Action<StateType, StateType>,
        addGlobalParams: Action<StateType, StateType>,
        addGlobalVars: Action<StateType, StateType>,
        handleGlobalVarsChange: Action<StateType, StateType>,
        handleGlobalParamsChange: Action<StateType, StateType>,
        getUserOptionsList: Action<StateType, StateType>,
        saveStoreServe: Action<StateType, StateType>,
        deleteStoreServe: Action<StateType, StateType>,
        copyStoreServe: Action<StateType, StateType>,
        disabledStoreServe: Action<StateType, StateType>,
        getSchemaList: Action<StateType, StateType>,
        copySchema: Action<StateType, StateType>,
        deleteSchema: Action<StateType, StateType>, 
        saveSchema: Action<StateType, StateType>,
        generateSchema: Action<StateType, StateType>,
        generateExample: Action<StateType, StateType>,
        getSecurityList: Action<StateType, StateType>,
        deleteSecurity: Action<StateType, StateType>,
        
    }
}

const initState: StateType = {
    envList: [],
    serviceOptions: [],
    globalParamsData: {
        header: [],
        cookie: [],
        body: [],
        query: []
    },
    globalVarsData: [],
    userListOptions: [],
    schemaList: [],
    securityList:[]
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'ProjectSetting',
    state: {
        ...initState
    },
    mutations: {
        setEnvsList(state, payload) {
            state.envList = payload;
        },
        setGlobalParamsList(state, payload) {
            state.globalParamsData = payload;
        },
        setGlobalVarsList(state, payload) {
            state.globalVarsData = payload;
        },
        setServersList(state, payload) {
            state.serviceOptions = payload;
        },
        setUserList(state, payload) {
            state.userListOptions = payload;
        },
        setSchemaList(state, payload) {
            state.schemaList = payload;
        },
        setSecurityList(state, payload) {
            state.securityList = payload;
        },
    },
    actions: {
        async getEnvsList({ commit, dispatch, state }, { projectId }: { projectId: number | string }) {
            const res = await getEnvList({
                projectId
            });
            res.data.forEach((item) => {
                item.displayName = item.name;
            })
            if (res.code === 0) {
                commit('setEnvsList', res.data);
                return true;
            } else {
                return false;
            }
        },
        async getServersList({ commit, dispatch, state }, { projectId, page, pageSize, name }: any) {
            const res = await getServeList({
                projectId,
                page: page || 0,
                pageSize: pageSize || 100,
                name: name || ''
            });
            if (res.code === 0) {
                res.data.result.forEach((item: any) => {
                    item.label = item.name;
                    item.value = item.id;
                    item.statusDesc = serveStatus.get(item.status);
                    item.statusTag = serveStatusTagColor.get(item.status);
                    item.createdAt = momentUtc(item.createdAt)
                    item.updatedAt = momentUtc(item.updatedAt)
                })
                commit('setServersList', res.data.result);
                return true;
            } else {
                return false;
            }
        },
        async addEnvData({ commit, dispatch, state }, { id, projectId, name, serveServers, vars }: any) {
            const res = await saveEnv({
                id,
                projectId,
                name,
                serveServers,
                vars,
            });
            if (res.code === 0) {
                message.success('保存环境成功');
                dispatch('getEnvsList', { projectId });
                return true;
            } else {
                return false;
            }
        },
        async deleteEnvData({ dispatch }, { activeEnvId, projectId }: any) {
            const res = await deleteEnv({
                id: activeEnvId,
            });
            if (res.code === 0) {
                message.success('删除环境成功');
                dispatch('getEnvsList', { projectId });
                return true;
            }
        },
        async copyEnvData({ commit, dispatch, state }, { activeEnvId, projectId }: any) {
            const res = await copyEnv({
                id: activeEnvId,
            });
            if (res.code === 0) {
                message.success('复制环境成功');
                return res.data;
            } else {
                return false;
            }
        },
        async getEnvironmentsParamList({ commit, state }, { projectId }: any) {
            const res = await getEnvironmentsParamList({
                projectId
            });
            if (res.code === 0) {
                const paramsData = res.data;
                if (paramsData.projectId) {
                    delete paramsData.projectId;
                }
                commit('setGlobalParamsList', paramsData);
                return true;
            } else {
                return false;
            }
        },
        async getGlobalVarsList({ commit, state, dispatch }, { projectId }: any) {
            const res = await getGlobalVarsList({
                projectId
            });
            if (res.code === 0) {
                commit('setGlobalVarsList', res.data);
                return true;
            } else {
                return false;
            }
        },
        async saveEnvironmentsParam({ commit, state }, { projectId }) {
            // 校验
            try {
                Object.keys(state.globalParamsData).forEach(key => {
                    const hasEmptyParams = state.globalParamsData[key].some(e => e.name === '');
                    if (hasEmptyParams) {
                        throw Error('全局参数不能为空');
                    }
                });
                const res = await saveEnvironmentsParam({ ...state.globalParamsData, projectId });
                if (res.code === 0) {
                    message.success('保存全局参数成功');
                    return true;
                } else {
                    return false;
                }
            } catch (e: any) {
                message.error(e.message);
                return false;
            }
        },
        async saveGlobalVars({ commit, state }) {
            // 校验
            const hasEmptyVarsData = state.globalVarsData.some(e => e.name === '');
            if (hasEmptyVarsData) {
                message.error('全局变量不能为空');
                return false;
            }
            const res = await saveGlobalVars(state.globalVarsData);
            if (res.code === 0) {
                message.success('保存全局变量成功');
                return true;
            } else {
                return false;
            }
        },
        addGlobalVars({ commit, state }) {
            const globalVarsData = state.globalVarsData;
            globalVarsData.push({
                "name": "",
                "rightValue": "",
                "localValue": "",
                "remoteValue": ""
            });
            commit('setGlobalVarsList', globalVarsData);
        },
        addGlobalParams({ commit, state }, { globalParamsActiveKey }) {
            const globalParamsData = state.globalParamsData;
            globalParamsData[globalParamsActiveKey.value].push({
                "name": "",
                "type": "string",
                "defaultValue": "",
                "description": "",
                "required": false
            })
            commit('setGlobalParamsList', globalParamsData);
        },
        handleGlobalVarsChange({ commit, state }, { field, index, e, action }: VarsChangeState) {
            // 删除
            const storeGlobalVarsData = JSON.parse(JSON.stringify(state.globalVarsData));
            if (action && action === 'delete') {
                storeGlobalVarsData.splice(index, 1);
            } else {
                storeGlobalVarsData[index][field] = e.target.value;
            }
            commit('setGlobalVarsList', JSON.parse(JSON.stringify(storeGlobalVarsData)));
        },
        handleGlobalParamsChange({ commit, state }, { type, field, index, e, action }: ParamsChangeState) {
            const storeGlobalParamsData = JSON.parse(JSON.stringify(state.globalParamsData));
            if (action === 'delete') {
                storeGlobalParamsData[type].splice(index, 1);
            } else {
                storeGlobalParamsData[type][index][field] = ["string", "boolean"].includes(typeof e) ? e : e.target.value;
            }
            commit('setGlobalParamsList', JSON.parse(JSON.stringify(storeGlobalParamsData)));
        },
        async getUserOptionsList({ commit }, params: any) {
            const res = await getUserList('');
            if (res.code === 0) {
                res.data.result.forEach((item) => {
                    item.label = item.name;
                    item.value = item.id
                })
                commit('setUserList', res.data.result);
            }
        },
        async saveStoreServe({ commit, state, dispatch }, params: any) {
            const res = await saveServe(params);
            if (res.code === 0) {
                message.success('新建服务成功');
                await dispatch('getServersList', {
                    projectId: params.projectId
                })
            } else {
                message.error('新建服务失败');
            }
        },
        async deleteStoreServe({ dispatch }, params: any) {
            const res = await deleteServe(params.id);
            if (res.code === 0) {
                message.success('删除成功');
                await dispatch('getServersList', {
                    projectId: params.projectId
                })
            } else {
                message.error('删除失败');
            }
        },
        async copyStoreServe({ dispatch }, params: any) {
            const res = await copyServe(params.id);
            if (res.code === 0) {
                message.success('复制服务成功');
                await dispatch('getServersList', {
                    projectId: params.projectId
                })
            } else {
                message.error('复制服务失败');
            }
        },
        async disabledStoreServe({ dispatch }, params: any) {
            const res = await disableServe(params.id);
            if (res.code === 0) {
                message.success('禁用服务成功');
                await dispatch('getServersList', {
                    projectId: params.projectId
                })
            } else {
                message.error('禁用服务失败');
            }
        },
        async getSchemaList({ commit }, params: SchemaListReqParams) {
            const reqParams = { ...params, page: 1, pageSize: 20 };
            const res = await getSchemaList(reqParams);
            if (res.code === 0) {
                console.log('%c getSchemaList request success===== sucessData', 'color: red', res);
                commit('setSchemaList', res.data.result);
            } else {
                console.log('%c getSchemaList request failed===== failedData', 'color: green', res);
            }
        },
        
        async deleteSchema({ dispatch }, data: BasicSchemaParams) {
            const { id, serveId, name } = data;
            const res = await deleteSchema(id);
            if (res.code === 0) {
                message.success('删除成功');
                await dispatch('getSchemaList', { serveId, name });
            } else {
                message.error('删除失败');
            }
        },
        async copySchema({ dispatch }, params: BasicSchemaParams) {
            const { id, serveId, name } = params;
            const res = await copySchema(id);
            if (res.code === 0) {
                message.success('复制成功');
                await dispatch('getSchemaList', { serveId, name });
            } else {
                message.error('复制失败');
            }
        },
        async saveSchema({ dispatch }, data: SaveSchemaReqParams) {
            const { schemaInfo, action, serveId, name } = data;
            const tips = { delete: '删除', update: '修改' };
            const res = await saveSchema(schemaInfo);
            if (res.code === 0) {
                message.success(`${tips[action]}组件成功`);
                return true;
            } else {
                message.error(`${tips[action]}组件失败`);
                return false;
            }
        },
        async generateSchema({ dispatch }, { data }: BasicSchemaParams) {
            const res = await example2schema({ data });
            if (res.code === 0) {
                return res.data;
            }
            return null;
        },
        async generateExample({ dispatch }, { data }: BasicSchemaParams) {
            const res = await schema2example({ data });
            if (res.code === 0) {
                return res.data;
            }
            return null;
        },

        async getSecurityList({ commit }, params: SecurityListReqParams) {
            const reqParams = { ...params, page: 1, pageSize: 20 };
            const res = await getSecurityList(reqParams);
            if (res.code === 0) {
                console.log('%c getSecurityList request success===== sucessData', 'color: red', res);
                commit('setSecurityList', res.data.result);
            } else {
                console.log('%c getSecurityList request failed===== failedData', 'color: green', res);
            }
        },
        async deleteSecurity({ dispatch }, data: SecurityListReqParams) {
            const { id, serveId } = data;
            const res = await deleteSecurity(id);
            if (res.code === 0) {
                message.success('删除成功');
                await dispatch('getSecurityList', { serveId });
            } else {
                message.error('删除失败');
            }
        },
    }
};

export default StoreModel;
