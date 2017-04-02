import { Routes, RouterModule }  from '@angular/router';

import { Role } from './role.component';
import { RoleList } from './list/list.component';
import { RoleEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Role,
    children: [
      { path: 'list', component: RoleList },
      { path: 'edit/:id', component: RoleEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
