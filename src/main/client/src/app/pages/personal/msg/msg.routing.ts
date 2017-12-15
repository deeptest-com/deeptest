import { Routes, RouterModule }  from '@angular/router';

import { Msg } from './msg.component';
import { MsgList } from './list/list.component';

// noinspection TypeScriptValidateTypes
const routes: Routes = [
  {
    path: '',
    component: Msg,
    children: [
      { path: 'list', component: MsgList }
    ]
  }
];

export const routing = RouterModule.forChild(routes);
