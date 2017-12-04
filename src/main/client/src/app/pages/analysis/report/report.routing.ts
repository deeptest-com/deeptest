import { Routes, RouterModule }  from '@angular/router';

import { Report } from './report.component';
import { ReportList } from './list/list.component';
import { ReportEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Report,
    children: [
      { path: 'list', component: ReportList },
      { path: ':reportId/edit', component: ReportEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
