import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Project, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query, save, remove, detail,
} from './service';

export interface StateType {
    queryResult: QueryResult;
    detailResult: Project;
    queryParams: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
        setItem: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;
    };
    actions: {
        queryProject: Action<StateType, StateType>;
        getProject: Action<StateType, StateType>;
        saveProject: Action<StateType, StateType>;
        removeProject: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    queryResult: {
        list: [],
        pagination: {
            total: 0,
            current: 1,
            pageSize: 10,
            showSizeChanger: true,
            showQuickJumper: true,
        },
    },
    detailResult: {} as Project,
    queryParams: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Project',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.queryResult = payload;
        },
        setItem(state, payload) {
            state.detailResult = payload;
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
    },
    actions: {
        async queryProject({ commit }, params: QueryParams ) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;

                commit('setList',{
                    ...initState.queryResult,
                    list: data.result || [],
                    pagination: {
                        ...initState.queryResult.pagination,
                        current: params.page,
                        pageSize: params.pageSize,
                        total: data.total || 0,
                    },
                });
                commit('setQueryParams', params);
                return true;
            } catch (error) {
                return false;
            }
        },
        async getProject({ commit }, id: number ) {
            if (id === 0) {
                commit('setItem',{
                    ...initState.detailResult,
                })
                return
            }
            try {
                const response: ResponseData = await detail(id);
                const { data } = response;
                commit('setItem',{
                    ...initState.detailResult,
                    ...data,
                });
                return true;
            } catch (error) {
                return false;
            }
        },
        async saveProject({ commit }, payload: Pick<Project, "name" | "desc"> ) {
            try {
                await save(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeProject({ commit, dispatch, state }, payload: number ) {
            try {
                await remove(payload);
                await dispatch('queryProject', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
