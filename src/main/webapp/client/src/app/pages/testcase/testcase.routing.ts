import { Routes, RouterModule }  from '@angular/router';

import { Testcase } from './testcase.component';
import { TestcaseList } from './testcase-list/list.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Testcase,
    children: [
      { path: 'list', component: TestcaseList }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
