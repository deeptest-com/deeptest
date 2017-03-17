import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: 'login',
    loadChildren: () => System.import('./user/login/login.module')
  },
  {
    path: 'register',
    loadChildren: () => System.import('./user/register/register.module')
  },
  {
    path: 'forgot-password',
    loadChildren: () => System.import('./user/forgot-password/forgot-password.module')
  },
  {
    path: 'pages',

    component: Pages,
    children: [
      { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
      { path: 'dashboard', loadChildren: () => System.import('./dashboard/dashboard.module') },

      { path: 'project', loadChildren: () => System.import('./project/project/project.module') },
      { path: 'case', loadChildren: () => System.import('./design/case/case.module') },

      { path: 'event', loadChildren: () => System.import('./event/event.module') },
      { path: 'business', loadChildren: () => System.import('./business/business.module') },
      { path: 'personal', loadChildren: () => System.import('./personal/personal.module') },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
