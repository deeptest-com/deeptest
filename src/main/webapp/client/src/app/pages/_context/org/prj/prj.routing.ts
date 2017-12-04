import { Routes, RouterModule }  from '@angular/router';
import { ModuleWithProviders } from '@angular/core';

import { Prj } from './prj.component';

// noinspection TypeScriptValidateTypes

export const routes: Routes = [
  {
    path: ':prjId',
    component: Prj,
    children: [
      { path: 'design', loadChildren: '../../../design/design.module#DesignModule' },
      { path: 'implement', loadChildren: '../../../implement/implement.module#ImplementModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
