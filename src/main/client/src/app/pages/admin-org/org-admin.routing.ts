import { Routes, RouterModule }  from '@angular/router';

import { OrgAdmin } from './org-admin.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: OrgAdmin,
    children: [
      { path: 'org', loadChildren: './org/org.module#OrgModule' },
      { path: 'user', loadChildren: './user/user.module#UserModule'  },
      { path: 'group', loadChildren: './group/group.module#GroupModule' },
      { path: 'org-role', loadChildren: './org-role/org-role.module#OrgRoleModule' },
      { path: 'project-role', loadChildren: './project-role/project-role.module#ProjectRoleModule' },
      { path: 'property', loadChildren: './property/property.module#PropertyModule' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
