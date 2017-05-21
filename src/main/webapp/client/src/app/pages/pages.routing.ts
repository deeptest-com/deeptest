import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
import { ModuleWithProviders } from '@angular/core';
// noinspection TypeScriptValidateTypes

// export function loadChildren(path) { return System.import(path); };

export const routes: Routes = [
  {
    path: 'login',
    loadChildren: './account/login/login.module'
  },
  {
    path: 'register',
    loadChildren: './account/register/register.module'
  },
  {
    path: 'forgot-password',
    loadChildren: './account/forgot-password/forgot-password.module'
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
