import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    title: 'workbench',
    path: '/workbench',
    redirect: '/workbench/index',
    component: BlankLayout,
    children: [
      {
        icon: 'home',
        title: 'workbench',
        path: 'index',
        component: () => import('@/views/workbench/index.vue'),
        hidden: false,
        meta: {
          code: 'WORKBENCH'
        }
      },
      {
        icon: 'home',
        title: 'workplace',
        path: ':id',
        component: () => import('@/views/workbench/index.vue'),
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
        title: 'endpoint-management',
        path: 'index',
        component: () => import('@/views/endpoint/index.vue'),
        hidden: false,
        meta: {
          code: 'ENDPOINT'
        }
      }
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/docs',
    redirect: '/docs',
    component: BlankLayout,
    children: [
      {
        icon: 'doc',
        title: 'endpoint-docs',
        path: 'index',
        component: () => import('@/views/docs/index.vue'),
        hidden: false,
        meta: {
          code: 'ENDPOINT'
        }
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/diagnose',
    redirect: '/diagnose/index',
    component: BlankLayout,
    children: [
      {
        icon: 'diagnose',
        title: 'diagnose',
        path: 'index',
        component: () => import('@/views/diagnose/index.vue'),
        hidden: false,
        meta: {
          code: 'DIAGNOSE'
        }
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
        meta: {
          code: 'SCENARIO'
        }
      }
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
        meta: {
          code: 'PLAN'
        }
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
    component: () => import('@/views/report/index.vue'),
    children: [
      {
        icon: 'report',
        title: 'report',
        path: 'index',
        component: () => import('@/views/report/index.vue'),
        hidden: false,
        meta: {
          code: 'REPORT'
        }
      }
    ],
  },

   // 项目管理
  {
    title: '项目设置',
    path: '/project-setting',
    redirect: '/project-setting/index',
    component: () => import('@/views/project-settings/index.vue'),
    children: [
      {
        icon:'set',
        title: 'projectSetting',
        path: 'index',
        component: BlankLayout,
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING'
        }
      },
      {
        icon:'set',
        title: 'projectSetting.enviroment',
        path: 'enviroment',
        name: 'enviroment',
        component: () => import('@/views/project-settings/components/EnvSetting/index.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-ENVIRONMENT'
        },
        children: [
          {
            icon: 'set',
            title: 'envsetting.var',
            path: 'var',
            name: 'enviroment.var',
            component: () => import('@/views/project-settings/components/EnvSetting/GlobalVar.vue'),
            hidden: true
          },
          {
            icon: 'set',
            title: 'envsetting.params',
            path: 'params',
            name: 'enviroment.params',
            component: () => import('@/views/project-settings/components/EnvSetting/GlobalParams.vue'),
            hidden: true
          },
          {
            icon: 'set',
            title: 'envsetting.envdetail',
            path: 'envdetail/:id(\\d+)?',
            name: 'enviroment.envdetail',
            component: () => import('@/views/project-settings/components/EnvSetting/EnvDetail.vue'),
            hidden: true
          }
        ]
      },
      {
        icon:'set',
        title: 'projectSetting.datapool',
        path: 'datapool',
        name: 'datapool',
        component: () => import('@/views/project-settings/components/DataPool/index.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-DATA-POOL'
        }
      },
      {
        icon:'set',
        title: 'projectSetting.service',
        path: 'service-setting',
        name: 'service-setting',
        component: () => import('@/views/project-settings/components/ServiceSetting/index.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-SERVICE-SETTING'
        }
      },
      {
        icon: 'members',
        title: 'project.members',
        path: 'members',
        component: () => import('@/views/project/edit/members.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-MEMBERS'
        }
      },
      {
        icon: 'mock',
        title: 'projectSetting.mock',
        path: 'mock',
        component: () => import('@/views/project-settings/components/Mock/index.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-MOCK'
        }
      },
      {
        icon: 'swaggerSync',
        title: 'projectSetting.swaggerSync',
        path: 'swaggerSync',
        component: () => import('@/views/project-settings/components/SwaggerSync/index.vue'),
        hidden: false,
        meta: {
          code: 'PROJECT-SETTING-SWAGGERSYNC'
        }
      },
    ],
  },

] as Array<RoutesDataItem>;

export default IndexLayoutRoutes;
