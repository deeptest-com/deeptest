import { Routes, RouterModule }  from '@angular/router';

import { User } from './user.component';
import { UserList } from './list/list.component';

import { UserEdit } from './edit/edit.component';
import { UserEditInfo } from './edit/edit-info.component';
import { UserEditGroups } from './edit/edit-groups.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: User,
    children: [
      { path: 'list', component: UserList },
      { path: 'edit/:pid', component: UserEdit,
        children: [
          { path: 'info/:id', component: UserEditInfo },
          { path: 'groups/:id', component: UserEditGroups },
        ]
      }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
