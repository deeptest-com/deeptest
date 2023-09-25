/**
 * 路由入口
 * @author LiQingSong
 */
import NProgress from 'nprogress'; // progress bar
import 'nprogress/nprogress.css'; // progress bar style
NProgress.configure({ showSpinner: false, easing: 'ease', speed: 1000 }); // NProgress Configuration

import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router';
import { RoutesDataItem } from "@/utils/routes";
import settings from "@/config/settings";

import SecurityLayout from '@/layouts/SecurityLayout.vue';

import IndexLayoutRoutes from '@/layouts/IndexLayout/routes';
import IndexLayout from '@/layouts/IndexLayout/index.vue';

import UserLayoutRoutes from '@/layouts/UserLayout/routes';
import UserLayout from '@/layouts/UserLayout/index.vue';
import BlankLayout from "@/layouts/BlankLayout.vue";

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
        title: '分享文档',
        path: 'share',
        component: () => import('@/views/docs/index.vue'),
        hidden: true,
      },
    ],
  },
  {
    title: 'empty',
    path: '/docs',
    component: SecurityLayout,
    children: [
      {
        title: '查看文档',
        path: 'view',
        component: () => import('@/views/docs/index.vue'),
        hidden: true,
        meta: {
          title: '查看文档'
        }
      },
    ],
  },
  {
    title: 'empty',
    path: '/',
    redirect: '/',
    component: SecurityLayout,
    children: [
      {
        icon: 'home',
        title: '首页',
        path: '',
        component: () => import('@/views/home/index.vue'),
        meta: {
          title: '首页'
        }
      },
      {
        icon: 'profile',
        title: '个人信息',
        path: 'profile',
        component: () => import('@/views/user/info/profile.vue'),
        hidden: true,
        meta: {
          title: '个人信息'
        }
      },
      {
        icon: 'message',
        title: '消息',
        path: 'message',
        component: () => import('@/views/user/info/message.vue'),
        hidden: true,
        meta: {
          title: '消息'
        }
      }
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
  history: createWebHistory(process.env.BASE_URL),
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
