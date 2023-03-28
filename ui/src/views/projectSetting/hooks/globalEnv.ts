import { computed, ref } from "vue";
import { useStore } from "vuex";
import {StateType as ProjectSettingStateType} from "@/views/ProjectSetting/store";
import {StateType as ProjectStateType} from "@/store/project";
import { message } from "ant-design-vue";

export function useGlobalEnv({ isShowGlobalParams, isShowGlobalVars }) {
    const store = useStore<{ ProjectSetting: ProjectSettingStateType, ProjectGlobal: ProjectStateType }>();
    const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
    const envList = computed<any>(() => store.state.ProjectSetting.envList);
    const isShowEnvDetail = ref(false);
    const isShowAddEnv = ref(false);
    const activeEnvDetail: any = ref(null);

    // 请求环境列表
    async function getEnvsList() {
        console.log('%c[GET ENV LIST] --  currProject [globalEnv.ts -- 16]', 'color: red', currProject.value);
        await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id });
    }

    function showEnvDetail(item:any, isAdd?: boolean) {
        if (isAdd) {
            isShowAddEnv.value = true;
            isShowEnvDetail.value = true;
            activeEnvDetail.value = {
                displayName: "新建环境",
                name: "",
                serveServers: [],
                vars: [],
            };
        } else {
            isShowEnvDetail.value = true;
            isShowAddEnv.value = false;
            activeEnvDetail.value = item;
            activeEnvDetail.value.name = item.name || '';
            activeEnvDetail.value.displayName = item.name || '';
        }
        isShowGlobalParams.value = false;
        isShowGlobalVars.value = false;
    }

    function addVar() {
        activeEnvDetail.value.vars.push({
            "name": "",
            "rightValue": "",
            "localValue": "",
            "remoteValue": "",
            // "environmentId": 7
        })
    }

    /**
     * 增加环境变量
     */
    async function addEnvData() {
        console.log('%c[ADD ENV DATA] --  envVars [globalEnv.ts -- 90]', 'color: red', activeEnvDetail.value.vars);
        const envVars = activeEnvDetail.value?.vars || [];
        const hasEmptyVars = envVars.some((e: any) => e.name === '');
        if (hasEmptyVars) {
            message.error('变量名参数不能为空');
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
            showEnvDetail(null, true)
        }
    }

    /**
     * 删除环境变量
     */
    async function deleteEnvData() {
        const result = await store.dispatch('ProjectSetting/deleteEnvData', {
            activeEnvId: activeEnvDetail.value?.id,
            projectId: currProject.value.id
        })
        if (result) {
            showEnvDetail(null, true)
        }
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
            await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id })
            const newEnv = envList.value.find((item: any) => {
                return item.id === result;
            })
            showEnvDetail(newEnv, false)
        }
    }

    /**
     * 切换环境
     * @param type 
     * @param field 
     * @param index 
     * @param e 
     * @param action 
     */
    function handleEnvChange(type, field, index, e, action?:any) {
        if (action === 'delete') {
            activeEnvDetail.value[type].splice(index, 1);
        } else {
            activeEnvDetail.value[type][index][field] = e.target.value;
        }

    }

    function handleEnvNameChange(e) {
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