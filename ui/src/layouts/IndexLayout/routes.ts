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
        title: 'index-layout.menu.project',
        path: 'list',
        component: () => import('@/views/project/list/index.vue'),
        hidden: false,
      },
      {
        title: 'index-layout.menu.project.edit',
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
        title: 'index-layout.menu.interface',
        path: 'index',
        component: () => import('@/views/interface/index.vue'),
        hidden: false,
      },
    ],
  },

  {
    title: 'index-layout.menu',
    path: '/scene',
    redirect: '/scene/index',
    component: BlankLayout,
    children: [
      {
        icon: 'scene',
        title: 'index-layout.menu.scene',
        path: 'scene',
        component: () => import('@/views/interface/index.vue'),
        hidden: false,
      },
    ],
  },
];

export default IndexLayoutRoutes;