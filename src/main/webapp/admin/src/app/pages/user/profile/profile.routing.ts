import { Routes, RouterModule }  from '@angular/router';

import { Profile } from './profile.component';
import { ProfileEdit } from './profile-edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Profile,
    children: [
      { path: 'edit', component: ProfileEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
