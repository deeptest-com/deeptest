import { RoutesDataItem } from "@/utils/routes";

const UserLayoutRoutes: RoutesDataItem[] = [
    {
        title: 'user-layout.menu.login',
        path: 'login',
        component: () => import('@/views/user/login/index.vue'),
    },
    {
        title: 'user-layout.menu.register',
        path: 'register',
        component: () => import('@/views/user/register/index.vue'),
    },
    {
        title: 'user-layout.menu.forgotPassword',
        path: 'forgotPassword',
        component: () => import('@/views/user/forgotPassword/index.vue'),
    },

];

export default UserLayoutRoutes;