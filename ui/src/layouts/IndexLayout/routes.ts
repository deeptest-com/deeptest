import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    title: 'index-layout.menu',
    path: '/project',
    redirect: '/project/index',
    component: BlankLayout,
    children: [
      {
        icon: 'project',
        title: 'project.management',
        path: 'index',
        component: () => import('@/views/project/list/index.vue'),
        hidden: false,
      },
      {
        icon: 'project',
        title: 'project.edit',
        path: 'edit/:id',
        component: () => import('@/views/project/edit/edit.vue'),
        hidden: true,
      },
      {
        icon: 'members',
        title: 'project.members',
        path: 'members/:id',
        component: () => import('@/views/project/edit/members.vue'),
        hidden: true,
      },
      {
        icon: 'project',
        title: 'project.invite',
        path: 'invite/:id',
        component: () => import('@/views/project/edit/invite.vue'),
        hidden: true,
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/endpoint',
    redirect: '/endpoint/index',
    component: BlankLayout,
    children: [
      {
        icon: 'endpoint',
        title: 'endpoint',
        path: 'index',
        component: () => import('@/views/endpoint/index.vue'),
        hidden: false,
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/interface1',
    redirect: '/interface1/index',
    component: BlankLayout,
    children: [
      {
        icon: 'interface',
        title: 'interface1',
        path: 'index',
        component: () => import('@/views/interface1/index.vue'),
        hidden: false,
      },
    ],
  },

  // 项目管理
  {
    title: '项目设置',
    path: '/projectSetting',
    redirect: '/projectSetting/index',
    component: BlankLayout,
    children: [
      {
        icon:'set',
        title: 'projectSetting',
        path: 'index',
        component: () => import('@/views/project-setting/index.vue'),
        hidden: false,
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/scenario',
    redirect: '/scenario/index',
    component: BlankLayout,
    children: [
      {
        icon: 'scenario',
        title: 'scenario',
        path: 'index',
        component: () => import('@/views/scenario/index.vue'),
        hidden: false,
      },
      {
        icon: 'scenario',
        title: 'scenario.exec',
        path: 'exec/:id',
        component: () => import('@/views/scenario/exec/index.vue'),
        hidden: true,
      },
      {
        icon: 'scenario',
        title: 'scenario.design',
        path: 'design/:id',
        component: () => import('@/views/scenario/design/index.vue'),
        hidden: true,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/plan',
    redirect: '/plan/index',
    component: BlankLayout,
    children: [
      {
        icon: 'plan',
        title: 'plan',
        path: 'index',
        component: () => import('@/views/plan/index.vue'),
        hidden: false,
      },
      {
        icon: 'plan',
        title: 'plan.exec',
        path: 'exec/:id',
        component: () => import('@/views/plan/exec/index.vue'),
        hidden: true,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/report',
    redirect: '/report/index',
    component: BlankLayout,
    children: [
      {
        icon: 'report',
        title: 'report',
        path: 'index',
        component: () => import('@/views/report/list/index.vue'),
        hidden: false,
      },
      {
        icon: 'report',
        title: 'report.detail',
        path: ':id',
        component: () => import('@/views/report/detail/index.vue'),
        hidden: true,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/user',
    component: BlankLayout,
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
    ],
  },

] as Array<RoutesDataItem>;

export default IndexLayoutRoutes;
