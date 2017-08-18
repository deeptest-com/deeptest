import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
import { ModuleWithProviders } from '@angular/core';
// noinspection TypeScriptValidateTypes

// export function loadChildren(path) { return System.import(path); };

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
    path: 'pages',
    component: Pages,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadChildren: './dashboard/dashboard.module#DashboardModule' },

      { path: 'project', loadChildren: './project/project/project.module#ProjectModule' },
      { path: 'design', loadChildren: './design/design.module#DesignModule' },
      { path: 'implement', loadChildren: './implement/implement.module#ImplementModule' },
      { path: 'analysis', loadChildren: './analysis/analysis.module#AnalysisModule' },

      { path: 'org-admin', loadChildren: './admin-org/org-admin.module#OrgAdminModule' },
      { path: 'sys-admin', loadChildren: './admin-sys/sys-admin.module#AdminModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);

