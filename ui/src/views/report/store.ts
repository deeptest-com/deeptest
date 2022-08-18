import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Report, QueryResult, QueryParams, PaginationConfig } from './data';
import {
    query,
    get,
} from './service';

export interface StateType {
    ReportId: number;

    listResult: QueryResult;
    detailResult: Report;

    scenarios: any[];
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setReportId: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
    };
    actions: {
        listReport: Action<StateType, StateType>;
        getReport: Action<StateType, StateType>;
    };
}
const initState: StateType = {
    ReportId: 0,

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
    detailResult: {} as Report,

    scenarios: [],
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Report',
    state: {
        ...initState
    },
    mutations: {
        setReportId(state, id) {
            state.ReportId = id;
        },

        setList(state, payload) {
            state.listResult = payload;
        },
        setDetail(state, payload) {
            state.detailResult = payload;
        },
    },
    actions: {
        async listReport({ commit, dispatch }, params: QueryParams ) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;

                commit('setList',{
                    ...initState.listResult,
                    list: data.result || [],
                    pagination: {
                        ...initState.listResult.pagination,
                        current: params.page,
                        pageSize: params.pageSize,
                        total: data.total || 0,
                    },
                });

                return true;
            } catch (error) {
                return false;
            }
        },

        async getReport({ commit }, id: number ) {
            if (id === 0) {
                commit('setDetail',{
                    ...initState.detailResult,
                })
                return
            }
            try {
                const response: ResponseData = await get(id);
                const { data } = response;
                commit('setDetail',{
                    ...initState.detailResult,
                    ...data,
                });
                return true;
            } catch (error) {
                return false;
            }
        },

    }
};

export default StoreModel;
