import { Routes, RouterModule }  from '@angular/router';

import { Design } from './design.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Design,
    children: [
      { path: 'case', loadChildren: './case/case.module#CaseModule' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);

