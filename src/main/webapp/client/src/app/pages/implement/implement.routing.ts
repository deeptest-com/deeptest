import { Routes, RouterModule }  from '@angular/router';

import { Implement } from './implement.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Implement,
    children: [
      { path: ':projectId/plan', loadChildren: './plan/plan.module#PlanModule' },
      { path: ':projectId/run', loadChildren: './run/run.module#RunModule' },
      { path: ':projectId/plan/:planId/execution', loadChildren: './execution/execution.module#ExecutionModule' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);

