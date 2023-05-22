import { computed, createVNode, reactive, ref } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { Modal } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import { StateType as ProjectSettingStateType } from "@/views/projectSetting/store";
import { StateType as ProjectStateType } from "@/store/project";
import { EnvReturnData } from "../data";

export function useGlobalEnv(formRef?: any): EnvReturnData {
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
            router.replace('/project-setting/enviroment/envdetail')
        }

        if (isAdd) {
            store.dispatch('ProjectSetting/setEnvDetail', null);
        } else {
            item.displayName = item.name;
            store.dispatch('ProjectSetting/setEnvDetail', item);
        }
    }

    function addVar() {
        store.dispatch('ProjectSetting/addEnvVar', {
            "name": "",
            "rightValue": "",
            "localValue": "",
            "remoteValue": "",
            // "environmentId": 7
        });
    }

    async function setShowEnvDetail(result: string | number, needRedirect: boolean) {
        await store.dispatch('ProjectSetting/getEnvsList', { projectId: currProject.value.id })
        if (needRedirect) {
            const newEnv = envList.value.find((item: any) => {
                return item.id === result;
            })
            showEnvDetail(newEnv, false)
        }
    }

    /**
     * 增加环境变量
     */
    async function addEnvData() {
        try {
            const success = await formRef.value.validateFields();
            console.log('addEnvData validate success---', success);
            const envVars = activeEnvDetail.value?.vars || [];
            const result = await store.dispatch('ProjectSetting/addEnvData', {
                id: activeEnvDetail.value?.id,
                projectId: currProject.value.id,
                name: activeEnvDetail.value?.name,
                "serveServers": activeEnvDetail.value?.serveServers || [],
                "vars": envVars,
            })
            if (result) {
                setShowEnvDetail(result, true);
            }
        } catch (err) {
            console.log('addEnvData validate validateFiled--', err);
        }
    }

    /**
     * 删除环境变量
     */
    async function deleteEnvData(env: any) {
        const successCallBack = async () => {
            const result = await store.dispatch('ProjectSetting/deleteEnvData', {
                activeEnvId: env.id,
                projectId: currProject.value.id
            })
            // 如果删除的环境id和当前选中环境id一样，则删除才会跳转新建环境页面
            if (result && env.id === activeEnvDetail.value.id) {
                // showEnvDetail(null, true)
                const oldEnvList = [...envList.value];
                const index = oldEnvList.findIndex(e => {
                    return e.id === env.id;
                });
                const newIndex = oldEnvList.length - 1 > index ? index + 1 : 0;
                const newEnv = oldEnvList[newIndex];
                showEnvDetail(newEnv, false);
            }
        }
        Modal.confirm({
            title: '确认要删除该环境吗',
            icon: createVNode(ExclamationCircleOutlined),
            okText:'确定',
            cancelText:'取消',
            onOk() {
                successCallBack();
            },
        });
    }

    /**
     * 复制环境变量
     */
    async function copyEnvData(env: any) {
        const result = await store.dispatch('ProjectSetting/copyEnvData', {
            activeEnvId: env.id,
            projectId: currProject.value.id
        })
        // 如果复制的原环境id和当前选中环境id一样，则跳转到复制以后的target环境页面上
        if (result) {
            setShowEnvDetail(result, true);
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
