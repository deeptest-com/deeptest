import { Routes, RouterModule }  from '@angular/router';

import { Run } from './run.component';
import { Execution } from './execution/execution.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Run,
    children: [
      { path: 'execution/:id', component: Execution }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
