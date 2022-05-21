import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    title: 'index-layout.menu',
    path: '/project',
    redirect: '/project/list',
    component: BlankLayout,
    children: [
      {
        icon: 'project',
        title: 'project',
        path: 'list',
        component: () => import('@/views/project/list/index.vue'),
        hidden: false,
      },
      {
        icon: 'project',
        title: 'project.edit',
        path: 'edit/:id',
        component: () => import('@/views/project/edit/index.vue'),
        hidden: true,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/interface',
    redirect: '/interface/index',
    component: BlankLayout,
    children: [
      {
        icon: 'interface',
        title: 'interface',
        path: 'index',
        component: () => import('@/views/interface/index.vue'),
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
        path: 'list',
        component: () => import('@/views/scenario/list/index.vue'),
        hidden: false,
      },
      {
        icon: 'scenario',
        title: 'scenario.edit',
        path: 'edit/:id',
        component: () => import('@/views/scenario/edit/index.vue'),
        hidden: true,
      },
    ],
  },
];

export default IndexLayoutRoutes;