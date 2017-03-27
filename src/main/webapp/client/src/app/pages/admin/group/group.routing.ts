import { Routes, RouterModule }  from '@angular/router';

import { Group } from './group.component';
import { GroupList } from './list/list.component';
import { GroupEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Group,
    children: [
      { path: 'list', component: GroupList },
      { path: 'edit/:id', component: GroupEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
