import { Routes, RouterModule }  from '@angular/router';

import { Personal } from './personal.component';

import { MsgList } from './msg/list';
import { ProfileEdit } from './profile/profile-edit';
import { SettingsEdit } from './settings/settings-edit';

const routes: Routes = [
  {
    path: '',
    component: Personal,
    children: [
      { path: 'msg', loadChildren: './msg/msg.module#MsgModule' },
      { path: 'profile', component: ProfileEdit },
      { path: 'settings', component: SettingsEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
