import { Routes, RouterModule }  from '@angular/router';
import { ModuleWithProviders } from '@angular/core';

import { AutoTest } from './autotest.component';
import { Report } from './report/report.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: AutoTest,
    children: [
      { path: 'report', loadChildren: './report/report.module#ReportModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
