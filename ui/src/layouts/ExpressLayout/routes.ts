import { RoutesDataItem } from "@/utils/routes";

const ExpressLayoutRoutes: RoutesDataItem[] = [
  {
    icon: 'empty',
    title: 'empty',
    path: 'index',
    component: () => import('@/views/express/index.vue')
  },
]

export default ExpressLayoutRoutes;
