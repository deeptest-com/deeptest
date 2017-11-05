import { Routes, RouterModule }  from '@angular/router';

import { Personal } from './personal.component';

import { Profile } from './profile';
import { ProfileEdit } from './profile/profile-edit';

import { Settings } from './settings';
import { SettingsEdit } from './settings/settings-edit';

const routes: Routes = [
  {
    path: '',
    component: Personal,
    children: [
      { path: 'profile', component: ProfileEdit },
      { path: 'settings', component: SettingsEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
