import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {  QueryResult,QueryParams } from './data.d';
import {
    query
} from './service';

export interface StateType {
    queryResult: QueryResult;
 
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
      
    };
    actions: {
        queryRanking: Action<StateType, StateType>;
       
      
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
  
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'workbench',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.queryResult = payload;
        },
     
        
        
    },
    actions: {
        async queryRanking({ commit }, params: QueryParams ) {
            console.log('~~~~~~params',params)
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;

                commit('setList',{
                    ...initState.queryResult,
                    list: data.user_ranking_list || [],
                    pagination: {
                        ...initState.queryResult.pagination,
                        current: params.page,
                        pageSize: params.pageSize,
                        // total: data.project_total || 0,
                    },
                });
                // commit('setQueryParams', params);
                return true;
            } catch (error) {
                return false;
            }
        },
    
     
 
    }
};

export default StoreModel;
