import { Routes, RouterModule }  from '@angular/router';

import { Project } from './project.component';
import { ProjectList } from './list/list.component';
import { ProjectEdit } from './edit/edit.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Project,
    children: [
      { path: 'list', component: ProjectList },
      { path: 'edit/:type/:id', component: ProjectEdit },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
