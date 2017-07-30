import { Routes, RouterModule }  from '@angular/router';

import { Execution } from './execution.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Execution,
    children: [

    ]
  }
];

export const routing = RouterModule.forChild(routes);
