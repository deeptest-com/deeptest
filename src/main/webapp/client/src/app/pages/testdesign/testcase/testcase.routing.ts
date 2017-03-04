import { Routes, RouterModule }  from '@angular/router';

import { Testcase } from './testcase.component';
import { TestcaseList } from './list/list.component';
import { TestcaseEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Testcase,
    children: [
      
    ]
  }
];

export const routing = RouterModule.forChild(routes);
