import { Routes, RouterModule }  from '@angular/router';

import { CaseType } from './case-type.component';
import { CaseTypeList } from './list/list.component';

import { CaseTypeEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: CaseType,
    children: [
      { path: 'list', component: CaseTypeList },
      { path: 'edit/:id', component: CaseTypeEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
