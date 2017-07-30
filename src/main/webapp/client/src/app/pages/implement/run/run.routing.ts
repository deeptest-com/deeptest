import { Routes, RouterModule }  from '@angular/router';

import { Run } from './run.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Run,
    children: [
    ]
  }
];

export const routing = RouterModule.forChild(routes);
