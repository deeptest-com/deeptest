import { Routes, RouterModule }  from '@angular/router';

import { Account } from './account.component';
import { AccountList } from './account-list/list.component';
import { AccountEdit } from './account-edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Account,
    children: [
      { path: 'list', component: AccountList },
      { path: 'edit/:id', component: AccountEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
