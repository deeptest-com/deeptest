import { Routes, RouterModule }  from '@angular/router';
import { ModuleWithProviders } from '@angular/core';

import { Org } from './org.component';

// noinspection TypeScriptValidateTypes

export const routes: Routes = [
  {
    path: ':orgId',
    component: Org,

    children: [
      { path: 'prjs', loadChildren: '../../project/project/project.module#ProjectModule' },
      { path: 'prj', loadChildren: './prj/prj.module#PrjModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
