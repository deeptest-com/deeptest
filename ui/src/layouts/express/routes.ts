import { RoutesDataItem } from "@/utils/routes";
import BlankLayout from '@/layouts/BlankLayout.vue';

const IndexLayoutRoutes: Array<RoutesDataItem> = [
  {
    title: 'express-layout.deeptest',
    path: '/express',
    redirect: '/express/index',
    component: BlankLayout,
    children: [
      {
        icon: 'interface',
        title: 'interface',
        path: 'index',
        component: () => import('@/views/express/index.vue'),
        hidden: false,
      },
    ],
  }
] as Array<RoutesDataItem>;

export default IndexLayoutRoutes;
