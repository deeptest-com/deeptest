import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import { Report, ReportLog, ReportLogDetail, QueryResult, QueryParams, PaginationConfig } from './data';
import { query, get, remove} from './service';

//定义
export interface StateType {
    ReportId: number;
    listResult: QueryResult;
    //这个report是场景的执行结果
    detailResult: Report;
    queryParams: any;
    //这个logs是场景执行结果里的logs，是个array
    reportlog: ReportLog;
    reportlogdetail: ReportLogDetail;
}

//初始化state
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
    queryParams: {},
    reportlog:{
        id:0,
        name:"",
        resultStatus:"",
        logs: [],
    },
    reportlogdetail: {
        name: "",
        resultStatus: "",
        exectime: 0,
    },
};

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setReportId: Mutation<StateType>;

        setList: Mutation<StateType>;
        setDetail: Mutation<StateType>;
        setQueryParams: Mutation<StateType>;
        setReportLog: Mutation<StateType>;
        setReportLogDetail:  Mutation<StateType>;
    };
    actions: {
        list: Action<StateType, StateType>;
        get: Action<StateType, StateType>;
        remove: Action<StateType, StateType>;
    };
}

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
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
        setReportLog(state, reportlog){
            state.reportlog = reportlog;
        },
        setReportLogDetail(state, logdetails){
            state.reportlogdetail = logdetails;
        },
    },
    actions: {
        async list({ commit, dispatch }, params: QueryParams ) {
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
                commit('setQueryParams', params);

                return true;
            } catch (error) {
                return false;
            }
        },

        async get({ commit }, id: number ) {
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
                commit('setReportLog',{
                    ...initState.detailResult.logs[0],
                    ...data.reportlog[0],
                });
                commit('setReportLogDetail',{
                    ...initState.detailResult.logs[0].logs,
                    ...data.reportlog[0].logs,
                });
                return true;
            } catch (error) {
                return false;
            }
        },


        async remove({ commit, dispatch, state }, payload: number ) {
            try {
                await remove(payload);

                await dispatch('list', state.queryParams)
                return true;
            } catch (error) {
                return false;
            }
        },

    }
};




export default StoreModel;
