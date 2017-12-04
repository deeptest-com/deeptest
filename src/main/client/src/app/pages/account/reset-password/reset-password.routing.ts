import { Routes, RouterModule }  from '@angular/router';

import { ResetPassword } from './reset-password.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: ResetPassword
  }
];

export const routing = RouterModule.forChild(routes);
