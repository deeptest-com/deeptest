<template>
    <a-result v-if="(!isLogin && loading) || !isReady">
        <template #icon>
            <a-spin size="large" />
        </template>
    </a-result>
    <router-view v-if="isLogin" />
</template>
<script lang="ts">
import { computed, ComputedRef, defineComponent, onMounted, Ref, ref, unref, watch } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { StateType as UserStateType, CurrentUser } from "@/store/user";
import { RouteMenuType } from "@/types/permission";
import { notifyError } from "@/utils/notify";

interface SecurityLayoutSetupData {
    isLogin: ComputedRef<boolean>;
    loading: Ref<boolean>;
    getUser: () => Promise<void>;
    isReady: Ref<boolean>;
}

export default defineComponent({
    name: 'SecurityLayout',
    setup(): SecurityLayoutSetupData {
        const router = useRouter();
        const store = useStore<{User: UserStateType}>();

        // 获取当前登录用户信息
        const currentUser = computed<CurrentUser>(()=> store.state.User.currentUser);

        // 判断是否登录
        const isLogin = computed<boolean>(()=> currentUser.value ? currentUser.value.id > 0 : false);

        // 读取当前用户信息func
        const isReady = ref<boolean>(false); // 是否读取过用户信息
        const loading = ref<boolean>(false);
        const getUser = async () => {
            loading.value = true;
            await store.dispatch('User/fetchCurrent');
            loading.value = false;
            if(!isLogin.value && router.currentRoute.value.path !== '/user/login') {
                router.replace({
                    path: '/user/login',
                    query: {
                        redirect: router.currentRoute.value.path,
                        ...router.currentRoute.value.query
                    }
                })
                return;
            }
            
            const { params: { projectNameAbbr }, meta: { code } } = router.currentRoute.value;
            // 查看具体项目页面时刷新才会校验 权限按钮以及权限路由
            if (projectNameAbbr) {
                // 校验项目权限
                const result = await store.dispatch('ProjectGlobal/checkProjectAndUser', { project_code: projectNameAbbr });
                if (!result) {
                    isReady.value = true;
                    router.replace('/');
                    return;
                }

                // 获取权限路由
                const { menuData } = await store.dispatch('Global/getPermissionList', { projectId: result.id });

                // 校验路由权限
                if (!menuData[RouteMenuType[`${code}`]]) {
                    isReady.value = true;
                    notifyError('权限不足');
                    router.replace('/');
                    return;
                }
            }

            isReady.value = true;
        }

        onMounted(() => {
            getUser().catch((err)=>{
                console.log('getUser',err)

            });
        })

        watch(() => {
            return unref(isReady);
        }, val => {
            if (!val) return;
            const appLoadingEl = document.getElementsByClassName('app-loading');
            if (appLoadingEl[0]) {
                appLoadingEl[0].classList.add('hide');
                setTimeout(() => {
                    document.body.removeChild(appLoadingEl[0]);
                }, 600);
            }
        })

        return {
            isLogin,
            loading,
            getUser,
            isReady
        }


    }
})
</script>