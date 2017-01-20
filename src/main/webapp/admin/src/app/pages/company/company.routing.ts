import { Routes, RouterModule }  from '@angular/router';

import { Company } from './company.component';
import { CompanyEdit } from './company-edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Company,
    children: [
      { path: 'edit/:id', component: CompanyEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
