import {Action, Mutation} from 'vuex';
import {StoreModuleType} from "@/utils/store";
import {loadSpec} from './service';

export interface StateType {
    specData: any;
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setSpecData: Mutation<StateType>;
    };
    actions: {
        loadSpec: Action<StateType, StateType>;
    };
}

const initState: StateType = {
    specData: {info: {}},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Spec',
    state: {
        ...initState
    },
    mutations: {
        setSpecData(state, payload) {
            state.specData = payload;
        }
    },
    actions: {
        async loadSpec({commit, dispatch, state}, payload: any) {
            loadSpec(payload).then((json) => {
                if (json.code === 0) {
                    commit('setSpecData', json.data);
                    return true;
                } else {
                    return false
                }
            })
        },
    }
};

export default StoreModel;
