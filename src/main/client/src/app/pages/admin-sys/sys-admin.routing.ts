import { Routes, RouterModule }  from '@angular/router';

import { SysAdmin } from './sys-admin.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: SysAdmin,
    // children: [
    //   { path: 'user', loadChildren: () => System.import('./user/user.module')  },
    //   { path: 'group', loadChildren: () => System.import('./group/group.module')  },
    //   { path: 'role', loadChildren: () => System.import('./role/role.module')  }
    // ]
  }
];

export const routing = RouterModule.forChild(routes);
