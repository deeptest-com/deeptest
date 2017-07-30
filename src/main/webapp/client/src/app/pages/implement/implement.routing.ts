import { Routes, RouterModule }  from '@angular/router';

import { Implement } from './implement.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Implement,
    children: [
      { path: 'plan', loadChildren: './plan/plan.module' },
      { path: 'run', loadChildren: './run/run.module' },
      { path: 'plan/:planId/execution/:runId', loadChildren: './execution/execution.module' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);

