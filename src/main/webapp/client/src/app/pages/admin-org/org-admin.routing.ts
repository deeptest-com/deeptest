import { Routes, RouterModule }  from '@angular/router';

import { OrgAdmin } from './org-admin.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: OrgAdmin,
    children: [
      { path: 'org', loadChildren: './org/org.module' },
      { path: 'user', loadChildren: './user/user.module'  },
      { path: 'group', loadChildren: './group/group.module' },
      { path: 'org-role', loadChildren: './org-role/org-role.module' },
      { path: 'project-role', loadChildren: './project-role/project-role.module' },
      { path: 'property', loadChildren: './property/property.module' }
    ]
  }
];
 
export const routing = RouterModule.forChild(routes);
