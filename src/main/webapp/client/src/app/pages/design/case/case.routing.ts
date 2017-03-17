import { Routes, RouterModule }  from '@angular/router';

import { Case } from './case.component';
import { CaseList } from './list/list.component';
import { CaseEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Case,
    children: [
      
    ]
  }
];

export const routing = RouterModule.forChild(routes);
