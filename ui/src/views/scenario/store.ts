import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Scenario, QueryResult, QueryParams, PaginationConfig } from './data.d';
import {
    query, save, remove, detail,
} from './service';

export interface StateType {
    queryResult: QueryResult;
    detailResult: Scenario;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
        setItem: Mutation<StateType>;
    };
    actions: {
        queryScenario: Action<StateType, StateType>;
        getScenario: Action<StateType, StateType>;
        saveScenario: Action<StateType, StateType>;
        removeScenario: Action<StateType, StateType>;
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
    detailResult: {} as Scenario,
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Scenario',
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
    },
    actions: {
        async queryScenario({ commit }, params: QueryParams ) {
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
                return true;
            } catch (error) {
                return false;
            }
        },
        async getScenario({ commit }, id: number ) {
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
        async saveScenario({ commit }, payload: Pick<Scenario, "name" | "desc"> ) {
            try {
                await save(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
        async removeScenario({ commit }, payload: number ) {
            try {
                await remove(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
    }
};

export default StoreModel;
