import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    title: 'workbench',
    path: '/:projectId/workspace',
    redirect: '/:projectId/workspace',
    component: BlankLayout,
    children: [
      {
        icon: 'workspace',
        title: 'workbench',
        path: '',
        component: () => import('@/views/workbench/index.vue'),
        hidden: false,
        meta: {
          code: 'WORKSPACE'
        }
      },
      {
        icon: 'workspace',
        title: 'workplace',
        path: ':id',
        component: () => import('@/views/workbench/index.vue'),
        hidden: true,
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/:projectId/IM',
    redirect: '/:projectId/IM',
    component: BlankLayout,
    children: [
      {
        icon: 'define',
        title: 'endpoint-management',
        path: '',
        component: () => import('@/views/endpoint/index.vue'),
        hidden: false,
        meta: {
          code: 'IM'
        }
      }
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/:projectId/docs',
    redirect: '/:projectId/docs',
    component: BlankLayout,
    children: [
      {
        icon: 'docs',
        title: 'endpoint-docs',
        path: '',
        component: () => import('@/views/docs/index.vue'),
        hidden: false,
        meta: {
          code: 'DOCS'
        }
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/:projectId/debug',
    redirect: '/:projectId/debug',
    component: BlankLayout,
    children: [
      {
        icon: 'debug',
        title: 'diagnose',
        path: '',
        component: () => import('@/views/diagnose/index.vue'),
        hidden: false,
        meta: {
          code: 'DEBUG'
        }
      },
    ],
  },
  {
    title: 'index-layout.menu',
    path: '/:projectId/TS',
    redirect: '/:projectId/TS',
    component: BlankLayout,
    children: [
      {
        icon: 'test',
        title: 'scenario',
        path: '',
        component: () => import('@/views/scenario/index.vue'),
        hidden: false,
        meta: {
          code: 'TS'
        }
      }
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/:projectId/TP',
    redirect: '/:projectId/TP',
    component: BlankLayout,
    children: [
      {
        icon: 'tp',
        title: 'plan',
        path: '',
        component: () => import('@/views/plan/index.vue'),
        hidden: false,
        meta: {
          code: 'TP'
        }
      },
      {
        icon: 'tp',
        title: 'plan.exec',
        path: 'exec/:id',
        component: () => import('@/views/plan/exec/index.vue'),
        hidden: true,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/:projectId/TR',
    redirect: '/:projectId/TR',
    component: () => import('@/views/report/index.vue'),
    children: [
      {
        icon: 'tr',
        title: 'report',
        path: '',
        component: () => import('@/views/report/index.vue'),
        hidden: false,
        meta: {
          code: 'TR'
        }
      }
    ],
  },

   // 项目管理
  {
    title: '项目设置',
    path: '/:projectId/project-setting',
    redirect: '/:projectId/project-setting',
    component: () => import('@/views/project-settings/index.vue'),
    children: [
      {
        icon:'set',
        title: 'projectSetting',
        path: '',
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
