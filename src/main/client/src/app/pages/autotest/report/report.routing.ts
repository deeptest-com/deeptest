import { Routes, RouterModule }  from '@angular/router';

import { Report } from './report.component';
import { ReportList } from './list/list.component';
import { ReportView } from './view/view.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Report,
    children: [
      { path: '', component: ReportList },
      { path: ':id', component: ReportView },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
