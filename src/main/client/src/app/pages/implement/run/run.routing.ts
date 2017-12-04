import { Routes, RouterModule }  from '@angular/router';

import { Run } from './run.component';
import { RunEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Run,
    children: [
      { path: ':id/edit', component: RunEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
