import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';
import {
    query,
    get,
    save,
    remove,
    move,
    clone, saveDiagnoseDebugData, importInterfaces, importCurl,
} from './service';
import {serverList} from "@/views/project-settings/service";
import {genNodeMap, getNodeMap} from "@/services/tree";

export interface StateType {
    interfaceId: number;
    interfaceData: any;
    interfaceTabs: any[];

    queryParams: any;
    serveServers: [],

    treeData: any[] | null;
    treeDataMap: any,
}

export interface ModuleType extends StoreModuleType<StateType> {
    state: StateType;
    mutations: {
        setInterfaceId: Mutation<StateType>;
        setInterfaceData: Mutation<StateType>;

        setQueryParams: Mutation<StateType>;
        setServeServers: Mutation<StateType>;

        setTreeData: Mutation<StateType>;
        setTreeDataMap: Mutation<StateType>;
        changeTreeDataMapItem: Mutation<StateType>;
        changeTreeDataMapItemProp: Mutation<StateType>;

        setInterfaceTabs: Mutation<StateType>;
        updateTabName: Mutation<StateType>;
    };
    actions: {
        loadTree: Action<StateType, StateType>;
        getInterface: Action<StateType, StateType>;
        saveInterface: Action<StateType, StateType>;
        removeInterface: Action<StateType, StateType>;
        moveInterface: Action<StateType, StateType>;
        cloneInterface: Action<StateType, StateType>;

        importInterfaces: Action<StateType, StateType>;
        importCurl: Action<StateType, StateType>;

        openInterfaceTab: Action<StateType, StateType>;
        removeInterfaceTab: Action<StateType, StateType>;
        removeInterfaceTabs: Action<StateType, StateType>;

        getServeServers: Action<StateType, StateType>;
        saveDiagnoseDebugData: Action<StateType, StateType>;
    }
}

