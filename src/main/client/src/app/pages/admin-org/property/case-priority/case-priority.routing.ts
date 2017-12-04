import { Routes, RouterModule }  from '@angular/router';

import { CasePriority } from './case-priority.component';
import { CasePriorityList } from './list/list.component';

import { CasePriorityEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: CasePriority,
    children: [
      { path: 'list', component: CasePriorityList },
      { path: 'edit/:id', component: CasePriorityEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
