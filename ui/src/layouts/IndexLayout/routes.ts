import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    icon: 'project',
    title: 'index-layout.menu.project',
    path: '/project',
    redirect: '/project/list',
    component: BlankLayout,
    children: [
      {
        title: 'index-layout.menu.project.list',
        path: 'list',
        component: () => import('@/views/project/list/index.vue'),
        hidden: true,
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
    icon: 'interface',
    title: 'index-layout.menu.interface',
    path: '/interface',
    redirect: '/interface/index',
    component: BlankLayout,
    children: [
      {
        title: 'index-layout.menu.interface',
        path: 'index',
        component: () => import('@/views/interface/index.vue'),
        hidden: true,
      },
    ],
  },
];

export default IndexLayoutRoutes;