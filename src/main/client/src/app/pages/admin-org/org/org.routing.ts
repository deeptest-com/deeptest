import { Routes, RouterModule }  from '@angular/router';

import { Org } from './org.component';
import { OrgList } from './list/list.component';

import { OrgEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Org,
    children: [
      { path: 'list', component: OrgList },
      { path: 'edit/:id', component: OrgEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
