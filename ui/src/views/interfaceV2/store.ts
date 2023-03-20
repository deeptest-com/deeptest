import {Mutation, Action, useStore} from 'vuex';
import {StoreModuleType} from "@/utils/store";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import {InterfaceListReqParams, SaveInterfaceReqParams} from './data.d.ts';
import {
    copyInterface,
    deleteInterface,
    expireInterface,
    getInterfaceList,
    saveInterface,
} from './service';


import {momentUtc} from "@/utils/datetime";
import {QueryResult} from "@/views/interface/data";

export interface StateType {
    listResult: QueryResult;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
    };
    actions: {
        loadList: Action<StateType, StateType>;
        createApi: Action<StateType, StateType>;
        disabled: Action<StateType, StateType>;
        del: Action<StateType, StateType>;
        copy: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    listResult: {
        list: [],
        pagination: {
            total: 0,
            current: 1,
            pageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
        },
    },
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'InterfaceV2',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.listResult = payload;
        },
    },
    actions: {
        async loadList({commit, dispatch, state}, {currProjectId, page, pageSize}: any) {
            page = page || state.listResult.pagination.current;
            pageSize = pageSize || state.listResult.pagination.pageSize;
            const res = await getInterfaceList({
                "projectId": currProjectId,
                "page": page,
                "pageSize": pageSize,
            });
            if (res.code === 0) {
                const {result, total} = res.data;
                result.forEach((item, index) => {
                    item.index = index + 1;
                    item.key = `${index + 1}`;
                    item.updatedAt = momentUtc(item.updatedAt);
                })
                commit('setList', {
                    list: result || [],
                    pagination: {
                        ...initState.listResult.pagination,
                        "current": page,
                        "pageSize": pageSize,
                        total: total || 0,
                    },
                });
                return true;
            } else {
                return false
            }
        },
        async createApi({commit, dispatch, state}, params: any) {
            const res = await saveInterface({
                ...params
            });
            if (res.code === 0) {
                dispatch('loadList', {currProjectId: params.projectId});
            } else {
                return false
            }
        },
        async disabled({commit, dispatch, state}, payload: any) {
            const res = await expireInterface(payload.id);
            if (res.code === 0) {
                dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        async del({commit, dispatch, state}, payload: any) {
            const res = await deleteInterface(payload.id);
            if (res.code === 0) {
                dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
        async copy({commit, dispatch, state}, payload: any) {
            const res = await copyInterface(payload.id);
            if (res.code === 0) {
                dispatch('loadList', {currProjectId: payload.projectId});
            } else {
                return false
            }
        },
    }
};

export default StoreModel;
