import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import {notifyError, notifySuccess} from "@/utils/notify";
import {
    listJslib,
    deleteJslib,
    disableJslib,
    getJslib,
    saveJslib, updateJsLibName
} from './service';

export interface StateType {
    jslibModels: any[];
    jslibModel: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setJslibs: Mutation<StateType>,
        setJslib: Mutation<StateType>,
    };
    actions: {
        listJslib: Action<StateType, StateType>,
        getJslib: Action<StateType, StateType>,
        saveJslib: Action<StateType, StateType>,
        updateJsLibName: Action<StateType, StateType>,
        deleteJslib: Action<StateType, StateType>,
        disableJslib: Action<StateType, StateType>,
    }
}

const initState: StateType = {
    jslibModels: [],
    jslibModel: {}
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
    }
};

export default StoreModel;
