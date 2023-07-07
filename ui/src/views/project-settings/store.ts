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
    getSecurityList,
    getServeVersionList,
    saveServeVersion,
    deleteServeVersion,
    disableServeVersions,
    sortEnv, listDatapool, saveDatapool, deleteDatapool, disableDatapool, getDatapool
} from './service';
import { message } from 'ant-design-vue';
import {
    BasicSchemaParams,
    SecurityListReqParams,
    EnvDataItem,
    EnvReqParams,
    ParamsChangeState,
    SaveSchemaReqParams,
    SaveVersionParams,
    SchemaListReqParams,
    ServeDetail,
    ServeListParams,
    ServeReqParams,
    StoreServeParams,
    VarsChangeState,
    VarsReqParams,
    DatapoolListParams, DatapoolReqParams, StoreDatapoolParams
} from './data';
import {disabledStatus, disabledStatusTagColor, serveStatus, serveStatusTagColor} from '@/config/constant';
import { momentUtc } from '@/utils/datetime';

export interface StateType {
    envList: any;
    serviceOptions: any;
    globalVarsData: any;
    globalParamsData: any;
    userListOptions: any;
    schemaList: any;
    securityList: any;
    datapoolList: any;
    datapoolDetail: any;

    selectServiceDetail: any;
    serveVersionsList: any;
    activeEnvDetail: any;
    selectEnvId: number | null;
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
        setServiceDetail: Mutation<StateType>,
        setVersionList: Mutation<StateType>,
        setEnvDetail: Mutation<StateType>,
        setSelectEnvId: Mutation<StateType>

