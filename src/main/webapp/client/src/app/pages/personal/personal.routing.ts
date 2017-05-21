import { Routes, RouterModule }  from '@angular/router';

import { Personal } from './personal.component';

import { Password } from './password';
import { PasswordEdit } from './password/password-edit';

import { Profile } from './profile';
import { ProfileEdit } from './profile/profile-edit';

const routes: Routes = [
  {
    path: '',
    component: Personal,
    children: [
      { path: 'password', component: PasswordEdit },
      { path: 'profile', component: ProfileEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
