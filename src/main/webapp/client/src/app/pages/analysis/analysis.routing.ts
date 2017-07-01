import { Routes, RouterModule }  from '@angular/router';

import { Analysis } from './analysis.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Analysis,
    children: [
      { path: 'report', loadChildren: './report/report.module' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);

