import { Routes, RouterModule }  from '@angular/router';

import { Personal } from './personal.component';

import { ProfileEdit } from './profile/profile-edit';

const routes: Routes = [
  {
    path: '',
    component: Personal,
    children: [
      { path: 'msg', loadChildren: './msg/msg.module#MsgModule' },
      { path: 'profile', component: ProfileEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
