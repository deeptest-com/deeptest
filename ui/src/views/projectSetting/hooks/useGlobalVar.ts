import { computed, createVNode } from "vue";
import { Modal } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import { useStore } from "vuex";
import { StateType as ProjectSettingStateType } from "@/views/projectSetting/store";
import { StateType as ProjectStateType } from "@/store/project";
import { GlobalVarsProps, VarsReturnData } from "../data";

export function useGlobalVarAndParams(props: GlobalVarsProps): VarsReturnData {
    const { isShowAddEnv, isShowEnvDetail, activeEnvDetail, isShowGlobalParams, isShowGlobalVars, globalParamsActiveKey } = props;
    const store = useStore<{ ProjectSetting: ProjectSettingStateType, ProjectGlobal: ProjectStateType }>();
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
        await store.dispatch('ProjectSetting/getEnvironmentsParamList', { projectId: currProject.value.id });
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
        await store.dispatch('ProjectSetting/getGlobalVarsList', { projectId: currProject.value.id });
    }

    /**
     * 前端模拟添加全局变量
     */
    function addGlobalVar() {
        store.dispatch('ProjectSetting/addGlobalVars');
    }

    /**
     * 前端模拟添加全局参数
     */
    function addGlobalParams(data: { globalParamsActiveKey: string }) {
        store.dispatch('ProjectSetting/addGlobalParams', data);
    }

    /**
     * 保存全局参数
     */
    async function handleSaveGlobalParams() {
        await store.dispatch('ProjectSetting/saveEnvironmentsParam', { projectId: currProject.value.id });
    }

    async function handleSaveGlobalVars() {
        await store.dispatch('ProjectSetting/saveGlobalVars');
    }

    function handleGlobalVarsChange(field: string, index: number, e: any, action?: string) {
        const confirmCallBack = () => store.dispatch('ProjectSetting/handleGlobalVarsChange', { field, index, e, action });
        if (action && action === 'delete') {
            Modal.confirm({
                title: '确认要删除该全局变量吗',
                icon: createVNode(ExclamationCircleOutlined),
                onOk() {
                    confirmCallBack()
                },
            });
        } else {
            confirmCallBack();
        }
    }


    function handleGlobalParamsChange(type: string, field: string, index: number, e: any, action?: string) {
        const confirmCallBack = () => store.dispatch('ProjectSetting/handleGlobalParamsChange', { type, field, index, e, action });
        if (action && action === 'delete') {
            Modal.confirm({
                title: '确认要删除该参数吗',
                icon: createVNode(ExclamationCircleOutlined),
                onOk() {
                    confirmCallBack();
                },
            });
        } else {
            confirmCallBack();
        }
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
