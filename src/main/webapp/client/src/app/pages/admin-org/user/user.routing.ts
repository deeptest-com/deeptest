import { Routes, RouterModule }  from '@angular/router';

import { User } from './user.component';
import { UserList } from './list/list.component';

import { UserEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: User,
    children: [
      { path: 'list', component: UserList },
      { path: 'edit/:id', component: UserEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
