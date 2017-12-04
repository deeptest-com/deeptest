import { Routes, RouterModule }  from '@angular/router';

import { OrgRole } from './org-role.component';
import { OrgRoleList } from './list/list.component';
import { OrgRoleEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: OrgRole,
    children: [
      { path: 'list', component: OrgRoleList },
      { path: 'edit/:id', component: OrgRoleEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
