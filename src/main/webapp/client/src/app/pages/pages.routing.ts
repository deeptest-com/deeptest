import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
import { ModuleWithProviders } from '@angular/core';
// noinspection TypeScriptValidateTypes

// export function loadChildren(path) { return System.import(path); };

export const routes: Routes = [
  {
    path: 'login',
    loadChildren: 'app/pages/login/login.module#LoginModule'
  },
  {
    path: 'register',
    loadChildren: 'app/pages/register/register.module#RegisterModule'
  },
  {
    path: 'pages',
    component: Pages,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadChildren: './dashboard/dashboard.module#DashboardModule' },

      { path: 'project', loadChildren: './project/project/project.module' },
      { path: 'case', loadChildren: './design/case/case.module' },

      { path: 'org-admin', loadChildren: './admin-org/org-admin.module' },
      { path: 'sys-admin', loadChildren: './admin-sys/sys-admin.module' }

    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
