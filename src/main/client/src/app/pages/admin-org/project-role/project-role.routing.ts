import { Routes, RouterModule }  from '@angular/router';

import { ProjectRole } from './project-role.component';
import { ProjectRoleList } from './list/list.component';
import { ProjectRoleEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: ProjectRole,
    children: [
      { path: 'list', component: ProjectRoleList },
      { path: 'edit/:id', component: ProjectRoleEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
