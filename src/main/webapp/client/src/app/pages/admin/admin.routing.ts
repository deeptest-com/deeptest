import { Routes, RouterModule }  from '@angular/router';

import { Admin } from './admin.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Admin,
    children: [
      { path: 'user', loadChildren: () => System.import('./user/user.module')  },
      { path: 'group', loadChildren: () => System.import('./group/group.module')  },
      { path: 'role', loadChildren: () => System.import('./role/role.module')  }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
