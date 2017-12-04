import { Routes, RouterModule }  from '@angular/router';

import { Property } from './property.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Property,
    children: [
      { path: 'custom-field', loadChildren: './custom-field/custom-field.module#CustomFieldModule' },
      { path: 'case-type', loadChildren: './case-type/case-type.module#CaseTypeModule' },
      { path: 'case-priority', loadChildren: './case-priority/case-priority.module#CasePriorityModule' },
      { path: 'case-exe-status', loadChildren: './case-exe-status/case-exe-status.module#CaseExeStatusModule' }
    ]
  }
];

export const routing = RouterModule.forChild(routes);