const initState: StateType = {
    interfaceId: 0,
    interfaceData: null,
    interfaceTabs: [],

    queryParams: {},
    serveServers: [],

    treeData: [],
    treeDataMap: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'DiagnoseInterface',
    state: {
        ...initState
    },
    mutations: {
        setInterfaceId(state, id) {
            state.interfaceId = id;
        },
        setInterfaceData(state, payload) {
            state.interfaceData = payload;
        },

        setTreeData(state, data) {
            state.treeData = data
        },
        setTreeDataMap(state, payload) {
            state.treeDataMap = payload
        },
        changeTreeDataMapItem(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id] = payload
        },
        changeTreeDataMapItemProp(state, payload) {
            if (!state.treeDataMap[payload.id]) return
            state.treeDataMap[payload.id][payload.prop] = payload.value
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
        setServeServers(state, payload) {
            state.serveServers = payload;
        },

        setInterfaceTabs(state, payload) {
            state.interfaceTabs = payload;
        },
        updateTabName(state, payload) {
            state.interfaceTabs.forEach(function(item) {
                console.log(item)
                if (item.id === payload.id) {
                    item.title = payload.title
                }
            });
        },
    },
    actions: {
        async loadTree({ commit, state, dispatch }, params: any) {
            try {
                const response: ResponseData = await query(params);
                if (response.code != 0) return;

                commit('setQueryParams', params);
                commit('setTreeData', response.data);

                const data = {id: 0, children: response.data} // covert arr to obj
                const mp = genNodeMap(data)
                commit('setTreeDataMap', mp);

                return true;
            } catch (error) {
                return false;
            }
        },
        async getInterface({ commit }, node: any) {
            if (!node || node.type !== 'interface') {
                commit('setInterfaceData', null)
                return
            }

            try {
                const resp: ResponseData = await get(node.id);
                const { data } = resp;
                commit('setInterfaceData', {
                    ...data,
                });
                return true;
            } catch (error) {
                return false;
            }
        },

        async saveInterface({ commit, state, dispatch }, payload: any) {
            const jsn = await save(payload)
            if (jsn.code === 0) {
                dispatch('loadTree', state.queryParams);
                commit('updateTabName', {id: payload.id, title: payload.title})
                return true;
            } else {
                return false
            }
        },
        async removeInterface({ commit, dispatch, state }, payload: any) {
            try {
                const jsn = await remove(payload.id, payload.type);
                if (jsn.code === 0) {
                    if (payload.type == 'interface') {
                        await dispatch('removeInterfaceTab', payload.id)
                    } else if (payload.type == 'dir') {
                        await dispatch('removeInterfaceTabs', payload.id)
                    }
                    await dispatch('loadTree', state.queryParams)
                    return true;
                }
                return false;
            } catch (error) {
                return false;
            }
        },
        async moveInterface({commit, dispatch, state}, payload: any) {
            try {
                await move(payload);
                dispatch('loadTree', state.queryParams);
                return true;
            } catch (error) {
                return false;
            }
        },
        async cloneInterface({ dispatch, state }, payload: number) {
            try {
                const jsn = await clone(payload);
                if (jsn.code === 0) {
                    dispatch('listInterface', state.queryParams);
                    return true;
                }
                return false;
            } catch (error) {
                return false;
            }
        },
        async importInterfaces({commit, dispatch, state}, payload: any) {
            try {
                const resp = await importInterfaces(payload);

                await dispatch('loadTree', state.queryParams);
                return resp.data;
            } catch (error) {
                return false;
            }
        },
        async importCurl({commit, dispatch, state}, payload: any) {
            try {
                const resp = await importCurl(payload);

                await dispatch('loadTree', state.queryParams);
                return resp.data;
            } catch (error) {
                return false;
            }
        },

        async getServeServers({commit}, payload: any) {
            const res = await serverList({
                serveId: payload.id
            });
            if (res.code === 0) {
                const servers = (res.data.servers || []).map((item: any) => {
                    item.label = item.environmentName;
                    item.value = item.environmentId;
                    return item;
                })
                commit('setServeServers', servers);
            } else {
                return false
            }
        },

        async openInterfaceTab({commit, dispatch, state}, payload: any) {
            await dispatch('getInterface', payload)
            if (state.interfaceData) {
                const tabs = state.interfaceTabs

                const found = state.interfaceTabs.find(function (item, index, arr) {
                    return item.id === state.interfaceData.id
                })

                if (!found) {
                    tabs.push({
                        id: state.interfaceData.id,
                        title: state.interfaceData.title,
                        type: state.interfaceData.type
                    })
                    commit('setInterfaceTabs', tabs);
                }

                commit('setInterfaceId', state.interfaceData.id);
            }
        },
        async removeInterfaceTab({commit, dispatch, state}, id: number) {
            console.log('removeInterfaceTab', id)

            const needReload = id === state.interfaceId

            let index = 0;
            state.interfaceTabs.forEach((tab, i) => {
                if (tab.id === id) {
                    index = i;
                }
            });

            const interfaceTabs = state.interfaceTabs.filter(tab => tab.id !== id);
            console.log('after remove ', interfaceTabs)
            commit('setInterfaceTabs', interfaceTabs)

            let openTab = {} as any
            if (state.interfaceTabs.length && state.interfaceId === id) { // close curr tab
                openTab = state.interfaceTabs[0]
                commit('setInterfaceId', openTab.id)
            }

            if (needReload && openTab.id) {
                dispatch('openInterfaceTab', openTab);
            }
        },
        async removeInterfaceTabs({commit, dispatch, state}, id: number) {
            const removeInterfaceIds = [] as number[]
            state.treeDataMap[id].children?.forEach((item: any) => {
                removeInterfaceIds.push(item.id)
            })

            const needReload = removeInterfaceIds.indexOf(state.interfaceId) > -1

            const interfaceTabs = state.interfaceTabs.filter(tab => removeInterfaceIds.indexOf(tab.id) == -1);
            commit('setInterfaceTabs', interfaceTabs)

            let openTab = {} as any
            if (state.interfaceTabs.length && removeInterfaceIds.indexOf(state.interfaceId) > -1) { // close curr tab
                openTab = state.interfaceTabs[0]
                commit('setInterfaceId', openTab.id)
            }

            if (needReload && openTab.id) {
                await dispatch('openInterfaceTab', openTab);
            }
        },

        async saveDiagnoseDebugData({commit}, payload: any) {
            const resp = await  saveDiagnoseDebugData(payload)
            return resp.code === 0;
        },
    }
};

export default StoreModel;
