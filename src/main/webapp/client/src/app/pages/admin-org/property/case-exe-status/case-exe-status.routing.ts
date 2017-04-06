import { Routes, RouterModule }  from '@angular/router';

import { CaseExeStatus } from './case-exe-status.component';
import { CaseExeStatusList } from './list/list.component';

import { CaseExeStatusEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: CaseExeStatus,
    children: [
      { path: 'list', component: CaseExeStatusList },
      { path: 'edit/:id', component: CaseExeStatusEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
