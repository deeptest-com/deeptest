import { Routes, RouterModule }  from '@angular/router';

import { Plan } from './plan.component';
import { PlanList } from './list/list.component';
import { PlanEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Plan,
    children: [
      { path: 'list', component: PlanList },
      { path: 'edit/:id', component: PlanEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
