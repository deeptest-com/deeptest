import {Mutation, Action, useStore} from 'vuex';
import {StoreModuleType} from "@/utils/store";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import {InterfaceListReqParams, SaveInterfaceReqParams} from './data.d.ts';
import {
    copyInterface,
    deleteInterface,
    expireInterface, getInterfaceDetail,
    getInterfaceList,
    saveInterface,

} from './service';
import {momentUtc} from "@/utils/datetime";
import {QueryResult} from "@/views/interface/data";
import {filterFormState} from "./data";
import {loadCategory} from "@/services/category";
import {getNodeMap} from "@/services/tree";

export interface StateType {
    listResult: QueryResult;
    filterState: filterFormState;
    interFaceCategoryOpt: any,
    interfaceDetail:any,
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
        setFilterState: Mutation<StateType>;
        setInterfaceCategory: Mutation<StateType>;
        setInterfaceDetail: Mutation<StateType>;
    };
    actions: {
        loadList: Action<StateType, StateType>;
        createApi: Action<StateType, StateType>;
        disabled: Action<StateType, StateType>;
        del: Action<StateType, StateType>;
        copy: Action<StateType, StateType>;
        loadCategory: Action<StateType, StateType>;
        getInterfaceDetail: Action<StateType, StateType>;
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
    filterState: {
        "status": null,
        "createUser": null,
        "title": null,
    },
    interFaceCategoryOpt: [],
    interfaceDetail:null,
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
        setFilterState(state, payload) {
            state.filterState = payload;
        },
        setInterfaceCategory(state, payload) {
            state.interFaceCategoryOpt = payload;
        },
        setInterfaceDetail(state, payload) {
            state.interfaceDetail = payload;
        },
    },
    actions: {
        async loadList({commit, dispatch, state}, {currProjectId, page, pageSize, opts}: any) {
            page = page || state.listResult.pagination.current;
            pageSize = pageSize || state.listResult.pagination.pageSize;
            const otherParams = {...state.filterState, ...opts};

            console.log(832, 832, opts, state.filterState)
            const res = await getInterfaceList({
                "projectId": currProjectId,
                "page": page,
                "pageSize": pageSize,
                ...otherParams,
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
                commit('setFilterState', {
                    ...otherParams
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
        // 用于新建接口时选择接口分类
        async loadCategory({commit}) {
            const res = await loadCategory('interface');
            if (res.code === 0) {
                commit('setInterfaceCategory', res.data || null);
            } else {
                return false
            }
        },

        // 用于新建接口时选择接口分类
        async getInterfaceDetail({commit},payload: any) {
            const res = await getInterfaceDetail(payload.id);
            res.data.createdAt = momentUtc(res.data.createdAt);
            res.data.updatedAt = momentUtc(res.data.updatedAt);
            if (res.code === 0) {
                commit('setInterfaceDetail', res.data || null);
            } else {
                return false
            }
        },
    }
};

export default StoreModel;
