import { Routes, RouterModule }  from '@angular/router';

import { Run } from './run.component';
import { RunList } from './list/list.component';
import { RunEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Run,
    children: [
      { path: 'list', component: RunList },
      { path: 'edit/:id', component: RunEdit }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
