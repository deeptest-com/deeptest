import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: 'login',
    loadChildren: () => System.import('./account/login/login.module')
  },
  {
    path: 'register',
    loadChildren: () => System.import('./account/register/register.module')
  },
  {
    path: 'forgot-password',
    loadChildren: () => System.import('./account/forgot-password/forgot-password.module')
  },
  {
    path: 'pages',

    component: Pages,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadChildren: () => System.import('./dashboard/dashboard.module') },

      { path: 'project', loadChildren: () => System.import('./project/project/project.module') },
      { path: 'case', loadChildren: () => System.import('./design/case/case.module') },

      { path: 'admin', loadChildren: () => System.import('./admin/admin.module') }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
