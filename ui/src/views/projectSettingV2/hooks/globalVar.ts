import { getEnvironmentsParamList, getGlobalVarsList, saveEnvironmentsParam, saveGlobalVars } from "@/views/projectSetting/service";
import { message } from "ant-design-vue";
import { computed, ref } from "vue";
import { useStore } from "vuex";
import {StateType as ProjectSettingStateType} from "@/views/projectSettingV2/store";
import {StateType as ProjectStateType} from "@/store/project";

interface GlobaleProps {
    isShowAddEnv: any,
    isShowEnvDetail: any,
    activeEnvDetail: any,
    isShowGlobalParams: any,
    isShowGlobalVars: any,
    globalParamsActiveKey: any
}

export function useGlobalVarAndParams({ isShowAddEnv, isShowEnvDetail, activeEnvDetail, isShowGlobalParams, isShowGlobalVars, globalParamsActiveKey }: GlobaleProps) {
    const store = useStore<{ ProjectSettingV2: ProjectSettingStateType, ProjectGlobal: ProjectStateType }>();
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
    console.log('%c[GET ENV LIST] --  currProject [gloablVars.ts -- 21]', 'color: red', currProject.value);
    /**
     * 展示全局参数
     */
    async function showGlobalParams() {
        isShowGlobalParams.value = true;
        isShowGlobalVars.value = false;
        isShowAddEnv.value = false;
        isShowEnvDetail.value = false;

        activeEnvDetail.value = null;
        await store.dispatch('ProjectSettingV2/getEnvironmentsParamList', { projectId: currProject.value.id });
    }

    /**
     * 全局变量列表
     */
    async function showGlobalVars() {
        isShowGlobalParams.value = false;
        isShowGlobalVars.value = true;
        isShowAddEnv.value = false;
        isShowEnvDetail.value = false;
        console.log(currProject);
        await store.dispatch('ProjectSettingV2/getGlobalVarsList', { projectId: currProject.value.id });
    }

    /**
     * 前端模拟添加全局变量
     */
    function addGlobalVar() {
        store.dispatch('ProjectSettingV2/addGlobalVars');
    }

    /**
     * 前端模拟添加全局参数
     */
    function addGlobalParams() {
        store.dispatch('ProjectSettingV2/addGlobalParams', { globalParamsActiveKey });
    }

    /**
     * 保存全局参数
     */
    async function handleSaveGlobalParams() {
        await store.dispatch('ProjectSettingV2/saveEnvironmentsParam', { projectId: currProject.value.id });
    }

    async function handleSaveGlobalVars() {
        await store.dispatch('ProjectSettingV2/saveGlobalVars');
    }

    function handleGlobalVarsChange(field, index, e, action?: string) {
        store.dispatch('ProjectSettingV2/handleGlobalVarsChange', { field, index, e, action });
    }


    function handleGlobalParamsChange(type: string, field: string, index: number, e: any, action?: string) {
        store.dispatch('ProjectSettingV2/handleGlobalParamsChange', { type, field, index, e, action });
    }

    return {
        showGlobalParams,
        showGlobalVars,
        addGlobalVar,
        addGlobalParams,
        handleSaveGlobalParams,
        handleSaveGlobalVars,
        handleGlobalVarsChange,
        handleGlobalParamsChange
    }
}