import { Routes, RouterModule }  from '@angular/router';

import { ForgotPassword } from './forgot-password.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: ForgotPassword
  }
];

export const routing = RouterModule.forChild(routes);
