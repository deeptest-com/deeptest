import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import {notifyError, notifySuccess} from "@/utils/notify";
import {
    listJslib, deleteJslib, disableJslib, getJslib, saveJslib, updateJsLibName,
    listAgent, deleteAgent, disableAgent, getAgent, saveAgent, updateAgentName
} from './service';

export interface StateType {
    jslibModels: any[];
    jslibModel: any;

    agentModels: any[];
    agentModel: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setJslibs: Mutation<StateType>,
        setJslib: Mutation<StateType>,

        setAgents: Mutation<StateType>,
        setAgent: Mutation<StateType>,
    };
    actions: {
        listJslib: Action<StateType, StateType>,
        getJslib: Action<StateType, StateType>,
        saveJslib: Action<StateType, StateType>,
        updateJsLibName: Action<StateType, StateType>,
        deleteJslib: Action<StateType, StateType>,
        disableJslib: Action<StateType, StateType>,

        listAgent: Action<StateType, StateType>,
        getAgent: Action<StateType, StateType>,
        saveAgent: Action<StateType, StateType>,
        updateAgentName: Action<StateType, StateType>,
        deleteAgent: Action<StateType, StateType>,
        disableAgent: Action<StateType, StateType>,
    }
}

const initState: StateType = {
    jslibModels: [],
    jslibModel: {},

    agentModels: [],
    agentModel: {}
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'SysSetting',
    state: {
        ...initState
    },
    mutations: {
        setJslibs(state, payload) {
            state.jslibModels = payload;
        },
        setJslib(state, payload) {
            state.jslibModel = payload;
        },

        setAgents(state, payload) {
            state.agentModels = payload;
        },
        setAgent(state, payload) {
            state.agentModel = payload;
        },
    },
    actions: {
        async listJslib({ commit }, params) {
           const res = await listJslib(params)
            if (res.code === 0) {
                commit('setJslibs', res.data);
                return true;
            } else {
                return false;
            }
        },
        async getJslib({ commit, dispatch }, id: number) {
            const res = await getJslib(id);
            if (res.code === 0) {
                commit('setJslib', res.data);
            } else {
                notifyError(`获取自定义脚本库失败`);
            }
        },
        async saveJslib({ dispatch }, data) {
            const res = await saveJslib(data);

            if (res.code === 0) {
                notifySuccess('保存成功');
                dispatch('listJslib')
            } else {
                notifyError('删除自定义脚本库失败');
            }
            return res.msgKey
        },
        async updateJsLibName({ dispatch }, data) {
            const res = await updateJsLibName(data);

            if (res.code === 0) {
                dispatch('listJslib')
            } else {
                notifyError('修改自定义脚本库名称失败');
            }
            return res.msgKey
        },
        async deleteJslib({ dispatch }, id) {
            const res = await deleteJslib(id);
            if (res.code === 0) {
                notifySuccess('删除自定义脚本库成功');
                dispatch('listJslib')
            } else {
                notifyError('删除自定义脚本库失败');
            }
        },
        async disableJslib({ dispatch }, id) {
            const res = await disableJslib(id);
            if (res.code === 0) {
                dispatch('listJslib')
            } else {
                notifyError('删除自定义脚本库失败');
            }
        },

        async listAgent({ commit }, params) {
            const res = await listAgent(params)
            if (res.code === 0) {
                commit('setAgents', res.data);
                return true;
            } else {
                return false;
            }
        },
        async getAgent({ commit, dispatch }, id: number) {
            const res = await getAgent(id);
            if (res.code === 0) {
                commit('setAgent', res.data);
            } else {
                notifyError(`获取自定义脚本库失败`);
            }
        },
        async saveAgent({ dispatch }, data) {
            const res = await saveAgent(data);

            if (res.code === 0) {
                notifySuccess('保存成功');
                dispatch('listAgent')
            } else {
                notifyError('删除自定义脚本库失败');
            }
            return res.msgKey
        },
        async updateAgentName({ dispatch }, data) {
            const res = await updateAgentName(data);

            if (res.code === 0) {
                dispatch('listAgent')
            } else {
                notifyError('修改自定义脚本库名称失败');
            }
            return res.msgKey
        },
        async deleteAgent({ dispatch }, id) {
            const res = await deleteAgent(id);
            if (res.code === 0) {
                notifySuccess('删除自定义脚本库成功');
                dispatch('listAgent')
            } else {
                notifyError('删除自定义脚本库失败');
            }
        },
        async disableAgent({ dispatch }, id) {
            const res = await disableAgent(id);
            if (res.code === 0) {
                dispatch('listAgent')
            } else {
                notifyError('删除自定义脚本库失败');
            }
        },
    }
};

export default StoreModel;
