import { Routes, RouterModule }  from '@angular/router';

import { OrgAdmin } from './org-admin.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: OrgAdmin,
    children: [
      { path: 'org', loadChildren: () => System.import('./org/org.module')  },
      { path: 'user', loadChildren: () => System.import('./user/user.module')  },
      { path: 'group', loadChildren: () => System.import('./group/group.module')  },
      { path: 'role', loadChildren: () => System.import('./role/role.module')  }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