        setDatapoolList: Mutation<StateType>,
        setDatapoolDetail: Mutation<StateType>,
    };
    actions: {
        // 环境-全局变量-全局参数相关
        getEnvsList: Action<StateType, StateType>,
        sortEnvList: Action<StateType, StateType>,
        getServersList: Action<StateType, StateType>,

        addEnvData: Action<StateType, StateType>,
        saveEnvId: Action<StateType, StateType>,
        deleteEnvData: Action<StateType, StateType>,
        copyEnvData: Action<StateType, StateType>,
        setEnvDetail: Action<StateType, StateType>,
        addEnvServe: Action<StateType, StateType>,
        addEnvVar: Action<StateType, StateType>,
        getEnvironmentsParamList: Action<StateType, StateType>,
        getGlobalVarsList: Action<StateType, StateType>,
        saveEnvironmentsParam: Action<StateType, StateType>,
        saveGlobalVars: Action<StateType, StateType>,
        addGlobalParams: Action<StateType, StateType>,
        addGlobalVars: Action<StateType, StateType>,
        handleGlobalVarsChange: Action<StateType, StateType>,
        handleGlobalParamsChange: Action<StateType, StateType>,
        // 用户相关
        getUserOptionsList: Action<StateType, StateType>,
        // 服务相关
        saveStoreServe: Action<StateType, StateType>,
        deleteStoreServe: Action<StateType, StateType>,
        copyStoreServe: Action<StateType, StateType>,
        disabledStoreServe: Action<StateType, StateType>,
        // 服务组件相关
        getSchemaList: Action<StateType, StateType>,
        copySchema: Action<StateType, StateType>,
        deleteSchema: Action<StateType, StateType>,
        saveSchema: Action<StateType, StateType>,
        generateSchema: Action<StateType, StateType>,
        generateExample: Action<StateType, StateType>,
        // security相关
        getSecurityList: Action<StateType, StateType>,
        deleteSecurity: Action<StateType, StateType>,

        setServiceDetail: Action<StateType, StateType>,
        // 服务版本相关
        getVersionList: Action<StateType, StateType>,
        deleteVersion: Action<StateType, StateType>,
        disabledVersion: Action<StateType, StateType>,
        saveVersion: Action<StateType, StateType>

        // 数据池相关
        listDatapool: Action<StateType, StateType>,
        getDatapool: Action<StateType, StateType>,
        saveDatapool: Action<StateType, StateType>,
        deleteDatapool: Action<StateType, StateType>,
        disableDatapool: Action<StateType, StateType>,
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
    securityList:[],
    datapoolList: [],
    datapoolDetail: [],
    selectServiceDetail: {
        name: '',
        description: '',
        serveId: ''
    },
    serveVersionsList: [],
    activeEnvDetail: {
        displayName: "新建环境",
        name: "",
        serveServers: [],
        vars: [],
    },
    selectEnvId: null
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
        setServiceDetail(state, payload) {
            state.selectServiceDetail = payload;
        },
        setVersionList(state, payload) {
            state.serveVersionsList = payload;
        },
        setEnvDetail(state, payload) {
            state.activeEnvDetail = payload;
        },
        setSelectEnvId(state, payload) {
            state.selectEnvId = payload;
        },
        setDatapoolList(state, payload) {
            state.datapoolList = payload;
        },
        setDatapoolDetail(state, payload) {
            state.datapoolDetail = payload;
        },
    },
    actions: {
        async getEnvsList({ commit }, { projectId }: EnvReqParams) {
            const res = await getEnvList({
                projectId
            });
            res.data.forEach((item) => {
                item.displayName = item.name;
                item.label = item.name;
                item.value = item.id;
            })
            if (res.code === 0) {
                commit('setEnvsList', res.data);
                commit('setSelectEnvId', res.data[0].id);
                return true;
            } else {
                return false;
            }
        },
        async saveEnvId({ commit }, payload: number) {
            commit('setSelectEnvId', payload);
            return true;
        },
        async sortEnvList({ dispatch }, { data, projectId }: { data: number[], projectId: string | number }) {
            const res = await sortEnv(data);
            if (res.code === 0) {
                dispatch('getEnvsList', { projectId });
            }
        },
        async getServersList({ commit }, { projectId, page, pageSize, name }: ServeListParams) {
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
        async addEnvData({ dispatch }, { id, projectId, name, serveServers, vars }: EnvDataItem) {
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
                return res.data;
            } else {
                return false;
            }
        },
        async deleteEnvData({ dispatch }, { activeEnvId, projectId }: EnvReqParams) {
            const res = await deleteEnv({
                id: activeEnvId,
            });
            if (res.code === 0) {
                message.success('删除环境成功');
                dispatch('getEnvsList', { projectId });
                return true;
            }
        },
        async copyEnvData({ commit, dispatch, state }, { activeEnvId }: EnvReqParams) {
            const res = await copyEnv({
                id: activeEnvId,
            });
            if (res.code === 0) {
                console.log(res);
                message.success('复制环境成功');
                return res.data;
            } else {
                return false;
            }
        },
        async getEnvironmentsParamList({ commit, state }, { projectId }: VarsReqParams) {
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
        async getGlobalVarsList({ commit, state, dispatch }) {
            const res = await getGlobalVarsList();
            if (res.code === 0) {
                commit('setGlobalVarsList', res.data);
                return true;
            } else {
                return false;
            }
        },
        async saveEnvironmentsParam({ state }, { projectId }: VarsReqParams) {
            const res = await saveEnvironmentsParam({ ...state.globalParamsData, projectId });
            if (res.code === 0) {
                message.success('保存全局参数成功');
                return true;
            } else {
                return false;
            }
        },
        async saveGlobalVars({ state }) {
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
            if (!globalParamsData[globalParamsActiveKey.value])
                globalParamsData[globalParamsActiveKey.value] = []
            globalParamsData[globalParamsActiveKey.value].push({
                "name": "",
                "type": "string",
                "defaultValue": "",
                "description": "",
                "required": false
            })
            commit('setGlobalParamsList', globalParamsData);
        },
        handleGlobalVarsChange({ commit, state }, payload: VarsChangeState) {
            const { field, index, e, action } = payload;
            // 删除
            const storeGlobalVarsData = JSON.parse(JSON.stringify(state.globalVarsData));
            if (action && action === 'delete') {
                storeGlobalVarsData.splice(index, 1);
            } else {
                storeGlobalVarsData[index][field] = e.target.value;
            }
            commit('setGlobalVarsList', JSON.parse(JSON.stringify(storeGlobalVarsData)));
        },
        handleGlobalParamsChange({ commit, state }, payload: ParamsChangeState) {
            const { type, field, index, e, action } = payload;
            const storeGlobalParamsData = JSON.parse(JSON.stringify(state.globalParamsData));
            if (action === 'delete') {
                storeGlobalParamsData[type].splice(index, 1);
            } else {
                storeGlobalParamsData[type][index][field] = ["string", "boolean"].includes(typeof e) ? e : e.target.value;
            }
            commit('setGlobalParamsList', JSON.parse(JSON.stringify(storeGlobalParamsData)));
        },
        async getUserOptionsList({ commit }) {
            const res = await getUserList();
            if (res.code === 0) {
                res.data.result.forEach((item) => {
                    item.label = item.username;
                    item.value = item.username
                })
                commit('setUserList', res.data.result);
            }
        },
        async saveStoreServe({ dispatch }, params: StoreServeParams) {
            const { formState, projectId, action = 'create' } = params;
            const tips = { 'create': '新建服务', 'update': '修改服务' };
            const res = await saveServe({ ...formState, projectId });
            if (res.code === 0) {
                message.success(`${tips[action]}成功`);
                await dispatch('getServersList', {
                    projectId
                })
            } else {
                message.error(`${tips[action]}失败`);
            }
        },
        async deleteStoreServe({ dispatch }, params: ServeReqParams) {
            const { id, projectId } = params;
            const res = await deleteServe(id);
            if (res.code === 0) {
                message.success('删除成功');
                await dispatch('getServersList', {
                    projectId
                })
            } else {
                message.error('删除失败');
            }
        },
        async copyStoreServe({ dispatch }, params: ServeReqParams) {
            const { id, projectId } = params;
            const res = await copyServe(id);
            if (res.code === 0) {
                message.success('复制服务成功');
                await dispatch('getServersList', {
                    projectId
                })
            } else {
                message.error('复制服务失败');
            }
        },
        async disabledStoreServe({ dispatch }, params: ServeReqParams) {
            const { id, projectId } = params;
            const res = await disableServe(id);
            if (res.code === 0) {
                message.success('禁用服务成功');
                await dispatch('getServersList', {
                    projectId
                })
            } else {
                message.error('禁用服务失败');
            }
        },

        async getSchemaList({ commit }, params: SchemaListReqParams) {
            // const reqParams = { ...params, page: 1, pageSize: 20 };
            const res = await getSchemaList(params);
            if (res.code === 0) {
                commit('setSchemaList', res.data.result);
                return res;
            } else {
                message.error('获取schema列表失败');
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
            const { schemaInfo, action } = data;
            const tips = { create: '新建', update: '修改' };
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
        async generateExample({ dispatch }, { data,serveId }: BasicSchemaParams) {
            const res = await schema2example({ data,serveId });
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
        setServiceDetail({ commit }, payload: ServeDetail) {
            commit('setServiceDetail', payload);
        },
        async getVersionList({ commit, state }) {
            const res = await getServeVersionList({
                "serveId": state.selectServiceDetail.id,
                "createUser": "",
                "version": "",
                "page": 1,
                "pageSize": 100
            });
            if (res.code === 0) {
                commit('setVersionList', res.data.result);
                return true;
            }
            return false;
        },
        async saveVersion({ commit }, payload: SaveVersionParams) {
            const res = await saveServeVersion(payload);
            if (res.code === 0) {
                commit('setVersionList', res.data.result);
                return true;
            }
            return false;
        },
        async deleteVersion({ commit }, id: string | number) {
            const res = await deleteServeVersion(id);
            if (res.code === 0) {
                commit('getVersionList');
                return true;
            }
            return false;
        },
        async disabledVersion({ commit }, id: string | number) {
            const res = await disableServeVersions(id);
            if (res.code === 0) {
                commit('getVersionList');
                return true;
            }
            return false;
        },
        async setEnvDetail({ commit }, envData: any) {
            const initEnvData = {
                displayName: "新建环境",
                name: "",
                serveServers: [],
                vars: []
            };
            commit('setEnvDetail', envData || initEnvData);
        },
        addEnvServe({ commit, state }, serveData: any) {
            const newEnvDetail = JSON.parse(JSON.stringify(state.activeEnvDetail));
            newEnvDetail.serveServers.push(serveData);
            commit('setEnvDetail', newEnvDetail);
        },
        addEnvVar({ commit, state }, varData: any) {
            const newEnvDetail = JSON.parse(JSON.stringify(state.activeEnvDetail));
            newEnvDetail.vars.push(varData);
            commit('setEnvDetail', newEnvDetail);
        },

        async listDatapool({ commit }, { projectId, page, pageSize, name }: DatapoolListParams) {
            const res = await listDatapool({
                projectId,
                page: page || 0,
                pageSize: pageSize || 100,
                name: name || ''
            });
            if (res.code === 0) {
                res.data.result.forEach((item: any) => {
                    item.label = item.name;
                    item.value = item.id;
                    item.statusDesc = disabledStatus.get(item.disabled ? 1 : 0);
                    item.statusTag = disabledStatusTagColor.get(item.disabled ? 1 : 0);
                    item.createdAt = momentUtc(item.createdAt)
                    item.updatedAt = momentUtc(item.updatedAt)
                })
                console.log('res.data.result', res.data.result)
                commit('setDatapoolList', res.data.result);
                return true;
            } else {
                return false;
            }
        },
        async getDatapool({ commit, dispatch }, id: number) {
            const res = await getDatapool(id);
            if (res.code === 0) {
                commit('setDatapoolDetail', res.data);
            } else {
                message.error(`获取数据池失败`);
            }
        },
        async saveDatapool({ dispatch }, params: StoreDatapoolParams) {
            const { formState, projectId, action = 'create' } = params;
            const tips = { 'create': '新建数据池', 'update': '修改数据池' };

            const res = await saveDatapool({ ...formState, projectId });

            if (res.code === 0) {
                message.success(`${tips[action]}成功`);
                await dispatch('listDatapool', {
                    projectId
                })
            } else {
                message.error(`${tips[action]}失败`);
            }
        },
        async deleteDatapool({ dispatch }, params: DatapoolReqParams) {
            const { id, projectId } = params;
            const res = await deleteDatapool(id);
            if (res.code === 0) {
                message.success('删除数据池成功');
                await dispatch('listDatapool', {
                    projectId
                })
            } else {
                message.error('删除数据池失败');
            }
        },
        async disableDatapool({ dispatch }, params: ServeReqParams) {
            const { id, projectId } = params;
            const res = await disableDatapool(id);
            if (res.code === 0) {
                await dispatch('listDatapool', {
                    projectId
                })
            } else {
                message.error('禁用数据池失败');
            }
        },
    }
};

export default StoreModel;
