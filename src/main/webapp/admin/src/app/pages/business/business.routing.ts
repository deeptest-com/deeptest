import { Routes, RouterModule }  from '@angular/router';

import { Business } from './business.component';

import { Account } from './account';
import { AccountList } from './account/account-list';
import { AccountEdit } from './account/account-edit';

import { Company } from './company';
import { CompanyEdit } from './company/company-edit';

import { Profile } from './profile';
import { ProfileEdit } from './profile/profile-edit';

const routes: Routes = [
  {
    path: '',
    component: Business,
    children: [
      { path: 'account-list', component: AccountList },
      { path: 'account-edit/:id', component: AccountEdit },
      { path: 'company-edit', component: CompanyEdit },
      { path: 'profile-edit', component: ProfileEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
