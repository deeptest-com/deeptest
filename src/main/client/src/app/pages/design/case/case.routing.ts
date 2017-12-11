import { Routes, RouterModule }  from '@angular/router';

import { Case } from './case.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Case,
    children: [
    ]
  },
  {
    path: ':caseId',
    component: Case,
    children: [
    ]
  },
];

export const routing = RouterModule.forChild(routes);
