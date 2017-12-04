import { Routes, RouterModule }  from '@angular/router';

import { Project } from './project.component';
import { ProjectList } from './list/list.component';
import { ProjectEdit } from './edit/edit.component';
import { ProjectView } from './view/view.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Project,
    children: [
      { path: '', component: ProjectList },
      { path: ':id/edit/:type', component: ProjectEdit },
      { path: ':id/view', component: ProjectView },
    ]
  }
];

export const routing = RouterModule.forChild(routes);
