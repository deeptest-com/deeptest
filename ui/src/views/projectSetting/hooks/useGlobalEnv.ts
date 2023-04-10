import { computed, createVNode, reactive, ref } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { message, Modal } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import { StateType as ProjectSettingStateType } from "@/views/projectSetting/store";
import { StateType as ProjectStateType } from "@/store/project";
import { EnvReturnData, VarDataItem } from "../data";

export function useGlobalEnv(): EnvReturnData {
    const store = useStore<{ ProjectSetting: ProjectSettingStateType, ProjectGlobal: ProjectStateType }>();
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
    const envList = computed<any>(() => store.state.ProjectSetting.envList);
    const activeEnvDetail = computed<any>(() => store.state.ProjectSetting.activeEnvDetail);
    const isShowEnvDetail = ref(false);
    const isShowAddEnv = ref(false);
    const router = useRouter();

    // 请求环境列表
    async function getEnvsList() {
        console.log('%c[GET ENV LIST] --  currProject [globalEnv.ts -- 16]', 'color: red', currProject.value);
        await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id });
    }

    function showEnvDetail(item: any, isAdd?: boolean) {
        if (item) {
            router.push(`/project-setting/enviroment/envdetail/${item.id}`)
        } else {
            router.replace('/project-setting/enviroment/envdetail/-1')
        }
        
        if (isAdd) {
            store.dispatch('ProjectSetting/setEnvDetail', null);
        } else {
            item.displayName = item.name;
            store.dispatch('ProjectSetting/setEnvDetail', item);
        }
    }

    function addVar() {
        const newActiveEnvDetail = JSON.parse(JSON.stringify(activeEnvDetail.value));
        newActiveEnvDetail.vars.push({
            "name": "",
            "rightValue": "",
            "localValue": "",
            "remoteValue": "",
            // "environmentId": 7
        })
        store.dispatch('ProjectSetting/setEnvDetail', newActiveEnvDetail);
    }

    async function setShowEnvDetail(result) {
        await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id })
        const newEnv = envList.value.find((item: any) => {
            return item.id === result;
        })
        showEnvDetail(newEnv, false)
    }

    /**
     * 增加环境变量
     */
    async function addEnvData() {
        console.log('%c[ADD ENV DATA] --  envVars [globalEnv.ts -- 90]', 'color: red', activeEnvDetail.value.vars);
        if (!activeEnvDetail.value?.name) {
            return;
        }
        const envVars = activeEnvDetail.value?.vars || [];
        const hasEmptyVars = envVars.some((e: VarDataItem) => e.name === '' || e.remoteValue === '' || e.localValue === '');
        if (hasEmptyVars) {
            message.error('变量名参数/远程值/本地值不能为空');
            return;
        }
        const result = await store.dispatch('ProjectSetting/addEnvData', {
            id: activeEnvDetail.value?.id,
            projectId: currProject.value.id,
            name: activeEnvDetail.value?.name,
            "serveServers": activeEnvDetail.value?.serveServers || [],
            "vars": envVars,
        })
        if (result) {
            setShowEnvDetail(result);
        }
    }

    /**
     * 删除环境变量
     */
    async function deleteEnvData() {
        const successCallBack = async () => {
            const result = await store.dispatch('ProjectSetting/deleteEnvData', {
                activeEnvId: activeEnvDetail.value?.id,
                projectId: currProject.value.id
            })
            if (result) {
                showEnvDetail(null, true)
            }
        }
        Modal.confirm({
            title: '确认要删除该环境吗',
            icon: createVNode(ExclamationCircleOutlined),
            onOk() {
                successCallBack();
            },
        });
    }

    /**
     * 复制环境变量
     */
    async function copyEnvData() {
        const result = await store.dispatch('ProjectSetting/copyEnvData', {
            activeEnvId: activeEnvDetail.value?.id,
            projectId: currProject.value.id
        })
        if (result) {
            setShowEnvDetail(result);
        }
    }

    function handleEnvChange(type: string, field: string, index: number, e: any, action?: string) {
        if (action === 'delete') {
            activeEnvDetail.value[type].splice(index, 1);
        } else {
            activeEnvDetail.value[type][index][field] = e.target.value;
        }

    }

    function handleEnvNameChange(e: any) {
        activeEnvDetail.value.name = e.target.value;
    }

    return {
        isShowAddEnv,
        isShowEnvDetail,
        activeEnvDetail,
        getEnvsList,
        showEnvDetail,
        addVar,
        addEnvData,
        deleteEnvData,
        copyEnvData,
        handleEnvChange,
        handleEnvNameChange
    }
}
