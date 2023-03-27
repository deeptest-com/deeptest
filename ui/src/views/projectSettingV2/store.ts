import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import {
    copyEnv,
    deleteEnv,
    getEnvironmentsParamList,
    getEnvList,
    getGlobalVarsList,
    getServeList,
    getUserList,
    saveEnv,
    saveEnvironmentsParam,
    saveGlobalVars,
    saveServe
} from './service';
import { message } from 'ant-design-vue';
import { ParamsChangeState, VarsChangeState } from './data';
import { serveStatus, serveStatusTagColor } from '@/config/constant';
import {momentUtc} from '@/utils/datetime';

export interface StateType {
    envList: any;
    serviceOptions: any;
    globalVarsData: any;
    globalParamsData: any;
    userListOptions: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setEnvsList: Mutation<StateType>,
        setGlobalVarsList: Mutation<StateType>,
        setGlobalParamsList: Mutation<StateType>,
        setServersList: Mutation<StateType>,
        setUserList: Mutation<StateType>,
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
        createServe: Action<StateType, StateType>,
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
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'ProjectSettingV2',
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
        }
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
        async createServe({ commit, state, dispatch }, { projectId, formState }: any) {
            const res = await saveServe({
                "projectId": projectId,
                "name": formState.name,
                "description": formState.description,
                "userId": formState.userId,
            });
            if (res.code === 0) {
                message.success('新建服务成功');
                await dispatch('getServersList', {
                    projectId
                })
            } else {
                message.error('新建服务失败');
            }
        }

    }
};

export default StoreModel;
