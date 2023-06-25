/**
 * 路由入口
 * @author LiQingSong
 */
import NProgress from 'nprogress'; // progress bar
import 'nprogress/nprogress.css'; // progress bar style
NProgress.configure({ showSpinner: false, easing: 'ease', speed: 1000 }); // NProgress Configuration

import { createRouter, createWebHashHistory } from 'vue-router';
import { RoutesDataItem } from "@/utils/routes";
import settings from "@/config/settings";

import SecurityLayout from '@/layouts/SecurityLayout.vue';

import IndexLayoutRoutes from '@/layouts/IndexLayout/routes';
import IndexLayout from '@/layouts/IndexLayout/index.vue';

import UserLayoutRoutes from '@/layouts/UserLayout/routes';
import UserLayout from '@/layouts/UserLayout/index.vue';
import BlankLayout from "@/layouts/BlankLayout.vue";
import HomeLayout from "@/layouts/HomeLayout.vue";

const routes: RoutesDataItem[] = [
  {
    title: 'empty',
    path: '/mock',
    component: BlankLayout,
    children: [
      {
        title: 'mock.oauth2.callback',
        path: 'oauth2/callback',
        component: () => import('@/views/mock/oauth2-callback.vue'),
        hidden: true,
      },
    ],
  },
  {
    title: 'empty',
    path: '/docs',
    component: BlankLayout,
    children: [
      {
        title: '接口文档',
        path: 'share',
        component: () => import('@/views/docs/index.vue'),
        hidden: true,
      },
    ],
  },
  {
    title: 'empty',
    path: '/',
    component: SecurityLayout,
    children: [
      {
        title: '首页',
        path: '/',
        redirect: '/home',
        component: HomeLayout,
        children: [
          {
            icon: 'home',
            title: 'home',
            path: 'home',
            component: () => import('@/views/home/index.vue'),
            hidden: false,
          }
        ],
      },
      {
        title: 'index-layout.menu',
        path: '/user-manage',
        component: HomeLayout,
        children: [
          {
            icon: 'profile',
            title: 'profile',
            path: 'profile',
            component: () => import('@/views/user/info/profile.vue'),
            hidden: true,
          },
          {
            icon: 'message',
            title: 'message',
            path: 'message',
            component: () => import('@/views/user/info/message.vue'),
            hidden: true,
          },
          {
            icon: 'user',
            title: 'user.management',
            path: 'index',
            component: () => import('@/views/user/list/index.vue'),
            hidden: true
          }
        ],
      },
    ]
  },
  {
    title: 'empty',
    path: '/project',
    component: SecurityLayout,
    children: [
      {
        title: 'empty',
        path: '/project',
        redirect: settings.homeRouteItem.path,
        component: IndexLayout,
        children: IndexLayoutRoutes
      },
      {
        title: 'empty',
        path: '/refresh',
        component: () => import('@/views/refresh/index.vue'),
      },
    ]
  },
  {
    title: 'empty',
    path: '/user',
    component: UserLayout,
    children: UserLayoutRoutes
  },
  {
    title: 'app.global.menu.notfound',
    path: '/:pathMatch(.*)*',
    component: () => import('@/views/404/index.vue'),
  },
]

const router = createRouter({
  scrollBehavior(/* to, from, savedPosition */) {
    return { top: 0 }
  },
  history: createWebHashHistory(process.env.BASE_URL),
  routes: routes,
});

router.beforeEach((/* to, from */) => {
  // start progress bar
  NProgress.start();
});

router.afterEach(() => {
  // finish progress bar
  NProgress.done();
});

export default router;
