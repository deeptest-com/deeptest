import { Routes, RouterModule }  from '@angular/router';

import { Property } from './property.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Property,
    children: [
      { path: 'custom-field', loadChildren: './custom-field/custom-field.module' },
      { path: 'case-type', loadChildren: './case-type/case-type.module' },
      { path: 'case-priority', loadChildren: './case-priority/case-priority.module' },
      { path: 'case-exe-status', loadChildren: './case-exe-status/case-exe-status.module' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
