import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {  QueryResult,QueryParams } from './data.d';
import {
    query
} from './service';

export interface StateType {
    queryResult: QueryResult;
    mode:string,
    loading:boolean
 
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setList: Mutation<StateType>;
        setMode: Mutation<StateType>;
        setLoading: Mutation<StateType>;
      
    };
    actions: {
        queryProject: Action<StateType, StateType>;
        changemode:Action<StateType, StateType>;
      
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
    mode:'list',
    loading:false
  
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Home',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.queryResult = payload;
        },
        
        setMode(state, payload) {
            console.log('~~~~~~~~~setMode',state,payload)
            state.mode = payload;
        },
        setLoading(state, payload) {
            state.loading = payload;
        },
     
        
        
    },
    actions: {
        async queryProject({ commit }, params: QueryParams ) {
            console.log('~~~~~~params',params)
           
            commit('setLoading',{
                loading:true
            })
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                const data = response.data;
              
                commit('setList',{
                    ...initState.queryResult,
                    list: data || [],
                 
                });
                commit('setLoading',{
                    loading:false
                })
                return true;
            } catch (error) {
                commit('setLoading',{
                    loading:false
                })
                return false;
            }
        },
         changemode({ commit }, params: any ) {

                  commit('setMode', params.mode);
        }
    
     
 
    }
};

export default StoreModel;
