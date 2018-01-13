import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
import { ModuleWithProviders } from '@angular/core';
// noinspection TypeScriptValidateTypes

// export function loadChildren(path) { return System.import(path); };

import { PagesResolve } from './pages.resolve';

export const routes: Routes = [
  {
    path: 'login',
    loadChildren: './account/login/login.module#LoginModule'
  },
  {
    path: 'register',
    loadChildren: './account/register/register.module#RegisterModule'
  },
  {
    path: 'forgot-password',
    loadChildren: './account/forgot-password/forgot-password.module#ForgotPasswordModule'
  },
  {
    path: 'reset-password/:vcode',
    loadChildren: './account/reset-password/reset-password.module#ResetPasswordModule'
  },
  {
    path: 'pages',
    component: Pages,
    resolve: {
      data: PagesResolve
    },
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadChildren: './dashboard/dashboard.module#DashboardModule' },
      { path: 'personal', loadChildren: './personal/personal.module#PersonalModule' },

      { path: 'org-admin', loadChildren: './admin-org/org-admin.module#OrgAdminModule' },
      { path: 'sys-admin', loadChildren: './admin-sys/sys-admin.module#AdminModule' },

      { path: 'org', loadChildren: './_context/org/org.module#OrgModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);

